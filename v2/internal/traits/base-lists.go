package traits

import (
	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/bval"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/field"
	"github.com/veupathdb/lib-go-blast/v2/pkg/cli"
)

type WithListContainer struct {
	GIList                 field.GIList
	SequenceIDList         field.SequenceIDList
	NegativeGIList         field.NegativeGIList
	NegativeSequenceIDList field.NegativeSequenceIDList
	TaxIDs                 field.TaxIDs
	NegativeTaxIDs         field.NegativeTaxIDs
	TaxIDList              field.TaxIDList
	NegativeTaxIDList      field.NegativeTaxIDList
}

func (b *WithListContainer) UnmarshalJSONObject(dec *gojay.Decoder, key string) (bool, error) {
	switch key {
	case consts.FlagGIList:
		return true, field.DecodeJSONGIList(dec, &b.GIList)
	case consts.FlagSequenceIDList:
		return true, field.DecodeJSONSequenceIDList(dec, &b.SequenceIDList)
	case consts.FlagNegativeGIList:
		return true, field.DecodeJSONNegativeGIList(dec, &b.NegativeGIList)
	case consts.FlagNegativeSequenceIDList:
		return true, field.DecodeJSONNegativeSequenceIDList(dec, &b.NegativeSequenceIDList)
	case consts.FlagTaxIDs:
		return true, field.DecodeJSONTaxIDs(dec, &b.TaxIDs)
	case consts.FlagNegativeTaxIDs:
		return true, field.DecodeJSONNegativeTaxIDs(dec, &b.NegativeTaxIDs)
	case consts.FlagTaxIDList:
		return true, field.DecodeJSONTaxIDList(dec, &b.TaxIDList)
	case consts.FlagNegativeTaxIDList:
		return true, field.DecodeJSONNegativeTaxIDList(dec, &b.NegativeTaxIDList)
	default:
		return false, nil
	}
}

func (b *WithListContainer) ToCLI(builder *cli.Builder) {
	builder.
		Append(&b.GIList).
		Append(&b.SequenceIDList).
		Append(&b.NegativeGIList).
		Append(&b.NegativeSequenceIDList).
		Append(&b.TaxIDs).
		Append(&b.NegativeTaxIDs).
		Append(&b.TaxIDList).
		Append(&b.NegativeTaxIDList)
}

func (b *WithListContainer) Validate(errors bval.ValidationBuilder) {
	_ = errors.IncompatibleList(
		&b.GIList,
		&b.NegativeGIList,
		&b.SequenceIDList,
		&b.NegativeSequenceIDList,
		&b.TaxIDs,
		&b.NegativeTaxIDs,
		&b.TaxIDList,
		&b.NegativeTaxIDList,
	)
}

type WithIPGListContainer struct {
	WithListContainer

	IPGList         field.IPGList
	NegativeIPGList field.NegativeIPGList
}

func (b *WithIPGListContainer) UnmarshalJSONObject(dec *gojay.Decoder, key string) (bool, error) {
	if tmp, err := b.WithListContainer.UnmarshalJSONObject(dec, key); err != nil || tmp {
		return tmp, err
	}

	switch key {
	case consts.FlagIPGList:
		return true, field.DecodeJSONIPGList(dec, &b.IPGList)
	case consts.FlagNegativeIPGList:
		return true, field.DecodeJSONNegativeIPGList(dec, &b.NegativeIPGList)
	default:
		return false, nil
	}
}

func (b *WithIPGListContainer) ToCLI(build *cli.Builder) {
	b.WithListContainer.ToCLI(build)
	build.Append(&b.IPGList)
	build.Append(&b.NegativeIPGList)
}

func (b *WithIPGListContainer) Validate(errors bval.ValidationBuilder) {
	b.WithListContainer.Validate(errors)
}
