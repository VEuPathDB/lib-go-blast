package field

import (
	"errors"
	"strconv"

	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v1/pkg/blast/consts"
	"github.com/veupathdb/lib-go-blast/v1/pkg/cli"
)

func NewYesSeg() Seg {
	return yesSeg{}
}

func NewNoSeg() Seg {
	return noSeg{}
}

func NewSeg(window int32, locut, hicut float64) Seg {
	return wlhSeg{window: window, locut: locut, hicut: hicut}
}

func DecodeJSONSeg(dec *gojay.Decoder) (Seg, error) {
	var tmp1 string

	if dec.String(&tmp1) == nil {
		switch tmp1 {
		case "yes":
			return yesSeg{}, nil
		case "no":
			return noSeg{}, nil
		default:
			return nil, errors.New("Invalid seg value.")
		}
	}

	var tmp2 wlhSeg

	if dec.Object(&tmp2) == nil {
		return tmp2, nil
	}

	return nil, errors.New("Invalid seg value.")
}

type Seg interface {
	cli.Appender

	IsYes() bool
	IsNo() bool
	IsWLH() bool

	Window() int32
	Locut() float64
	Hicut() float64
}

type yesSeg struct{}

func (yesSeg) Flag() string {
	return consts.FlagSeg
}

func (yesSeg) IsDefault() bool {
	return false
}

func (y yesSeg) FlagString() string {
	return y.Flag() + "=yes"
}

func (yesSeg) IsYes() bool {
	return true
}

func (yesSeg) IsNo() bool {
	return false
}

func (yesSeg) IsWLH() bool {
	return false
}

func (yesSeg) Window() int32 {
	panic("Cannot get window value from a 'yes' seg flag.")
}

func (yesSeg) Locut() float64 {
	panic("Cannot get locut value from a 'yes' seg flag.")
}

func (yesSeg) Hicut() float64 {
	panic("Cannot get hicut value from a 'yes' seg flag.")
}

type noSeg struct{}

func (noSeg) Flag() string {
	return consts.FlagSeg
}

func (noSeg) IsDefault() bool {
	return true
}

func (y noSeg) FlagString() string {
	return y.Flag() + "=no"
}

func (noSeg) IsYes() bool {
	return false
}

func (noSeg) IsNo() bool {
	return true
}

func (noSeg) IsWLH() bool {
	return false
}

func (noSeg) Window() int32 {
	panic("Cannot get window value from a 'no' seg flag.")
}

func (noSeg) Locut() float64 {
	panic("Cannot get locut value from a 'no' seg flag.")
}

func (noSeg) Hicut() float64 {
	panic("Cannot get hicut value from a 'no' seg flag.")
}

type wlhSeg struct{
	window int32
	locut float64
	hicut float64
}

func (w wlhSeg) Flag() string {
	return consts.FlagSeg
}

func (w wlhSeg) IsDefault() bool {
	return false
}

func (w wlhSeg) FlagString() string {
	return w.Flag() + "='" +
		strconv.Itoa(int(w.window)) + " " +
		strconv.FormatFloat(w.Locut(), 'f', -1, 64) + " " +
		strconv.FormatFloat(w.Hicut(), 'f', -1, 64) + "'"
}

func (w wlhSeg) IsYes() bool {
	return false
}

func (w wlhSeg) IsNo() bool {
	return false
}

func (w wlhSeg) IsWLH() bool {
	return true
}

func (w wlhSeg) Window() int32 {
	return w.window
}

func (w wlhSeg) Locut() float64 {
	return w.locut
}

func (w wlhSeg) Hicut() float64 {
	return w.hicut
}

func (w *wlhSeg) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	switch key {
	case consts.KeyWindow:
		return dec.Int32(&w.window)
	case consts.KeyLocut:
		return dec.Float64(&w.locut)
	case consts.KeyHicut:
		return dec.Float64(&w.hicut)
	default:
		return nil
	}
}

func (w *wlhSeg) NKeys() int {
	return 0
}

