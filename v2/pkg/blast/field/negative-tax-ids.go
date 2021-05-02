package field

import (
	"strings"

	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
)

func DecodeJSONNegativeTaxIDs(dec *gojay.Decoder, val *NegativeTaxIDs) error {
	return dec.Array(val)
}

type NegativeTaxIDs []string

func (t NegativeTaxIDs) Flag() string {
	return consts.FlagNegativeTaxIDs
}

func (t NegativeTaxIDs) IsDefault() bool {
	return len(t) == 0
}

func (t NegativeTaxIDs) FlagString() string {
	return t.Flag() + "='" + strings.Join(t, ",") + "'"
}

func (t *NegativeTaxIDs) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var tmp string
	if err := dec.String(&tmp); err != nil {
		return err
	}

	*t = append(*t, tmp)

	return nil
}
