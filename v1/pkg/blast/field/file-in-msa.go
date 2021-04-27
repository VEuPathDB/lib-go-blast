package field

import (
	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v1/pkg/blast/consts"
)

func NewInMSAFile(val string) InMSAFile {
	return InMSAFile{true, val}
}

func NewEmptyInMSAFile() InMSAFile {
	return InMSAFile{}
}

func DecodeJSONInMSAFile(dec *gojay.Decoder, val *InMSAFile) error {
	val.set = true
	return dec.DecodeString(&val.val)
}

type InMSAFile struct {
	set bool
	val string
}

func (q InMSAFile) IsSet() bool {
	return q.set
}

func (q InMSAFile) Get() string {
	if !q.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return q.val
}

func (q *InMSAFile) Set(val string) {
	q.set = true
	q.val = val
}

func (q *InMSAFile) Unset() {
	q.set = false
	q.val = ""
}

func (q InMSAFile) Flag() string {
	return consts.FlagInMSAFile
}

func (q InMSAFile) IsDefault() bool {
	return !q.set
}

func (q InMSAFile) FlagString() string {
	return q.Flag() + "='" + q.Get() + "'"
}

