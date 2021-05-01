package field

import (
	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
)

func DecodeJSONNegativeSequenceIDList(dec *gojay.Decoder, val *NegativeSequenceIDList) error {
	val.set = true
	return dec.DecodeString(&val.val)
}

type NegativeSequenceIDList struct {
	set bool
	val string
}

func (q *NegativeSequenceIDList) IsSet() bool {
	return q.set
}

func (q *NegativeSequenceIDList) Get() string {
	if !q.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return q.val
}

func (q *NegativeSequenceIDList) Set(val string) {
	q.set = true
	q.val = val
}

func (q *NegativeSequenceIDList) Unset() {
	q.set = false
	q.val = ""
}

func (q *NegativeSequenceIDList) Flag() string {
	return consts.FlagNegativeSequenceIDList
}

func (q *NegativeSequenceIDList) IsDefault() bool {
	return !q.set || len(q.val) == 0
}

func (q *NegativeSequenceIDList) FlagString() string {
	return q.Flag() + "='" + q.Get() + "'"
}

