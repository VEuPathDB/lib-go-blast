package blast

import (
	"os/exec"

	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v2/internal/traits"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/bval"
	. "github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/field"
	"github.com/veupathdb/lib-go-blast/v2/pkg/cli"
)

type RPSBlast struct {
	traits.QueryConfig
	traits.WithBestHit

	CompBasedStats               field.ShortCompBasedStats
	Seg                          field.Seg
	CullingLimit                 field.CullingLimit
	SumStats                     field.SumStats
	ExtensionDropoffPrelimGapped field.ExtensionDropoffPrelimGapped
	ExtensionDropoffFinalGapped  field.ExtensionDropoffFinalGapped
	NumThreads                   field.ThreadMode
	MTMode                       field.MTMode
	UseSmithWatermanTraceback    field.UseSmithWatermanTraceback
}

func (*RPSBlast) Tool() Tool {
	return ToolRPSBlast
}

func (r *RPSBlast) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	if v, e := r.QueryConfig.UnmarshalJSONObject(dec, key); v || e != nil {
		return e
	}
	if v, e := r.WithBestHit.UnmarshalJSONObject(dec, key); v || e != nil {
		return e
	}

	switch key {
	case FlagCompBasedStats:
		return field.DecodeJSONShortCompBasedStats(dec, &r.CompBasedStats)
	case FlagSeg:
		if v, e := field.DecodeJSONSeg(dec); e != nil {
			return e
		} else {
			r.Seg = v
			return nil
		}
	case FlagSumStats:
		return field.DecodeJSONSumStats(dec, &r.SumStats)
	case FlagExtensionDropoffPrelimGapped:
		return field.DecodeJSONExtensionDropoffPrelimGapped(dec, &r.ExtensionDropoffPrelimGapped)
	case FlagExtensionDropoffFinalGapped:
		return field.DecodeJSONExtensionDropoffFinalGapped(dec, &r.ExtensionDropoffFinalGapped)
	case FlagNumThreads:
		return field.DecodeJSONThreadMode(dec, &r.NumThreads)
	case FlagMTMode:
		return field.DecodeJSONMTMode(dec, &r.MTMode)
	case FlagUseSmithWatermanTraceback:
		return field.DecodeJSONUseSmithWatermanTraceback(dec, &r.UseSmithWatermanTraceback)
	default:
		return nil
	}
}

func (r *RPSBlast) NKeys() int {
	return 0
}

func (r *RPSBlast) ToCLI() *exec.Cmd {
	b := cli.NewBuilder().AppendFlag(ToolRPSBlast.String())

	r.QueryConfig.ToCLI(b)
	r.WithBestHit.ToCLI(b)

	b.Append(&r.CompBasedStats).
		Append(r.Seg).
		Append(&r.SumStats).
		Append(&r.ExtensionDropoffPrelimGapped).
		Append(&r.ExtensionDropoffFinalGapped).
		Append(&r.NumThreads).
		Append(&r.MTMode).
		Append(&r.UseSmithWatermanTraceback)

	out := new(exec.Cmd)
	out.Args = *b

	return out
}

func (r *RPSBlast) Validate() bval.ValidationError {
	em := make(bval.ValidationBuilder)

	em.Validate(&r.QueryConfig, &r.CompBasedStats, &r.NumThreads, &r.MTMode)

	if r.Format.Type == field.FormatTypeSequenceAlignmentMap {
		em.AppendError(r.Format.Flag(), "Format type 17 is not supported with rpsblast")
	}

	em.IncompatibleWith(&r.CullingLimit, &r.BestHitOverhang, &r.BestHitScoreEdge)

	if len(em) == 0 {
		return nil
	}

	return bval.ValidationError(em)
}
