package field

import (
	"regexp"

	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/bval"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
)

// TODO: EValue should be a newtype in v2

const DefaultExpectValue = "10"

func NewExpectValue(val string) ExpectValue {
	return ExpectValue{true, val}
}

func NewEmptyExpectValue() ExpectValue {
	return ExpectValue{}
}

func DecodeJSONExpectValue(dec *gojay.Decoder, val *ExpectValue) error {
	val.set = true
	return dec.String(&val.val)
}

type ExpectValue struct {
	set bool
	val string
}

func (q *ExpectValue) IsSet() bool {
	return q.set
}

func (q *ExpectValue) Get() string {
	if !q.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return q.val
}

func (q *ExpectValue) Set(val string) {
	q.set = true
	q.val = val
}

func (q *ExpectValue) Unset() {
	q.set = false
	q.val = ""
}

func (q *ExpectValue) Flag() string {
	return consts.FlagExpectValue
}

func (q *ExpectValue) IsDefault() bool {
	return !q.set || q.val == DefaultExpectValue
}

func (q *ExpectValue) FlagString() string {
	return q.Flag() + "=" + q.Get()
}

var evalPattern = regexp.MustCompile("^[-+]?\\d+(?:\\.\\d+)?(?:[eE][-+]?\\d+)?$")

func (q *ExpectValue) Validate(em bval.ValidationBuilder) {
	if !q.IsDefault() && !evalPattern.MatchString(q.Get()) {
		_ = em.AppendError(
			q.Flag(),
			q.Flag() + " must be a number formatted as: 1, -1.0, 1e10, 1E-10, etc..",
		)
	}
}

