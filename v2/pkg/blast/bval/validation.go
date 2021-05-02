package bval

import (
	"fmt"
	"strings"

	"github.com/foxcapades/go-bytify/v0/bytify"
)

type Validator interface {
	Validate() ValidationError
}

type Validatable interface {
	Validate(ValidationBuilder)
}

type ValidationError map[string][]string

func (e ValidationError) Error() string {
	l := 0
	for _, v := range e {
		for i := range v {
			l += len(v[i])
		}
	}

	return fmt.Sprintf("%d validation error(s)", l)
}

// ////////////////////////////////////////////////////////////////////////////////////////////// //

const (
	ErrFloatExclusiveRange = `Invalid %s value "%f".  Value must be in the range (%f..%f).`
	ErrFloatInclusiveRange = `Invalid %s value "%f".  Value must be in the range [%f..%f].`
	ErrIncompatibleWith    = "Flag %s is incompatible with Flag %s."
	ErrRequires            = "Flag %s requires flag %s to also be set."
	ErrIntLTEQ             = `Invalid %s value "%d".  Value must be <= %d.`
	ErrIntGTEQ             = `Invalid %s value "%d".  Value must be >= %d.`
	ErrFloatGTEQ           = `Invalid %s value "%f".  Value must be >= %f.`

	errInvalidEnum         = `Invalid  value "".  Must be one of [, ].`
	errReportTypeLTEQ      = `%s is not applicable for format types > %d`
	errReportTypeEQ        = `%s is not applicable for format types != %d`
)

type ValidationBuilder map[string][]string

// Validate collects the validation output for all passed args into this
// ValidationBuilder.
func (e ValidationBuilder) Validate(vals ...Validatable) ValidationBuilder {
	for i := range vals {
		vals[i].Validate(e)
	}

	return e
}

// Requires appends an error if `this` is set, but `that` is not.
func (e ValidationBuilder) Requires(this, that Defaulter) ValidationBuilder {
	if !this.IsDefault() && that.IsDefault() {
		return e.AppendError(this.Flag(), fmt.Sprintf(ErrRequires, this.Flag(), that.Flag()))
	}

	return e
}

func (e ValidationBuilder) AppendError(key, val string) ValidationBuilder {
	if _, ok := e[key]; ok {
		e[key] = append(e[key], val)
	} else {
		e[key] = []string{val}
	}

	return e
}

// Incompatible appends an error if both `v1` and `v2` are set.
func (e ValidationBuilder) Incompatible(v1 Defaulter, v2 Defaulter) ValidationBuilder {
	def1, def2 := v1.IsDefault(), v2.IsDefault()

	if def1 == def2 && !def1 {
		k1, k2 := v1.Flag(), v2.Flag()

		_ = e.AppendError(k1, fmt.Sprintf(ErrIncompatibleWith, k1, k2))
		_ = e.AppendError(k2, fmt.Sprintf(ErrIncompatibleWith, k2, k1))
	}

	return e
}

// IncompatibleWith calls Incompatible on tgt and each entry in others.
func (e ValidationBuilder) IncompatibleWith(tgt Defaulter, others ...Defaulter) ValidationBuilder {
	for i := range others {
		_ = e.Incompatible(tgt, others[i])
	}

	return e
}

// IncompatibleList verifies that only one of the given fields is set.
//
// If more than one field is set, an error will be inserted for each set field.
func (e ValidationBuilder) IncompatibleList(fields ...Defaulter) ValidationBuilder {
	set := make([]Defaulter, 0, 3)

	for _, v := range fields {
		if !v.IsDefault() {
			set = append(set, v)
		}
	}

	ln := len(set)

	if ln < 2 {
		return e
	}

	for i := 0; i < ln; i++ {
		for j := i + 1; j < ln; j++ {
			_ = e.Incompatible(set[i], set[j])
		}
	}

	return e
}

func (e ValidationBuilder) InvalidStrEnum(key string, val string, allowed ...string) ValidationBuilder {
	l := len(errInvalidEnum) + len(key) + len(val) + 2*(len(allowed)-1)
	for i := range allowed {
		l += len(allowed[i])
	}

	o := strings.Builder{}
	o.Grow(l)
	o.Reset()

	o.WriteString(errInvalidEnum[:8])
	o.WriteString(key)
	o.WriteString(errInvalidEnum[8:16])
	o.WriteString(val)
	o.WriteString(errInvalidEnum[16:36])

	s := false
	for i := range allowed {
		if s {
			o.WriteString(errInvalidEnum[36:38])
		} else {
			s = true
		}

		o.WriteString(allowed[i])
	}

	o.WriteString(errInvalidEnum[38:])

	return e.AppendError(key, o.String())
}

func (e ValidationBuilder) InvalidUint8Enum(key string, val uint8, allowed ...uint8) ValidationBuilder {
	l := len(errInvalidEnum) +
		len(key) +
		int(bytify.Uint8StringSize(val)) +
		2*(len(allowed)-1)
	for _, v := range allowed {
		l += int(bytify.Uint8StringSize(v))
	}

	o := strings.Builder{}
	o.Grow(l)
	o.Reset()

	o.WriteString(errInvalidEnum[:8])
	o.WriteString(key)
	o.WriteString(errInvalidEnum[8:16])
	bytify.Uint8ToBuf(val, &o)
	o.WriteString(errInvalidEnum[16:36])

	s := false
	for i := range allowed {
		if s {
			o.WriteString(errInvalidEnum[36:38])
		} else {
			s = true
		}

		bytify.Uint8ToBuf(allowed[i], &o)
	}

	o.WriteString(errInvalidEnum[38:])

	return e.AppendError(key, o.String())
}

func (e ValidationBuilder) F64GTEQ(val F64Defaulter, min float64) ValidationBuilder {
	if val.IsDefault() || val.Get() >= min {
		return e
	}

	return e.AppendError(val.Flag(), fmt.Sprintf(ErrFloatGTEQ, val.Flag(), val.Get(), min))
}

func (e ValidationBuilder) I32LTEQ(val I32Defaulter, max int32) ValidationBuilder {
	if val.IsDefault() || val.Get() <= max {
		return e
	}

	return e.AppendError(val.Flag(), fmt.Sprintf(ErrIntLTEQ, val.Flag(), val.Get(), max))
}

// Uint8GTEQ inserts a new error for a uint8 field value if it is below the
// given minimum.
func (e ValidationBuilder) Uint8GTEQ(val U8Defaulter, min uint8) ValidationBuilder {
	if val.IsDefault() || val.Get() >= min {
		return e
	}

	return e.AppendError(val.Flag(), fmt.Sprintf(ErrIntGTEQ, val.Flag(), val.Get(), min))
}

// Uint16GTEQ inserts a new error for a uint16 field value if it is below the
// given minimum.
func (e ValidationBuilder) Uint16GTEQ(val U16Defaulter, min uint16) ValidationBuilder {
	if val.IsDefault() || val.Get() >= min {
		return e
	}

	return e.AppendError(val.Flag(), fmt.Sprintf(ErrIntGTEQ, val.Flag(), val.Get(), min))
}

// U32GTEQ inserts a new error for a uint32 field value if it is below the
// given minimum.
func (e ValidationBuilder) U32GTEQ(val U32Defaulter, min uint32) ValidationBuilder {
	if val.IsDefault() || val.Get() >= min {
		return e
	}

	return e.AppendError(val.Flag(), fmt.Sprintf(ErrIntGTEQ, val.Flag(), val.Get(), min))
}

func (e ValidationBuilder) Uint8LTEQ(val U8Defaulter, max uint8) ValidationBuilder {
	if val.IsDefault() || val.Get() <= max {
		return e
	}

	return e.AppendError(
		val.Flag(),
		fmt.Sprintf(ErrIntLTEQ, val.Flag(), val.Get(), max),
	)
}

// F64InclusiveRange inserts a new error for a float64 value if it falls outside
// the given min and max values.
func (e ValidationBuilder) F64InclusiveRange(flag F64Defaulter, min, max float64) ValidationBuilder {
	if flag.IsDefault() {
		return e
	}

	if flag.Get() < min || flag.Get() > max {
		_ = e.AppendError(
			flag.Flag(),
			fmt.Sprintf(ErrFloatInclusiveRange, flag.Flag(), flag.Get(), min, max),
		)
	}

	return e
}

// F64ExclusiveRange inserts a new error for a float64 value if it falls outside
// the given min and max values.
func (e ValidationBuilder) F64ExclusiveRange(flag F64Defaulter, min, max float64) ValidationBuilder {
	if flag.IsDefault() {
		return e
	}

	if flag.Get() <= min || flag.Get() >= max {
		_ = e.AppendError(
			flag.Flag(),
			fmt.Sprintf(ErrFloatExclusiveRange, flag.Flag(), flag.Get(), min, max),
		)
	}

	return e
}

func (e ValidationBuilder) ReportTypeEQ(val, exact uint8, keys ...Defaulter) ValidationBuilder {
	if val != exact {
		for i := range keys {
			if !keys[i].IsDefault() {
				_ = e.AppendError(keys[i].Flag(), fmt.Sprintf(errReportTypeEQ, keys[i].Flag(), exact))
			}
		}
	}

	return e
}

func (e ValidationBuilder) ReportTypeLTEQ(val, max uint8, keys ...Defaulter) ValidationBuilder {
	if val > max {
		for i := range keys {
			if !keys[i].IsDefault() {
				_ = e.AppendError(keys[i].Flag(), fmt.Sprintf(errReportTypeLTEQ, keys[i].Flag(), max))
			}
		}
	}

	return e
}

func (e ValidationBuilder) PutAll(other ValidationBuilder) ValidationBuilder {
	for k, v := range other {
		if _, ok := e[k]; ok {
			e[k] = append(e[k], v...)
		} else {
			e[k] = make([]string, len(v))
			copy(e[k], v)
		}
	}

	return e
}
