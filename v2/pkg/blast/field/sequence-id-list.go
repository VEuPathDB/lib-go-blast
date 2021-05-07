package field

import (
	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
)

func DecodeJSONSequenceIDList(dec *gojay.Decoder, val *SequenceIDList) error {
	val.set = true
	return dec.String(&val.val)
}

type SequenceIDList struct {
	set bool
	val string
}

func (q *SequenceIDList) IsSet() bool {
	return q.set
}

func (q *SequenceIDList) Get() string {
	if !q.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return q.val
}

func (q *SequenceIDList) Set(val string) {
	q.set = true
	q.val = val
}

func (q *SequenceIDList) Unset() {
	q.set = false
	q.val = ""
}

func (q *SequenceIDList) Flag() string {
	return consts.FlagSequenceIDList
}

func (q *SequenceIDList) IsDefault() bool {
	return !q.set || len(q.val) == 0
}

func (q *SequenceIDList) FlagString() string {
	return q.Flag() + "=" + q.Get()
}

