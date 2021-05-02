package field

import (
	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
)

func NewPhiPatternFile(val string) PhiPatternFile {
	return PhiPatternFile{true, val}
}

func NewEmptyPhiPatternFile() PhiPatternFile {
	return PhiPatternFile{}
}

func DecodeJSONPhiPatternFile(dec *gojay.Decoder, val *PhiPatternFile) error {
	val.set = true
	return dec.String(&val.val)
}

type PhiPatternFile struct {
	set bool
	val string
}

func (p *PhiPatternFile) IsSet() bool {
	return p.set
}

func (p *PhiPatternFile) Get() string {
	if !p.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return p.val
}

func (p *PhiPatternFile) Set(val string) {
	p.set = true
	p.val = val
}

func (p *PhiPatternFile) Unset() {
	p.set = false
	p.val = ""
}

func (p *PhiPatternFile) Flag() string {
	return consts.FlagPhiPatternFile
}

func (p *PhiPatternFile) IsDefault() bool {
	return !p.set || len(p.val) == 0
}

func (p *PhiPatternFile) FlagString() string {
	return p.Flag() + "='" + p.Get() + "'"
}

