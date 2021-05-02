package field

import (
	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
)

func NewImportSearchStrategy(val string) ImportSearchStrategy {
	return ImportSearchStrategy{true, val}
}

func NewEmptyImportSearchStrategy() ImportSearchStrategy {
	return ImportSearchStrategy{}
}

func DecodeJSONImportSearchStrategy(dec *gojay.Decoder, val *ImportSearchStrategy) error {
	val.set = true
	return dec.String(&val.val)
}

type ImportSearchStrategy struct {
	set bool
	val string
}

func (q *ImportSearchStrategy) IsSet() bool {
	return q.set
}

func (q *ImportSearchStrategy) Get() string {
	if !q.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return q.val
}

func (q *ImportSearchStrategy) Set(val string) {
	q.set = true
	q.val = val
}

func (q *ImportSearchStrategy) Unset() {
	q.set = false
	q.val = ""
}

func (q *ImportSearchStrategy) Flag() string {
	return consts.FlagImportSearchStrategy
}

func (q *ImportSearchStrategy) IsDefault() bool {
	return !q.set
}

func (q *ImportSearchStrategy) FlagString() string {
	return q.Flag() + "='" + q.Get() + "'"
}

