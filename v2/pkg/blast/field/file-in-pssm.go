package field

import (
	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
)

func NewInPSSMFile(val string) InPSSMFile {
	return InPSSMFile{true, val}
}

func NewEmptyInPSSMFile() InPSSMFile {
	return InPSSMFile{}
}

func DecodeJSONInPSSMFile(dec *gojay.Decoder, val *InPSSMFile) error {
	val.set = true
	return dec.DecodeString(&val.val)
}

type InPSSMFile struct {
	set bool
	val string
}

func (q *InPSSMFile) IsSet() bool {
	return q.set
}

func (q *InPSSMFile) Get() string {
	if !q.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return q.val
}

func (q *InPSSMFile) Set(val string) {
	q.set = true
	q.val = val
}

func (q *InPSSMFile) Unset() {
	q.set = false
	q.val = ""
}

func (q *InPSSMFile) Flag() string {
	return consts.FlagInPSSMFile
}

func (q *InPSSMFile) IsDefault() bool {
	return !q.set || len(q.val) == 0
}

func (q *InPSSMFile) FlagString() string {
	return q.Flag() + "='" + q.Get() + "'"
}

