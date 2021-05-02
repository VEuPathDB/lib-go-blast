package traits

import (
	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/bval"
	. "github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/field"
	"github.com/veupathdb/lib-go-blast/v2/pkg/cli"
)

// CLIConfig defines a container for universal blast tool options across the
// supported blast+ tools.
//
// In practice this type shouldn't need to be used directly, instead see one of
// the tool specific CLI configs, such as BlastN, BlastP, BlastX, etc..
type CLIConfig struct {
	ShortHelp          field.ShortHelp
	LongHelp           field.LongHelp
	Version            field.Version
	Format             field.Format
	ShowGIs            field.ShowGIs
	NumDescriptions    field.NumDescriptions
	NumAlignments      field.NumAlignments
	LineLength         field.LineLength
	HTML               field.HTML
	SortHits           field.HitSorting
	SortHSPs           field.HSPSorting
	MaxTargetSequences field.MaxTargetSequences
	OutFile            field.OutFile
	ParseDefLines      field.ParseDefLines
}

func (b *CLIConfig) UnmarshalJSONObject(dec *gojay.Decoder, key string) (bool, error) {
	switch key {
	case FlagShortHelp:
		return true, field.DecodeJSONShortHelp(dec, &b.ShortHelp)
	case FlagLongHelp:
		return true, field.DecodeJSONLongHelp(dec, &b.LongHelp)
	case FlagVersion:
		return true, field.DecodeJSONVersion(dec, &b.Version)
	case FlagOutFormat:
		return true, field.DecodeJSONOutFormat(dec, &b.Format)
	case FlagShowGIs:
		return true, field.DecodeJSONShowGIs(dec, &b.ShowGIs)
	case FlagNumDescriptions:
		return true, field.DecodeJSONNumDescriptions(dec, &b.NumDescriptions)
	case FlagNumAlignments:
		return true, field.DecodeJSONNumAlignments(dec, &b.NumAlignments)
	case FlagLineLength:
		return true, field.DecodeJSONLineLength(dec, &b.LineLength)
	case FlagHTML:
		return true, field.DecodeJSONHTML(dec, &b.HTML)
	case FlagSortHits:
		return true, field.DecodeJSONHitSorting(dec, &b.SortHits)
	case FlagSortHSPs:
		return true, field.DecodeJSONHSPSorting(dec, &b.SortHSPs)
	case FlagMaxTargetSequences:
		return true, field.DecodeJSONMaxTargetSequences(dec, &b.MaxTargetSequences)
	case FlagOutFile:
		return true, field.DecodeJSONOutFile(dec, &b.OutFile)
	case FlagParseDefLines:
		return true, field.DecodeJSONParseDefLines(dec, &b.ParseDefLines)
	default:
		return false, nil
	}
}

func (b *CLIConfig) ToCLI(builder *cli.Builder) {
	builder.Append(&b.ShortHelp).
		Append(&b.LongHelp).
		Append(&b.Version).
		Append(&b.Format).
		Append(&b.ShowGIs).
		Append(&b.NumDescriptions).
		Append(&b.NumAlignments).
		Append(&b.LineLength).
		Append(&b.HTML).
		Append(&b.SortHits).
		Append(&b.SortHSPs).
		Append(&b.MaxTargetSequences).
		Append(&b.OutFile).
		Append(&b.ParseDefLines)
}

func (b *CLIConfig) Validate(errors bval.ValidationBuilder) {
	_ = errors.Validate(&b.Format).
		Incompatible(&b.NumDescriptions, &b.MaxTargetSequences).
		Incompatible(&b.NumAlignments, &b.MaxTargetSequences).
		Validate(&b.LineLength).
		ReportTypeLTEQ(
			uint8(b.Format.Type),
			uint8(field.FormatTypeFlatQueryAnchoredNoIdentities),
			&b.NumDescriptions,
			&b.LineLength,
			&b.SortHits,
		).
		ReportTypeEQ(
			uint8(b.Format.Type),
			uint8(field.FormatTypePairwise),
			&b.SortHSPs,
		).
		Validate(&b.SortHits).
		Validate(&b.SortHSPs).
		Validate(&b.MaxTargetSequences)
}

func (b *CLIConfig) IsShortHelpEnabled() bool {
	return !b.ShortHelp.IsDefault()
}

func (b *CLIConfig) SetShortHelpEnabled(val bool) {
	b.ShortHelp.Set(val)
}

func (b *CLIConfig) GetShortHelp() *field.ShortHelp {
	return &b.ShortHelp
}

func (b *CLIConfig) SetShortHelp(val field.ShortHelp) {
	b.ShortHelp = val
}

func (b *CLIConfig) IsLongHelpEnabled() bool {
	return !b.LongHelp.IsDefault()
}

func (b *CLIConfig) SetLongHelpEnabled(val bool) {
	b.LongHelp.Set(val)
}

func (b *CLIConfig) GetLongHelp() *field.LongHelp {
	return &b.LongHelp
}

func (b *CLIConfig) SetLongHelp(val field.LongHelp) {
	b.LongHelp = val
}

func (b *CLIConfig) GetFormat() *field.Format {
	return &b.Format
}

func (b *CLIConfig) SetFormat(fmt field.Format) {
	b.Format = fmt
}

func (b *CLIConfig) IsShowGIsEnabled() bool {
	return !b.ShowGIs.IsDefault()
}

func (b *CLIConfig) SetShowGIsEnabled(val bool) {
	b.ShowGIs.Set(val)
}

func (b *CLIConfig) GetShowGIs() *field.ShowGIs {
	return &b.ShowGIs
}

func (b *CLIConfig) SetShowGIs(val field.ShowGIs) {
	b.ShowGIs = val
}

func (b *CLIConfig) GetNumDescriptions() *field.NumDescriptions {
	return &b.NumDescriptions
}

func (b *CLIConfig) SetNumDescriptions(val field.NumDescriptions) {
	b.NumDescriptions = val
}

func (b *CLIConfig) GetNumAlignments() *field.NumAlignments {
	return &b.NumAlignments
}

func (b *CLIConfig) SetNumAlignments(val field.NumAlignments) {
	b.NumAlignments = val
}

func (b *CLIConfig) GetLineLength() *field.LineLength {
	return &b.LineLength
}

func (b *CLIConfig) SetLineLength(val field.LineLength) {
	b.LineLength = val
}

func (b *CLIConfig) IsHTMLOutputEnabled() bool {
	return !b.HTML.IsDefault()
}

func (b *CLIConfig) SetHTMLOutputEnabled(val bool) {
	b.HTML.Set(val)
}

func (b *CLIConfig) GetHTMLOutput() *field.HTML {
	return &b.HTML
}

func (b *CLIConfig) SetHTMLOutput(val field.HTML) {
	b.HTML = val
}

func (b *CLIConfig) GetSortHits() *field.HitSorting {
	return &b.SortHits
}

func (b *CLIConfig) SetSortHits(val field.HitSorting) {
	b.SortHits = val
}

func (b *CLIConfig) GetSortHSPs() *field.HSPSorting {
	return &b.SortHSPs
}

func (b *CLIConfig) SetSortHSPs(val field.HSPSorting) {
	b.SortHSPs = val
}

func (b *CLIConfig) GetMaxTargetSequences() *field.MaxTargetSequences {
	return &b.MaxTargetSequences
}

func (b *CLIConfig) SetMaxTargetSequences(val field.MaxTargetSequences) {
	b.MaxTargetSequences = val
}

func (b *CLIConfig) GetOutFile() *field.OutFile {
	return &b.OutFile
}

func (b *CLIConfig) SetOutFile(fmt field.OutFile) {
	b.OutFile = fmt
}

func (b *CLIConfig) IsParseDefLinesEnabled() bool {
	return !b.ParseDefLines.IsDefault()
}

func (b *CLIConfig) SetParseDefLinesEnabled(val bool) {
	b.ParseDefLines.Set(val)
}

func (b *CLIConfig) GetParseDefLines() *field.ParseDefLines {
	return &b.ParseDefLines
}

func (b *CLIConfig) SetParseDefLines(val field.ParseDefLines) {
	b.ParseDefLines = val
}
