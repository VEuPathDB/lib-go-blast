package field

import (
	"strconv"

	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/bval"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
)

const MinThreshold float64 = 0

func NewThreshold(val float64) Threshold {
	return Threshold{true, val}
}

func NewEmptyThreshold() Threshold {
	return Threshold{}
}

func DecodeJSONThreshold(dec *gojay.Decoder, val *Threshold) error {
	val.set = true
	return dec.DecodeFloat64(&val.val)
}

type Threshold struct {
	set bool
	val float64
}

func (t *Threshold) IsSet() bool {
	return t.set
}

func (t *Threshold) Get() float64 {
	if !t.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return t.val
}

func (t *Threshold) Set(val float64) {
	t.set = true
	t.val = val
}

func (t *Threshold) Unset() {
	t.set = false
	t.val = 0
}

func (t *Threshold) Flag() string {
	return consts.FlagThreshold
}

func (t *Threshold) IsDefault() bool {
	return !t.IsSet()
}

func (t *Threshold) ArgString() string {
	return strconv.FormatFloat(t.Get(), 'f', -1, 64)
}

func (t *Threshold) FlagString() string {
	return t.Flag() + "=" + t.ArgString()
}

func (t *Threshold) Validate(b bval.ValidationBuilder) {
	b.F64GTEQ(t, 0)
}
