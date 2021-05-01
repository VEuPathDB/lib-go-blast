package blast_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	. "github.com/veupathdb/lib-go-blast/v2/pkg/blast"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/field"
)

func TestRPSBlast_ToCLI(t *testing.T) {
	Convey("blast.RPSBlast.ToCLI", t, func() {
		Convey("Base Field Serialization", func() {
			tests := []struct {
				mod func(n *RPSBlast)
				out []string
			}{
				{
					func(*RPSBlast) {},
					[]string{ToolRPSBlast.String()},
				},
				{
					func(n *RPSBlast) {
						n.ShortHelp.Set(true)
					},
					[]string{ToolRPSBlast.String(), consts.FlagShortHelp},
				},
				{
					func(n *RPSBlast) {
						n.LongHelp.Set(true)
					},
					[]string{ToolRPSBlast.String(), consts.FlagLongHelp},
				},
				{
					func(n *RPSBlast) {
						n.Version.Set(true)
					},
					[]string{ToolRPSBlast.String(), consts.FlagVersion},
				},
				{
					func(n *RPSBlast) {
						n.QueryFile.Set("hello")
					},
					[]string{ToolRPSBlast.String(), "-query='hello'"},
				},
				{
					func(n *RPSBlast) {
						n.QueryLocation.Start = 1
						n.QueryLocation.Stop = 2
					},
					[]string{ToolRPSBlast.String(), "-query_loc=1-2"},
				},
				{
					func(n *RPSBlast) {
						n.DBFile.Set("hello")
					},
					[]string{ToolRPSBlast.String(), "-db='hello'"},
				},
				{
					func(n *RPSBlast) {
						n.OutFile.Set("hello")
					},
					[]string{ToolRPSBlast.String(), "-out='hello'"},
				},
				{
					func(n *RPSBlast) {
						n.ExpectValue.Set("hello")
					},
					[]string{ToolRPSBlast.String(), "-evalue=hello"},
				},
				{
					func(n *RPSBlast) {
						n.Format.Type = field.FormatTypeBlastXML
					},
					[]string{ToolRPSBlast.String(), "-outfmt=5"},
				},
				{
					func(n *RPSBlast) {
						n.Format.Type = field.FormatTypeCommaSeparatedValues
						n.Format.Delimiter = '@'
					},
					[]string{ToolRPSBlast.String(), "-outfmt='10 delim=@'"},
				},
				{
					func(n *RPSBlast) {
						n.Format.Type = field.FormatTypeCommaSeparatedValues
						n.Format.Delimiter = '@'
						n.Format.Fields = field.FormatFieldSlice{field.FormatFieldQueryAccession}
					},
					[]string{ToolRPSBlast.String(), "-outfmt='10 delim=@ qacc'"},
				},
				{
					func(n *RPSBlast) {
						n.ShowGIs.Set(true)
					},
					[]string{ToolRPSBlast.String(), "-show_gis"},
				},
				{
					func(n *RPSBlast) {
						n.NumDescriptions.Set(3)
					},
					[]string{ToolRPSBlast.String(), "-num_descriptions=3"},
				},
				{
					func(n *RPSBlast) {
						n.NumAlignments.Set(3)
					},
					[]string{ToolRPSBlast.String(), "-num_alignments=3"},
				},
				{
					func(n *RPSBlast) {
						n.LineLength.Set(3)
					},
					[]string{ToolRPSBlast.String(), "-line_length=3"},
				},
				{
					func(n *RPSBlast) {
						n.HTML.Set(true)
					},
					[]string{ToolRPSBlast.String(), "-html"},
				},
				{
					func(n *RPSBlast) {
						n.SortHits.Set(field.HitSortingTypeByBitScore)
					},
					[]string{ToolRPSBlast.String(), "-sorthits=1"},
				},
				{
					func(n *RPSBlast) {
						n.SortHSPs.Set(field.HSPSortingTypeByHSPPercentIdentity)
					},
					[]string{ToolRPSBlast.String(), "-sorthsps=3"},
				},
				{
					func(n *RPSBlast) {
						n.SoftMasking.Set(true)
					},
					[]string{ToolRPSBlast.String(), "-soft_masking=true"},
				},
				{
					func(n *RPSBlast) {
						n.LowercaseMasking.Set(true)
					},
					[]string{ToolRPSBlast.String(), "-lcase_masking"},
				},
				{
					func(n *RPSBlast) {
						n.EntrezQuery.Set("goodbye")
					},
					[]string{ToolRPSBlast.String(), "-entrez_query='goodbye'"},
				},
				{
					func(n *RPSBlast) {
						n.QueryCoverageHSPPercent.Set(3.333333333333)
					},
					[]string{ToolRPSBlast.String(), "-qcov_hsp_perc=3.333333333333"},
				},
				{
					func(n *RPSBlast) {
						n.MaxHSPs.Set(5)
					},
					[]string{ToolRPSBlast.String(), "-max_hsps=5"},
				},
				{
					func(n *RPSBlast) {
						n.MaxTargetSequences.Set(5)
					},
					[]string{ToolRPSBlast.String(), "-max_target_seqs=5"},
				},
				{
					func(n *RPSBlast) {
						n.DBSize.Set(12)
					},
					[]string{ToolRPSBlast.String(), "-dbsize=12"},
				},
				{
					func(n *RPSBlast) {
						n.SearchSpace.Set(16)
					},
					[]string{ToolRPSBlast.String(), "-searchsp=16"},
				},
				{
					func(n *RPSBlast) {
						n.ImportSearchStrategy.Set("hi")
					},
					[]string{ToolRPSBlast.String(), "-import_search_strategy='hi'"},
				},
				{
					func(n *RPSBlast) {
						n.ExportSearchStrategy.Set("bi")
					},
					[]string{ToolRPSBlast.String(), "-export_search_strategy='bi'"},
				},
				{
					func(n *RPSBlast) {
						n.ExtensionDropoffUngapped.Set(5.555)
					},
					[]string{ToolRPSBlast.String(), "-xdrop_ungap=5.555"},
				},
				{
					func(n *RPSBlast) {
						n.WindowSize.Set(5)
					},
					[]string{ToolRPSBlast.String(), "-window_size=5"},
				},
				{
					func(n *RPSBlast) {
						n.ParseDefLines.Set(true)
					},
					[]string{ToolRPSBlast.String(), "-parse_deflines"},
				},
				{
					func(n *RPSBlast) {
						n.Remote.Set(true)
					},
					[]string{ToolRPSBlast.String(), "-remote"},
				},
			}

			for _, test := range tests {
				tgt := RPSBlast{}
				test.mod(&tgt)
				So(tgt.ToCLI().Args, ShouldResemble, test.out)
			}
		})

		Convey("Best Hit Field Serialization", func() {
			tests := []struct {
				mod func(n *RPSBlast)
				out []string
			}{
				{
					func(n *RPSBlast) {
						n.BestHitOverhang.Set(1.125)
					},
					[]string{ToolRPSBlast.String(), "-best_hit_overhang=1.125"},
				},
				{
					func(n *RPSBlast) {
						n.BestHitScoreEdge.Set(1.25)
					},
					[]string{ToolRPSBlast.String(), "-best_hit_score_edge=1.25"},
				},
				{
					func(n *RPSBlast) {
						n.SubjectBestHit.Set(true)
					},
					[]string{ToolRPSBlast.String(), "-subject_besthit"},
				},
			}

			for _, test := range tests {
				tgt := RPSBlast{}
				test.mod(&tgt)
				So(tgt.ToCLI().Args, ShouldResemble, test.out)
			}
		})
	})
}
