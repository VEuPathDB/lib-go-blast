package field

import (
	"strconv"

	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v1/pkg/blast/consts"
)

const DefaultMTModeValue = MTModeSplitByDBVolumes

func DecodeJSONMTMode(dec *gojay.Decoder, val *MTMode) error {
	val.set = true
	return dec.Uint8((*uint8)(&val.val))
}

type MTMode struct {
	set bool
	val MTModeValue
}

func (m MTMode) IsSet() bool {
	return m.set
}

func (m MTMode) Get() MTModeValue {
	if !m.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return m.val
}

func (m *MTMode) Set(val MTModeValue) {
	m.set = true
	m.val = val
}

func (m *MTMode) Unset() {
	m.set = false
}

func (m MTMode) Flag() string {
	return consts.FlagMTMode
}

func (m MTMode) IsDefault() bool {
	return !m.set || m.val == DefaultMTModeValue
}

func (m MTMode) FlagString() string {
	return m.Flag() + "=" + m.Get().String()
}

type MTModeValue uint8

const (
	MTModeSplitByDBVolumes MTModeValue = iota
	MTModeSplitByQueries
)

func (m MTModeValue) String() string {
	return strconv.Itoa(int(m))
}
