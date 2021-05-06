package field

import (
	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v1/pkg/blast/consts"
)

func NewEntrezQuery(val string) EntrezQuery {
	return EntrezQuery{true, val}
}

func NewEmptyEntrezQuery() EntrezQuery {
	return EntrezQuery{}
}

func DecodeJSONEntrezQuery(dec *gojay.Decoder, val *EntrezQuery) error {
	val.set = true
	return dec.String(&val.val)
}

type EntrezQuery struct {
	set bool
	val string
}

func (q EntrezQuery) IsSet() bool {
	return q.set
}

func (q EntrezQuery) Get() string {
	if !q.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return q.val
}

func (q *EntrezQuery) Set(val string) {
	q.set = true
	q.val = val
}

func (q *EntrezQuery) Unset() {
	q.set = false
	q.val = ""
}

func (q EntrezQuery) Flag() string {
	return consts.FlagEntrezQuery
}

func (q EntrezQuery) IsDefault() bool {
	return !q.set
}

func (q EntrezQuery) FlagString() string {
	return q.Flag() + "='" + q.Get() + "'"
}

