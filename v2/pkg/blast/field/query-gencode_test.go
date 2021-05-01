package field_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/bval"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/field"
)

func TestQueryGenCode_Validate(t *testing.T) {
	Convey("field.QueryGenCode.Validate", t, func() {
		Convey("invalid values", func() {
			tests := []uint8{0, 7, 8, 17, 18, 19, 20, 32, 34}

			for _, v := range tests {
				err := make(bval.ValidationBuilder)
				tgt := field.NewQueryGenCode(v)

				tgt.Validate(err)

				So(len(err), ShouldEqual, 1)

				_, ok := err[consts.FlagQueryGenCode]
				So(ok, ShouldBeTrue)

				So(len(err[consts.FlagQueryGenCode]), ShouldEqual, 1)
			}
		})

		Convey("valid values", func() {
			tests := []uint8{1, 2, 3, 4, 5, 6, 9, 10, 11, 12, 13, 14, 15, 16, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 33}

			for _, v := range tests {
				err := make(bval.ValidationBuilder)
				tgt := field.NewQueryGenCode(v)

				tgt.Validate(err)

				So(len(err), ShouldEqual, 0)
			}
		})
	})
}
