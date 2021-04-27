package blast

import (
	"os/exec"

	"github.com/francoispqt/gojay"
	. "lib-go-blast/v1/pkg/blast/consts"
	"lib-go-blast/v1/pkg/blast/field"
	"lib-go-blast/v1/pkg/cli"
)

type RPSBlast struct {
	baseBlast
	bestHit

	CompBasedStats               field.ShortCompBasedStats
	Seg                          field.Seg
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
	if v, e := r.baseBlast.UnmarshalJSONObject(dec, key); v || e != nil {
		return e
	}
	if v, e := r.bestHit.UnmarshalJSONObject(dec, key); v || e != nil {
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

	r.baseBlast.ToCLI(b)
	r.bestHit.ToCLI(b)

	b.Append(r.CompBasedStats).
		Append(r.Seg).
		Append(r.SumStats).
		Append(r.ExtensionDropoffPrelimGapped).
		Append(r.ExtensionDropoffFinalGapped).
		Append(r.NumThreads).
		Append(r.MTMode).
		Append(r.UseSmithWatermanTraceback)

	out := new(exec.Cmd)
	out.Args = *b

	return out
}
