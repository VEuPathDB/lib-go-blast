package field

import (
	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
)

func NewDBSoftMask(val string) DBSoftMask {
	return DBSoftMask{true, val}
}

func NewEmptyDBSoftMask() DBSoftMask {
	return DBSoftMask{}
}

func DecodeJSONDBSoftMask(dec *gojay.Decoder, val *DBSoftMask) error {
	val.set = true
	return dec.DecodeString(&val.val)
}

type DBSoftMask struct {
	set bool
	val string
}

func (q *DBSoftMask) IsSet() bool {
	return q.set
}

func (q *DBSoftMask) Get() string {
	if !q.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return q.val
}

func (q *DBSoftMask) Set(val string) {
	q.set = true
	q.val = val
}

func (q *DBSoftMask) Unset() {
	q.set = false
	q.val = ""
}

func (q *DBSoftMask) Flag() string {
	return consts.FlagDBSoftMask
}

func (q *DBSoftMask) IsDefault() bool {
	return !q.set
}

func (q *DBSoftMask) FlagString() string {
	return q.Flag() + "='" + q.Get() + "'"
}

