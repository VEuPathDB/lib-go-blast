package field

import (
	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v1/pkg/blast/consts"
)

const DefaultOutFile = "-"

func NewOutFile(val string) OutFile {
	return OutFile{true, val}
}

func NewEmptyOutFile() OutFile {
	return OutFile{}
}

func DecodeJSONOutFile(dec *gojay.Decoder, val *OutFile) error {
	val.set = true
	return dec.DecodeString(&val.val)
}

type OutFile struct {
	set bool
	val string
}

func (q OutFile) IsSet() bool {
	return q.set
}

func (q OutFile) Get() string {
	if !q.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return q.val
}

func (q *OutFile) Set(val string) {
	q.set = true
	q.val = val
}

func (q *OutFile) Unset() {
	q.set = false
	q.val = ""
}

func (q OutFile) Flag() string {
	return consts.FlagOutFile
}

func (q OutFile) IsDefault() bool {
	return !q.set || q.val == DefaultOutFile
}

func (q OutFile) FlagString() string {
	return q.Flag() + "='" + q.Get() + "'"
}

