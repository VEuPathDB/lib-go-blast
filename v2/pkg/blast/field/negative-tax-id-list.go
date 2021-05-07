package field

import (
	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
)

func DecodeJSONNegativeTaxIDList(dec *gojay.Decoder, val *NegativeTaxIDList) error {
	val.set = true
	return dec.String(&val.val)
}

type NegativeTaxIDList struct {
	set bool
	val string
}

func (q *NegativeTaxIDList) IsSet() bool {
	return q.set
}

func (q *NegativeTaxIDList) Get() string {
	if !q.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return q.val
}

func (q *NegativeTaxIDList) Set(val string) {
	q.set = true
	q.val = val
}

func (q *NegativeTaxIDList) Unset() {
	q.set = false
	q.val = ""
}

func (q *NegativeTaxIDList) Flag() string {
	return consts.FlagNegativeTaxIDList
}

func (q *NegativeTaxIDList) IsDefault() bool {
	return !q.set || len(q.val) == 0
}

func (q *NegativeTaxIDList) FlagString() string {
	return q.Flag() + "=" + q.Get()
}

