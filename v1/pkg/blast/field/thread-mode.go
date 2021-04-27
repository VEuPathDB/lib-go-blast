package field

import (
	"strconv"

	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v1/pkg/blast/consts"
)

const DefaultThreadModeValue = ThreadModeDisable

func DecodeJSONThreadMode(dec *gojay.Decoder, val *ThreadMode) error {
	val.set = true
	return dec.DecodeUint8((*uint8)(&val.val))
}

type ThreadMode struct {
	set bool
	val ThreadModeValue
}

func (t ThreadMode) IsSet() bool {
	return t.set
}

func (t ThreadMode) Get() ThreadModeValue {
	if !t.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return t.val
}

func (t *ThreadMode) Set(val ThreadModeValue) {
	t.set = true
	t.val = val
}

func (t *ThreadMode) Unset() {
	t.set = false
}

func (t ThreadMode) Flag() string {
	return consts.FlagNumThreads
}

func (t ThreadMode) IsDefault() bool {
	return !t.set || t.val == DefaultThreadModeValue
}

func (t ThreadMode) FlagString() string {
	return t.Flag() + "=" + t.Get().String()
}

type ThreadModeValue uint8

const (
	ThreadModeAuto ThreadModeValue = iota
	ThreadModeDisable
)

func (t ThreadModeValue) String() string {
	return strconv.Itoa(int(t))
}
