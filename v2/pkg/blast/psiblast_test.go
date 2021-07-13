package blast_test

import (
	"testing"

	"github.com/francoispqt/gojay"
	. "github.com/smartystreets/goconvey/convey"
	. "github.com/veupathdb/lib-go-blast/v2/pkg/blast"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/field"
)

func TestPSIBlast_ToCLI(t *testing.T) {
	Convey("blast.PSIBlast.ToCLI", t, func() {
		Convey("Base Field Serialization", func() {
			tests := []struct {
				mod func(n *PSIBlast)
				out []string
			}{
				{
					func(*PSIBlast) {},
					[]string{ToolPSIBlast.String()},
				},
				{
					func(n *PSIBlast) {
						n.ShortHelp.Set(true)
					},
					[]string{ToolPSIBlast.String(), consts.FlagShortHelp},
				},
				{
					func(n *PSIBlast) {
						n.LongHelp.Set(true)
					},
					[]string{ToolPSIBlast.String(), consts.FlagLongHelp},
				},
				{
					func(n *PSIBlast) {
						n.Version.Set(true)
					},
					[]string{ToolPSIBlast.String(), consts.FlagVersion},
				},
				{
					func(n *PSIBlast) {
						n.QueryFile.Set("hello")
					},
					[]string{ToolPSIBlast.String(), "-query=hello"},
				},
				{
					func(n *PSIBlast) {
						n.QueryLocation.Start = 1
						n.QueryLocation.Stop = 2
					},
					[]string{ToolPSIBlast.String(), "-query_loc=1-2"},
				},
				{
					func(n *PSIBlast) {
						n.DBFile.Set("hello")
					},
					[]string{ToolPSIBlast.String(), "-db=hello"},
				},
				{
					func(n *PSIBlast) {
						n.OutFile.Set("hello")
					},
					[]string{ToolPSIBlast.String(), "-out=hello"},
				},
				{
					func(n *PSIBlast) {
						n.ExpectValue.Set("hello")
					},
					[]string{ToolPSIBlast.String(), "-evalue=hello"},
				},
				{
					func(n *PSIBlast) {
						n.Format.Type = field.FormatTypeBlastXML
					},
					[]string{ToolPSIBlast.String(), "-outfmt=5"},
				},
				{
					func(n *PSIBlast) {
						n.Format.Type = field.FormatTypeCommaSeparatedValues
						n.Format.Delimiter = '@'
					},
					[]string{ToolPSIBlast.String(), "-outfmt=10 delim=@"},
				},
				{
					func(n *PSIBlast) {
						n.Format.Type = field.FormatTypeCommaSeparatedValues
						n.Format.Delimiter = '@'
						n.Format.Fields = field.FormatFieldSlice{field.FormatFieldQueryAccession}
					},
					[]string{ToolPSIBlast.String(), "-outfmt=10 delim=@ qacc"},
				},
				{
					func(n *PSIBlast) {
						n.ShowGIs.Set(true)
					},
					[]string{ToolPSIBlast.String(), "-show_gis"},
				},
				{
					func(n *PSIBlast) {
						n.NumDescriptions.Set(3)
					},
					[]string{ToolPSIBlast.String(), "-num_descriptions=3"},
				},
				{
					func(n *PSIBlast) {
						n.NumAlignments.Set(3)
					},
					[]string{ToolPSIBlast.String(), "-num_alignments=3"},
				},
				{
					func(n *PSIBlast) {
						n.LineLength.Set(3)
					},
					[]string{ToolPSIBlast.String(), "-line_length=3"},
				},
				{
					func(n *PSIBlast) {
						n.HTML.Set(true)
					},
					[]string{ToolPSIBlast.String(), "-html"},
				},
				{
					func(n *PSIBlast) {
						n.SortHits.Set(field.HitSortingTypeByBitScore)
					},
					[]string{ToolPSIBlast.String(), "-sorthits=1"},
				},
				{
					func(n *PSIBlast) {
						n.SortHSPs.Set(field.HSPSortingTypeByHSPPercentIdentity)
					},
					[]string{ToolPSIBlast.String(), "-sorthsps=3"},
				},
				{
					func(n *PSIBlast) {
						n.SoftMasking.Set(true)
					},
					[]string{ToolPSIBlast.String(), "-soft_masking=true"},
				},
				{
					func(n *PSIBlast) {
						n.LowercaseMasking.Set(true)
					},
					[]string{ToolPSIBlast.String(), "-lcase_masking"},
				},
				{
					func(n *PSIBlast) {
						n.EntrezQuery.Set("goodbye")
					},
					[]string{ToolPSIBlast.String(), "-entrez_query=goodbye"},
				},
				{
					func(n *PSIBlast) {
						n.QueryCoverageHSPPercent.Set(3.333333333333)
					},
					[]string{ToolPSIBlast.String(), "-qcov_hsp_perc=3.333333333333"},
				},
				{
					func(n *PSIBlast) {
						n.MaxHSPs.Set(5)
					},
					[]string{ToolPSIBlast.String(), "-max_hsps=5"},
				},
				{
					func(n *PSIBlast) {
						n.MaxTargetSequences.Set(5)
					},
					[]string{ToolPSIBlast.String(), "-max_target_seqs=5"},
				},
				{
					func(n *PSIBlast) {
						n.DBSize.Set(12)
					},
					[]string{ToolPSIBlast.String(), "-dbsize=12"},
				},
				{
					func(n *PSIBlast) {
						n.SearchSpace.Set(16)
					},
					[]string{ToolPSIBlast.String(), "-searchsp=16"},
				},
				{
					func(n *PSIBlast) {
						n.ImportSearchStrategy.Set("hi")
					},
					[]string{ToolPSIBlast.String(), "-import_search_strategy=hi"},
				},
				{
					func(n *PSIBlast) {
						n.ExportSearchStrategy.Set("bi")
					},
					[]string{ToolPSIBlast.String(), "-export_search_strategy=bi"},
				},
				{
					func(n *PSIBlast) {
						n.ExtensionDropoffUngapped.Set(5.555)
					},
					[]string{ToolPSIBlast.String(), "-xdrop_ungap=5.555"},
				},
				{
					func(n *PSIBlast) {
						n.WindowSize.Set(5)
					},
					[]string{ToolPSIBlast.String(), "-window_size=5"},
				},
				{
					func(n *PSIBlast) {
						n.ParseDefLines.Set(true)
					},
					[]string{ToolPSIBlast.String(), "-parse_deflines"},
				},
				{
					func(n *PSIBlast) {
						n.Remote.Set(true)
					},
					[]string{ToolPSIBlast.String(), "-remote"},
				},
			}

			for _, test := range tests {
				tgt := PSIBlast{}
				test.mod(&tgt)
				So(tgt.ToCLI().Args, ShouldResemble, test.out)
			}
		})

		Convey("IPG List Field Serialization", func() {
			tests := []struct {
				mod func(n *PSIBlast)
				out []string
			}{
				{
					func(n *PSIBlast) {
						n.GIList.Set("hope")
					},
					[]string{ToolPSIBlast.String(), "-gilist=hope"},
				},
				{
					func(n *PSIBlast) {
						n.SequenceIDList.Set("hell")
					},
					[]string{ToolPSIBlast.String(), "-seqidlist=hell"},
				},
				{
					func(n *PSIBlast) {
						n.NegativeGIList.Set("love")
					},
					[]string{ToolPSIBlast.String(), "-negative_gilist=love"},
				},
				{
					func(n *PSIBlast) {
						n.NegativeSequenceIDList.Set("pain")
					},
					[]string{ToolPSIBlast.String(), "-negative_seqidlist=pain"},
				},
				{
					func(n *PSIBlast) {
						n.TaxIDs = append(n.TaxIDs, "care", "kill")
					},
					[]string{ToolPSIBlast.String(), "-taxids=care,kill"},
				},
				{
					func(n *PSIBlast) {
						n.NegativeTaxIDs = append(n.NegativeTaxIDs, "luck", "hate")
					},
					[]string{ToolPSIBlast.String(), "-negative_taxids=luck,hate"},
				},
				{
					func(n *PSIBlast) {
						n.TaxIDList.Set("protect")
					},
					[]string{ToolPSIBlast.String(), "-taxidlist=protect"},
				},
				{
					func(n *PSIBlast) {
						n.NegativeTaxIDList.Set("destroy")
					},
					[]string{ToolPSIBlast.String(), "-negative_taxidlist=destroy"},
				},
				{
					func(n *PSIBlast) {
						n.IPGList.Set("build")
					},
					[]string{ToolPSIBlast.String(), "-ipglist=build"},
				},
				{
					func(n *PSIBlast) {
						n.NegativeIPGList.Set("break")
					},
					[]string{ToolPSIBlast.String(), "-negative_ipglist=break"},
				},
			}

			for _, test := range tests {
				tgt := PSIBlast{}
				test.mod(&tgt)
				So(tgt.ToCLI().Args, ShouldResemble, test.out)
			}
		})

		Convey("Best Hit Field Serialization", func() {
			tests := []struct {
				mod func(n *PSIBlast)
				out []string
			}{
				{
					func(n *PSIBlast) {
						n.BestHitOverhang.Set(1.125)
					},
					[]string{ToolPSIBlast.String(), "-best_hit_overhang=1.125"},
				},
				{
					func(n *PSIBlast) {
						n.BestHitScoreEdge.Set(1.25)
					},
					[]string{ToolPSIBlast.String(), "-best_hit_score_edge=1.25"},
				},
				{
					func(n *PSIBlast) {
						n.SubjectBestHit.Set(true)
					},
					[]string{ToolPSIBlast.String(), "-subject_besthit"},
				},
			}

			for _, test := range tests {
				tgt := PSIBlast{}
				test.mod(&tgt)
				So(tgt.ToCLI().Args, ShouldResemble, test.out)
			}
		})
	})
}

func TestPSIBlast_UnmarshalJSONObject(t *testing.T) {
	Convey("blast.PSIBlast.UnmarshalJSONObject", t, func() {
		// Regression coverage (Seg deserialization failed)
		Convey("Seg Deserialization", func() {
			Convey("Yes Seg", func() {
				json := []byte(`{"-seg":"yes"}`)

				test := new(PSIBlast)

				err := gojay.UnmarshalJSONObject(json, test)

				So(err, ShouldBeNil)
				So(test.Seg.IsYes(), ShouldBeTrue)
			})

			Convey("No Seg", func() {
				json := []byte(`{"-seg":"no"}`)

				test := new(PSIBlast)

				err := gojay.UnmarshalJSONObject(json, test)

				So(err, ShouldBeNil)
				So(test.Seg.IsNo(), ShouldBeTrue)
			})

			Convey("Value Seg", func() {
				json := []byte(`{"-seg":{"window":1,"locut":2.2,"hicut":3.3}}`)

				test := new(PSIBlast)

				err := gojay.UnmarshalJSONObject(json, test)

				So(err, ShouldBeNil)
				So(test.Seg.Window(), ShouldEqual, 1)
				So(test.Seg.Locut(), ShouldEqual, 2.2)
				So(test.Seg.Hicut(), ShouldEqual, 3.3)
			})
		})
	})
}
