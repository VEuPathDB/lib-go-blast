package field

import (
	"github.com/francoispqt/gojay"
	"lib-go-blast/v1/pkg/blast/consts"
)

func DecodeJSONTemplateType(dec *gojay.Decoder, val *TemplateType) error {
	val.set = true
	return dec.DecodeString((*string)(&val.val))
}

type TemplateType struct {
	set bool
	val TemplateTypeEnum
}

func (t TemplateType) IsSet() bool {
	return t.set
}

func (t TemplateType) Get() TemplateTypeEnum {
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

func (t TemplateType) Flag() string {
	return consts.FlagTemplateType
}

func (t TemplateType) IsDefault() bool {
	return !t.set
}

func (t TemplateType) FlagString() string {
	return t.Flag() + "=" + string(t.Get())
}

type TemplateTypeEnum string

const (
	TemplateTypeEnumCoding           TemplateTypeEnum = "coding"
	TemplateTypeEnumCodingAndOptimal TemplateTypeEnum = "coding_and_optimal"
	TemplateTypeEnumOptimal          TemplateTypeEnum = "optimal"
)