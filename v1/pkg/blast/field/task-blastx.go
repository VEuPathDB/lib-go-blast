package field

import (
	"github.com/francoispqt/gojay"
	"lib-go-blast/v1/pkg/blast/consts"
)

const DefaultBlastXTaskType = BlastXTaskTypeBlastX

func DecodeJSONBlastXTask(dec *gojay.Decoder, val *BlastXTask) error {
	val.set = true
	return dec.DecodeString((*string)(&val.val))
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

type BlastXTaskType string

const (
	BlastXTaskTypeBlastX     BlastXTaskType = "blastx"
	BlastXTaskTypeBlastXFast BlastXTaskType = "blastx-fast"
)

func (b BlastXTaskType) String() string {
	return string(b)
}
