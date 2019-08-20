package connectionmanager

import (
	"fmt"
	"net/http"
)

type httpDialect struct {
	ConnectionDialect
}

func (dialect *httpDialect) WriteConnection(connection interface{}, buf []byte, accept AcceptType) {
	w := connection.(http.ResponseWriter)
	var f func(http.ResponseWriter, []byte)
	if _, ok := WriteFunction[accept]; !ok {
		f = WriteFunction[AcceptJson]
		dialect.WriteHeader(connection, "Content-Type", AcceptJson)
		//dialect.WriteHeader(connection, "Content-Type", AcceptJson)
	} else {
		f = WriteFunction[accept]
		dialect.WriteHeader(connection, "Content-Type", string(accept))
		//dialect.WriteHeader(connection, "Content-Type", string(accept))
	}
	f(w, buf)
}

func (dialect *httpDialect) WriteHeader(writer interface{}, key string, keyValue string) {
	w := writer.(http.ResponseWriter)
	w.Header().Set(key, keyValue)
}

var WriteFunction = map[AcceptType]func(http.ResponseWriter, []byte){
	AcceptJson: func(w http.ResponseWriter, buf []byte) {
		fmt.Fprint(w, string(buf))
	},
	AcceptMsgpack: func(w http.ResponseWriter, buf []byte) {
		w.Write(buf)
	},
}
