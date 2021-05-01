package bval_test

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/bval"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/field"
)

func TestErrorMap_InvalidStrEnum(t *testing.T) {
	Convey("bval.ValidationBuilder.InvalidStrEnum", t, func() {
		tgt := bval.ValidationBuilder{}

		tgt.InvalidStrEnum("hi", "ho", "hi", "ho")

		So(len(tgt), ShouldEqual, 1)

		_, ok := tgt["hi"]

		So(ok, ShouldBeTrue)
		So(len(tgt["hi"]), ShouldEqual, 1)

		So(tgt["hi"][0], ShouldEqual, "Invalid hi value \"ho\".  Must be one of [hi, ho].")
	})
}

func TestErrorMap_InvalidUint8Enum(t *testing.T) {
	Convey("bval.ValidationBuilder.InvalidUint8Enum", t, func() {
		tgt := bval.ValidationBuilder{}

		tgt.InvalidUint8Enum("hi", 1, 2, 3, 4, 5)

		So(len(tgt), ShouldEqual, 1)

		_, ok := tgt["hi"]

		So(ok, ShouldBeTrue)
		So(len(tgt["hi"]), ShouldEqual, 1)

		So(tgt["hi"][0], ShouldEqual, "Invalid hi value \"1\".  Must be one of [2, 3, 4, 5].")
	})
}

func TestErrorMap_Incompatible(t *testing.T) {
	Convey("bval.ValidationBuilder.Incompatible", t, func() {
		tgt := bval.ValidationBuilder{}
		in1 := field.IPGList{}
		in2 := field.NegativeTaxIDList{}

		in1.Set("goodbye")
		in2.Set("hello")

		_ = tgt.Incompatible(&in1, &in2)

		So(len(tgt), ShouldEqual, 2)

		_, ok1 := tgt[in1.Flag()]
		So(ok1, ShouldBeTrue)

		_, ok2 := tgt[in2.Flag()]
		So(ok2, ShouldBeTrue)

		So(len(tgt[in1.Flag()]), ShouldEqual, 1)
		So(len(tgt[in2.Flag()]), ShouldEqual, 1)

		So(tgt[in1.Flag()][0], ShouldEqual, fmt.Sprintf(bval.ErrIncompatibleWith, in1.Flag(), in2.Flag()))
		So(tgt[in2.Flag()][0], ShouldEqual, fmt.Sprintf(bval.ErrIncompatibleWith, in2.Flag(), in1.Flag()))
	})
}

func TestErrorMap_IncompatibleList(t *testing.T) {
	Convey("bval.ValidationBuilder.IncompatibleList", t, func() {
		tgt := bval.ValidationBuilder{}
		in1 := field.IPGList{}
		in2 := field.NegativeTaxIDList{}
		in3 := field.SequenceIDList{}

		in1.Set("goodbye")
		in2.Set("hello")
		in3.Set("willow")

		_ = tgt.IncompatibleList(&in1, &in2, &in3)

		So(len(tgt), ShouldEqual, 3)

		_, ok1 := tgt[in1.Flag()]
		So(ok1, ShouldBeTrue)

		_, ok2 := tgt[in2.Flag()]
		So(ok2, ShouldBeTrue)

		_, ok3 := tgt[in3.Flag()]
		So(ok3, ShouldBeTrue)

		So(len(tgt[in1.Flag()]), ShouldEqual, 2)
		So(len(tgt[in2.Flag()]), ShouldEqual, 2)
		So(len(tgt[in3.Flag()]), ShouldEqual, 2)

		So(tgt[in1.Flag()][0], ShouldEqual, fmt.Sprintf(bval.ErrIncompatibleWith, in1.Flag(), in2.Flag()))
		So(tgt[in1.Flag()][1], ShouldEqual, fmt.Sprintf(bval.ErrIncompatibleWith, in1.Flag(), in3.Flag()))
		So(tgt[in2.Flag()][0], ShouldEqual, fmt.Sprintf(bval.ErrIncompatibleWith, in2.Flag(), in1.Flag()))
		So(tgt[in2.Flag()][1], ShouldEqual, fmt.Sprintf(bval.ErrIncompatibleWith, in2.Flag(), in3.Flag()))
		So(tgt[in3.Flag()][0], ShouldEqual, fmt.Sprintf(bval.ErrIncompatibleWith, in3.Flag(), in1.Flag()))
		So(tgt[in3.Flag()][1], ShouldEqual, fmt.Sprintf(bval.ErrIncompatibleWith, in3.Flag(), in2.Flag()))
	})
}