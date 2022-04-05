package field_test

import (
	"strings"
	"testing"

	"github.com/francoispqt/gojay"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/field"
)

func TestDecodeJSONPenalty(t *testing.T) {
	Convey("DecodeJSONPenalty()", t, func() {

		Convey("when given a non-integral json input", func() {

			Convey("returns an error", func() {
				inp := `"hello"`
				dec := gojay.BorrowDecoder(strings.NewReader(inp))
				defer dec.Release()
				tgt := new(field.Penalty)
				err := field.DecodeJSONPenalty(dec, tgt)

				ShouldNotBeNil(err)
			})
		})

		Convey("when given an int json input", func() {

			Convey("returns a Penalty instance wrapping the parsed value", func() {

				inp := "21"
				dec := gojay.BorrowDecoder(strings.NewReader(inp))
				defer dec.Release()
				tgt := new(field.Penalty)
				err := field.DecodeJSONPenalty(dec, tgt)

				ShouldBeNil(err)
				ShouldEqual(21, tgt.Get())
			})
		})
	})
}
