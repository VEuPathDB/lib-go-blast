package field

import (
	"strconv"

	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
)

func NewMinRawGappedScore(val int32) MinRawGappedScore {
	return MinRawGappedScore{true, val}
}

func NewEmptyMinRawGappedScore() MinRawGappedScore {
	return MinRawGappedScore{}
}

func DecodeJSONMinRawGappedScore(dec *gojay.Decoder, val *MinRawGappedScore) error {
	return dec.DecodeInt32(&val.val)
}

type MinRawGappedScore struct {
	set bool
	val int32
}

func (m *MinRawGappedScore) IsSet() bool {
	return m.set
}

func (m *MinRawGappedScore) Get() int32 {
	if !m.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return m.val
}

func (m *MinRawGappedScore) Set(val int32) {
	m.set = true
	m.val = val
}

func (m *MinRawGappedScore) Unset() {
	m.set = false
	m.val = 0
}

func (m *MinRawGappedScore) Flag() string {
	return consts.FlagMinRawGappedScore
}

func (m *MinRawGappedScore) IsDefault() bool {
	return !m.set
}

func (m *MinRawGappedScore) ArgString() string {
	return strconv.Itoa(int(m.Get()))
}

func (m *MinRawGappedScore) FlagString() string {
	return m.Flag() + "=" + m.ArgString()
}
