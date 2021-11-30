package domain

type ReqType string
type Signal int

const (
	TYPE1   ReqType = "type1"
	TYPE2           = "type2"
	TYPE3           = "type3"
	TYPE4           = "type4"
	ONGOING Signal  = 1
	STOP            = 0
	ERROR           = -1
)

type Request struct {
	ReqType ReqType `json:"reqType"`
	ReqData ReqData `json:"reqData"`
	Signal  Signal  `json:"signal"`
}

type ReqData struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
