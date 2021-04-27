package field

import (
	"github.com/francoispqt/gojay"
	"lib-go-blast/v1/pkg/blast/consts"
)

func NewPhiPatternFile(val string) PhiPatternFile {
	return PhiPatternFile{true, val}
}

func NewEmptyPhiPatternFile() PhiPatternFile {
	return PhiPatternFile{}
}

func DecodeJSONPhiPatternFile(dec *gojay.Decoder, val *PhiPatternFile) error {
	val.set = true
	return dec.DecodeString(&val.val)
}

type PhiPatternFile struct {
	set bool
	val string
}

func (q PhiPatternFile) IsSet() bool {
	return q.set
}

func (q PhiPatternFile) Get() string {
	if !q.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return q.val
}

func (q *PhiPatternFile) Set(val string) {
	q.set = true
	q.val = val
}

func (q *PhiPatternFile) Unset() {
	q.set = false
	q.val = ""
}

func (q PhiPatternFile) Flag() string {
	return consts.FlagPhiPatternFile
}

func (q PhiPatternFile) IsDefault() bool {
	return !q.set || len(q.val) == 0
}

func (q PhiPatternFile) FlagString() string {
	return q.Flag() + "='" + q.Get() + "'"
}

