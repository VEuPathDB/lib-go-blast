package field

import (
	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v1/pkg/blast/consts"
)

func NewArchiveFile(val string) ArchiveFile {
	return ArchiveFile{true, val}
}

func NewEmptyArchiveFile() ArchiveFile {
	return ArchiveFile{}
}

func DecodeJSONArchiveFile(dec *gojay.Decoder, val *ArchiveFile) error {
	val.set = true
	return dec.String(&val.val)
}

type ArchiveFile struct {
	set bool
	val string
}

func (q ArchiveFile) IsSet() bool {
	return q.set
}

func (q ArchiveFile) Get() string {
	if !q.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return q.val
}

func (q *ArchiveFile) Set(val string) {
	q.set = true
	q.val = val
}

func (q *ArchiveFile) Unset() {
	q.set = false
	q.val = ""
}

func (q ArchiveFile) Flag() string {
	return consts.FlagArchiveFile
}

func (q ArchiveFile) IsDefault() bool {
	return !q.set || len(q.val) == 0
}

func (q ArchiveFile) FlagString() string {
	return q.Flag() + "='" + q.Get() + "'"
}

