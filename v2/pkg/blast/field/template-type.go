package field

import (
	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/bval"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
)

func DecodeJSONTemplateType(dec *gojay.Decoder, val *TemplateType) error {
	val.set = true
	return dec.DecodeString((*string)(&val.val))
}

type TemplateType struct {
	set bool
	val TemplateTypeEnum
}

func (t *TemplateType) IsSet() bool {
	return t.set
}

func (t *TemplateType) Get() TemplateTypeEnum {
	if !t.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return t.val
}

func (t *TemplateType) Set(enum TemplateTypeEnum) {
	t.set = true
	t.val = enum
}

func (t *TemplateType) Unset() {
	t.set = false
}

func (t *TemplateType) Flag() string {
	return consts.FlagTemplateType
}

func (t *TemplateType) IsDefault() bool {
	return !t.set
}

func (t *TemplateType) FlagString() string {
	return t.Flag() + "=" + string(t.Get())
}

func (t *TemplateType) Validate(em bval.ValidationBuilder) {
	if !t.IsDefault() && !t.Get().IsValid() {
		_ = em.InvalidStrEnum(
			t.Flag(),
			t.Get().String(),
			TemplateTypeEnumCoding.String(),
			TemplateTypeEnumCodingAndOptimal.String(),
			TemplateTypeEnumOptimal.String(),
		)
	}
}

type TemplateTypeEnum string

const (
	TemplateTypeEnumCoding           TemplateTypeEnum = "coding"
	TemplateTypeEnumCodingAndOptimal TemplateTypeEnum = "coding_and_optimal"
	TemplateTypeEnumOptimal          TemplateTypeEnum = "optimal"
)

func (t TemplateTypeEnum) Value() string {
	return string(t)
}

func (t TemplateTypeEnum) String() string {
	return string(t)
}

func (t TemplateTypeEnum) IsValid() bool {
	switch t {
	case TemplateTypeEnumCoding:
		return true
	case TemplateTypeEnumCodingAndOptimal:
		return true
	case TemplateTypeEnumOptimal:
		return true
	default:
		return false
	}
}
