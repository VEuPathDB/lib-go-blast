package field

import (
	"strconv"

	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/bval"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
)

const DefaultMTModeValue = MTModeSplitByDBVolumes

func DecodeJSONMTMode(dec *gojay.Decoder, val *MTMode) error {
	val.set = true
	return dec.DecodeUint8((*uint8)(&val.val))
}

type MTMode struct {
	set bool
	val MTModeEnum
}

func (m *MTMode) IsSet() bool {
	return m.set
}

func (m *MTMode) Get() MTModeEnum {
	if !m.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return m.val
}

func (m *MTMode) Set(val MTModeEnum) {
	m.set = true
	m.val = val
}

func (m *MTMode) Unset() {
	m.set = false
}

func (m *MTMode) Flag() string {
	return consts.FlagMTMode
}

func (m *MTMode) IsDefault() bool {
	return !m.set || m.val == DefaultMTModeValue
}

func (m *MTMode) FlagString() string {
	return m.Flag() + "=" + m.Get().String()
}

func (m *MTMode) Validate(em bval.ValidationBuilder) {
	if !m.IsDefault() && !m.Get().IsValid() {
		_ = em.InvalidUint8Enum(
			m.Flag(),
			uint8(m.Get()),
			uint8(MTModeSplitByDBVolumes),
			uint8(MTModeSplitByQueries),
		)
	}
}

// ////////////////////////////////////////////////////////////////////////////////////////////// //

type MTModeEnum uint8

const (
	MTModeSplitByDBVolumes MTModeEnum = iota
	MTModeSplitByQueries
)

func (m MTModeEnum) Value() uint8 {
	return uint8(m)
}

func (m MTModeEnum) String() string {
	return strconv.Itoa(int(m))
}

func (m MTModeEnum) IsValid() bool {
	return m <= MTModeSplitByQueries
}
