package field

import (
	"strconv"

	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v1/pkg/blast/consts"
)

func NewMSAMasterIDX(val uint32) MSAMasterIDX {
	return MSAMasterIDX{true, val}
}

func NewEmptyMSAMasterIDX() MSAMasterIDX {
	return MSAMasterIDX{}
}

func DecodeJSONMSAMasterIDX(dec *gojay.Decoder, val *MSAMasterIDX) error {
	val.set = true
	return dec.DecodeUint32(&val.val)
}

type MSAMasterIDX struct {
	set bool
	val uint32
}

func (m MSAMasterIDX) IsSet() bool {
	return m.set
}

func (m MSAMasterIDX) Get() uint32 {
	if !m.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return m.val
}

func (m *MSAMasterIDX) Set(val uint32) {
	m.set = true
	m.val = val
}

func (m *MSAMasterIDX) Unset() {
	m.set = false
	m.val = 0
}

func (MSAMasterIDX) Flag() string {
	return consts.FlagMSAMasterIDX
}

func (m MSAMasterIDX) IsDefault() bool {
	return !m.set
}

func (m MSAMasterIDX) ArgString() string {
	return strconv.FormatUint(uint64(m.Get()), 10)
}

func (m MSAMasterIDX) FlagString() string {
	return m.Flag() + "=" + m.ArgString()
}

