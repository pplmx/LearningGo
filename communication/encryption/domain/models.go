package domain

type RespType string

const (
    TYPE1 RespType = "type1"
    TYPE2          = "type2"
    TYPE3          = "type3"
    TYPE4          = "type4"
)

type Response struct {
    RespType RespType `json:"respType"`
    RespData RespData `json:"respData"`
}

type RespData struct {
    Id   string `json:"id,omitempty"`
    Name string `json:"name,omitempty"`
}
