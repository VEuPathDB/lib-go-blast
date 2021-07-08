package blast_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	. "github.com/veupathdb/lib-go-blast/v2/pkg/blast"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/field"
)

func TestRPSTBlastN_ToCLI(t *testing.T) {
	Convey("blast.RPSTBlastN.ToCLI", t, func() {
		Convey("Base Field Serialization", func() {
			tests := []struct {
				mod func(n *RPSTBlastN)
				out []string
			}{
				{
					func(*RPSTBlastN) {},
					[]string{ToolRPSTBlastN.String()},
				},
				{
					func(n *RPSTBlastN) {
						n.ShortHelp.Set(true)
					},
					[]string{ToolRPSTBlastN.String(), consts.FlagShortHelp},
				},
				{
					func(n *RPSTBlastN) {
						n.LongHelp.Set(true)
					},
					[]string{ToolRPSTBlastN.String(), consts.FlagLongHelp},
				},
				{
					func(n *RPSTBlastN) {
						n.Version.Set(true)
					},
					[]string{ToolRPSTBlastN.String(), consts.FlagVersion},
				},
				{
					func(n *RPSTBlastN) {
						n.QueryFile.Set("hello")
					},
					[]string{ToolRPSTBlastN.String(), "-query=hello"},
				},
				{
					func(n *RPSTBlastN) {
						n.QueryLocation.Start = 1
						n.QueryLocation.Stop = 2
					},
					[]string{ToolRPSTBlastN.String(), "-query_loc=1-2"},
				},
				{
					func(n *RPSTBlastN) {
						n.DBFile.Set("hello")
					},
					[]string{ToolRPSTBlastN.String(), "-db=hello"},
				},
				{
					func(n *RPSTBlastN) {
						n.OutFile.Set("hello")
					},
					[]string{ToolRPSTBlastN.String(), "-out=hello"},
				},
				{
					func(n *RPSTBlastN) {
						n.ExpectValue.Set("hello")
					},
					[]string{ToolRPSTBlastN.String(), "-evalue=hello"},
				},
				{
					func(n *RPSTBlastN) {
						n.Format.Type = field.FormatTypeBlastXML
					},
					[]string{ToolRPSTBlastN.String(), "-outfmt=5"},
				},
				{
					func(n *RPSTBlastN) {
						n.Format.Type = field.FormatTypeCommaSeparatedValues
						n.Format.Delimiter = '@'
					},
					[]string{ToolRPSTBlastN.String(), "-outfmt=10 delim=@"},
				},
				{
					func(n *RPSTBlastN) {
						n.Format.Type = field.FormatTypeCommaSeparatedValues
						n.Format.Delimiter = '@'
						n.Format.Fields = field.FormatFieldSlice{field.FormatFieldQueryAccession}
					},
					[]string{ToolRPSTBlastN.String(), "-outfmt=10 delim=@ qacc"},
				},
				{
					func(n *RPSTBlastN) {
						n.ShowGIs.Set(true)
					},
					[]string{ToolRPSTBlastN.String(), "-show_gis"},
				},
				{
					func(n *RPSTBlastN) {
						n.NumDescriptions.Set(3)
					},
					[]string{ToolRPSTBlastN.String(), "-num_descriptions=3"},
				},
				{
					func(n *RPSTBlastN) {
						n.NumAlignments.Set(3)
					},
					[]string{ToolRPSTBlastN.String(), "-num_alignments=3"},
				},
				{
					func(n *RPSTBlastN) {
						n.LineLength.Set(3)
					},
					[]string{ToolRPSTBlastN.String(), "-line_length=3"},
				},
				{
					func(n *RPSTBlastN) {
						n.HTML.Set(true)
					},
					[]string{ToolRPSTBlastN.String(), "-html"},
				},
				{
					func(n *RPSTBlastN) {
						n.SortHits.Set(field.HitSortingTypeByBitScore)
					},
					[]string{ToolRPSTBlastN.String(), "-sorthits=1"},
				},
				{
					func(n *RPSTBlastN) {
						n.SortHSPs.Set(field.HSPSortingTypeByHSPPercentIdentity)
					},
					[]string{ToolRPSTBlastN.String(), "-sorthsps=3"},
				},
				{
					func(n *RPSTBlastN) {
						n.SoftMasking.Set(true)
					},
					[]string{ToolRPSTBlastN.String(), "-soft_masking=true"},
				},
				{
					func(n *RPSTBlastN) {
						n.LowercaseMasking.Set(true)
					},
					[]string{ToolRPSTBlastN.String(), "-lcase_masking"},
				},
				{
					func(n *RPSTBlastN) {
						n.EntrezQuery.Set("goodbye")
					},
					[]string{ToolRPSTBlastN.String(), "-entrez_query=goodbye"},
				},
				{
					func(n *RPSTBlastN) {
						n.QueryCoverageHSPPercent.Set(3.333333333333)
					},
					[]string{ToolRPSTBlastN.String(), "-qcov_hsp_perc=3.333333333333"},
				},
				{
					func(n *RPSTBlastN) {
						n.MaxHSPs.Set(5)
					},
					[]string{ToolRPSTBlastN.String(), "-max_hsps=5"},
				},
				{
					func(n *RPSTBlastN) {
						n.MaxTargetSequences.Set(5)
					},
					[]string{ToolRPSTBlastN.String(), "-max_target_seqs=5"},
				},
				{
					func(n *RPSTBlastN) {
						n.DBSize.Set(12)
					},
					[]string{ToolRPSTBlastN.String(), "-dbsize=12"},
				},
				{
					func(n *RPSTBlastN) {
						n.SearchSpace.Set(16)
					},
					[]string{ToolRPSTBlastN.String(), "-searchsp=16"},
				},
				{
					func(n *RPSTBlastN) {
						n.ImportSearchStrategy.Set("hi")
					},
					[]string{ToolRPSTBlastN.String(), "-import_search_strategy=hi"},
				},
				{
					func(n *RPSTBlastN) {
						n.ExportSearchStrategy.Set("bi")
					},
					[]string{ToolRPSTBlastN.String(), "-export_search_strategy=bi"},
				},
				{
					func(n *RPSTBlastN) {
						n.ExtensionDropoffUngapped.Set(5.555)
					},
					[]string{ToolRPSTBlastN.String(), "-xdrop_ungap=5.555"},
				},
				{
					func(n *RPSTBlastN) {
						n.WindowSize.Set(5)
					},
					[]string{ToolRPSTBlastN.String(), "-window_size=5"},
				},
				{
					func(n *RPSTBlastN) {
						n.ParseDefLines.Set(true)
					},
					[]string{ToolRPSTBlastN.String(), "-parse_deflines"},
				},
				{
					func(n *RPSTBlastN) {
						n.Remote.Set(true)
					},
					[]string{ToolRPSTBlastN.String(), "-remote"},
				},
			}

			for _, test := range tests {
				tgt := RPSTBlastN{}
				test.mod(&tgt)
				So(tgt.ToCLI().Args, ShouldResemble, test.out)
			}
		})
	})
}
