package field

import (
	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
)

func NewMatrix(val string) Matrix {
	return Matrix{true, val}
}

func NewEmptyMatrix() Matrix {
	return Matrix{}
}

func DecodeJSONMatrix(dec *gojay.Decoder, val *Matrix) error {
	val.set = true
	return dec.String(&val.val)
}

type Matrix struct {
	set bool
	val string
}

func (q *Matrix) IsSet() bool {
	return q.set
}

func (q *Matrix) Get() string {
	if !q.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return q.val
}

func (q *Matrix) Set(val string) {
	q.set = true
	q.val = val
}

func (q *Matrix) Unset() {
	q.set = false
	q.val = ""
}

func (q *Matrix) Flag() string {
	return consts.FlagMatrix
}

func (q *Matrix) IsDefault() bool {
	return !q.set || len(q.val) == 0
}

func (q *Matrix) FlagString() string {
	return q.Flag() + "='" + q.Get() + "'"
}

type MatrixType string

const (
	MatrixTypeBlosum45 MatrixType = "BLOSUM45"
	MatrixTypeBlosum50 MatrixType = "BLOSUM50"
	MatrixTypeBlosum62 MatrixType = "BLOSUM62"
	MatrixTypeBlosum80 MatrixType = "BLOSUM80"
	MatrixTypeBlosum90 MatrixType = "BLOSUM90"
	MatrixTypePam30    MatrixType = "PAM30"
	MatrixTypePam70    MatrixType = "PAM70"
	MatrixTypePam250   MatrixType = "PAM250"
	MatrixTypeIdentity MatrixType = "IDENTITY"
)
