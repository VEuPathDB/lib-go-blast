package blast_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	. "github.com/veupathdb/lib-go-blast/v1/pkg/blast"
	"github.com/veupathdb/lib-go-blast/v1/pkg/blast/consts"
	"github.com/veupathdb/lib-go-blast/v1/pkg/blast/field"
)

func TestTBlastN_ToCLI(t *testing.T) {
	Convey("blast.TBlastN.ToCLI", t, func() {
		Convey("Base Field Serialization", func() {
			tests := []struct {
				mod func(n *TBlastN)
				out []string
			}{
				{
					func(*TBlastN) {},
					[]string{ToolTBlastN.String()},
				},
				{
					func(n *TBlastN) {
						n.ShortHelp.Set(true)
					},
					[]string{ToolTBlastN.String(), consts.FlagShortHelp},
				},
				{
					func(n *TBlastN) {
						n.LongHelp.Set(true)
					},
					[]string{ToolTBlastN.String(), consts.FlagLongHelp},
				},
				{
					func(n *TBlastN) {
						n.Version.Set(true)
					},
					[]string{ToolTBlastN.String(), consts.FlagVersion},
				},
				{
					func(n *TBlastN) {
						n.QueryFile.Set("hello")
					},
					[]string{ToolTBlastN.String(), "-query='hello'"},
				},
				{
					func(n *TBlastN) {
						n.QueryLocation.Start = 1
						n.QueryLocation.Stop = 2
					},
					[]string{ToolTBlastN.String(), "-query_loc=1-2"},
				},
				{
					func(n *TBlastN) {
						n.DBFile.Set("hello")
					},
					[]string{ToolTBlastN.String(), "-db='hello'"},
				},
				{
					func(n *TBlastN) {
						n.OutFile.Set("hello")
					},
					[]string{ToolTBlastN.String(), "-out='hello'"},
				},
				{
					func(n *TBlastN) {
						n.ExpectValue.Set("hello")
					},
					[]string{ToolTBlastN.String(), "-evalue=hello"},
				},
				{
					func(n *TBlastN) {
						n.OutFormat.Type = field.FormatTypeBlastXML
					},
					[]string{ToolTBlastN.String(), "-outfmt=5"},
				},
				{
					func(n *TBlastN) {
						n.OutFormat.Type = field.FormatTypeCommaSeparatedValues
						n.OutFormat.Delimiter = '@'
					},
					[]string{ToolTBlastN.String(), "-outfmt='10 delim=@'"},
				},
				{
					func(n *TBlastN) {
						n.OutFormat.Type = field.FormatTypeCommaSeparatedValues
						n.OutFormat.Delimiter = '@'
						n.OutFormat.Fields = field.FormatFieldSlice{field.FormatFieldQueryAccession}
					},
					[]string{ToolTBlastN.String(), "-outfmt='10 delim=@ qacc'"},
				},
				{
					func(n *TBlastN) {
						n.ShowGIs.Set(true)
					},
					[]string{ToolTBlastN.String(), "-show_gis"},
				},
				{
					func(n *TBlastN) {
						n.NumDescriptions.Set(3)
					},
					[]string{ToolTBlastN.String(), "-num_descriptions=3"},
				},
				{
					func(n *TBlastN) {
						n.NumAlignments.Set(3)
					},
					[]string{ToolTBlastN.String(), "-num_alignments=3"},
				},
				{
					func(n *TBlastN) {
						n.LineLength.Set(3)
					},
					[]string{ToolTBlastN.String(), "-line_length=3"},
				},
				{
					func(n *TBlastN) {
						n.HTML.Set(true)
					},
					[]string{ToolTBlastN.String(), "-html"},
				},
				{
					func(n *TBlastN) {
						n.SortHits.Set(field.HitSortingTypeByBitScore)
					},
					[]string{ToolTBlastN.String(), "-sorthits=1"},
				},
				{
					func(n *TBlastN) {
						n.SortHSPs.Set(field.HSPSortingTypeByHSPPercentIdentity)
					},
					[]string{ToolTBlastN.String(), "-sorthsps=3"},
				},
				{
					func(n *TBlastN) {
						n.SoftMasking.Set(true)
					},
					[]string{ToolTBlastN.String(), "-soft_masking=true"},
				},
				{
					func(n *TBlastN) {
						n.LowercaseMasking.Set(true)
					},
					[]string{ToolTBlastN.String(), "-lcase_masking"},
				},
				{
					func(n *TBlastN) {
						n.EntrezQuery.Set("goodbye")
					},
					[]string{ToolTBlastN.String(), "-entrez_query='goodbye'"},
				},
				{
					func(n *TBlastN) {
						n.QueryCoverageHSPPercent.Set(3.333333333333)
					},
					[]string{ToolTBlastN.String(), "-qcov_hsp_perc=3.333333333333"},
				},
				{
					func(n *TBlastN) {
						n.MaxHSPs.Set(5)
					},
					[]string{ToolTBlastN.String(), "-max_hsps=5"},
				},
				{
					func(n *TBlastN) {
						n.MaxTargetSequences.Set(5)
					},
					[]string{ToolTBlastN.String(), "-max_target_seqs=5"},
				},
				{
					func(n *TBlastN) {
						n.DBSize.Set(12)
					},
					[]string{ToolTBlastN.String(), "-dbsize=12"},
				},
				{
					func(n *TBlastN) {
						n.SearchSpace.Set(16)
					},
					[]string{ToolTBlastN.String(), "-searchsp=16"},
				},
				{
					func(n *TBlastN) {
						n.ImportSearchStrategy.Set("hi")
					},
					[]string{ToolTBlastN.String(), "-import_search_strategy='hi'"},
				},
				{
					func(n *TBlastN) {
						n.ExportSearchStrategy.Set("bi")
					},
					[]string{ToolTBlastN.String(), "-export_search_strategy='bi'"},
				},
				{
					func(n *TBlastN) {
						n.ExtensionDropoffUngapped.Set(5.555)
					},
					[]string{ToolTBlastN.String(), "-xdrop_ungap=5.555"},
				},
				{
					func(n *TBlastN) {
						n.WindowSize.Set(5)
					},
					[]string{ToolTBlastN.String(), "-window_size=5"},
				},
				{
					func(n *TBlastN) {
						n.ParseDefLines.Set(true)
					},
					[]string{ToolTBlastN.String(), "-parse_deflines"},
				},
				{
					func(n *TBlastN) {
						n.Remote.Set(true)
					},
					[]string{ToolTBlastN.String(), "-remote"},
				},
			}

			for _, test := range tests {
				tgt := TBlastN{}
				test.mod(&tgt)
				So(tgt.ToCLI().Args, ShouldResemble, test.out)
			}
		})

		Convey("Base List Field Serialization", func() {
			tests := []struct {
				mod func(n *TBlastN)
				out []string
			}{
				{
					func(n *TBlastN) {
						n.GIList.Set("hope")
					},
					[]string{ToolTBlastN.String(), "-gilist='hope'"},
				},
				{
					func(n *TBlastN) {
						n.SequenceIDList.Set("hell")
					},
					[]string{ToolTBlastN.String(), "-seqidlist='hell'"},
				},
				{
					func(n *TBlastN) {
						n.NegativeGIList.Set("love")
					},
					[]string{ToolTBlastN.String(), "-negative_gilist='love'"},
				},
				{
					func(n *TBlastN) {
						n.NegativeSequenceIDList.Set("pain")
					},
					[]string{ToolTBlastN.String(), "-negative_seqidlist='pain'"},
				},
				{
					func(n *TBlastN) {
						n.TaxIDs = append(n.TaxIDs, "care", "kill")
					},
					[]string{ToolTBlastN.String(), "-taxids='care,kill'"},
				},
				{
					func(n *TBlastN) {
						n.NegativeTaxIDs = append(n.NegativeTaxIDs, "luck", "hate")
					},
					[]string{ToolTBlastN.String(), "-negative_taxids='luck,hate'"},
				},
				{
					func(n *TBlastN) {
						n.TaxIDList.Set("protect")
					},
					[]string{ToolTBlastN.String(), "-taxidlist='protect'"},
				},
				{
					func(n *TBlastN) {
						n.NegativeTaxIDList.Set("destroy")
					},
					[]string{ToolTBlastN.String(), "-negative_taxidlist='destroy'"},
				},
			}

			for _, test := range tests {
				tgt := TBlastN{}
				test.mod(&tgt)
				So(tgt.ToCLI().Args, ShouldResemble, test.out)
			}
		})

		Convey("Best Hit Field Serialization", func() {
			tests := []struct {
				mod func(n *TBlastN)
				out []string
			}{
				{
					func(n *TBlastN) {
						n.BestHitOverhang.Set(1.125)
					},
					[]string{ToolTBlastN.String(), "-best_hit_overhang=1.125"},
				},
				{
					func(n *TBlastN) {
						n.BestHitScoreEdge.Set(1.25)
					},
					[]string{ToolTBlastN.String(), "-best_hit_score_edge=1.25"},
				},
				{
					func(n *TBlastN) {
						n.SubjectBestHit.Set(true)
					},
					[]string{ToolTBlastN.String(), "-subject_besthit"},
				},
			}

			for _, test := range tests {
				tgt := TBlastN{}
				test.mod(&tgt)
				So(tgt.ToCLI().Args, ShouldResemble, test.out)
			}
		})
	})
}
