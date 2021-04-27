package field

import (
	"github.com/francoispqt/gojay"
	"lib-go-blast/v1/pkg/blast/consts"
)

const DefaultTBlastNTaskType = TBlastNTaskTBlastN

func DecodeJSONTBlastNTask(dec *gojay.Decoder, val *TBlastNTask) error {
	val.set = true
	return dec.DecodeString((*string)(&val.val))
}

type TBlastNTask struct {
	set bool
	val TBlastNTaskType
}

func (t TBlastNTask) IsSet() bool {
	return t.set
}

func (t TBlastNTask) Get() TBlastNTaskType {
	if !t.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return t.val
}

func (t *TBlastNTask) Set(val TBlastNTaskType)  {
	t.set = true
	t.val = val
}

func (t *TBlastNTask) Unset() {
	t.set = false
}

func (t TBlastNTask) Flag() string {
	return consts.FlagTask
}

func (t TBlastNTask) IsDefault() bool {
	return !t.set || t.val == DefaultTBlastNTaskType
}

func (t TBlastNTask) FlagString() string {
	return t.Flag() + "=" + t.Get().String()
}

type TBlastNTaskType string

const (
	TBlastNTaskTBlastN     TBlastNTaskType = "tblastn"
	TBlastNTaskTBlastNFast TBlastNTaskType = "tblastn-fast"
)

func (t TBlastNTaskType) String() string {
	return string(t)
}
