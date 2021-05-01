package blast

type Tool string

const (
	ToolBlastN         Tool = "blastn"
	ToolBlastP         Tool = "blastp"
	ToolBlastX         Tool = "blastx"
	ToolDeltaBlast     Tool = "deltablast"
	ToolPSIBlast       Tool = "psiblast"
	ToolRPSBlast       Tool = "rpsblast"
	ToolRPSTBlastN     Tool = "rpstblastn"
	ToolTBlastN        Tool = "tblastn"
	ToolTBlastX        Tool = "tblastx"
	ToolBlastFormatter Tool = "blast_formatter"
)

func (t Tool) String() string {
	return string(t)
}
