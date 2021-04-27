package field

import (
	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v1/pkg/blast/consts"
)

func NewRequestID(val string) RequestID {
	return RequestID{true, val}
}

func NewEmptyRequestID() RequestID {
	return RequestID{}
}

func DecodeJSONRequestID(dec *gojay.Decoder, val *RequestID) error {
	val.set = true
	return dec.DecodeString(&val.val)
}

type RequestID struct {
	set bool
	val string
}

func (q RequestID) IsSet() bool {
	return q.set
}

func (q RequestID) Get() string {
	if !q.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return q.val
}

func (q *RequestID) Set(val string) {
	q.set = true
	q.val = val
}

func (q *RequestID) Unset() {
	q.set = false
	q.val = ""
}

func (q RequestID) Flag() string {
	return consts.FlagRequestID
}

func (q RequestID) IsDefault() bool {
	return !q.set
}

func (q RequestID) FlagString() string {
	return q.Flag() + "=" + q.Get()
}

