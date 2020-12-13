package admin_api

import (
	libs_http "api/libs/http"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"reflect"
	"time"
)

/*
	Token验证成功，保存operator到gin.Context上下文
*/
func SetOperator(ctx *gin.Context, op *Operator) {
	ctx.Set("operator", op)
}

/*
	gin.Context上下文读取 operator 结构
*/
func GetOperator(ctx *gin.Context) (*Operator, error) {
	val, exists := ctx.Get("operator")
	if !exists {
		return nil, errors.New("Lose the operator info\n")
	}
	operator, ok := val.(*Operator)
	if !ok {
		return nil, errors.New("Assert operator type failed\n")
	}
	return operator, nil
}

/*
	operator数据结构
*/
type Operator struct {
	UserId    uint64
	UserType  uint64
	Username  string
	Nickname  string
	LoginTime time.Time
}

func (c *Operator) encrypt(key []byte) (string, error) {
	j, err := json.Marshal(c)
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

func (c *Operator) decrypt(key []byte, str string) error {
	token, err := jwt.Parse(str, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v\n", token.Header["alg"])
		}
		return key, nil
	})
	if err != nil {
		return err
	}

	if err := c.decode(token.Claims.(jwt.MapClaims), c); err != nil {
		return err
	}

	//if err := mapstructure.WeakDecode(token.Claims.(jwt.MapClaims), c); err != nil {
	//	return err
	//}
	return nil
}

func (c *Operator) toTimeHookFunc() mapstructure.DecodeHookFunc {
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
func (c *Operator) decode(input map[string]interface{}, result interface{}) error {
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Metadata:   nil,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(c.toTimeHookFunc()),
		Result:     result,
	})
	if err != nil {
		return err
	}

	if err := decoder.Decode(input); err != nil {
		return err
	}
	return err
}

/*
	获取操作者信息
	@method GET
*/
func (c *Operator) Info(ctx *gin.Context) {
	operator, err := GetOperator(ctx)
	if err != nil {
		libs_http.RspState(ctx, 1000, err)
		return
	}
	libs_http.RspData(ctx, 0, "获取成功", operator)
}
