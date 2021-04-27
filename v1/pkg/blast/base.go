package blast

import (
	"github.com/francoispqt/gojay"

	. "github.com/veupathdb/lib-go-blast/v1/pkg/blast/consts"
	"github.com/veupathdb/lib-go-blast/v1/pkg/blast/field"
	"lib-go-blast/v1/pkg/cli"
)

type BlastConfig interface {
	cli.CommandSerializable

	Tool() Tool

	SetQuery(query string)

	SetFormat(fmt field.Format)

	SetOut(out string)
}

type baseBlast struct {
	ShortHelp                field.ShortHelp
	LongHelp                 field.LongHelp
	Version                  field.Version
	QueryFile                field.QueryFile
	QueryLocation            field.QueryLocation
	DBFile                   field.DBFile
	OutFile                  field.OutFile
	ExpectValue              field.ExpectValue
	OutFormat                field.Format
	ShowGIs                  field.ShowGIs
	NumDescriptions          field.NumDescriptions
	NumAlignments            field.NumAlignments
	LineLength               field.LineLength
	HTML                     field.HTML
	SortHits                 field.HitSorting
	SortHSPs                 field.HSPSorting
	SoftMasking              field.SoftMasking
	LowercaseMasking         field.LowercaseMasking
	EntrezQuery              field.EntrezQuery
	QueryCoverageHSPPercent  field.QueryCoverageHSPPercent
	MaxHSPs                  field.MaxHSPs
	MaxTargetSequences       field.MaxTargetSequences
	DBSize                   field.DBSize
	SearchSpace              field.SearchSpace
	ImportSearchStrategy     field.ImportSearchStrategy
	ExportSearchStrategy     field.ExportSearchStrategy
	ExtensionDropoffUngapped field.ExtensionDropoffUngapped
	WindowSize               field.WindowSize
	ParseDefLines            field.ParseDefLines
	Remote                   field.Remote
}

func (b *baseBlast) SetQuery(q string) {
	b.QueryFile.Set(q)
}

func (b *baseBlast) SetFormat(f field.Format) {
	b.OutFormat = f
}

func (b *baseBlast) SetOut(out string) {
	b.OutFile.Set(out)
}

func (b *baseBlast) UnmarshalJSONObject(dec *gojay.Decoder, key string) (bool, error) {
	switch key {
	case FlagShortHelp:
		return true, field.DecodeJSONShortHelp(dec, &b.ShortHelp)
	case FlagLongHelp:
		return true, field.DecodeJSONLongHelp(dec, &b.LongHelp)
	case FlagVersion:
		return true, field.DecodeJSONVersion(dec, &b.Version)
	case FlagQueryFile:
		return true, field.DecodeJSONQueryFile(dec, &b.QueryFile)
	case FlagQueryLocation:
		return true, field.DecodeJSONQueryLocation(dec, &b.QueryLocation)
	case FlagDBFile:
		return true, field.DecodeJSONDBFile(dec, &b.DBFile)
	case FlagOutFile:
		return true, field.DecodeJSONOutFile(dec, &b.OutFile)
	case FlagExpectValue:
		return true, field.DecodeJSONExpectValue(dec, &b.ExpectValue)
	case FlagOutFormat:
		return true, field.DecodeJSONOutFormat(dec, &b.OutFormat)
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
	case FlagSoftMasking:
		return true, field.DecodeJSONSoftMasking(dec, &b.SoftMasking)
	case FlagLowercaseMasking:
		return true, field.DecodeJSONLowercaseMasking(dec, &b.LowercaseMasking)
	case FlagEntrezQuery:
		return true, field.DecodeJSONEntrezQuery(dec, &b.EntrezQuery)
	case FlagQueryCoverageHSPPercent:
		return true, field.DecodeJSONQueryCoverageHSPPercent(dec, &b.QueryCoverageHSPPercent)
	case FlagMaxHSPs:
		return true, field.DecodeJSONMaxHSPs(dec, &b.MaxHSPs)
	case FlagMaxTargetSequences:
		return true, field.DecodeJSONMaxTargetSequences(dec, &b.MaxTargetSequences)
	case FlagDBSize:
		return true, field.DecodeJSONDBSize(dec, &b.DBSize)
	case FlagSearchSpace:
		return true, field.DecodeJSONSearchSpace(dec, &b.SearchSpace)
	case FlagImportSearchStrategy:
		return true, field.DecodeJSONImportSearchStrategy(dec, &b.ImportSearchStrategy)
	case FlagExportSearchStrategy:
		return true, field.DecodeJSONExportSearchStrategy(dec, &b.ExportSearchStrategy)
	case FlagExtensionDropoffUngapped:
		return true, field.DecodeJSONExtensionDropoffUngapped(dec, &b.ExtensionDropoffUngapped)
	case FlagWindowSize:
		return true, field.DecodeJSONWindowSize(dec, &b.WindowSize)
	case FlagParseDefLines:
		return true, field.DecodeJSONParseDefLines(dec, &b.ParseDefLines)
	case FlagRemote:
		return true, field.DecodeJSONRemote(dec, &b.Remote)
	default:
		return false, nil
	}
}

func (b *baseBlast) ToCLI(builder *cli.Builder) {
	builder.Append(b.ShortHelp).
		Append(b.LongHelp).
		Append(b.Version).
		Append(b.QueryFile).
		Append(b.QueryLocation).
		Append(b.DBFile).
		Append(b.OutFile).
		Append(b.ExpectValue).
		Append(b.OutFormat).
		Append(b.ShowGIs).
		Append(b.NumDescriptions).
		Append(b.NumAlignments).
		Append(b.LineLength).
		Append(b.HTML).
		Append(b.SortHits).
		Append(b.SortHSPs).
		Append(b.SoftMasking).
		Append(b.LowercaseMasking).
		Append(b.EntrezQuery).
		Append(b.QueryCoverageHSPPercent).
		Append(b.MaxHSPs).
		Append(b.MaxTargetSequences).
		Append(b.DBSize).
		Append(b.SearchSpace).
		Append(b.ImportSearchStrategy).
		Append(b.ExportSearchStrategy).
		Append(b.ExtensionDropoffUngapped).
		Append(b.WindowSize).
		Append(b.ParseDefLines).
		Append(b.Remote)
}

type baseListContainer struct {
	GIList                 field.GIList
	SequenceIDList         field.SequenceIDList
	NegativeGIList         field.NegativeGIList
	NegativeSequenceIDList field.NegativeSequenceIDList
	TaxIDs                 field.TaxIDs
	NegativeTaxIDs         field.NegativeTaxIDs
	TaxIDList              field.TaxIDList
	NegativeTaxIDList      field.NegativeTaxIDList
}

func (b *baseListContainer) UnmarshalJSONObject(dec *gojay.Decoder, key string) (bool, error) {
	switch key {
	case FlagGIList:
		return true, field.DecodeJSONGIList(dec, &b.GIList)
	case FlagSequenceIDList:
		return true, field.DecodeJSONSequenceIDList(dec, &b.SequenceIDList)
	case FlagNegativeGIList:
		return true, field.DecodeJSONNegativeGIList(dec, &b.NegativeGIList)
	case FlagNegativeSequenceIDList:
		return true, field.DecodeJSONNegativeSequenceIDList(dec, &b.NegativeSequenceIDList)
	case FlagTaxIDs:
		return true, field.DecodeJSONTaxIDs(dec, &b.TaxIDs)
	case FlagNegativeTaxIDs:
		return true, field.DecodeJSONNegativeTaxIDs(dec, &b.NegativeTaxIDs)
	case FlagTaxIDList:
		return true, field.DecodeJSONTaxIDList(dec, &b.TaxIDList)
	case FlagNegativeTaxIDList:
		return true, field.DecodeJSONNegativeTaxIDList(dec, &b.NegativeTaxIDList)
	default:
		return false, nil
	}
}

func (b *baseListContainer) ToCLI(builder *cli.Builder) {
	builder.
		Append(b.GIList).
		Append(b.SequenceIDList).
		Append(b.NegativeGIList).
		Append(b.NegativeSequenceIDList).
		Append(b.TaxIDs).
		Append(b.NegativeTaxIDs).
		Append(b.TaxIDList).
		Append(b.NegativeTaxIDList)
}

type ipgListContainer struct {
	baseListContainer

	IPGList         field.IPGList
	NegativeIPGList field.NegativeIPGList
}

func (b *ipgListContainer) UnmarshalJSONObject(dec *gojay.Decoder, key string) (bool, error) {
	if tmp, err := b.baseListContainer.UnmarshalJSONObject(dec, key); err != nil || tmp {
		return tmp, err
	}

	switch key {
	case FlagIPGList:
		return true, field.DecodeJSONIPGList(dec, &b.IPGList)
	case FlagNegativeIPGList:
		return true, field.DecodeJSONNegativeIPGList(dec, &b.NegativeIPGList)
	default:
		return false, nil
	}
}

func (b *ipgListContainer) ToCLI(build *cli.Builder) {
	b.baseListContainer.ToCLI(build)
	build.Append(b.IPGList)
	build.Append(b.NegativeIPGList)
}

type bestHit struct {
	BestHitOverhang  field.BestHitOverhang
	BestHitScoreEdge field.BestHitScoreEdge
	SubjectBestHit   field.SubjectBestHit
}

func (b *bestHit) UnmarshalJSONObject(dec *gojay.Decoder, key string) (bool, error) {
	switch key {
	case FlagBestHitOverhang:
		return true, field.DecodeJSONBestHitOverhang(dec, &b.BestHitOverhang)
	case FlagBestHitScoreEdge:
		return true, field.DecodeJSONBestHitScoreEdge(dec, &b.BestHitScoreEdge)
	case FlagSubjectBestHit:
		return true, field.DecodeJSONSubjectBestHit(dec, &b.SubjectBestHit)
	default:
		return false, nil
	}
}

func (b *bestHit) ToCLI(build *cli.Builder) {
	build.Append(b.BestHitOverhang)
	build.Append(b.BestHitScoreEdge)
	build.Append(b.SubjectBestHit)
}
