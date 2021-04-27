package field

import (
	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v1/pkg/blast/consts"
)

func NewDBHardMask(val string) DBHardMask {
	return DBHardMask{true, val}
}

func NewEmptyDBHardMask() DBHardMask {
	return DBHardMask{}
}

func DecodeJSONDBHardMask(dec *gojay.Decoder, val *DBHardMask) error {
	val.set = true
	return dec.DecodeString(&val.val)
}

type DBHardMask struct {
	set bool
	val string
}

func (q DBHardMask) IsSet() bool {
	return q.set
}

func (q DBHardMask) Get() string {
	if !q.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return q.val
}

func (q *DBHardMask) Set(val string) {
	q.set = true
	q.val = val
}

func (q *DBHardMask) Unset() {
	q.set = false
	q.val = ""
}

func (q DBHardMask) Flag() string {
	return consts.FlagDBHardMask
}

func (q DBHardMask) IsDefault() bool {
	return !q.set
}

func (q DBHardMask) FlagString() string {
	return q.Flag() + "='" + q.Get() + "'"
}

