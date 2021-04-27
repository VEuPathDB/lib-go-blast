package field

import (
	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v1/pkg/blast/consts"
)

func DecodeJSONTaxIDList(dec *gojay.Decoder, val *TaxIDList) error {
	val.set = true
	return dec.DecodeString(&val.val)
}

type TaxIDList struct {
	set bool
	val string
}

func (q TaxIDList) IsSet() bool {
	return q.set
}

func (q TaxIDList) Get() string {
	if !q.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return q.val
}

func (q *TaxIDList) Set(val string) {
	q.set = true
	q.val = val
}

func (q *TaxIDList) Unset() {
	q.set = false
	q.val = ""
}

func (q TaxIDList) Flag() string {
	return consts.FlagTaxIDList
}

func (q TaxIDList) IsDefault() bool {
	return !q.set || len(q.val) == 0
}

func (q TaxIDList) FlagString() string {
	return q.Flag() + "='" + q.Get() + "'"
}

