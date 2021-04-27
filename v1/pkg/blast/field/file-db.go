package field

import (
	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v1/pkg/blast/consts"
)

func NewDBFile(val string) DBFile {
	return DBFile{true, val}
}

func NewEmptyDBFile() DBFile {
	return DBFile{}
}

func DecodeJSONDBFile(dec *gojay.Decoder, val *DBFile) error {
	val.set = true
	return dec.DecodeString(&val.val)
}

type DBFile struct {
	set bool
	val string
}

func (q DBFile) IsSet() bool {
	return q.set
}

func (q DBFile) Get() string {
	if !q.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return q.val
}

func (q *DBFile) Set(val string) {
	q.set = true
	q.val = val
}

func (q *DBFile) Unset() {
	q.set = false
	q.val = ""
}

func (q DBFile) Flag() string {
	return consts.FlagDBFile
}

func (q DBFile) IsDefault() bool {
	return !q.set || len(q.val) == 0
}

func (q DBFile) FlagString() string {
	return q.Flag() + "='" + q.Get() + "'"
}

