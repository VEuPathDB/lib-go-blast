package blast

import (
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/field"
	"github.com/veupathdb/lib-go-blast/v2/pkg/cli"
)

type BlastConfig interface {
	cli.CommandSerializable

	// Tool returns the specific BlastTool that this config is for.
	Tool() Tool

	// IsShortHelpEnabled returns whether the ShortHelp flag enabled.
	IsShortHelpEnabled() bool

	// SetShortHelpEnabled sets whether the ShortHelp flag is enabled.
	SetShortHelpEnabled(val bool)

	// GetShortHelp returns the ShortHelp field itself.
	GetShortHelp() *field.ShortHelp

	// SetShortHelp overwrites the ShortHelp field itself.
	SetShortHelp(val field.ShortHelp)

	// IsLongHelpEnabled returns whether the LongFlag field is set to true.
	IsLongHelpEnabled() bool

	// SetLongHelpEnabled sets the longform help flag to the given value.
	SetLongHelpEnabled(val bool)

	// GetLongHelp returns the LongHelp field itself.
	GetLongHelp() *field.LongHelp

	// SetLongHelp overwrites the LongHelp field itself.
	SetLongHelp(val field.LongHelp)

	GetFormat() *field.Format
	SetFormat(fmt field.Format)

	IsShowGIsEnabled() bool
	SetShowGIsEnabled(val bool)
	GetShowGIs() *field.ShowGIs
	SetShowGIs(val field.ShowGIs)

	GetNumDescriptions() *field.NumDescriptions
	SetNumDescriptions(val field.NumDescriptions)

	GetNumAlignments() *field.NumAlignments
	SetNumAlignments(val field.NumAlignments)

	GetLineLength() *field.LineLength
	SetLineLength(val field.LineLength)

	IsHTMLOutputEnabled() bool
	SetHTMLOutputEnabled(val bool)
	GetHTMLOutput() *field.HTML
	SetHTMLOutput(val field.HTML)

	GetSortHits() *field.HitSorting
	SetSortHits(val field.HitSorting)

	GetSortHSPs() *field.HSPSorting
	SetSortHSPs(val field.HSPSorting)

	GetMaxTargetSequences() *field.MaxTargetSequences
	SetMaxTargetSequences(val field.MaxTargetSequences)

	GetOutFile() *field.OutFile
	SetOutFile(fmt field.OutFile)

	IsParseDefLinesEnabled() bool
	SetParseDefLinesEnabled(val bool)
	GetParseDefLines() *field.ParseDefLines
	SetParseDefLines(val field.ParseDefLines)
}

type BlastQueryConfig interface {
	BlastConfig

	GetQueryFile() *field.QueryFile
	SetQueryFile(v field.QueryFile)

	GetQueryLocation() *field.QueryLocation
	SetQueryLocation(v field.QueryLocation)

	GetDBFile() *field.DBFile
	SetDBFile(v field.DBFile)

	GetExpectValue() *field.ExpectValue
	SetExpectValue(v field.ExpectValue)

	IsSoftMaskingEnabled() bool
	SetSoftMaskingEnabled(v bool)
	GetSoftMasking() *field.SoftMasking
	SetSoftMasking(v field.SoftMasking)

	IsLowercaseMaskingEnabled() bool
	SetLowercaseMaskingEnabled(v bool)
	GetLowercaseMasking() *field.LowercaseMasking
	SetLowercaseMasking(v field.LowercaseMasking)

	GetEntrezQuery() *field.EntrezQuery
	SetEntrezQuery(v field.EntrezQuery)

	GetQueryCoverageHSPPercent() *field.QueryCoverageHSPPercent
	SetQueryCoverageHSPPercent(v field.QueryCoverageHSPPercent)

	GetMaxHSPs() *field.MaxHSPs
	SetMaxHSPs(v field.MaxHSPs)

	GetDBSize() *field.DBSize
	SetDBSize(v field.DBSize)

	GetSearchSpace() *field.SearchSpace
	SetSearchSpace(v field.SearchSpace)

	GetImportSearchStrategy() *field.ImportSearchStrategy
	SetImportSearchStrategy(v field.ImportSearchStrategy)

	GetExportSearchStrategy() *field.ExportSearchStrategy
	SetExportSearchStrategy(v field.ExportSearchStrategy)

	GetExtensionDropoffUngapped() *field.ExtensionDropoffUngapped
	SetExtensionDropoffUngapped(v field.ExtensionDropoffUngapped)

	GetWindowSize() *field.WindowSize
	SetWindowSize(v field.WindowSize)

	IsRemoteEnabled() bool
	SetRemoteEnabled(v bool)
	GetRemote() *field.Remote
	SetRemote(v field.Remote)
}
