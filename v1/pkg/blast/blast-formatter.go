package blast

import (
	"os/exec"

	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v1/pkg/blast/consts"
	"github.com/veupathdb/lib-go-blast/v1/pkg/blast/field"
	"lib-go-blast/v1/pkg/cli"
)

type BlastFormatter struct {
	ShortHelp          field.ShortHelp
	LongHelp           field.LongHelp
	Version            field.Version
	RequestID          field.RequestID
	ArchiveFile        field.ArchiveFile
	Format             field.Format
	ShowGIs            field.ShowGIs
	NumDescriptions    field.NumDescriptions
	NumAlignments      field.NumAlignments
	LineLength         field.LineLength
	HTML               field.HTML
	HitSorting         field.HitSorting
	HSPSorting         field.HSPSorting
	MaxTargetSequences field.MaxTargetSequences
	OutFile            field.OutFile
	ParseDefLines      field.ParseDefLines
}

func (b *BlastFormatter) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	switch key {
	case consts.FlagShortHelp:
		return field.DecodeJSONShortHelp(dec, &b.ShortHelp)
	case consts.FlagLongHelp:
		return field.DecodeJSONLongHelp(dec, &b.LongHelp)
	case consts.FlagVersion:
		return field.DecodeJSONVersion(dec, &b.Version)
	case consts.FlagRequestID:
		return field.DecodeJSONRequestID(dec, &b.RequestID)
	case consts.FlagArchiveFile:
		return field.DecodeJSONArchiveFile(dec, &b.ArchiveFile)
	case consts.FlagOutFormat:
		return field.DecodeJSONOutFormat(dec, &b.Format)
	case consts.FlagShowGIs:
		return field.DecodeJSONShowGIs(dec, &b.ShowGIs)
	case consts.FlagNumDescriptions:
		return field.DecodeJSONNumDescriptions(dec, &b.NumDescriptions)
	case consts.FlagNumAlignments:
		return field.DecodeJSONNumAlignments(dec, &b.NumAlignments)
	case consts.FlagLineLength:
		return field.DecodeJSONLineLength(dec, &b.LineLength)
	case consts.FlagHTML:
		return field.DecodeJSONHTML(dec, &b.HTML)
	case consts.FlagSortHits:
		return field.DecodeJSONHitSorting(dec, &b.HitSorting)
	case consts.FlagSortHSPs:
		return field.DecodeJSONHSPSorting(dec, &b.HSPSorting)
	case consts.FlagMaxTargetSequences:
		return field.DecodeJSONMaxTargetSequences(dec, &b.MaxTargetSequences)
	case consts.FlagOutFile:
		return field.DecodeJSONOutFile(dec, &b.OutFile)
	case consts.FlagParseDefLines:
		return field.DecodeJSONParseDefLines(dec, &b.ParseDefLines)
	default:
		return nil
	}
}

func (b *BlastFormatter) NKeys() int {
	return 0
}

func (b *BlastFormatter) ToCLI() *exec.Cmd {
	a := cli.NewBuilder().AppendFlag(ToolBlastFormatter.String())

	a.Append(b.ShortHelp).
		Append(b.LongHelp).
		Append(b.Version).
		Append(b.RequestID).
		Append(b.ArchiveFile).
		Append(b.Format).
		Append(b.ShowGIs).
		Append(b.NumDescriptions).
		Append(b.NumAlignments).
		Append(b.LineLength).
		Append(b.HTML).
		Append(b.HitSorting).
		Append(b.HSPSorting).
		Append(b.MaxTargetSequences).
		Append(b.OutFile).
		Append(b.ParseDefLines)

	out := new(exec.Cmd)
	out.Args = *a

	return out
}
