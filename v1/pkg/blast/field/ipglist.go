package field

import (
	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v1/pkg/blast/consts"
)

func DecodeJSONIPGList(dec *gojay.Decoder, val *IPGList) error {
	val.set = true
	return dec.DecodeString(&val.val)
}

type IPGList struct {
	set bool
	val string
}

func (q IPGList) IsSet() bool {
	return q.set
}

func (q IPGList) Get() string {
	if !q.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return q.val
}

func (q *IPGList) Set(val string) {
	q.set = true
	q.val = val
}

func (q *IPGList) Unset() {
	q.set = false
	q.val = ""
}

func (q IPGList) Flag() string {
	return consts.FlagIPGList
}

func (q IPGList) IsDefault() bool {
	return !q.set || len(q.val) == 0
}

func (q IPGList) FlagString() string {
	return q.Flag() + "='" + q.Get() + "'"
}

