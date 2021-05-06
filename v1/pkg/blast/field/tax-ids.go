package field

import (
	"strings"

	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v1/pkg/blast/consts"
)

func DecodeJSONTaxIDs(dec *gojay.Decoder, val *TaxIDs) error {
	return dec.Array(val)
}

type TaxIDs []string

func (t TaxIDs) Flag() string {
	return consts.FlagTaxIDs
}

func (t TaxIDs) IsDefault() bool {
	return len(t) == 0
}

func (t TaxIDs) FlagString() string {
	return t.Flag() + "='" + strings.Join(t, ",") + "'"
}

func (t *TaxIDs) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var tmp string
	if err := dec.String(&tmp); err != nil {
		return err
	}

	*t = append(*t, tmp)

	return nil
}
