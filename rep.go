package cast

import (
	"encoding/json"
	"net/http"
	"time"
)

type Reply struct {
	statusCode int
	body       []byte
	cost       time.Duration
}

func (rep *Reply) DecodeFromJson(v interface{}) error {
	return json.Unmarshal(rep.body, &v)
}

func (rep *Reply) Body() []byte {
	return rep.body
}

func (rep *Reply) StatusOk() bool {
	return rep.statusCode == http.StatusOK
}

func (rep *Reply) StatusCode() int {
	return rep.statusCode
}

func (rep *Reply) Cost() time.Duration {
	return rep.cost
}
