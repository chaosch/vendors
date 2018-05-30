package connectionmanager

import (
	"net/http"
	"github.com/gorilla/websocket"
	"time"
	"egov/serializer"
	"io"
)

const writeWait = 10 * time.Second

type websocketDialect struct {
	ConnectionDialect
	MessageType map[AcceptType]interface{}
}


//func (jd *jsonDialect) WriteWebSocket(c *WsClient, result *ResultTemplate) {
//	message, _ := ffjson.Marshal(result)
//	c.conn.SetWriteDeadline(time.Now().Add(writeWait))
//	w, err := c.conn.NextWriter(websocket.TextMessage)
//	if err != nil {
//		c.conn.WriteMessage(websocket.CloseMessage, []byte{})
//		return
//	}
//	w.Write(message)
//
//	if err := w.Close(); err != nil {
//		return
//	}
//}


func (dialect *websocketDialect) WriteConnection(connection interface{}, buf []byte, accept AcceptType) {
	conn := connection.(*websocket.Conn)
	conn.SetWriteDeadline(time.Now().Add(writeWait))
	var w io.WriteCloser
	var err error
	if accept == serializer.ContentMsgpack {
		w, err = conn.NextWriter(websocket.BinaryMessage)
	} else {
		w, err = conn.NextWriter(websocket.TextMessage)
	}
	if err != nil {
		conn.WriteMessage(websocket.CloseMessage, []byte{})
		return
	}

	w.Write(buf)


	if err := w.Close(); err != nil {
		return
	}

}

func (dialect *websocketDialect) WriteHeader(writer interface{}, key string, keyValue string) {
	w := writer.(http.ResponseWriter)
	w.Header().Set(key, keyValue)
}

//func (dialect *websocketDialect)InitMessageType(){
//	dialect.MessageType=make(map[AcceptType]interface{})
//	dialect.MessageType[AcceptJson]=websocket.TextMessage
//	dialect.MessageType[AcceptMsgpack]=websocket.BinaryMessage
//}
