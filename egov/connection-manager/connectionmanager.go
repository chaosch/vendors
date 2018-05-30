package connectionmanager

const (
	ConnectionHttp      = "http"
	ConnectionWebsocket = "websocket"
	//添加连接种类
)

const (
	AcceptJson    = "application/json"
	AcceptMsgpack = "application/x-msgpack"
	//添加Accept种类
)

type ConnectionType string
type ContentType string
type AcceptType string

var (
	connectionDialects = map[ConnectionType]func() ConnectionDialect{}
)

func init() {
	regConnectionDialects()
}

func regConnectionDialects() bool {
	providedConnectionDialects := map[ConnectionType]struct {
		ConnectionType ConnectionType
		getDialect     func() ConnectionDialect
	}{
		ConnectionHttp:      {ConnectionHttp, func() ConnectionDialect { return &httpDialect{} }},
		ConnectionWebsocket: {ConnectionWebsocket, func() ConnectionDialect { return &websocketDialect{} }},
		//此处可添加其他种类的连接 需实现一个对应的 ConnectionDialect ，目前实现两种，见 http.go和websocket.go
	}

	for ct, v := range providedConnectionDialects {
		//if td := QueryConnectionDialect(ct); td == nil {
			RegisterConnectionDialect(ct, v.getDialect)
		//}
	}
	return true
}

func RegisterConnectionDialect(connectionType ConnectionType, dialectFunc func() ConnectionDialect) {
	if dialectFunc == nil {
		panic("core: Register dialect is nil")
	}
	connectionDialects[connectionType] = dialectFunc // !nashtsai! allow override dialect
}

func QueryConnectionDialect(connectionType ConnectionType) ConnectionDialect {
	if d, ok := connectionDialects[connectionType]; ok {
		return d()
	}
	return nil
}

type ConnectionDialect interface {
	WriteConnection(connection interface{},buf []byte,acceptType AcceptType)
	WriteHeader(writor interface{},key string,value string)
}
