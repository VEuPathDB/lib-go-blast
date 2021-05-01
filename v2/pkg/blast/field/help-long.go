package field

import (
	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
)

func NewLongHelp(val bool) LongHelp {
	return LongHelp{true, val}
}

func NewEmptyLongHelp() LongHelp {
	return LongHelp{}
}

func DecodeJSONLongHelp(dec *gojay.Decoder, val *LongHelp) error {
	val.set = true
	return dec.DecodeBool(&val.val)
}

type LongHelp struct {
	set bool
	val bool
}

func (l *LongHelp) IsSet() bool {
	return l.set
}

func (l *LongHelp) Get() bool {
	if !l.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return l.val
}

func (l *LongHelp) Set(val bool) {
	l.set = true
	l.val = val
}

func (l *LongHelp) Unset() {
	l.set = false
}

func (l *LongHelp) Flag() string {
	return consts.FlagLongHelp
}

func (l *LongHelp) IsDefault() bool {
	return !l.set || !l.val
}

func (l *LongHelp) FlagString() string {
	return l.Flag()
}

