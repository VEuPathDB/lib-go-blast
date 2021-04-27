package field

import (
	"strconv"

	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v1/pkg/blast/consts"
)

const DefaultQueryGenCode uint8 = 1

func DecodeJSONQueryGenCode(dec *gojay.Decoder, val *QueryGenCode) error {
	val.set = true
	return dec.DecodeUint8(&val.val)
}

type QueryGenCode struct {
	set bool
	val uint8
}

func (q QueryGenCode) IsSet() bool {
	return q.set
}

func (q QueryGenCode) Get() uint8 {
	if !q.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return q.val
}

func (q *QueryGenCode) Set(val uint8) {
	q.set = true
	q.val = val
}

func (q *QueryGenCode) Unset() {
	q.set = false
}

func (q QueryGenCode) Flag() string {
	return consts.FlagQueryGenCode
}

func (q QueryGenCode) IsDefault() bool {
	return !q.IsSet() || q.val == DefaultQueryGenCode
}

func (q QueryGenCode) FlagString() string {
	return q.Flag() + "=" + strconv.Itoa(int(q.Get()))
}

