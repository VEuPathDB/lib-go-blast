package field

import (
	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v1/pkg/blast/consts"
)

func DecodeJSONNegativeGIList(dec *gojay.Decoder, val *NegativeGIList) error {
	val.set = true
	return dec.DecodeString(&val.val)
}

type NegativeGIList struct {
	set bool
	val string
}

func (q NegativeGIList) IsSet() bool {
	return q.set
}

func (q NegativeGIList) Get() string {
	if !q.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return q.val
}

func (q *NegativeGIList) Set(val string) {
	q.set = true
	q.val = val
}

func (q *NegativeGIList) Unset() {
	q.set = false
	q.val = ""
}

func (q NegativeGIList) Flag() string {
	return consts.FlagNegativeGIList
}

func (q NegativeGIList) IsDefault() bool {
	return !q.set || len(q.val) == 0
}

func (q NegativeGIList) FlagString() string {
	return q.Flag() + "='" + q.Get() + "'"
}

