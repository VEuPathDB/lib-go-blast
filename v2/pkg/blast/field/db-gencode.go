package field

import (
	"fmt"
	"strconv"

	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/bval"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
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

func (q *DBGenCode) IsSet() bool {
	return q.set
}

func (q *DBGenCode) Get() uint8 {
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

func (q *DBGenCode) Flag() string {
	return consts.FlagDBGenCode
}

func (q *DBGenCode) IsDefault() bool {
	return !q.IsSet() || q.val == DefaultDBGenCode
}

func (q *DBGenCode) FlagString() string {
	return q.Flag() + "=" + strconv.Itoa(int(q.Get()))
}

const errInvalidDBGenCode = `Invalid %s value "%d".  Valid values are [1-6], [9-16], [21-31], 33.`

func (q *DBGenCode) Validate(b bval.ValidationBuilder) {
	if q.IsDefault() {
		return
	}

	switch true {
	case q.Get() > 0 && q.Get() < 7:
	case q.Get() > 8 && q.Get() < 17:
	case q.Get() > 20 && q.Get() < 32:
	case q.Get() == 33:
	default:
		b.AppendError(q.Flag(), fmt.Sprintf(errInvalidDBGenCode, q.Flag(), q.Get()))
	}
}
