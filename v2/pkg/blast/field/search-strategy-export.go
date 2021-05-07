package field

import (
	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
)

func NewExportSearchStrategy(val string) ExportSearchStrategy {
	return ExportSearchStrategy{true, val}
}

func NewEmptyExportSearchStrategy() ExportSearchStrategy {
	return ExportSearchStrategy{}
}

func DecodeJSONExportSearchStrategy(dec *gojay.Decoder, val *ExportSearchStrategy) error {
	val.set = true
	return dec.String(&val.val)
}

type ExportSearchStrategy struct {
	set bool
	val string
}

func (q *ExportSearchStrategy) IsSet() bool {
	return q.set
}

func (q *ExportSearchStrategy) Get() string {
	if !q.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return q.val
}

func (q *ExportSearchStrategy) Set(val string) {
	q.set = true
	q.val = val
}

func (q *ExportSearchStrategy) Unset() {
	q.set = false
	q.val = ""
}

func (q *ExportSearchStrategy) Flag() string {
	return consts.FlagExportSearchStrategy
}

func (q *ExportSearchStrategy) IsDefault() bool {
	return !q.set
}

func (q *ExportSearchStrategy) FlagString() string {
	return q.Flag() + "=" + q.Get()
}

