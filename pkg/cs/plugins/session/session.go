package session

import (
	"context"
	"cs/class/session"
	"cs/meta"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/mitchellh/mapstructure"
)

func keyOfSessionInfo(sessionId string) string {
	return fmt.Sprintf("cs:session:info:%s", sessionId)
}
func keyOfSessionClients(sessionId string) string {
	return fmt.Sprintf("cs:session:clients:%s", sessionId)
}

var _ session.Session = &Session{}

type Session struct {
	redisClient *redis.Client
}

func (s *Session) SetEnable(sessionId string, enable bool) (bool, error) {
	key := keyOfSessionInfo(sessionId)

	// 查询会话ID是否存在
	{
		exists, err := s.ExistsInfo(sessionId)
		if err != nil {
			return false, err
		}
		if !exists {
			return false, fmt.Errorf("会话信息不存在")
		}
	}

	// 设置Enable
	{
		_, err := s.redisClient.HSet(context.TODO(), key, "Enable", enable).Result()
		if err != nil {
			return false, err
		}
	}

	return true, nil
}

func (s *Session) GetEnable(sessionId string) (bool, error) {
	key := keyOfSessionInfo(sessionId)

	{
		res, err := s.redisClient.HGet(context.TODO(), key, "Enable").Result()
		if err != nil {
			return false, err
		}
		if res == "0" {
			return false, nil
		}
		if res == "1" {
			return true, nil
		}
		return false, fmt.Errorf("会话信息Enable值错误(%s)", res)
	}
}

func (s *Session) SaveInfo(sessionId string, info session.Info) error {
	key := keyOfSessionInfo(sessionId)

	sessionInfo := &Info{}

	{
		j, err := json.Marshal(info)
		if err != nil {
			return err
		}
		sessionInfo.Info = j
	}

	{
		m := make(map[string]interface{})
		j, err := json.Marshal(sessionInfo)
		if err != nil {
			return err
		}
		if err := json.Unmarshal(j, &m); err != nil {
			return err
		}
		_, err = s.redisClient.HMSet(context.TODO(), key, &m).Result()
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Session) ReadInfo(sessionId string, info session.Info) (bool, error) {
	key := keyOfSessionInfo(sessionId)

	sessionInfo := &Info{}

	// 读取所有信息
	res, err := s.redisClient.HGetAll(context.TODO(), key).Result()
	if err != nil {
		return false, err
	}

	// 格式转换
	if err := mapstructure.WeakDecode(res, sessionInfo); err != nil {
		return false, err
	}

	// decode info
	{
		if err := json.Unmarshal(sessionInfo.Info, info); err != nil {
			return false, err
		}
	}

	return true, nil
}

func (s *Session) ExistsInfo(sessionId string) (bool, error) {
	key := keyOfSessionInfo(sessionId)
	num, err := s.redisClient.Exists(context.TODO(), key).Result()
	if err != nil {
		return false, err
	}
	if num == 0 {
		return false, nil
	}
	if num == 1 {
		return true, nil
	}
	return false, fmt.Errorf("会话信息num(%d)超出判断值", num)
}

func (s *Session) AddClient(sessionId string, client meta.Client) (bool, error) {
	key := keyOfSessionInfo(sessionId)
	val := make([]byte, 0)

	// encode client
	{
		j, err := json.Marshal(client)
		if err != nil {
			return false, err
		}
		val = j
	}

	// add the client to session
	{
		_, err := s.redisClient.HSet(context.TODO(), key, client.GetUUID(), val).Result()
		if err != nil {
			return false, err
		}
	}

	return true, nil
}

func (s *Session) DelClient(sessionId string, client meta.Client) (bool, error) {
	key := keyOfSessionInfo(sessionId)

	{
		_, err := s.redisClient.HDel(context.TODO(), key, client.GetUUID()).Result()
		if err != nil {
			return false, err
		}
	}

	return true, nil
}

func (s *Session) ExistsClient(sessionId string, client meta.Client) (bool, error) {
	key := keyOfSessionInfo(sessionId)

	{
		exists, err := s.redisClient.HExists(context.TODO(), key, client.GetUUID()).Result()
		if err != nil {
			return false, err
		}
		return exists, nil
	}

	return false, nil
}
