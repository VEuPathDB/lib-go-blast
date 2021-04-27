package field

import (
	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v1/pkg/blast/consts"
)

func NewSubjectFile(val string) SubjectFile {
	return SubjectFile{true, val}
}

func NewEmptySubjectFile() SubjectFile {
	return SubjectFile{}
}

func DecodeJSONSubjectFile(dec *gojay.Decoder, val *SubjectFile) error {
	val.set = true
	return dec.DecodeString(&val.val)
}

type SubjectFile struct {
	set bool
	val string
}

func (q SubjectFile) IsSet() bool {
	return q.set
}

func (q SubjectFile) Get() string {
	if !q.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return q.val
}

func (q *SubjectFile) Set(val string) {
	q.set = true
	q.val = val
}

func (q *SubjectFile) Unset() {
	q.set = false
	q.val = ""
}

func (q SubjectFile) Flag() string {
	return consts.FlagSubjectFile
}

func (q SubjectFile) IsDefault() bool {
	return !q.set || len(q.val) == 0
}

func (q SubjectFile) FlagString() string {
	return q.Flag() + "='" + q.Get() + "'"
}

