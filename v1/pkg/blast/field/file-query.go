package field

import (
	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v1/pkg/blast/consts"
)

const DefaultQueryFile = "-"

func NewQueryFile(val string) QueryFile {
	return QueryFile{true, val}
}

func NewEmptyQueryFile() QueryFile {
	return QueryFile{}
}

func DecodeJSONQueryFile(dec *gojay.Decoder, val *QueryFile) error {
	val.set = true
	return dec.String(&val.val)
}

type QueryFile struct {
	set bool
	val string
}

func (q QueryFile) IsSet() bool {
	return q.set
}

func (q QueryFile) Get() string {
	if !q.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return q.val
}

func (q *QueryFile) Set(val string) {
	q.set = true
	q.val = val
}

func (q *QueryFile) Unset() {
	q.set = false
	q.val = ""
}

func (q QueryFile) Flag() string {
	return consts.FlagQueryFile
}

func (q QueryFile) IsDefault() bool {
	return !q.set || q.val == DefaultQueryFile
}

func (q QueryFile) FlagString() string {
	return q.Flag() + "='" + q.Get() + "'"
}

