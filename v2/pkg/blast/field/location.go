package field

import (
	"fmt"

	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
)

type location struct {
	Start uint32
	Stop  uint32
}

func (l *location) IsDefault() bool {
	return l.Start == 0 && l.Stop == 0
}

func (l *location) ArgString() string {
	return fmt.Sprintf("%d-%d", l.Start, l.Stop)
}

func (l *location) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	switch key {
	case consts.KeyStart:
		return dec.Uint32(&l.Start)
	case consts.KeyStop:
		return dec.Uint32(&l.Stop)
	default:
		return nil
	}
}

func (l *location) NKeys() int {
	return 2
}

func NewSubjectLocation(start, stop uint32) SubjectLocation {
	return SubjectLocation{location{start, stop}}
}

func NewEmptySubjectLocation() SubjectLocation {
	return SubjectLocation{}
}

func DecodeJSONSubjectLocation(dec *gojay.Decoder, val *SubjectLocation) error {
	return dec.Object(val)
}

type SubjectLocation struct{ location }

func (s SubjectLocation) Flag() string {
	return consts.FlagSubjectLocation
}

func (s SubjectLocation) FlagString() string {
	return s.Flag() + "=" + s.ArgString()
}

func NewQueryLocation(start, stop uint32) QueryLocation {
	return QueryLocation{location{start, stop}}
}

func NewEmptyQueryLocation() QueryLocation {
	return QueryLocation{}
}

func DecodeJSONQueryLocation(dec *gojay.Decoder, val *QueryLocation) error {
	return dec.Object(val)
}

type QueryLocation struct{ location }

func (q QueryLocation) Flag() string {
	return consts.FlagQueryLocation
}

func (q QueryLocation) FlagString() string {
	return q.Flag() + "=" + q.ArgString()
}


