package field

import (
	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v1/pkg/blast/consts"
)

func DecodeJSONGIList(dec *gojay.Decoder, val *GIList) error {
	val.set = true
	return dec.String(&val.val)
}

type GIList struct {
	set bool
	val string
}

func (q GIList) IsSet() bool {
	return q.set
}

func (q GIList) Get() string {
	if !q.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return q.val
}

func (q *GIList) Set(val string) {
	q.set = true
	q.val = val
}

func (q *GIList) Unset() {
	q.set = false
	q.val = ""
}

func (q GIList) Flag() string {
	return consts.FlagGIList
}

func (q GIList) IsDefault() bool {
	return !q.set || len(q.val) == 0
}

func (q GIList) FlagString() string {
	return q.Flag() + "='" + q.Get() + "'"
}

