package blast_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	. "github.com/veupathdb/lib-go-blast/v2/pkg/blast"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/field"
)

func TestDeltaBlast_ToCLI(t *testing.T) {
	Convey("blast.DeltaBlast.ToCLI", t, func() {
		Convey("Base Field Serialization", func() {
			tests := []struct {
				mod func(n *DeltaBlast)
				out []string
			}{
				{
					func(*DeltaBlast) {},
					[]string{ToolDeltaBlast.String()},
				},
				{
					func(n *DeltaBlast) {
						n.ShortHelp.Set(true)
					},
					[]string{ToolDeltaBlast.String(), consts.FlagShortHelp},
				},
				{
					func(n *DeltaBlast) {
						n.LongHelp.Set(true)
					},
					[]string{ToolDeltaBlast.String(), consts.FlagLongHelp},
				},
				{
					func(n *DeltaBlast) {
						n.Version.Set(true)
					},
					[]string{ToolDeltaBlast.String(), consts.FlagVersion},
				},
				{
					func(n *DeltaBlast) {
						n.QueryFile.Set("hello")
					},
					[]string{ToolDeltaBlast.String(), "-query='hello'"},
				},
				{
					func(n *DeltaBlast) {
						n.QueryLocation.Start = 1
						n.QueryLocation.Stop = 2
					},
					[]string{ToolDeltaBlast.String(), "-query_loc=1-2"},
				},
				{
					func(n *DeltaBlast) {
						n.DBFile.Set("hello")
					},
					[]string{ToolDeltaBlast.String(), "-db='hello'"},
				},
				{
					func(n *DeltaBlast) {
						n.OutFile.Set("hello")
					},
					[]string{ToolDeltaBlast.String(), "-out='hello'"},
				},
				{
					func(n *DeltaBlast) {
						n.ExpectValue.Set("hello")
					},
					[]string{ToolDeltaBlast.String(), "-evalue=hello"},
				},
				{
					func(n *DeltaBlast) {
						n.Format.Type = field.FormatTypeBlastXML
					},
					[]string{ToolDeltaBlast.String(), "-outfmt=5"},
				},
				{
					func(n *DeltaBlast) {
						n.Format.Type = field.FormatTypeCommaSeparatedValues
						n.Format.Delimiter = '@'
					},
					[]string{ToolDeltaBlast.String(), "-outfmt='10 delim=@'"},
				},
				{
					func(n *DeltaBlast) {
						n.Format.Type = field.FormatTypeCommaSeparatedValues
						n.Format.Delimiter = '@'
						n.Format.Fields = field.FormatFieldSlice{field.FormatFieldQueryAccession}
					},
					[]string{ToolDeltaBlast.String(), "-outfmt='10 delim=@ qacc'"},
				},
				{
					func(n *DeltaBlast) {
						n.ShowGIs.Set(true)
					},
					[]string{ToolDeltaBlast.String(), "-show_gis"},
				},
				{
					func(n *DeltaBlast) {
						n.NumDescriptions.Set(3)
					},
					[]string{ToolDeltaBlast.String(), "-num_descriptions=3"},
				},
				{
					func(n *DeltaBlast) {
						n.NumAlignments.Set(3)
					},
					[]string{ToolDeltaBlast.String(), "-num_alignments=3"},
				},
				{
					func(n *DeltaBlast) {
						n.LineLength.Set(3)
					},
					[]string{ToolDeltaBlast.String(), "-line_length=3"},
				},
				{
					func(n *DeltaBlast) {
						n.HTML.Set(true)
					},
					[]string{ToolDeltaBlast.String(), "-html"},
				},
				{
					func(n *DeltaBlast) {
						n.SortHits.Set(field.HitSortingTypeByBitScore)
					},
					[]string{ToolDeltaBlast.String(), "-sorthits=1"},
				},
				{
					func(n *DeltaBlast) {
						n.SortHSPs.Set(field.HSPSortingTypeByHSPPercentIdentity)
					},
					[]string{ToolDeltaBlast.String(), "-sorthsps=3"},
				},
				{
					func(n *DeltaBlast) {
						n.SoftMasking.Set(true)
					},
					[]string{ToolDeltaBlast.String(), "-soft_masking=true"},
				},
				{
					func(n *DeltaBlast) {
						n.LowercaseMasking.Set(true)
					},
					[]string{ToolDeltaBlast.String(), "-lcase_masking"},
				},
				{
					func(n *DeltaBlast) {
						n.EntrezQuery.Set("goodbye")
					},
					[]string{ToolDeltaBlast.String(), "-entrez_query='goodbye'"},
				},
				{
					func(n *DeltaBlast) {
						n.QueryCoverageHSPPercent.Set(3.333333333333)
					},
					[]string{ToolDeltaBlast.String(), "-qcov_hsp_perc=3.333333333333"},
				},
				{
					func(n *DeltaBlast) {
						n.MaxHSPs.Set(5)
					},
					[]string{ToolDeltaBlast.String(), "-max_hsps=5"},
				},
				{
					func(n *DeltaBlast) {
						n.MaxTargetSequences.Set(5)
					},
					[]string{ToolDeltaBlast.String(), "-max_target_seqs=5"},
				},
				{
					func(n *DeltaBlast) {
						n.DBSize.Set(12)
					},
					[]string{ToolDeltaBlast.String(), "-dbsize=12"},
				},
				{
					func(n *DeltaBlast) {
						n.SearchSpace.Set(16)
					},
					[]string{ToolDeltaBlast.String(), "-searchsp=16"},
				},
				{
					func(n *DeltaBlast) {
						n.ImportSearchStrategy.Set("hi")
					},
					[]string{ToolDeltaBlast.String(), "-import_search_strategy='hi'"},
				},
				{
					func(n *DeltaBlast) {
						n.ExportSearchStrategy.Set("bi")
					},
					[]string{ToolDeltaBlast.String(), "-export_search_strategy='bi'"},
				},
				{
					func(n *DeltaBlast) {
						n.ExtensionDropoffUngapped.Set(5.555)
					},
					[]string{ToolDeltaBlast.String(), "-xdrop_ungap=5.555"},
				},
				{
					func(n *DeltaBlast) {
						n.WindowSize.Set(5)
					},
					[]string{ToolDeltaBlast.String(), "-window_size=5"},
				},
				{
					func(n *DeltaBlast) {
						n.ParseDefLines.Set(true)
					},
					[]string{ToolDeltaBlast.String(), "-parse_deflines"},
				},
				{
					func(n *DeltaBlast) {
						n.Remote.Set(true)
					},
					[]string{ToolDeltaBlast.String(), "-remote"},
				},
			}

			for _, test := range tests {
				tgt := DeltaBlast{}
				test.mod(&tgt)
				So(tgt.ToCLI().Args, ShouldResemble, test.out)
			}
		})

		Convey("Base List Field Serialization", func() {
			tests := []struct {
				mod func(n *DeltaBlast)
				out []string
			}{
				{
					func(n *DeltaBlast) {
						n.GIList.Set("hope")
					},
					[]string{ToolDeltaBlast.String(), "-gilist='hope'"},
				},
				{
					func(n *DeltaBlast) {
						n.SequenceIDList.Set("hell")
					},
					[]string{ToolDeltaBlast.String(), "-seqidlist='hell'"},
				},
				{
					func(n *DeltaBlast) {
						n.NegativeGIList.Set("love")
					},
					[]string{ToolDeltaBlast.String(), "-negative_gilist='love'"},
				},
				{
					func(n *DeltaBlast) {
						n.NegativeSequenceIDList.Set("pain")
					},
					[]string{ToolDeltaBlast.String(), "-negative_seqidlist='pain'"},
				},
				{
					func(n *DeltaBlast) {
						n.TaxIDs = append(n.TaxIDs, "care", "kill")
					},
					[]string{ToolDeltaBlast.String(), "-taxids='care,kill'"},
				},
				{
					func(n *DeltaBlast) {
						n.NegativeTaxIDs = append(n.NegativeTaxIDs, "luck", "hate")
					},
					[]string{ToolDeltaBlast.String(), "-negative_taxids='luck,hate'"},
				},
				{
					func(n *DeltaBlast) {
						n.TaxIDList.Set("protect")
					},
					[]string{ToolDeltaBlast.String(), "-taxidlist='protect'"},
				},
				{
					func(n *DeltaBlast) {
						n.NegativeTaxIDList.Set("destroy")
					},
					[]string{ToolDeltaBlast.String(), "-negative_taxidlist='destroy'"},
				},
			}

			for _, test := range tests {
				tgt := DeltaBlast{}
				test.mod(&tgt)
				So(tgt.ToCLI().Args, ShouldResemble, test.out)
			}
		})

		Convey("Best Hit Field Serialization", func() {
			tests := []struct {
				mod func(n *DeltaBlast)
				out []string
			}{
				{
					func(n *DeltaBlast) {
						n.BestHitOverhang.Set(1.125)
					},
					[]string{ToolDeltaBlast.String(), "-best_hit_overhang=1.125"},
				},
				{
					func(n *DeltaBlast) {
						n.BestHitScoreEdge.Set(1.25)
					},
					[]string{ToolDeltaBlast.String(), "-best_hit_score_edge=1.25"},
				},
				{
					func(n *DeltaBlast) {
						n.SubjectBestHit.Set(true)
					},
					[]string{ToolDeltaBlast.String(), "-subject_besthit"},
				},
			}

			for _, test := range tests {
				tgt := DeltaBlast{}
				test.mod(&tgt)
				So(tgt.ToCLI().Args, ShouldResemble, test.out)
			}
		})
	})
}
