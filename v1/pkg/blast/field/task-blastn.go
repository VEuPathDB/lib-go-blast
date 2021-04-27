package field

import (
	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v1/pkg/blast/consts"
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

func (b BlastNTask) IsSet() bool {
	return b.set
}

func (b BlastNTask) Get() BlastNTaskType {
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

func (b BlastNTask) Flag() string {
	return consts.FlagTask
}

func (b BlastNTask) IsDefault() bool {
	return !b.IsSet() || b.val == DefaultBlastNTaskType
}

func (b BlastNTask) FlagString() string {
	return b.Flag() + "=" + b.Get().String()
}

type BlastNTaskType string

const (
	BlastNTaskTypeBlastN                 BlastNTaskType = "blastn"
	BlastNTaskTypeBlastNShort            BlastNTaskType = "blastn-short"
	BlastNTaskTypeDiscontiguousMegablast BlastNTaskType = "dc-megablast"
	BlastNTaskTypeMegablast              BlastNTaskType = "megablast"
	BlastNTaskTypeRMBlastN               BlastNTaskType = "rmblastn"
)

func (b BlastNTaskType) String() string {
	return string(b)
}
