package field

import (
	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v1/pkg/blast/consts"
)

func NewRPSDBFile(val string) RPSDBFile {
	return RPSDBFile{true, val}
}

func NewEmptyRPSDBFile() RPSDBFile {
	return RPSDBFile{}
}

func DecodeJSONRPSDBFile(dec *gojay.Decoder, val *RPSDBFile) error {
	val.set = true
	return dec.DecodeString(&val.val)
}

type RPSDBFile struct {
	set bool
	val string
}

func (q RPSDBFile) IsSet() bool {
	return q.set
}

func (q RPSDBFile) Get() string {
	if !q.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return q.val
}

func (q *RPSDBFile) Set(val string) {
	q.set = true
	q.val = val
}

func (q *RPSDBFile) Unset() {
	q.set = false
	q.val = ""
}

func (q RPSDBFile) Flag() string {
	return consts.FlagRPSDBFile
}

func (q RPSDBFile) IsDefault() bool {
	return !q.set || len(q.val) == 0
}

func (q RPSDBFile) FlagString() string {
	return q.Flag() + "='" + q.Get() + "'"
}

