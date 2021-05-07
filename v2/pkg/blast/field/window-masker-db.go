package field

import (
	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
)

func NewWindowMaskerDB(val string) WindowMaskerDB {
	return WindowMaskerDB{true, val}
}

func NewEmptyWindowMaskerDB() WindowMaskerDB {
	return WindowMaskerDB{}
}

func DecodeJSONWindowMaskerDB(dec *gojay.Decoder, val *WindowMaskerDB) error {
	val.set = true
	return dec.String(&val.val)
}

type WindowMaskerDB struct {
	set bool
	val string
}

func (q *WindowMaskerDB) IsSet() bool {
	return q.set
}

func (q *WindowMaskerDB) Get() string {
	if !q.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return q.val
}

func (q *WindowMaskerDB) Set(val string) {
	q.set = true
	q.val = val
}

func (q *WindowMaskerDB) Unset() {
	q.set = false
	q.val = ""
}

func (q *WindowMaskerDB) Flag() string {
	return consts.FlagWindowMaskerDB
}

func (q *WindowMaskerDB) IsDefault() bool {
	return !q.set || len(q.val) == 0
}

func (q *WindowMaskerDB) FlagString() string {
	return q.Flag() + "=" + q.Get()
}

