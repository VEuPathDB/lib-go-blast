package field

import (
	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
)

func DecodeJSONTaxIDList(dec *gojay.Decoder, val *TaxIDList) error {
	val.set = true
	return dec.String(&val.val)
}

type TaxIDList struct {
	set bool
	val string
}

func (t *TaxIDList) IsSet() bool {
	return t.set
}

func (t *TaxIDList) Get() string {
	if !t.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return t.val
}

func (t *TaxIDList) Set(val string) {
	t.set = true
	t.val = val
}

func (t *TaxIDList) Unset() {
	t.set = false
	t.val = ""
}

func (t *TaxIDList) Flag() string {
	return consts.FlagTaxIDList
}

func (t *TaxIDList) IsDefault() bool {
	return !t.set || len(t.val) == 0
}

func (t *TaxIDList) FlagString() string {
	return t.Flag() + "='" + t.Get() + "'"
}

