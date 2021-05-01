package blast

import (
	"os/exec"

	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v2/internal/traits"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/bval"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/field"
	"github.com/veupathdb/lib-go-blast/v2/pkg/cli"
)

type RPSTBlastN struct {
	traits.QueryConfig

	QueryGenCode                 field.QueryGenCode
	Strand                       field.Strand
	CompBasedStats               field.ShortCompBasedStats
	Seg                          field.Seg
	SumStats                     field.SumStats
	ExtensionDropoffPrelimGapped field.ExtensionDropoffPrelimGapped
	ExtensionDropoffFinalGapped  field.ExtensionDropoffFinalGapped
	UngappedAlignmentsOnly       field.UngappedAlignmentsOnly
	NumThreads                   field.ThreadMode
	MTMode                       field.MTMode
	UseSmithWatermanTraceback    field.UseSmithWatermanTraceback
}

func (*RPSTBlastN) Tool() Tool {
	return ToolRPSTBlastN
}

func (r *RPSTBlastN) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	if v, e := r.QueryConfig.UnmarshalJSONObject(dec, key); v || e != nil {
		return e
	}

	switch key {
	case consts.FlagQueryGenCode:
		return field.DecodeJSONQueryGenCode(dec, &r.QueryGenCode)
	case consts.FlagStrand:
		return field.DecodeJSONStrand(dec, &r.Strand)
	case consts.FlagCompBasedStats:
		return field.DecodeJSONShortCompBasedStats(dec, &r.CompBasedStats)
	case consts.FlagSeg:
		if v, e := field.DecodeJSONSeg(dec); e != nil {
			return e
		} else {
			r.Seg = v
			return nil
		}
	case consts.FlagSumStats:
		return field.DecodeJSONSumStats(dec, &r.SumStats)
	case consts.FlagExtensionDropoffPrelimGapped:
		return field.DecodeJSONExtensionDropoffPrelimGapped(dec, &r.ExtensionDropoffPrelimGapped)
	case consts.FlagExtensionDropoffFinalGapped:
		return field.DecodeJSONExtensionDropoffFinalGapped(dec, &r.ExtensionDropoffFinalGapped)
	case consts.FlagUngappedAlignmentsOnly:
		return field.DecodeJSONUngappedAlignmentsOnly(dec, &r.UngappedAlignmentsOnly)
	case consts.FlagNumThreads:
		return field.DecodeJSONThreadMode(dec, &r.NumThreads)
	case consts.FlagMTMode:
		return field.DecodeJSONMTMode(dec, &r.MTMode)
	case consts.FlagUseSmithWatermanTraceback:
		return field.DecodeJSONUseSmithWatermanTraceback(dec, &r.UseSmithWatermanTraceback)
	default:
		return nil
	}
}

func (r *RPSTBlastN) NKeys() int {
	return 0
}

func (r *RPSTBlastN) ToCLI() *exec.Cmd {
	b := cli.NewBuilder().AppendFlag(ToolRPSTBlastN.String())

	r.QueryConfig.ToCLI(b)

	b.Append(&r.QueryGenCode).
		Append(&r.Strand).
		Append(&r.CompBasedStats).
		Append(r.Seg).
		Append(&r.SumStats).
		Append(&r.ExtensionDropoffPrelimGapped).
		Append(&r.ExtensionDropoffFinalGapped).
		Append(&r.UngappedAlignmentsOnly).
		Append(&r.NumThreads).
		Append(&r.MTMode).
		Append(&r.UseSmithWatermanTraceback)

	out := new(exec.Cmd)
	out.Args = *b

	return out
}

func (r *RPSTBlastN) Validate() bval.ValidationError {
	em := make(bval.ValidationBuilder)

	em.Validate(
		&r.QueryConfig,
		&r.QueryGenCode,
		&r.Strand,
		&r.CompBasedStats,
		&r.NumThreads,
		&r.MTMode,
	)

	if r.Format.Type == field.FormatTypeSequenceAlignmentMap {
		em.AppendError(r.Format.Flag(), "Format type 17 is not supported with rpstblastn")
	}

	if len(em) == 0 {
		return nil
	}

	return bval.ValidationError(em)
}
