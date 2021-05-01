package field

import (
	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/bval"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
)

const DefaultBlastNTaskType = BlastNTaskTypeMegablast

func NewBlastNTask(val BlastNTaskType) BlastNTask {
	return BlastNTask{true, val}
}

func NewEmptyBlastNTask() BlastNTask {
	return BlastNTask{}
}

func DecodeJSONBlastNTask(dec *gojay.Decoder, val *BlastNTask) error {
	val.set = true
	return dec.DecodeString((*string)(&val.val))
}

type BlastNTask struct {
	set bool
	val BlastNTaskType
}

func (b *BlastNTask) IsSet() bool {
	return b.set
}

func (b *BlastNTask) Get() BlastNTaskType {
	if !b.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return b.val
}

func (b *BlastNTask) Set(val BlastNTaskType) {
	b.set = true
	b.val = val
}

func (b *BlastNTask) Unset() {
	b.set = false
}

func (b *BlastNTask) Flag() string {
	return consts.FlagTask
}

func (b *BlastNTask) IsDefault() bool {
	return !b.IsSet() || b.val == DefaultBlastNTaskType
}

func (b *BlastNTask) FlagString() string {
	return b.Flag() + "=" + b.Get().String()
}

func (b *BlastNTask) Validate(em bval.ValidationBuilder) {
	if !b.IsDefault() && !b.Get().IsValid() {
		_ = em.InvalidStrEnum(
			b.Flag(),
			string(b.Get()),
			string(BlastNTaskTypeBlastN),
			string(BlastNTaskTypeBlastNShort),
			string(BlastNTaskTypeDiscontiguousMegablast),
			string(BlastNTaskTypeMegablast),
			string(BlastNTaskTypeRMBlastN),
		)
	}
}

type BlastNTaskType string

const (
	BlastNTaskTypeBlastN                 BlastNTaskType = "blastn"
	BlastNTaskTypeBlastNShort            BlastNTaskType = "blastn-short"
	BlastNTaskTypeDiscontiguousMegablast BlastNTaskType = "dc-megablast"
	BlastNTaskTypeMegablast              BlastNTaskType = "megablast"
	BlastNTaskTypeRMBlastN               BlastNTaskType = "rmblastn"
)

func (b BlastNTaskType) Value() string {
	return string(b)
}

func (b BlastNTaskType) String() string {
	return string(b)
}

func (b BlastNTaskType) IsValid() bool {
	switch b {
	case BlastNTaskTypeBlastN:
		return true
	case BlastNTaskTypeBlastNShort:
		return true
	case BlastNTaskTypeDiscontiguousMegablast:
		return true
	case BlastNTaskTypeMegablast:
		return true
	case BlastNTaskTypeRMBlastN:
		return true
	default:
		return false
	}
}
