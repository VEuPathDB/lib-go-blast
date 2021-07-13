package field_test

import (
	"bytes"
	"testing"

	"github.com/francoispqt/gojay"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/field"
)

func TestNewYesDust(t *testing.T) {
	Convey("NewYesDust", t, func() {
		test := field.NewYesDust()

		So(test.IsYes(), ShouldBeTrue)
		So(test.IsNo(), ShouldBeFalse)
		So(test.IsLWL(), ShouldBeFalse)

		So(func() {test.Level()}, ShouldPanic)
		So(func() {test.Window()}, ShouldPanic)
		So(func() {test.Linker()}, ShouldPanic)

		So(test.Flag(), ShouldEqual, consts.FlagDust)
		So(test.IsDefault(), ShouldBeFalse)
		So(test.FlagString(), ShouldEqual, consts.FlagDust + "=yes")
	})
}

func TestNewNoDust(t *testing.T) {
	Convey("NewNoDust", t, func() {
		test := field.NewNoDust()

		So(test.IsYes(), ShouldBeFalse)
		So(test.IsNo(), ShouldBeTrue)
		So(test.IsLWL(), ShouldBeFalse)

		So(func() {test.Level()}, ShouldPanic)
		So(func() {test.Window()}, ShouldPanic)
		So(func() {test.Linker()}, ShouldPanic)

		So(test.Flag(), ShouldEqual, consts.FlagDust)
		So(test.IsDefault(), ShouldBeFalse)
		So(test.FlagString(), ShouldEqual, consts.FlagDust + "=no")
	})
}

func TestNewLWLDust(t *testing.T) {
	Convey("NewLWLDust", t, func() {
		test := field.NewLWLDust(10, 12, 14)

		So(test.IsYes(), ShouldBeFalse)
		So(test.IsNo(), ShouldBeFalse)
		So(test.IsLWL(), ShouldBeTrue)

		So(test.Level(), ShouldEqual, 10)
		So(test.Window(), ShouldEqual, 12)
		So(test.Linker(), ShouldEqual, 14)

		So(test.Flag(), ShouldEqual, consts.FlagDust)
		So(test.IsDefault(), ShouldBeFalse)
		So(test.FlagString(), ShouldEqual, consts.FlagDust + "=10 12 14")
	})
}

func TestDecodeJSONDust(t *testing.T) {
	Convey("DecodeJSONDust", t, func() {
		Convey("Yes JSON", func() {
			dec := gojay.NewDecoder(bytes.NewBufferString(`"yes"`))

			val, err := field.DecodeJSONDust(dec)

			So(err, ShouldBeNil)

			So(val.IsYes(), ShouldBeTrue)
		})

		Convey("No JSON", func() {
			dec := gojay.NewDecoder(bytes.NewBufferString(`"no"`))

			val, err := field.DecodeJSONDust(dec)

			So(err, ShouldBeNil)

			So(val.IsNo(), ShouldBeTrue)
		})

		Convey("LWL JSON", func() {
			buf := bytes.NewBuffer([]byte(`{"level": 10,"window":12,"linker":14}`))
			dec := gojay.NewDecoder(buf)

			val, err := field.DecodeJSONDust(dec)

			So(err, ShouldBeNil)

			So(val.IsLWL(), ShouldBeTrue)
			So(val.Level(), ShouldEqual, 10)
			So(val.Window(), ShouldEqual, 12)
			So(val.Linker(), ShouldEqual, 14)
		})
	})
}