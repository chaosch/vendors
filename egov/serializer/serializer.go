package serializer

type ContentType string

func init() {
	regTransDialects()
}

func regTransDialects() bool {
	providedTransDialects := map[ContentType]struct {
		ContentType ContentType
		getDialect  func() TransDialect
	}{
		ContentJson:    {ContentJson, func() TransDialect { return &jsonDialect{ContentType: ContentJson} }},
		ContentMsgpack: {ContentMsgpack, func() TransDialect { return &msgpackDialect{ContentType: ContentMsgpack} }},
	}

	for ct, v := range providedTransDialects {
		//if td := QueryTransDialect(ct); td == nil {
		RegisterTransDialect(ct, v.getDialect)
		//}
	}
	return true
}

const (
	ContentJson    = "application/json"
	ContentMsgpack = "application/x-msgpack"
)

var (
	transdialects = map[ContentType]func() TransDialect{}
)

// RegisterDialect register transdialects
func RegisterTransDialect(ContentType ContentType, dialectFunc func() TransDialect) {
	if dialectFunc == nil {
		panic("core: Register dialect is nil")
	}
	transdialects[ContentType] = dialectFunc // !nashtsai! allow override dialect
}

// QueryDialect query if registed database dialect
func QueryTransDialect(ContentType ContentType) TransDialect {
	if d, ok := transdialects[ContentType]; ok {
		return d()
	} else {
		return transdialects[ContentJson]()
	}

}

//实现如下方法
type TransDialect interface {
	GetValue([]byte, string) (string, error)
	Unmarshal([]byte, interface{}) error
	Marshal(interface{}) ([]byte, error)
}
