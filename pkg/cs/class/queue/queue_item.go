package queue

import (
	"cs/meta"
	"encoding/json"
	"time"
)

type UnmarshalSessionInfo func(info []byte) (meta.Session, error)

type Item struct {
	SessionId string
	JoinTime  time.Time
}

func (i *Item) SessionInfoMarshal(session meta.Session) error {
	j, err := json.Marshal(session)
	if err != nil {
		return err
	}
	i.SessionInfo = j
	return nil
}

func (i *Item) SessionInfoUnmarshal(session meta.Session) error {
	return json.Unmarshal(i.SessionInfo, session)
}
