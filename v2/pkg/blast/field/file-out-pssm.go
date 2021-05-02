package field

import (
	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
)

func NewOutPSSMFile(val string) OutPSSMFile {
	return OutPSSMFile{true, val}
}

func NewEmptyOutPSSMFile() OutPSSMFile {
	return OutPSSMFile{}
}

func DecodeJSONOutPSSMFile(dec *gojay.Decoder, val *OutPSSMFile) error {
	val.set = true
	return dec.String(&val.val)
}

type OutPSSMFile struct {
	set bool
	val string
}

func (q *OutPSSMFile) IsSet() bool {
	return q.set
}

func (q *OutPSSMFile) Get() string {
	if !q.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return q.val
}

func (q *OutPSSMFile) Set(val string) {
	q.set = true
	q.val = val
}

func (q *OutPSSMFile) Unset() {
	q.set = false
	q.val = ""
}

func (q *OutPSSMFile) Flag() string {
	return consts.FlagOutPSSMFile
}

func (q *OutPSSMFile) IsDefault() bool {
	return !q.set || len(q.val) == 0
}

func (q *OutPSSMFile) FlagString() string {
	return q.Flag() + "='" + q.Get() + "'"
}

