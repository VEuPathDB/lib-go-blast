package field

import (
	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
)

func NewFilteringDB(val string) FilteringDB {
	return FilteringDB{true, val}
}

func NewEmptyFilteringDB() FilteringDB {
	return FilteringDB{}
}

func DecodeJSONFilteringDB(dec *gojay.Decoder, val *FilteringDB) error {
	val.set = true
	return dec.DecodeString(&val.val)
}

type FilteringDB struct {
	set bool
	val string
}

func (q *FilteringDB) IsSet() bool {
	return q.set
}

func (q *FilteringDB) Get() string {
	if !q.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return q.val
}

func (q *FilteringDB) Set(val string) {
	q.set = true
	q.val = val
}

func (q *FilteringDB) Unset() {
	q.set = false
	q.val = ""
}

func (q *FilteringDB) Flag() string {
	return consts.FlagFilteringDB
}

func (q *FilteringDB) IsDefault() bool {
	return !q.set || len(q.val) == 0
}

func (q *FilteringDB) FlagString() string {
	return q.Flag() + "='" + q.Get() + "'"
}

