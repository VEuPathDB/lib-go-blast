package field

import (
	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/bval"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
)

const DefaultBlastPTaskType = BlastPTaskTypeBlastP

func DecodeJSONBlastPTask(dec *gojay.Decoder, val *BlastPTask) error {
	val.set = true
	return dec.String((*string)(&val.val))
}

type BlastPTask struct {
	set bool
	val BlastPTaskType
}

func (b *BlastPTask) IsSet() bool {
	return b.set
}

func (b *BlastPTask) Get() BlastPTaskType {
	if !b.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return b.val
}

func (b *BlastPTask) Set(val BlastPTaskType) {
	b.set = true
	b.val = val
}

func (b *BlastPTask) Unset() {
	b.set = false
}

func (b *BlastPTask) Flag() string {
	return consts.FlagTask
}

func (b *BlastPTask) IsDefault() bool {
	return !b.set || b.val == DefaultBlastPTaskType
}

func (b *BlastPTask) FlagString() string {
	return b.Flag() + "=" + b.Get().String()
}

func (b *BlastPTask) Validate(em bval.ValidationBuilder) {
	if !b.IsDefault() && !b.Get().IsValid() {
		_ = em.InvalidStrEnum(
			b.Flag(),
			string(b.Get()),
			string(BlastPTaskTypeBlastP),
			string(BlastPTaskTypeBlastPFast),
			string(BlastPTaskTypeBlastPShort),
		)
	}
}

type BlastPTaskType string

const (
	BlastPTaskTypeBlastP      BlastPTaskType = "blastp"
	BlastPTaskTypeBlastPFast  BlastPTaskType = "blastp-fast"
	BlastPTaskTypeBlastPShort BlastPTaskType = "blastp-short"
)

func (b BlastPTaskType) String() string {
	return string(b)
}

func (b BlastPTaskType) IsValid() bool {
	switch b {
	case BlastPTaskTypeBlastP:
		return true
	case BlastPTaskTypeBlastPFast:
		return true
	case BlastPTaskTypeBlastPShort:
		return true
	default:
		return false
	}
}
