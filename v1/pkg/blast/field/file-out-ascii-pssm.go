package field

import (
	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v1/pkg/blast/consts"
)

func NewOutASCIIPSSMFile(val string) OutASCIIPSSMFile {
	return OutASCIIPSSMFile{true, val}
}

func NewEmptyOutASCIIPSSMFile() OutASCIIPSSMFile {
	return OutASCIIPSSMFile{}
}

func DecodeJSONOutASCIIPSSMFile(dec *gojay.Decoder, val *OutASCIIPSSMFile) error {
	val.set = true
	return dec.DecodeString(&val.val)
}

type OutASCIIPSSMFile struct {
	set bool
	val string
}

func (q OutASCIIPSSMFile) IsSet() bool {
	return q.set
}

func (q OutASCIIPSSMFile) Get() string {
	if !q.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return q.val
}

func (q *OutASCIIPSSMFile) Set(val string) {
	q.set = true
	q.val = val
}

func (q *OutASCIIPSSMFile) Unset() {
	q.set = false
	q.val = ""
}

func (q OutASCIIPSSMFile) Flag() string {
	return consts.FlagOutASCIIPSSMFile
}

func (q OutASCIIPSSMFile) IsDefault() bool {
	return !q.set || len(q.val) == 0
}

func (q OutASCIIPSSMFile) FlagString() string {
	return q.Flag() + "='" + q.Get() + "'"
}

