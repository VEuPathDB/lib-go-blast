package field

import (
	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/bval"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
)

const DefaultBlastXTaskType = BlastXTaskTypeBlastX

func DecodeJSONBlastXTask(dec *gojay.Decoder, val *BlastXTask) error {
	val.set = true
	return dec.String((*string)(&val.val))
}

type BlastXTask struct {
	set bool
	val BlastXTaskType
}

func (b BlastXTask) IsSet() bool {
	return b.set
}

func (b BlastXTask) Get() BlastXTaskType {
	if !b.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return b.val
}

func (b *BlastXTask) Set(val BlastXTaskType) {
	b.set = true
	b.val = val
}

func (b *BlastXTask) Unset() {
	b.set = false
}

func (b BlastXTask) Flag() string {
	return consts.FlagTask
}

func (b BlastXTask) IsDefault() bool {
	return !b.set || b.val == DefaultBlastXTaskType
}

func (b BlastXTask) FlagString() string {
	return b.Flag() + "=" + b.Get().String()
}

func (b *BlastXTask) Validate(em bval.ValidationBuilder) {
	if !b.IsDefault() && !b.Get().IsValid() {
		_ = em.InvalidStrEnum(
			b.Flag(),
			string(b.Get()),
			string(BlastXTaskTypeBlastX),
			string(BlastXTaskTypeBlastXFast),
		)
	}
}

type BlastXTaskType string

const (
	BlastXTaskTypeBlastX     BlastXTaskType = "blastx"
	BlastXTaskTypeBlastXFast BlastXTaskType = "blastx-fast"
)

func (b BlastXTaskType) String() string {
	return string(b)
}

func (b BlastXTaskType) IsValid() bool {
	switch b {
	case BlastXTaskTypeBlastX:
		return true
	case BlastXTaskTypeBlastXFast:
		return true
	default:
		return false
	}
}
