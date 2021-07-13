package blast_test

import (
	"testing"

	"github.com/francoispqt/gojay"
	. "github.com/smartystreets/goconvey/convey"
	. "github.com/veupathdb/lib-go-blast/v2/pkg/blast"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/field"
)

func TestTBlastX_ToCLI(t *testing.T) {
	Convey("blast.TBlastX.ToCLI", t, func() {
		Convey("Base Field Serialization", func() {
			tests := []struct {
				mod func(n *TBlastX)
				out []string
			}{
				{
					func(*TBlastX) {},
					[]string{ToolTBlastX.String()},
				},
				{
					func(n *TBlastX) {
						n.ShortHelp.Set(true)
					},
					[]string{ToolTBlastX.String(), consts.FlagShortHelp},
				},
				{
					func(n *TBlastX) {
						n.LongHelp.Set(true)
					},
					[]string{ToolTBlastX.String(), consts.FlagLongHelp},
				},
				{
					func(n *TBlastX) {
						n.Version.Set(true)
					},
					[]string{ToolTBlastX.String(), consts.FlagVersion},
				},
				{
					func(n *TBlastX) {
						n.QueryFile.Set("hello")
					},
					[]string{ToolTBlastX.String(), "-query=hello"},
				},
				{
					func(n *TBlastX) {
						n.QueryLocation.Start = 1
						n.QueryLocation.Stop = 2
					},
					[]string{ToolTBlastX.String(), "-query_loc=1-2"},
				},
				{
					func(n *TBlastX) {
						n.DBFile.Set("hello")
					},
					[]string{ToolTBlastX.String(), "-db=hello"},
				},
				{
					func(n *TBlastX) {
						n.OutFile.Set("hello")
					},
					[]string{ToolTBlastX.String(), "-out=hello"},
				},
				{
					func(n *TBlastX) {
						n.ExpectValue.Set("hello")
					},
					[]string{ToolTBlastX.String(), "-evalue=hello"},
				},
				{
					func(n *TBlastX) {
						n.Format.Type = field.FormatTypeBlastXML
					},
					[]string{ToolTBlastX.String(), "-outfmt=5"},
				},
				{
					func(n *TBlastX) {
						n.Format.Type = field.FormatTypeCommaSeparatedValues
						n.Format.Delimiter = '@'
					},
					[]string{ToolTBlastX.String(), "-outfmt=10 delim=@"},
				},
				{
					func(n *TBlastX) {
						n.Format.Type = field.FormatTypeCommaSeparatedValues
						n.Format.Delimiter = '@'
						n.Format.Fields = field.FormatFieldSlice{field.FormatFieldQueryAccession}
					},
					[]string{ToolTBlastX.String(), "-outfmt=10 delim=@ qacc"},
				},
				{
					func(n *TBlastX) {
						n.ShowGIs.Set(true)
					},
					[]string{ToolTBlastX.String(), "-show_gis"},
				},
				{
					func(n *TBlastX) {
						n.NumDescriptions.Set(3)
					},
					[]string{ToolTBlastX.String(), "-num_descriptions=3"},
				},
				{
					func(n *TBlastX) {
						n.NumAlignments.Set(3)
					},
					[]string{ToolTBlastX.String(), "-num_alignments=3"},
				},
				{
					func(n *TBlastX) {
						n.LineLength.Set(3)
					},
					[]string{ToolTBlastX.String(), "-line_length=3"},
				},
				{
					func(n *TBlastX) {
						n.HTML.Set(true)
					},
					[]string{ToolTBlastX.String(), "-html"},
				},
				{
					func(n *TBlastX) {
						n.SortHits.Set(field.HitSortingTypeByBitScore)
					},
					[]string{ToolTBlastX.String(), "-sorthits=1"},
				},
				{
					func(n *TBlastX) {
						n.SortHSPs.Set(field.HSPSortingTypeByHSPPercentIdentity)
					},
					[]string{ToolTBlastX.String(), "-sorthsps=3"},
				},
				{
					func(n *TBlastX) {
						n.SoftMasking.Set(true)
					},
					[]string{ToolTBlastX.String(), "-soft_masking=true"},
				},
				{
					func(n *TBlastX) {
						n.LowercaseMasking.Set(true)
					},
					[]string{ToolTBlastX.String(), "-lcase_masking"},
				},
				{
					func(n *TBlastX) {
						n.EntrezQuery.Set("goodbye")
					},
					[]string{ToolTBlastX.String(), "-entrez_query=goodbye"},
				},
				{
					func(n *TBlastX) {
						n.QueryCoverageHSPPercent.Set(3.333333333333)
					},
					[]string{ToolTBlastX.String(), "-qcov_hsp_perc=3.333333333333"},
				},
				{
					func(n *TBlastX) {
						n.MaxHSPs.Set(5)
					},
					[]string{ToolTBlastX.String(), "-max_hsps=5"},
				},
				{
					func(n *TBlastX) {
						n.MaxTargetSequences.Set(5)
					},
					[]string{ToolTBlastX.String(), "-max_target_seqs=5"},
				},
				{
					func(n *TBlastX) {
						n.DBSize.Set(12)
					},
					[]string{ToolTBlastX.String(), "-dbsize=12"},
				},
				{
					func(n *TBlastX) {
						n.SearchSpace.Set(16)
					},
					[]string{ToolTBlastX.String(), "-searchsp=16"},
				},
				{
					func(n *TBlastX) {
						n.ImportSearchStrategy.Set("hi")
					},
					[]string{ToolTBlastX.String(), "-import_search_strategy=hi"},
				},
				{
					func(n *TBlastX) {
						n.ExportSearchStrategy.Set("bi")
					},
					[]string{ToolTBlastX.String(), "-export_search_strategy=bi"},
				},
				{
					func(n *TBlastX) {
						n.ExtensionDropoffUngapped.Set(5.555)
					},
					[]string{ToolTBlastX.String(), "-xdrop_ungap=5.555"},
				},
				{
					func(n *TBlastX) {
						n.WindowSize.Set(5)
					},
					[]string{ToolTBlastX.String(), "-window_size=5"},
				},
				{
					func(n *TBlastX) {
						n.ParseDefLines.Set(true)
					},
					[]string{ToolTBlastX.String(), "-parse_deflines"},
				},
				{
					func(n *TBlastX) {
						n.Remote.Set(true)
					},
					[]string{ToolTBlastX.String(), "-remote"},
				},
			}

			for _, test := range tests {
				tgt := TBlastX{}
				test.mod(&tgt)
				So(tgt.ToCLI().Args, ShouldResemble, test.out)
			}
		})

		Convey("Base List Field Serialization", func() {
			tests := []struct {
				mod func(n *TBlastX)
				out []string
			}{
				{
					func(n *TBlastX) {
						n.GIList.Set("hope")
					},
					[]string{ToolTBlastX.String(), "-gilist=hope"},
				},
				{
					func(n *TBlastX) {
						n.SequenceIDList.Set("hell")
					},
					[]string{ToolTBlastX.String(), "-seqidlist=hell"},
				},
				{
					func(n *TBlastX) {
						n.NegativeGIList.Set("love")
					},
					[]string{ToolTBlastX.String(), "-negative_gilist=love"},
				},
				{
					func(n *TBlastX) {
						n.NegativeSequenceIDList.Set("pain")
					},
					[]string{ToolTBlastX.String(), "-negative_seqidlist=pain"},
				},
				{
					func(n *TBlastX) {
						n.TaxIDs = append(n.TaxIDs, "care", "kill")
					},
					[]string{ToolTBlastX.String(), "-taxids=care,kill"},
				},
				{
					func(n *TBlastX) {
						n.NegativeTaxIDs = append(n.NegativeTaxIDs, "luck", "hate")
					},
					[]string{ToolTBlastX.String(), "-negative_taxids=luck,hate"},
				},
				{
					func(n *TBlastX) {
						n.TaxIDList.Set("protect")
					},
					[]string{ToolTBlastX.String(), "-taxidlist=protect"},
				},
				{
					func(n *TBlastX) {
						n.NegativeTaxIDList.Set("destroy")
					},
					[]string{ToolTBlastX.String(), "-negative_taxidlist=destroy"},
				},
			}

			for _, test := range tests {
				tgt := TBlastX{}
				test.mod(&tgt)
				So(tgt.ToCLI().Args, ShouldResemble, test.out)
			}
		})

		Convey("Best Hit Field Serialization", func() {
			tests := []struct {
				mod func(n *TBlastX)
				out []string
			}{
				{
					func(n *TBlastX) {
						n.BestHitOverhang.Set(1.125)
					},
					[]string{ToolTBlastX.String(), "-best_hit_overhang=1.125"},
				},
				{
					func(n *TBlastX) {
						n.BestHitScoreEdge.Set(1.25)
					},
					[]string{ToolTBlastX.String(), "-best_hit_score_edge=1.25"},
				},
				{
					func(n *TBlastX) {
						n.SubjectBestHit.Set(true)
					},
					[]string{ToolTBlastX.String(), "-subject_besthit"},
				},
			}

			for _, test := range tests {
				tgt := TBlastX{}
				test.mod(&tgt)
				So(tgt.ToCLI().Args, ShouldResemble, test.out)
			}
		})
	})
}

func TestTBlastX_UnmarshalJSONObject(t *testing.T) {
	Convey("blast.TBlastX.UnmarshalJSONObject", t, func() {
		// Regression coverage (Seg deserialization failed)
		Convey("Seg Deserialization", func() {
			Convey("Yes Seg", func() {
				json := []byte(`{"-seg":"yes"}`)

				test := new(TBlastX)

				err := gojay.UnmarshalJSONObject(json, test)

				So(err, ShouldBeNil)
				So(test.Seg.IsYes(), ShouldBeTrue)
			})

			Convey("No Seg", func() {
				json := []byte(`{"-seg":"no"}`)

				test := new(TBlastX)

				err := gojay.UnmarshalJSONObject(json, test)

				So(err, ShouldBeNil)
				So(test.Seg.IsNo(), ShouldBeTrue)
			})

			Convey("Value Seg", func() {
				json := []byte(`{"-seg":{"window":1,"locut":2.2,"hicut":3.3}}`)

				test := new(TBlastX)

				err := gojay.UnmarshalJSONObject(json, test)

				So(err, ShouldBeNil)
				So(test.Seg.Window(), ShouldEqual, 1)
				So(test.Seg.Locut(), ShouldEqual, 2.2)
				So(test.Seg.Hicut(), ShouldEqual, 3.3)
			})
		})
	})
}
