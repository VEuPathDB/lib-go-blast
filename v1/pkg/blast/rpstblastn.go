package blast

import (
	"os/exec"

	"github.com/francoispqt/gojay"
	"lib-go-blast/v1/pkg/blast/consts"
	"lib-go-blast/v1/pkg/blast/field"
	"lib-go-blast/v1/pkg/cli"
)

type RPSTBlastN struct {
	baseBlast

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
	if v, e := r.baseBlast.UnmarshalJSONObject(dec, key); v || e != nil {
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

	r.baseBlast.ToCLI(b)

	b.Append(r.QueryGenCode).
		Append(r.Strand).
		Append(r.CompBasedStats).
		Append(r.Seg).
		Append(r.SumStats).
		Append(r.ExtensionDropoffPrelimGapped).
		Append(r.ExtensionDropoffFinalGapped).
		Append(r.UngappedAlignmentsOnly).
		Append(r.NumThreads).
		Append(r.MTMode).
		Append(r.UseSmithWatermanTraceback)

	out := new(exec.Cmd)
	out.Args = *b

	return out
}


