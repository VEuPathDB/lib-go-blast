package field

import (
	"strconv"

	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/bval"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
)

const DefaultThreadModeValue = ThreadModeDisable

func DecodeJSONThreadMode(dec *gojay.Decoder, val *ThreadMode) error {
	val.set = true
	return dec.Uint8((*uint8)(&val.val))
}

type ThreadMode struct {
	set bool
	val ThreadModeEnum
}

func (t *ThreadMode) IsSet() bool {
	return t.set
}

func (t *ThreadMode) Get() ThreadModeEnum {
	if !t.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return t.val
}

func (t *ThreadMode) Set(val ThreadModeEnum) {
	t.set = true
	t.val = val
}

func (t *ThreadMode) Unset() {
	t.set = false
}

func (t *ThreadMode) Flag() string {
	return consts.FlagNumThreads
}

func (t *ThreadMode) IsDefault() bool {
	return !t.set || t.val == DefaultThreadModeValue
}

func (t *ThreadMode) FlagString() string {
	return t.Flag() + "=" + t.Get().String()
}

func (t *ThreadMode) Validate(b bval.ValidationBuilder) {
	if !t.IsDefault() && !t.Get().IsValid() {
		_ = b.InvalidUint8Enum(
			t.Flag(),
			t.Get().Value(),
			ThreadModeAuto.Value(),
			ThreadModeDisable.Value(),
		)
	}
}

// ////////////////////////////////////////////////////////////////////////////////////////////// //

type ThreadModeEnum uint8

const (
	ThreadModeAuto ThreadModeEnum = iota
	ThreadModeDisable
)

func (t ThreadModeEnum) Value() uint8 {
	return uint8(t)
}

func (t ThreadModeEnum) String() string {
	return strconv.Itoa(int(t))
}

func (t ThreadModeEnum) IsValid() bool {
	return t <= ThreadModeDisable
}

