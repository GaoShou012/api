package client

type Client interface {
	AddSession(clientInfo Info, sessionId string) (bool, error)
	DelSession(clientInfo Info, sessionId string) (bool, error)
	ExistsSession(clientInfo Info, sessionId string) (bool, error)

}
