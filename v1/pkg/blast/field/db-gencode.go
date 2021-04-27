package field

import (
	"strconv"

	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v1/pkg/blast/consts"
)

const DefaultDBGenCode uint8 = 1

func DecodeJSONDBGenCode(dec *gojay.Decoder, val *DBGenCode) error {
	val.set = true
	return dec.DecodeUint8(&val.val)
}

type DBGenCode struct {
	set bool
	val uint8
}

func (q DBGenCode) IsSet() bool {
	return q.set
}

func (q DBGenCode) Get() uint8 {
	if !q.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return q.val
}

func (q *DBGenCode) Set(val uint8) {
	q.set = true
	q.val = val
}

func (q *DBGenCode) Unset() {
	q.set = false
}

func (q DBGenCode) Flag() string {
	return consts.FlagDBGenCode
}

func (q DBGenCode) IsDefault() bool {
	return !q.IsSet() || q.val == DefaultDBGenCode
}

func (q DBGenCode) FlagString() string {
	return q.Flag() + "=" + strconv.Itoa(int(q.Get()))
}

