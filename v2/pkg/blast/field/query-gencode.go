package field

import (
	"fmt"
	"strconv"

	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/bval"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
)

const DefaultQueryGenCode uint8 = 1

func NewQueryGenCode(val uint8) QueryGenCode {
	return QueryGenCode{true, val}
}

func DecodeJSONQueryGenCode(dec *gojay.Decoder, val *QueryGenCode) error {
	val.set = true
	return dec.DecodeUint8(&val.val)
}

type QueryGenCode struct {
	set bool
	val uint8
}

func (q *QueryGenCode) IsSet() bool {
	return q.set
}

func (q *QueryGenCode) Get() uint8 {
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

func (q *QueryGenCode) Flag() string {
	return consts.FlagQueryGenCode
}

func (q *QueryGenCode) IsDefault() bool {
	return !q.IsSet() || q.val == DefaultQueryGenCode
}

func (q *QueryGenCode) FlagString() string {
	return q.Flag() + "=" + strconv.Itoa(int(q.Get()))
}

const errInvalidQueryGenCode = `Invalid %s value "%d".  Valid values are [1-6], [9-16], [21-31], 33.`

func (q *QueryGenCode) Validate(b bval.ValidationBuilder) {
	if q.IsDefault() {
		return
	}

	switch true {
	case q.Get() > 0 && q.Get() < 7:
	case q.Get() > 8 && q.Get() < 17:
	case q.Get() > 20 && q.Get() < 32:
	case q.Get() == 33:
	default:
		b.AppendError(q.Flag(), fmt.Sprintf(errInvalidQueryGenCode, q.Flag(), q.Get()))
	}
}

