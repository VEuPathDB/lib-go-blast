package blast_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	. "github.com/veupathdb/lib-go-blast/v2/pkg/blast"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/field"
)

func TestBlastX_ToCLI(t *testing.T) {
	Convey("blast.BlastX.ToCLI", t, func() {
		Convey("Base Field Serialization", func() {
			tests := []struct {
				mod func(n *BlastX)
				out []string
			}{
				{
					func(*BlastX) {},
					[]string{ToolBlastX.String()},
				},
				{
					func(n *BlastX) {
						n.ShortHelp.Set(true)
					},
					[]string{ToolBlastX.String(), consts.FlagShortHelp},
				},
				{
					func(n *BlastX) {
						n.LongHelp.Set(true)
					},
					[]string{ToolBlastX.String(), consts.FlagLongHelp},
				},
				{
					func(n *BlastX) {
						n.Version.Set(true)
					},
					[]string{ToolBlastX.String(), consts.FlagVersion},
				},
				{
					func(n *BlastX) {
						n.QueryFile.Set("hello")
					},
					[]string{ToolBlastX.String(), "-query='hello'"},
				},
				{
					func(n *BlastX) {
						n.QueryLocation.Start = 1
						n.QueryLocation.Stop = 2
					},
					[]string{ToolBlastX.String(), "-query_loc=1-2"},
				},
				{
					func(n *BlastX) {
						n.DBFile.Set("hello")
					},
					[]string{ToolBlastX.String(), "-db='hello'"},
				},
				{
					func(n *BlastX) {
						n.OutFile.Set("hello")
					},
					[]string{ToolBlastX.String(), "-out='hello'"},
				},
				{
					func(n *BlastX) {
						n.ExpectValue.Set("hello")
					},
					[]string{ToolBlastX.String(), "-evalue=hello"},
				},
				{
					func(n *BlastX) {
						n.Format.Type = field.FormatTypeBlastXML
					},
					[]string{ToolBlastX.String(), "-outfmt=5"},
				},
				{
					func(n *BlastX) {
						n.Format.Type = field.FormatTypeCommaSeparatedValues
						n.Format.Delimiter = '@'
					},
					[]string{ToolBlastX.String(), "-outfmt='10 delim=@'"},
				},
				{
					func(n *BlastX) {
						n.Format.Type = field.FormatTypeCommaSeparatedValues
						n.Format.Delimiter = '@'
						n.Format.Fields = field.FormatFieldSlice{field.FormatFieldQueryAccession}
					},
					[]string{ToolBlastX.String(), "-outfmt='10 delim=@ qacc'"},
				},
				{
					func(n *BlastX) {
						n.ShowGIs.Set(true)
					},
					[]string{ToolBlastX.String(), "-show_gis"},
				},
				{
					func(n *BlastX) {
						n.NumDescriptions.Set(3)
					},
					[]string{ToolBlastX.String(), "-num_descriptions=3"},
				},
				{
					func(n *BlastX) {
						n.NumAlignments.Set(3)
					},
					[]string{ToolBlastX.String(), "-num_alignments=3"},
				},
				{
					func(n *BlastX) {
						n.LineLength.Set(3)
					},
					[]string{ToolBlastX.String(), "-line_length=3"},
				},
				{
					func(n *BlastX) {
						n.HTML.Set(true)
					},
					[]string{ToolBlastX.String(), "-html"},
				},
				{
					func(n *BlastX) {
						n.SortHits.Set(field.HitSortingTypeByBitScore)
					},
					[]string{ToolBlastX.String(), "-sorthits=1"},
				},
				{
					func(n *BlastX) {
						n.SortHSPs.Set(field.HSPSortingTypeByHSPPercentIdentity)
					},
					[]string{ToolBlastX.String(), "-sorthsps=3"},
				},
				{
					func(n *BlastX) {
						n.SoftMasking.Set(true)
					},
					[]string{ToolBlastX.String(), "-soft_masking=true"},
				},
				{
					func(n *BlastX) {
						n.LowercaseMasking.Set(true)
					},
					[]string{ToolBlastX.String(), "-lcase_masking"},
				},
				{
					func(n *BlastX) {
						n.EntrezQuery.Set("goodbye")
					},
					[]string{ToolBlastX.String(), "-entrez_query='goodbye'"},
				},
				{
					func(n *BlastX) {
						n.QueryCoverageHSPPercent.Set(3.333333333333)
					},
					[]string{ToolBlastX.String(), "-qcov_hsp_perc=3.333333333333"},
				},
				{
					func(n *BlastX) {
						n.MaxHSPs.Set(5)
					},
					[]string{ToolBlastX.String(), "-max_hsps=5"},
				},
				{
					func(n *BlastX) {
						n.MaxTargetSequences.Set(5)
					},
					[]string{ToolBlastX.String(), "-max_target_seqs=5"},
				},
				{
					func(n *BlastX) {
						n.DBSize.Set(12)
					},
					[]string{ToolBlastX.String(), "-dbsize=12"},
				},
				{
					func(n *BlastX) {
						n.SearchSpace.Set(16)
					},
					[]string{ToolBlastX.String(), "-searchsp=16"},
				},
				{
					func(n *BlastX) {
						n.ImportSearchStrategy.Set("hi")
					},
					[]string{ToolBlastX.String(), "-import_search_strategy='hi'"},
				},
				{
					func(n *BlastX) {
						n.ExportSearchStrategy.Set("bi")
					},
					[]string{ToolBlastX.String(), "-export_search_strategy='bi'"},
				},
				{
					func(n *BlastX) {
						n.ExtensionDropoffUngapped.Set(5.555)
					},
					[]string{ToolBlastX.String(), "-xdrop_ungap=5.555"},
				},
				{
					func(n *BlastX) {
						n.WindowSize.Set(5)
					},
					[]string{ToolBlastX.String(), "-window_size=5"},
				},
				{
					func(n *BlastX) {
						n.ParseDefLines.Set(true)
					},
					[]string{ToolBlastX.String(), "-parse_deflines"},
				},
				{
					func(n *BlastX) {
						n.Remote.Set(true)
					},
					[]string{ToolBlastX.String(), "-remote"},
				},
			}

			for _, test := range tests {
				tgt := BlastX{}
				test.mod(&tgt)
				So(tgt.ToCLI().Args, ShouldResemble, test.out)
			}
		})

		Convey("IPG List Field Serialization", func() {
			tests := []struct {
				mod func(n *BlastX)
				out []string
			}{
				{
					func(n *BlastX) {
						n.GIList.Set("hope")
					},
					[]string{ToolBlastX.String(), "-gilist='hope'"},
				},
				{
					func(n *BlastX) {
						n.SequenceIDList.Set("hell")
					},
					[]string{ToolBlastX.String(), "-seqidlist='hell'"},
				},
				{
					func(n *BlastX) {
						n.NegativeGIList.Set("love")
					},
					[]string{ToolBlastX.String(), "-negative_gilist='love'"},
				},
				{
					func(n *BlastX) {
						n.NegativeSequenceIDList.Set("pain")
					},
					[]string{ToolBlastX.String(), "-negative_seqidlist='pain'"},
				},
				{
					func(n *BlastX) {
						n.TaxIDs = append(n.TaxIDs, "care", "kill")
					},
					[]string{ToolBlastX.String(), "-taxids='care,kill'"},
				},
				{
					func(n *BlastX) {
						n.NegativeTaxIDs = append(n.NegativeTaxIDs, "luck", "hate")
					},
					[]string{ToolBlastX.String(), "-negative_taxids='luck,hate'"},
				},
				{
					func(n *BlastX) {
						n.TaxIDList.Set("protect")
					},
					[]string{ToolBlastX.String(), "-taxidlist='protect'"},
				},
				{
					func(n *BlastX) {
						n.NegativeTaxIDList.Set("destroy")
					},
					[]string{ToolBlastX.String(), "-negative_taxidlist='destroy'"},
				},
				{
					func(n *BlastX) {
						n.IPGList.Set("build")
					},
					[]string{ToolBlastX.String(), "-ipglist='build'"},
				},
				{
					func(n *BlastX) {
						n.NegativeIPGList.Set("break")
					},
					[]string{ToolBlastX.String(), "-negative_ipglist='break'"},
				},
			}

			for _, test := range tests {
				tgt := BlastX{}
				test.mod(&tgt)
				So(tgt.ToCLI().Args, ShouldResemble, test.out)
			}
		})

		Convey("Best Hit Field Serialization", func() {
			tests := []struct {
				mod func(n *BlastX)
				out []string
			}{
				{
					func(n *BlastX) {
						n.BestHitOverhang.Set(1.125)
					},
					[]string{ToolBlastX.String(), "-best_hit_overhang=1.125"},
				},
				{
					func(n *BlastX) {
						n.BestHitScoreEdge.Set(1.25)
					},
					[]string{ToolBlastX.String(), "-best_hit_score_edge=1.25"},
				},
				{
					func(n *BlastX) {
						n.SubjectBestHit.Set(true)
					},
					[]string{ToolBlastX.String(), "-subject_besthit"},
				},
			}

			for _, test := range tests {
				tgt := BlastX{}
				test.mod(&tgt)
				So(tgt.ToCLI().Args, ShouldResemble, test.out)
			}
		})
	})
}
