package field

import (
	"errors"
	"fmt"

	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v1/pkg/blast/consts"
	"github.com/veupathdb/lib-go-blast/v1/pkg/cli"
)

func NewYesDust() Dust {
	return yesDust{}
}

func NewNoDust() Dust {
	return noDust{}
}

func NewLWLDust(level, window, linker int32) Dust {
	return lwlDust{level, window, linker}
}

func DecodeJSONDust(dec *gojay.Decoder) (Dust, error) {
	var tmp1 string

	if dec.String(&tmp1) == nil {
		switch tmp1 {
		case "yes":
			return yesDust{}, nil
		case "no":
			return noDust{}, nil
		default:
			return nil, errors.New("Invalid dust value.")
		}
	}

	var tmp2 lwlDust

	if dec.Object(&tmp2) == nil {
		return tmp2, nil
	}

	return nil, errors.New("Invalid dust value.")
}

type Dust interface {
	cli.Appender
	IsYes() bool
	IsNo() bool
	IsLWL() bool

	Level() int32
	Window() int32
	Linker() int32
}

type yesDust struct{}

func (yesDust) Flag() string {
	return consts.FlagDust
}

func (yesDust) IsDefault() bool {
	return false
}

func (y yesDust) FlagString() string {
	return y.Flag() + "=yes"
}

func (yesDust) IsYes() bool {
	return true
}

func (yesDust) IsNo() bool {
	return false
}

func (yesDust) IsLWL() bool {
	return false
}

func (yesDust) Level() int32 {
	panic("Cannot get level value from a 'yes' dust flag.")
}

func (yesDust) Window() int32 {
	panic("Cannot get window value from a 'yes' dust flag.")
}

func (yesDust) Linker() int32 {
	panic("Cannot get linker value from a 'yes' dust flag.")
}

type noDust struct{}

func (noDust) Flag() string {
	return consts.FlagDust
}

func (noDust) IsDefault() bool {
	return false
}

func (n noDust) FlagString() string {
	return n.Flag() + "=no"
}

func (noDust) IsYes() bool {
	return false
}

func (noDust) IsNo() bool {
	return true
}

func (noDust) IsLWL() bool {
	return false
}

func (noDust) Level() int32 {
	panic("Cannot get level value from a 'no' dust flag.")
}

func (noDust) Window() int32 {
	panic("Cannot get window value from a 'no' dust flag.")
}

func (noDust) Linker() int32 {
	panic("Cannot get linker value from a 'no' dust flag.")
}

type lwlDust struct {
	level  int32
	window int32
	linker int32
}

func (lwlDust) Flag() string {
	return consts.FlagDust
}

func (l lwlDust) IsDefault() bool {
	return l.Level() == 20 && l.Window() == 64 && l.Linker() == 1
}

func (l lwlDust) ArgString() string {
	return fmt.Sprintf("'%d %d %d'", l.Level(), l.Window(), l.Linker())
}

func (l lwlDust) FlagString() string {
	return l.Flag() + "=" + l.ArgString()
}

func (lwlDust) IsYes() bool {
	return false
}

func (lwlDust) IsNo() bool {
	return false
}

func (lwlDust) IsLWL() bool {
	return true
}

func (l lwlDust) Level() int32 {
	return l.level
}

func (l lwlDust) Window() int32 {
	return l.window
}

func (l lwlDust) Linker() int32 {
	return l.linker
}

func (l *lwlDust) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	switch key {
	case consts.KeyLevel:
		return dec.Int32(&l.level)
	case consts.KeyWindow:
		return dec.Int32(&l.window)
	case consts.KeyLinker:
		return dec.Int32(&l.linker)
	default:
		return nil
	}
}

func (l *lwlDust) NKeys() int {
	return 3
}

