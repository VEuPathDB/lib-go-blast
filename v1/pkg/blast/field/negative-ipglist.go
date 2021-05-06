package field

import (
	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v1/pkg/blast/consts"
)

func DecodeJSONNegativeIPGList(dec *gojay.Decoder, val *NegativeIPGList) error {
	val.set = true
	return dec.String(&val.val)
}

type NegativeIPGList struct {
	set bool
	val string
}

func (q NegativeIPGList) IsSet() bool {
	return q.set
}

func (q NegativeIPGList) Get() string {
	if !q.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return q.val
}

func (q *NegativeIPGList) Set(val string) {
	q.set = true
	q.val = val
}

func (q *NegativeIPGList) Unset() {
	q.set = false
	q.val = ""
}

func (q NegativeIPGList) Flag() string {
	return consts.FlagNegativeIPGList
}

func (q NegativeIPGList) IsDefault() bool {
	return !q.set || len(q.val) == 0
}

func (q NegativeIPGList) FlagString() string {
	return q.Flag() + "='" + q.Get() + "'"
}

