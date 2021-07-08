package blast_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	. "github.com/veupathdb/lib-go-blast/v2/pkg/blast"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/field"
)

func TestBlastN_ToCLI(t *testing.T) {
	Convey("blast.BlastN.ToCLI", t, func() {
		Convey("Base Field Serialization", func() {
			tests := []struct {
				mod func(n *BlastN)
				out []string
			}{
				{
					func(*BlastN) {},
					[]string{ToolBlastN.String()},
				},
				{
					func(n *BlastN) {
						n.ShortHelp.Set(true)
					},
					[]string{ToolBlastN.String(), consts.FlagShortHelp},
				},
				{
					func(n *BlastN) {
						n.LongHelp.Set(true)
					},
					[]string{ToolBlastN.String(), consts.FlagLongHelp},
				},
				{
					func(n *BlastN) {
						n.Version.Set(true)
					},
					[]string{ToolBlastN.String(), consts.FlagVersion},
				},
				{
					func(n *BlastN) {
						n.QueryFile.Set("hello")
					},
					[]string{ToolBlastN.String(), "-query=hello"},
				},
				{
					func(n *BlastN) {
						n.QueryLocation.Start = 1
						n.QueryLocation.Stop = 2
					},
					[]string{ToolBlastN.String(), "-query_loc=1-2"},
				},
				{
					func(n *BlastN) {
						n.DBFile.Set("hello")
					},
					[]string{ToolBlastN.String(), "-db=hello"},
				},
				{
					func(n *BlastN) {
						n.OutFile.Set("hello")
					},
					[]string{ToolBlastN.String(), "-out=hello"},
				},
				{
					func(n *BlastN) {
						n.ExpectValue.Set("hello")
					},
					[]string{ToolBlastN.String(), "-evalue=hello"},
				},
				{
					func(n *BlastN) {
						n.Format.Type = field.FormatTypeBlastXML
					},
					[]string{ToolBlastN.String(), "-outfmt=5"},
				},
				{
					func(n *BlastN) {
						n.Format.Type = field.FormatTypeCommaSeparatedValues
						n.Format.Delimiter = '@'
					},
					[]string{ToolBlastN.String(), "-outfmt=10 delim=@"},
				},
				{
					func(n *BlastN) {
						n.Format.Type = field.FormatTypeCommaSeparatedValues
						n.Format.Delimiter = '@'
						n.Format.Fields = field.FormatFieldSlice{field.FormatFieldQueryAccession}
					},
					[]string{ToolBlastN.String(), "-outfmt=10 delim=@ qacc"},
				},
				{
					func(n *BlastN) {
						n.ShowGIs.Set(true)
					},
					[]string{ToolBlastN.String(), "-show_gis"},
				},
				{
					func(n *BlastN) {
						n.NumDescriptions.Set(3)
					},
					[]string{ToolBlastN.String(), "-num_descriptions=3"},
				},
				{
					func(n *BlastN) {
						n.NumAlignments.Set(3)
					},
					[]string{ToolBlastN.String(), "-num_alignments=3"},
				},
				{
					func(n *BlastN) {
						n.LineLength.Set(3)
					},
					[]string{ToolBlastN.String(), "-line_length=3"},
				},
				{
					func(n *BlastN) {
						n.HTML.Set(true)
					},
					[]string{ToolBlastN.String(), "-html"},
				},
				{
					func(n *BlastN) {
						n.SortHits.Set(field.HitSortingTypeByBitScore)
					},
					[]string{ToolBlastN.String(), "-sorthits=1"},
				},
				{
					func(n *BlastN) {
						n.SortHSPs.Set(field.HSPSortingTypeByHSPPercentIdentity)
					},
					[]string{ToolBlastN.String(), "-sorthsps=3"},
				},
				{
					func(n *BlastN) {
						n.SoftMasking.Set(true)
					},
					[]string{ToolBlastN.String(), "-soft_masking=true"},
				},
				{
					func(n *BlastN) {
						n.LowercaseMasking.Set(true)
					},
					[]string{ToolBlastN.String(), "-lcase_masking"},
				},
				{
					func(n *BlastN) {
						n.EntrezQuery.Set("goodbye")
					},
					[]string{ToolBlastN.String(), "-entrez_query=goodbye"},
				},
				{
					func(n *BlastN) {
						n.QueryCoverageHSPPercent.Set(3.333333333333)
					},
					[]string{ToolBlastN.String(), "-qcov_hsp_perc=3.333333333333"},
				},
				{
					func(n *BlastN) {
						n.MaxHSPs.Set(5)
					},
					[]string{ToolBlastN.String(), "-max_hsps=5"},
				},
				{
					func(n *BlastN) {
						n.MaxTargetSequences.Set(5)
					},
					[]string{ToolBlastN.String(), "-max_target_seqs=5"},
				},
				{
					func(n *BlastN) {
						n.DBSize.Set(12)
					},
					[]string{ToolBlastN.String(), "-dbsize=12"},
				},
				{
					func(n *BlastN) {
						n.SearchSpace.Set(16)
					},
					[]string{ToolBlastN.String(), "-searchsp=16"},
				},
				{
					func(n *BlastN) {
						n.ImportSearchStrategy.Set("hi")
					},
					[]string{ToolBlastN.String(), "-import_search_strategy=hi"},
				},
				{
					func(n *BlastN) {
						n.ExportSearchStrategy.Set("bi")
					},
					[]string{ToolBlastN.String(), "-export_search_strategy=bi"},
				},
				{
					func(n *BlastN) {
						n.ExtensionDropoffUngapped.Set(5.555)
					},
					[]string{ToolBlastN.String(), "-xdrop_ungap=5.555"},
				},
				{
					func(n *BlastN) {
						n.WindowSize.Set(5)
					},
					[]string{ToolBlastN.String(), "-window_size=5"},
				},
				{
					func(n *BlastN) {
						n.ParseDefLines.Set(true)
					},
					[]string{ToolBlastN.String(), "-parse_deflines"},
				},
				{
					func(n *BlastN) {
						n.Remote.Set(true)
					},
					[]string{ToolBlastN.String(), "-remote"},
				},
			}

			for _, test := range tests {
				tgt := BlastN{}
				test.mod(&tgt)
				So(tgt.ToCLI().Args, ShouldResemble, test.out)
			}
		})

		Convey("Base List Field Serialization", func() {
			tests := []struct {
				mod func(n *BlastN)
				out []string
			}{
				{
					func(n *BlastN) {
						n.GIList.Set("hope")
					},
					[]string{ToolBlastN.String(), "-gilist=hope"},
				},
				{
					func(n *BlastN) {
						n.SequenceIDList.Set("hell")
					},
					[]string{ToolBlastN.String(), "-seqidlist=hell"},
				},
				{
					func(n *BlastN) {
						n.NegativeGIList.Set("love")
					},
					[]string{ToolBlastN.String(), "-negative_gilist=love"},
				},
				{
					func(n *BlastN) {
						n.NegativeSequenceIDList.Set("pain")
					},
					[]string{ToolBlastN.String(), "-negative_seqidlist=pain"},
				},
				{
					func(n *BlastN) {
						n.TaxIDs = append(n.TaxIDs, "care", "kill")
					},
					[]string{ToolBlastN.String(), "-taxids=care,kill"},
				},
				{
					func(n *BlastN) {
						n.NegativeTaxIDs = append(n.NegativeTaxIDs, "luck", "hate")
					},
					[]string{ToolBlastN.String(), "-negative_taxids=luck,hate"},
				},
				{
					func(n *BlastN) {
						n.TaxIDList.Set("protect")
					},
					[]string{ToolBlastN.String(), "-taxidlist=protect"},
				},
				{
					func(n *BlastN) {
						n.NegativeTaxIDList.Set("destroy")
					},
					[]string{ToolBlastN.String(), "-negative_taxidlist=destroy"},
				},
			}

			for _, test := range tests {
				tgt := BlastN{}
				test.mod(&tgt)
				So(tgt.ToCLI().Args, ShouldResemble, test.out)
			}
		})

		Convey("Best Hit Field Serialization", func() {
			tests := []struct {
				mod func(n *BlastN)
				out []string
			}{
				{
					func(n *BlastN) {
						n.BestHitOverhang.Set(1.125)
					},
					[]string{ToolBlastN.String(), "-best_hit_overhang=1.125"},
				},
				{
					func(n *BlastN) {
						n.BestHitScoreEdge.Set(1.25)
					},
					[]string{ToolBlastN.String(), "-best_hit_score_edge=1.25"},
				},
				{
					func(n *BlastN) {
						n.SubjectBestHit.Set(true)
					},
					[]string{ToolBlastN.String(), "-subject_besthit"},
				},
			}

			for _, test := range tests {
				tgt := BlastN{}
				test.mod(&tgt)
				So(tgt.ToCLI().Args, ShouldResemble, test.out)
			}
		})


		Convey("BlastN Field Serialization", func() {
			tests := []struct {
				mod func(n *BlastN)
				out []string
			}{
				{
					func(n *BlastN) {
						n.Strand.Set(field.StrandTypeMinus)
					},
					[]string{ToolBlastN.String(), "-strand=minus"},
				},
				{
					func(n *BlastN) {
						n.Task.Set(field.BlastNTaskTypeBlastN)
					},
					[]string{ToolBlastN.String(), "-task=blastn"},
				},
				{
					func(n *BlastN) {
						n.WordSize.Set(333)
					},
					[]string{ToolBlastN.String(), "-word_size=333"},
				},
				{
					func(n *BlastN) {
						n.GapOpen.Set(112)
					},
					[]string{ToolBlastN.String(), "-gapopen=112"},
				},
				{
					func(n *BlastN) {
						n.GapExtend.Set(221)
					},
					[]string{ToolBlastN.String(), "-gapextend=221"},
				},
				{
					func(n *BlastN) {
						n.Penalty.Set(-303)
					},
					[]string{ToolBlastN.String(), "-penalty=-303"},
				},
				{
					func(n *BlastN) {
						n.Reward.Set(808)
					},
					[]string{ToolBlastN.String(), "-reward=808"},
				},
				{
					func(n *BlastN) {
						n.UseIndex.Set(true)
					},
					[]string{ToolBlastN.String(), "-use_index=true"},
				},
				{
					func(n *BlastN) {
						n.IndexName.Set("tacos")
					},
					[]string{ToolBlastN.String(), "-index_name=tacos"},
				},
				{
					func(n *BlastN) {
						n.SubjectFile.Set("ambivalence")
					},
					[]string{ToolBlastN.String(), "-subject=ambivalence"},
				},
				{
					func(n *BlastN) {
						n.SubjectLocation.Start = 12
						n.SubjectLocation.Stop  = 15
					},
					[]string{ToolBlastN.String(), "-subject_loc=12-15"},
				},
				{
					func(n *BlastN) {
						n.Dust = field.NewLWLDust(10, 11, 12)
					},
					[]string{ToolBlastN.String(), "-dust=10 11 12"},
				},
				{
					func(n *BlastN) {
						n.FilteringDB.Set("grapes")
					},
					[]string{ToolBlastN.String(), "-filtering_db=grapes"},
				},
				{
					func(n *BlastN) {
						n.WindowMaskerTaxID.Set(42)
					},
					[]string{ToolBlastN.String(), "-window_masker_taxid=42"},
				},
				{
					func(n *BlastN) {
						n.WindowMaskerDB.Set("hairy")
					},
					[]string{ToolBlastN.String(), "-window_masker_db=hairy"},
				},
				{
					func(n *BlastN) {
						n.DBSoftMask.Set("ninety")
					},
					[]string{ToolBlastN.String(), "-db_soft_mask=ninety"},
				},
				{
					func(n *BlastN) {
						n.DBHardMask.Set("acorn")
					},
					[]string{ToolBlastN.String(), "-db_hard_mask=acorn"},
				},
				{
					func(n *BlastN) {
						n.PercentIdentity.Set(33.33)
					},
					[]string{ToolBlastN.String(), "-perc_identity=33.33"},
				},
				{
					func(n *BlastN) {
						n.CullingLimit.Set(2)
					},
					[]string{ToolBlastN.String(), "-culling_limit=2"},
				},
				{
					func(n *BlastN) {
						n.TemplateType.Set(field.TemplateTypeEnumCoding)
					},
					[]string{ToolBlastN.String(), "-template_type=coding"},
				},
				{
					func(n *BlastN) {
						n.TemplateLength.Set(field.TemplateLengthEnum18)
					},
					[]string{ToolBlastN.String(), "-template_length=18"},
				},
				{
					func(n *BlastN) {
						n.SumStats.Set(true)
					},
					[]string{ToolBlastN.String(), "-sum_stats=true"},
				},
				{
					func(n *BlastN) {
						n.ExtensionDropoffPrelimGapped.Set(66.6)
					},
					[]string{ToolBlastN.String(), "-xdrop_gap=66.6"},
				},
				{
					func(n *BlastN) {
						n.ExtensionDropoffFinalGapped.Set(6.66)
					},
					[]string{ToolBlastN.String(), "-xdrop_gap_final=6.66"},
				},
				{
					func(n *BlastN) {
						n.NonGreedy.Set(true)
					},
					[]string{ToolBlastN.String(), "-no_greedy"},
				},
				{
					func(n *BlastN) {
						n.MinRawGappedScore.Set(200)
					},
					[]string{ToolBlastN.String(), "-min_raw_gapped_score=200"},
				},
				{
					func(n *BlastN) {
						n.UngappedAlignmentsOnly.Set(true)
					},
					[]string{ToolBlastN.String(), "-ungapped"},
				},
				{
					func(n *BlastN) {
						n.OffDiagonalRange.Set(16)
					},
					[]string{ToolBlastN.String(), "-off_diagonal_range=16"},
				},
				{
					func(n *BlastN) {
						n.NumThreads.Set(15)
					},
					[]string{ToolBlastN.String(), "-num_threads=15"},
				},
			}

			for _, test := range tests {
				tgt := BlastN{}
				test.mod(&tgt)
				So(tgt.ToCLI().Args, ShouldResemble, test.out)
			}
		})
	})
}
