package traits

import (
	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/field"
	"github.com/veupathdb/lib-go-blast/v2/pkg/cli"
)

type WithBestHit struct {
	BestHitOverhang  field.BestHitOverhang
	BestHitScoreEdge field.BestHitScoreEdge
	SubjectBestHit   field.SubjectBestHit
}

func (b *WithBestHit) UnmarshalJSONObject(dec *gojay.Decoder, key string) (bool, error) {
	switch key {
	case consts.FlagBestHitOverhang:
		return true, field.DecodeJSONBestHitOverhang(dec, &b.BestHitOverhang)
	case consts.FlagBestHitScoreEdge:
		return true, field.DecodeJSONBestHitScoreEdge(dec, &b.BestHitScoreEdge)
	case consts.FlagSubjectBestHit:
		return true, field.DecodeJSONSubjectBestHit(dec, &b.SubjectBestHit)
	default:
		return false, nil
	}
}

func (b *WithBestHit) ToCLI(build *cli.Builder) {
	build.Append(&b.BestHitOverhang)
	build.Append(&b.BestHitScoreEdge)
	build.Append(&b.SubjectBestHit)
}
