package field

import (
	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v1/pkg/blast/consts"
)

func NewIndexName(val string) IndexName {
	return IndexName{true, val}
}

func NewEmptyIndexName() IndexName {
	return IndexName{}
}

func DecodeJSONIndexName(dec *gojay.Decoder, val *IndexName) error {
	val.set = true
	return dec.String(&val.val)
}

type IndexName struct {
	set bool
	val string
}

func (q IndexName) IsSet() bool {
	return q.set
}

func (q IndexName) Get() string {
	if !q.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return q.val
}

func (q *IndexName) Set(val string) {
	q.set = true
	q.val = val
}

func (q *IndexName) Unset() {
	q.set = false
	q.val = ""
}

func (q IndexName) Flag() string {
	return consts.FlagIndexName
}

func (q IndexName) IsDefault() bool {
	return !q.set || len(q.val) == 0
}

func (q IndexName) FlagString() string {
	return q.Flag() + "='" + q.Get() + "'"
}

