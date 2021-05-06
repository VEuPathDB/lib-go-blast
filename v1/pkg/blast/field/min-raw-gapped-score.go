package field

import (
	"strconv"

	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v1/pkg/blast/consts"
)

func NewMinRawGappedScore(val int32) MinRawGappedScore {
	return MinRawGappedScore{true, val}
}

func NewEmptyMinRawGappedScore() MinRawGappedScore {
	return MinRawGappedScore{}
}

func DecodeJSONMinRawGappedScore(dec *gojay.Decoder, val *MinRawGappedScore) error {
	return dec.Int32(&val.val)
}

type MinRawGappedScore struct {
	set bool
	val int32
}

func (b MinRawGappedScore) IsSet() bool {
	return b.set
}

func (b MinRawGappedScore) Get() int32 {
	if !b.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return b.val
}

func (b *MinRawGappedScore) Set(val int32) {
	b.set = true
	b.val = val
}

func (b *MinRawGappedScore) Unset() {
	b.set = false
	b.val = 0
}

func (b MinRawGappedScore) Flag() string {
	return consts.FlagMinRawGappedScore
}

func (b MinRawGappedScore) IsDefault() bool {
	return !b.set
}

func (b MinRawGappedScore) ArgString() string {
	return strconv.Itoa(int(b.Get()))
}

func (b MinRawGappedScore) FlagString() string {
	return b.Flag() + "=" + b.ArgString()
}
