package blast_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	. "github.com/veupathdb/lib-go-blast/v2/pkg/blast"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/field"
)

func TestBlastP_ToCLI(t *testing.T) {
	Convey("blast.BlastP.ToCLI", t, func() {
		Convey("Base Field Serialization", func() {
			tests := []struct {
				mod func(n *BlastP)
				out []string
			}{
				{
					func(*BlastP) {},
					[]string{ToolBlastP.String()},
				},
				{
					func(n *BlastP) {
						n.ShortHelp.Set(true)
					},
					[]string{ToolBlastP.String(), consts.FlagShortHelp},
				},
				{
					func(n *BlastP) {
						n.LongHelp.Set(true)
					},
					[]string{ToolBlastP.String(), consts.FlagLongHelp},
				},
				{
					func(n *BlastP) {
						n.Version.Set(true)
					},
					[]string{ToolBlastP.String(), consts.FlagVersion},
				},
				{
					func(n *BlastP) {
						n.QueryFile.Set("hello")
					},
					[]string{ToolBlastP.String(), "-query='hello'"},
				},
				{
					func(n *BlastP) {
						n.QueryLocation.Start = 1
						n.QueryLocation.Stop = 2
					},
					[]string{ToolBlastP.String(), "-query_loc=1-2"},
				},
				{
					func(n *BlastP) {
						n.DBFile.Set("hello")
					},
					[]string{ToolBlastP.String(), "-db='hello'"},
				},
				{
					func(n *BlastP) {
						n.OutFile.Set("hello")
					},
					[]string{ToolBlastP.String(), "-out='hello'"},
				},
				{
					func(n *BlastP) {
						n.ExpectValue.Set("hello")
					},
					[]string{ToolBlastP.String(), "-evalue=hello"},
				},
				{
					func(n *BlastP) {
						n.Format.Type = field.FormatTypeBlastXML
					},
					[]string{ToolBlastP.String(), "-outfmt=5"},
				},
				{
					func(n *BlastP) {
						n.Format.Type = field.FormatTypeCommaSeparatedValues
						n.Format.Delimiter = '@'
					},
					[]string{ToolBlastP.String(), "-outfmt='10 delim=@'"},
				},
				{
					func(n *BlastP) {
						n.Format.Type = field.FormatTypeCommaSeparatedValues
						n.Format.Delimiter = '@'
						n.Format.Fields = field.FormatFieldSlice{field.FormatFieldQueryAccession}
					},
					[]string{ToolBlastP.String(), "-outfmt='10 delim=@ qacc'"},
				},
				{
					func(n *BlastP) {
						n.ShowGIs.Set(true)
					},
					[]string{ToolBlastP.String(), "-show_gis"},
				},
				{
					func(n *BlastP) {
						n.NumDescriptions.Set(3)
					},
					[]string{ToolBlastP.String(), "-num_descriptions=3"},
				},
				{
					func(n *BlastP) {
						n.NumAlignments.Set(3)
					},
					[]string{ToolBlastP.String(), "-num_alignments=3"},
				},
				{
					func(n *BlastP) {
						n.LineLength.Set(3)
					},
					[]string{ToolBlastP.String(), "-line_length=3"},
				},
				{
					func(n *BlastP) {
						n.HTML.Set(true)
					},
					[]string{ToolBlastP.String(), "-html"},
				},
				{
					func(n *BlastP) {
						n.SortHits.Set(field.HitSortingTypeByBitScore)
					},
					[]string{ToolBlastP.String(), "-sorthits=1"},
				},
				{
					func(n *BlastP) {
						n.SortHSPs.Set(field.HSPSortingTypeByHSPPercentIdentity)
					},
					[]string{ToolBlastP.String(), "-sorthsps=3"},
				},
				{
					func(n *BlastP) {
						n.SoftMasking.Set(true)
					},
					[]string{ToolBlastP.String(), "-soft_masking=true"},
				},
				{
					func(n *BlastP) {
						n.LowercaseMasking.Set(true)
					},
					[]string{ToolBlastP.String(), "-lcase_masking"},
				},
				{
					func(n *BlastP) {
						n.EntrezQuery.Set("goodbye")
					},
					[]string{ToolBlastP.String(), "-entrez_query='goodbye'"},
				},
				{
					func(n *BlastP) {
						n.QueryCoverageHSPPercent.Set(3.333333333333)
					},
					[]string{ToolBlastP.String(), "-qcov_hsp_perc=3.333333333333"},
				},
				{
					func(n *BlastP) {
						n.MaxHSPs.Set(5)
					},
					[]string{ToolBlastP.String(), "-max_hsps=5"},
				},
				{
					func(n *BlastP) {
						n.MaxTargetSequences.Set(5)
					},
					[]string{ToolBlastP.String(), "-max_target_seqs=5"},
				},
				{
					func(n *BlastP) {
						n.DBSize.Set(12)
					},
					[]string{ToolBlastP.String(), "-dbsize=12"},
				},
				{
					func(n *BlastP) {
						n.SearchSpace.Set(16)
					},
					[]string{ToolBlastP.String(), "-searchsp=16"},
				},
				{
					func(n *BlastP) {
						n.ImportSearchStrategy.Set("hi")
					},
					[]string{ToolBlastP.String(), "-import_search_strategy='hi'"},
				},
				{
					func(n *BlastP) {
						n.ExportSearchStrategy.Set("bi")
					},
					[]string{ToolBlastP.String(), "-export_search_strategy='bi'"},
				},
				{
					func(n *BlastP) {
						n.ExtensionDropoffUngapped.Set(5.555)
					},
					[]string{ToolBlastP.String(), "-xdrop_ungap=5.555"},
				},
				{
					func(n *BlastP) {
						n.WindowSize.Set(5)
					},
					[]string{ToolBlastP.String(), "-window_size=5"},
				},
				{
					func(n *BlastP) {
						n.ParseDefLines.Set(true)
					},
					[]string{ToolBlastP.String(), "-parse_deflines"},
				},
				{
					func(n *BlastP) {
						n.Remote.Set(true)
					},
					[]string{ToolBlastP.String(), "-remote"},
				},
			}

			for _, test := range tests {
				tgt := BlastP{}
				test.mod(&tgt)
				So(tgt.ToCLI().Args, ShouldResemble, test.out)
			}
		})

		Convey("IPG List Field Serialization", func() {
			tests := []struct {
				mod func(n *BlastP)
				out []string
			}{
				{
					func(n *BlastP) {
						n.GIList.Set("hope")
					},
					[]string{ToolBlastP.String(), "-gilist='hope'"},
				},
				{
					func(n *BlastP) {
						n.SequenceIDList.Set("hell")
					},
					[]string{ToolBlastP.String(), "-seqidlist='hell'"},
				},
				{
					func(n *BlastP) {
						n.NegativeGIList.Set("love")
					},
					[]string{ToolBlastP.String(), "-negative_gilist='love'"},
				},
				{
					func(n *BlastP) {
						n.NegativeSequenceIDList.Set("pain")
					},
					[]string{ToolBlastP.String(), "-negative_seqidlist='pain'"},
				},
				{
					func(n *BlastP) {
						n.TaxIDs = append(n.TaxIDs, "care", "kill")
					},
					[]string{ToolBlastP.String(), "-taxids='care,kill'"},
				},
				{
					func(n *BlastP) {
						n.NegativeTaxIDs = append(n.NegativeTaxIDs, "luck", "hate")
					},
					[]string{ToolBlastP.String(), "-negative_taxids='luck,hate'"},
				},
				{
					func(n *BlastP) {
						n.TaxIDList.Set("protect")
					},
					[]string{ToolBlastP.String(), "-taxidlist='protect'"},
				},
				{
					func(n *BlastP) {
						n.NegativeTaxIDList.Set("destroy")
					},
					[]string{ToolBlastP.String(), "-negative_taxidlist='destroy'"},
				},
				{
					func(n *BlastP) {
						n.IPGList.Set("build")
					},
					[]string{ToolBlastP.String(), "-ipglist='build'"},
				},
				{
					func(n *BlastP) {
						n.NegativeIPGList.Set("break")
					},
					[]string{ToolBlastP.String(), "-negative_ipglist='break'"},
				},
			}

			for _, test := range tests {
				tgt := BlastP{}
				test.mod(&tgt)
				So(tgt.ToCLI().Args, ShouldResemble, test.out)
			}
		})

		Convey("Best Hit Field Serialization", func() {
			tests := []struct {
				mod func(n *BlastP)
				out []string
			}{
				{
					func(n *BlastP) {
						n.BestHitOverhang.Set(1.125)
					},
					[]string{ToolBlastP.String(), "-best_hit_overhang=1.125"},
				},
				{
					func(n *BlastP) {
						n.BestHitScoreEdge.Set(1.25)
					},
					[]string{ToolBlastP.String(), "-best_hit_score_edge=1.25"},
				},
				{
					func(n *BlastP) {
						n.SubjectBestHit.Set(true)
					},
					[]string{ToolBlastP.String(), "-subject_besthit"},
				},
			}

			for _, test := range tests {
				tgt := BlastP{}
				test.mod(&tgt)
				So(tgt.ToCLI().Args, ShouldResemble, test.out)
			}
		})

		Convey("BlastP Field Serialization", func() {
			tests := []struct {
				mod func(n *BlastP)
				out []string
			}{
				{
					func(n *BlastP) {
						n.Task.Set(field.BlastPTaskTypeBlastPFast)
					},
					[]string{ToolBlastP.String(), "-task=blastp-fast"},
				},
				{
					func(n *BlastP) {
						n.WordSize.Set(22)
					},
					[]string{ToolBlastP.String(), "-word_size=22"},
				},
				{
					func(n *BlastP) {
						n.GapOpen.Set(23)
					},
					[]string{ToolBlastP.String(), "-gapopen=23"},
				},
				{
					func(n *BlastP) {
						n.GapExtend.Set(24)
					},
					[]string{ToolBlastP.String(), "-gapextend=24"},
				},
				{
					func(n *BlastP) {
						n.Matrix.Set("carrots")
					},
					[]string{ToolBlastP.String(), "-matrix='carrots'"},
				},
				{
					func(n *BlastP) {
						n.Threshold.Set(1.23)
					},
					[]string{ToolBlastP.String(), "-threshold=1.23"},
				},
				{
					func(n *BlastP) {
						n.CompBasedStats.Set(field.FullCompBasedStatsStatistics)
					},
					[]string{ToolBlastP.String(), "-comp_based_stats=1"},
				},
				{
					func(n *BlastP) {
						n.SubjectFile.Set("radio")
					},
					[]string{ToolBlastP.String(), "-subject='radio'"},
				},
				{
					func(n *BlastP) {
						n.SubjectLocation.Start=56
						n.SubjectLocation.Stop=57
					},
					[]string{ToolBlastP.String(), "-subject_loc=56-57"},
				},
				{
					func(n *BlastP) {
						n.Seg = field.NewSeg(10, 12, 14)
					},
					[]string{ToolBlastP.String(), "-seg='10 12 14'"},
				},
				{
					func(n *BlastP) {
						n.DBSoftMask.Set("panini")
					},
					[]string{ToolBlastP.String(), "-db_soft_mask='panini'"},
				},
				{
					func(n *BlastP) {
						n.DBHardMask.Set("holiday")
					},
					[]string{ToolBlastP.String(), "-db_hard_mask='holiday'"},
				},
				{
					func(n *BlastP) {
						n.CullingLimit.Set(23)
					},
					[]string{ToolBlastP.String(), "-culling_limit=23"},
				},
				{
					func(n *BlastP) {
						n.ExtensionDropoffPrelimGapped.Set(2.34)
					},
					[]string{ToolBlastP.String(), "-xdrop_gap=2.34"},
				},
				{
					func(n *BlastP) {
						n.ExtensionDropoffFinalGapped.Set(3.45)
					},
					[]string{ToolBlastP.String(), "-xdrop_gap_final=3.45"},
				},
				{
					func(n *BlastP) {
						n.UngappedAlignmentsOnly.Set(true)
					},
					[]string{ToolBlastP.String(), "-ungapped"},
				},
				{
					func(n *BlastP) {
						n.NumThreads.Set(25)
					},
					[]string{ToolBlastP.String(), "-num_threads=25"},
				},
				{
					func(n *BlastP) {
						n.UseSmithWatermanTraceback.Set(true)
					},
					[]string{ToolBlastP.String(), "-use_sw_tback"},
				},
			}

			for _, test := range tests {
				tgt := BlastP{}
				test.mod(&tgt)
				So(tgt.ToCLI().Args, ShouldResemble, test.out)
			}
		})
	})
}
