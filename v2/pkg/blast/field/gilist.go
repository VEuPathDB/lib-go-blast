package field

import (
	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
)

func DecodeJSONGIList(dec *gojay.Decoder, val *GIList) error {
	val.set = true
	return dec.String(&val.val)
}

type GIList struct {
	set bool
	val string
}

func (g *GIList) IsSet() bool {
	return g.set
}

func (g *GIList) Get() string {
	if !g.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return g.val
}

func (g *GIList) Set(val string) {
	g.set = true
	g.val = val
}

func (g *GIList) Unset() {
	g.set = false
	g.val = ""
}

func (g *GIList) Flag() string {
	return consts.FlagGIList
}

func (g *GIList) IsDefault() bool {
	return !g.set || len(g.val) == 0
}

func (g *GIList) FlagString() string {
	return g.Flag() + "='" + g.Get() + "'"
}

