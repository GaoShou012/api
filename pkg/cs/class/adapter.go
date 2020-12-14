package class

import "cs/meta"

type Adapter interface {
	ExistsSession(session meta.Session) (bool, error)
	SaveSession(session meta.Session) error
	ReadSession(sessionId string, session meta.Session) (bool, error)

	ExistsClient(clientId string) (bool, error)
	SaveClient(client meta.Client) error
	ReadClient(clientId string, client meta.Client) (bool, error)
}
