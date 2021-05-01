package traits

import (
	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/bval"
	. "github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/field"
	"github.com/veupathdb/lib-go-blast/v2/pkg/cli"
)

// QueryConfig defines a container for shared blast tool options across the
// blast+ query tools, such as blastn, blastp, etc..
//
// In practice this type shouldn't need to be used directly, instead see one of
// the tool specific CLI configs, such as BlastN, BlastP, BlastX, etc..
type QueryConfig struct {
	CLIConfig

	QueryFile                field.QueryFile
	QueryLocation            field.QueryLocation
	DBFile                   field.DBFile
	ExpectValue              field.ExpectValue
	SoftMasking              field.SoftMasking
	LowercaseMasking         field.LowercaseMasking
	EntrezQuery              field.EntrezQuery
	QueryCoverageHSPPercent  field.QueryCoverageHSPPercent
	MaxHSPs                  field.MaxHSPs
	DBSize                   field.DBSize
	SearchSpace              field.SearchSpace
	ImportSearchStrategy     field.ImportSearchStrategy
	ExportSearchStrategy     field.ExportSearchStrategy
	ExtensionDropoffUngapped field.ExtensionDropoffUngapped
	WindowSize               field.WindowSize
	Remote                   field.Remote
}

func (b *QueryConfig) UnmarshalJSONObject(dec *gojay.Decoder, key string) (bool, error) {
	if o, e := b.CLIConfig.UnmarshalJSONObject(dec, key); o || e != nil {
		return o, e
	}

	switch key {
	case FlagQueryFile:
		return true, field.DecodeJSONQueryFile(dec, &b.QueryFile)
	case FlagQueryLocation:
		return true, field.DecodeJSONQueryLocation(dec, &b.QueryLocation)
	case FlagDBFile:
		return true, field.DecodeJSONDBFile(dec, &b.DBFile)
	case FlagExpectValue:
		return true, field.DecodeJSONExpectValue(dec, &b.ExpectValue)
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
	case FlagRemote:
		return true, field.DecodeJSONRemote(dec, &b.Remote)
	default:
		return false, nil
	}
}

func (b *QueryConfig) ToCLI(builder *cli.Builder) {
	b.CLIConfig.ToCLI(builder)
	builder.Append(&b.QueryFile).
		Append(&b.QueryLocation).
		Append(&b.DBFile).
		Append(&b.ExpectValue).
		Append(&b.SoftMasking).
		Append(&b.LowercaseMasking).
		Append(&b.EntrezQuery).
		Append(&b.QueryCoverageHSPPercent).
		Append(&b.MaxHSPs).
		Append(&b.DBSize).
		Append(&b.SearchSpace).
		Append(&b.ImportSearchStrategy).
		Append(&b.ExportSearchStrategy).
		Append(&b.ExtensionDropoffUngapped).
		Append(&b.WindowSize).
		Append(&b.Remote)
}

func (b *QueryConfig) Validate(errors bval.ValidationBuilder) {
	b.CLIConfig.Validate(errors)

	_ = errors.Validate(&b.ExpectValue).
		Requires(&b.EntrezQuery, &b.Remote).
		Validate(
			&b.QueryCoverageHSPPercent,
			&b.MaxHSPs,
			&b.SearchSpace,
		).
		Incompatible(&b.ImportSearchStrategy, &b.ExportSearchStrategy)
}

func (b *QueryConfig) GetQueryFile() *field.QueryFile {
	return &b.QueryFile
}

func (b *QueryConfig) SetQueryFile(v field.QueryFile) {
	b.QueryFile = v
}

func (b *QueryConfig) GetQueryLocation() *field.QueryLocation {
	return &b.QueryLocation
}

func (b *QueryConfig) SetQueryLocation(v field.QueryLocation) {
	b.QueryLocation = v
}

func (b *QueryConfig) GetDBFile() *field.DBFile {
	return &b.DBFile
}

func (b *QueryConfig) SetDBFile(v field.DBFile) {
	b.DBFile = v
}

func (b *QueryConfig) GetExpectValue() *field.ExpectValue {
	return &b.ExpectValue
}

func (b *QueryConfig) SetExpectValue(v field.ExpectValue) {
	b.ExpectValue = v
}

func (b *QueryConfig) IsSoftMaskingEnabled() bool {
	return !b.SoftMasking.IsDefault() && b.SoftMasking.Get()
}

func (b *QueryConfig) SetSoftMaskingEnabled(v bool) {
	b.SoftMasking.Set(v)
}

func (b *QueryConfig) GetSoftMasking() *field.SoftMasking {
	return &b.SoftMasking
}

func (b *QueryConfig) SetSoftMasking(v field.SoftMasking) {
	b.SoftMasking = v
}

func (b *QueryConfig) IsLowercaseMaskingEnabled() bool {
	return !b.LowercaseMasking.IsDefault()
}

func (b *QueryConfig) SetLowercaseMaskingEnabled(v bool) {
	b.LowercaseMasking.Set(v)
}

func (b *QueryConfig) GetLowercaseMasking() *field.LowercaseMasking {
	return &b.LowercaseMasking
}

func (b *QueryConfig) SetLowercaseMasking(v field.LowercaseMasking) {
	b.LowercaseMasking = v
}

func (b *QueryConfig) GetEntrezQuery() *field.EntrezQuery {
	return &b.EntrezQuery
}

func (b *QueryConfig) SetEntrezQuery(v field.EntrezQuery) {
	b.EntrezQuery = v
}

func (b *QueryConfig) GetQueryCoverageHSPPercent() *field.QueryCoverageHSPPercent {
	return &b.QueryCoverageHSPPercent
}

func (b *QueryConfig) SetQueryCoverageHSPPercent(v field.QueryCoverageHSPPercent) {
	b.QueryCoverageHSPPercent = v
}

func (b *QueryConfig) GetMaxHSPs() *field.MaxHSPs {
	return &b.MaxHSPs
}

func (b *QueryConfig) SetMaxHSPs(v field.MaxHSPs) {
	b.MaxHSPs = v
}

func (b *QueryConfig) GetDBSize() *field.DBSize {
	return &b.DBSize
}

func (b *QueryConfig) SetDBSize(v field.DBSize) {
	b.DBSize = v
}

func (b *QueryConfig) GetSearchSpace() *field.SearchSpace {
	return &b.SearchSpace
}

func (b *QueryConfig) SetSearchSpace(v field.SearchSpace) {
	b.SearchSpace = v
}

func (b *QueryConfig) GetImportSearchStrategy() *field.ImportSearchStrategy {
	return &b.ImportSearchStrategy
}

func (b *QueryConfig) SetImportSearchStrategy(v field.ImportSearchStrategy) {
	b.ImportSearchStrategy = v
}

func (b *QueryConfig) GetExportSearchStrategy() *field.ExportSearchStrategy {
	return &b.ExportSearchStrategy
}

func (b *QueryConfig) SetExportSearchStrategy(v field.ExportSearchStrategy) {
	b.ExportSearchStrategy = v
}

func (b *QueryConfig) GetExtensionDropoffUngapped() *field.ExtensionDropoffUngapped {
	return &b.ExtensionDropoffUngapped
}

func (b *QueryConfig) SetExtensionDropoffUngapped(v field.ExtensionDropoffUngapped) {
	b.ExtensionDropoffUngapped = v
}

func (b *QueryConfig) GetWindowSize() *field.WindowSize {
	return &b.WindowSize
}

func (b *QueryConfig) SetWindowSize(v field.WindowSize) {
	b.WindowSize = v
}

func (b *QueryConfig) IsRemoteEnabled() bool {
	return !b.Remote.IsDefault()
}

func (b *QueryConfig) SetRemoteEnabled(v bool) {
	b.Remote.Set(v)
}

func (b *QueryConfig) GetRemote() *field.Remote {
	return &b.Remote
}

func (b *QueryConfig) SetRemote(v field.Remote) {
	b.Remote = v
}
