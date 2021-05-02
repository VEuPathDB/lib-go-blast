package field

import (
	"strconv"

	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/bval"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
)

func DecodeJSONTemplateLength(dec *gojay.Decoder, val *TemplateLength) error {
	val.set = true
	return dec.Uint8((*uint8)(&val.val))
}

type TemplateLength struct {
	set bool
	val TemplateLengthEnum
}

func (t *TemplateLength) IsSet() bool {
	return t.set
}

func (t *TemplateLength) Get() TemplateLengthEnum {
	if !t.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return t.val
}

func (t *TemplateLength) Set(val TemplateLengthEnum) {
	t.set = true
	t.val = val
}

func (t *TemplateLength) Unset() {
	t.set = false
}

func (t *TemplateLength) Flag() string {
	return consts.FlagTemplateLength
}

func (t *TemplateLength) IsDefault() bool {
	return !t.set
}

func (t *TemplateLength) FlagString() string {
	return t.Flag() + "=" + strconv.Itoa(int(t.Get()))
}

func (t *TemplateLength) Validate(em bval.ValidationBuilder) {
	if !t.IsDefault() && !t.Get().IsValid() {
		_ = em.InvalidUint8Enum(
			t.Flag(),
			t.Get().Value(),
			TemplateLengthEnum16.Value(),
			TemplateLengthEnum18.Value(),
			TemplateLengthEnum21.Value(),
		)
	}
}

type TemplateLengthEnum uint8

const (
	TemplateLengthEnum16 TemplateLengthEnum = 16
	TemplateLengthEnum18 TemplateLengthEnum = 18
	TemplateLengthEnum21 TemplateLengthEnum = 21
)

func (t TemplateLengthEnum) Value() uint8 {
	return uint8(t)
}

func (t TemplateLengthEnum) String() string {
	return strconv.Itoa(int(t))
}

func (t TemplateLengthEnum) IsValid() bool {
	switch t {
	case TemplateLengthEnum16:
		return true
	case TemplateLengthEnum18:
		return true
	case TemplateLengthEnum21:
		return true
	default:
		return false
	}
}

