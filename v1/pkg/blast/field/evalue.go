package field

import (
	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v1/pkg/blast/consts"
)

const DefaultExpectValue = "10"

func NewExpectValue(val string) ExpectValue {
	return ExpectValue{true, val}
}

func NewEmptyExpectValue() ExpectValue {
	return ExpectValue{}
}

func DecodeJSONExpectValue(dec *gojay.Decoder, val *ExpectValue) error {
	val.set = true
	return dec.String(&val.val)
}

type ExpectValue struct {
	set bool
	val string
}

func (q ExpectValue) IsSet() bool {
	return q.set
}

func (q ExpectValue) Get() string {
	if !q.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return q.val
}

func (q *ExpectValue) Set(val string) {
	q.set = true
	q.val = val
}

func (q *ExpectValue) Unset() {
	q.set = false
	q.val = ""
}

func (q ExpectValue) Flag() string {
	return consts.FlagExpectValue
}

func (q ExpectValue) IsDefault() bool {
	return !q.set || q.val == DefaultExpectValue
}

func (q ExpectValue) FlagString() string {
	return q.Flag() + "=" + q.Get()
}

