package field

import (
	"strconv"

	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v1/pkg/blast/consts"
)

func NewReward(val uint32) Reward {
	return Reward{true, val}
}

func NewEmptyReward() Reward {
	return Reward{}
}

func DecodeJSONReward(dec *gojay.Decoder, val *Reward) error {
	val.set = true
	return dec.Uint32(&val.val)
}

type Reward struct {
	set bool
	val uint32
}

func (b Reward) IsSet() bool {
	return b.set
}

func (b Reward) Get() uint32 {
	if !b.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return b.val
}

func (b *Reward) Set(val uint32) {
	b.set = true
	b.val = val
}

func (b *Reward) Unset() {
	b.set = false
	b.val = 0
}

func (b Reward) Flag() string {
	return consts.FlagReward
}

func (b Reward) IsDefault() bool {
	return !b.set
}

func (b Reward) ArgString() string {
	return strconv.FormatUint(uint64(b.Get()), 10)
}

func (b Reward) FlagString() string {
	return b.Flag() + "=" + b.ArgString()
}

