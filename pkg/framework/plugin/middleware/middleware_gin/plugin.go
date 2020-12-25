package middleware_gin

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"framework/class/middleware"
	"framework/env"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/mitchellh/mapstructure"
	"reflect"
	"time"
)

func toTimeHookFunc() mapstructure.DecodeHookFunc {
	return func(
		f reflect.Type,
		t reflect.Type,
		data interface{}) (interface{}, error) {
		if t != reflect.TypeOf(time.Time{}) {
			return data, nil
		}

		switch f.Kind() {
		case reflect.String:
			return time.Parse(time.RFC3339, data.(string))
		case reflect.Float64:
			return time.Unix(0, int64(data.(float64))*int64(time.Millisecond)), nil
		case reflect.Int64:
			return time.Unix(0, data.(int64)*int64(time.Millisecond)), nil
		default:
			return data, nil
		}
		// Convert it by parsing
	}
}

var _ middleware.OperatorContext = &plugin{}

type plugin struct {
	model          middleware.Operator
	headerTokenKey string
	cipherKey      []byte
	redisClient    *redis.Client
	*Callback
	opts *Options
}

func (p *plugin) Init() error {
	if p.opts.model == nil {
		return errors.New("model is nil\n")
	}
	p.model = p.opts.model

	if p.opts.cipherKey == nil {
		return errors.New("cipher key is nil\n")
	}
	p.cipherKey = p.opts.cipherKey

	if p.opts.headerTokenKey == "" {
		p.opts.headerTokenKey = "X-API-TOKEN"
	}
	p.headerTokenKey = p.opts.headerTokenKey

	return nil
}

func (p *plugin) set(ctx *gin.Context, operator middleware.Operator) {
	ctx.Set("Operator", operator)
}
func (p *plugin) get(ctx *gin.Context) (middleware.Operator, error) {
	val, ok := ctx.Get("Operator")
	if !ok {
		return nil, errors.New("The operator info is not exists\n")
	}
	operator, ok := val.(middleware.Operator)
	if !ok {
		return nil, errors.New("Assert operator type failed\n")
	}
	return operator, nil
}

func (p *plugin) Get(args ...interface{}) (middleware.Operator, error) {
	ctx := args[0].(*gin.Context)
	return p.get(ctx)
}
func (p *plugin) encrypt(key []byte, operator middleware.Operator) (string, error) {
	j, err := json.Marshal(operator)
	if err != nil {
		return "", err
	}
	m := jwt.MapClaims{}
	if err := json.Unmarshal(j, &m); err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, m)
	str, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return str, nil
}
func (p *plugin) decrypt(key []byte, str string, operator middleware.Operator) error {
	// parse the string to be token
	token, err := jwt.Parse(str, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v\n", token.Header["alg"])
		}
		return key, nil
	})
	if err != nil {
		return err
	}

	// convert the claims to model
	{
		decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
			DecodeHook:       mapstructure.ComposeDecodeHookFunc(toTimeHookFunc()),
			ErrorUnused:      false,
			ZeroFields:       false,
			WeaklyTypedInput: false,
			Squash:           false,
			Metadata:         nil,
			Result:           operator,
			TagName:          "",
		})
		if err != nil {
			return err
		}

		input := token.Claims.(jwt.Claims)

		if err := decoder.Decode(input); err != nil {
			return err
		}
	}

	return nil
}

func (p *plugin) SignedString(args ...interface{}) (string, error) {
	operator := args[0].(middleware.Operator)
	return p.encrypt(p.cipherKey, operator)
}

func (p *plugin) Parse(args ...interface{}) interface{} {
	return func(ctx *gin.Context) {
		// read the token string
		str := ctx.GetHeader(p.headerTokenKey)
		if str == "" {
			desc := fmt.Sprintf("读取上下文(%s)失败", p.headerTokenKey)
			env.Logger.Error(desc)
			ctx.Abort()
			return
		}

		// create new model instance
		t := reflect.TypeOf(p.model)
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		}
		operator := reflect.New(t).Interface()

		// decrypt the string to the struct
		if err := p.decrypt(p.cipherKey, str, operator.(middleware.Operator)); err != nil {
			env.Logger.Error(err)
			ctx.Abort()
			return
		}

		// save the operator to context
		p.set(ctx, operator.(middleware.Operator))
	}
}

func (p *plugin) Expiration(args ...interface{}) interface{} {
	return func(ctx *gin.Context) {
		operator, err := p.get(ctx)
		if err != nil {
			env.Logger.Error(err)
			ctx.Abort()
			return
		}

		key := fmt.Sprintf("ctx:operator:%s", operator.GetContextId())
		num, err := p.redisClient.Exists(context.TODO(), key).Result()
		if err != nil {
			env.Logger.Error(err)
			ctx.Abort()
			return
		}
		if num == 0 {
			p.Callback.Expiration(ctx)
			ctx.Abort()
			return
		}
	}
}
