package blast

import (
	"os/exec"

	"github.com/francoispqt/gojay"
	. "github.com/veupathdb/lib-go-blast/v1/pkg/blast/consts"
	"github.com/veupathdb/lib-go-blast/v1/pkg/blast/field"
	"lib-go-blast/v1/pkg/cli"
)

type DeltaBlast struct {
	baseBlast
	baseListContainer
	bestHit

	WordSize                     field.WordSize
	GapOpen                      field.GapOpen
	GapExtend                    field.GapExtend
	Matrix                       field.Matrix
	Threshold                    field.Threshold
	CompBasedStats               field.ShortCompBasedStats
	SubjectFile                  field.SubjectFile
	SubjectLocation              field.SubjectLocation
	Seg                          field.Seg
	CullingLimit                 field.CullingLimit
	SumStats                     field.SumStats
	ExtensionDropoffPrelimGapped field.ExtensionDropoffPrelimGapped
	ExtensionDropoffFinalGapped  field.ExtensionDropoffFinalGapped
	GapTrigger                   field.GapTrigger
	NumThreads                   field.ThreadCount
	UseSmithWatermanTraceback    field.UseSmithWatermanTraceback
	NumIterations                field.NumIterations
	OutPSSMFile                  field.OutPSSMFile
	OutASCIIPSSMFile             field.OutASCIIPSSMFile
	SavePSSMAfterLastRound       field.SavePSSMAfterLastRound
	SaveEachPSSM                 field.SaveEachPSSM
	Pseudocount                  field.Pseudocount
	DomainInclusionEThreshold    field.DomainInclusionEThreshold
	InclusionEThreshold          field.InclusionEThreshold
	RPSDBFile                    field.RPSDBFile
	ShowDomainHits               field.ShowDomainHits
}

func (d *DeltaBlast) Tool() Tool {
	return ToolDeltaBlast
}

func (d *DeltaBlast) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	if val, err := d.baseBlast.UnmarshalJSONObject(dec, key); err != nil || val {
		return err
	}
	if val, err := d.baseListContainer.UnmarshalJSONObject(dec, key); err != nil || val {
		return err
	}
	if val, err := d.bestHit.UnmarshalJSONObject(dec, key); err != nil || val {
		return err
	}

	switch key {
	case FlagWordSize:
		return field.DecodeJSONWordSize(dec, &d.WordSize)
	case FlagGapOpen:
		return field.DecodeJSONGapOpen(dec, &d.GapOpen)
	case FlagGapExtend:
		return field.DecodeJSONGapExtend(dec, &d.GapExtend)
	case FlagMatrix:
		return field.DecodeJSONMatrix(dec, &d.Matrix)
	case FlagThreshold:
		return field.DecodeJSONThreshold(dec, &d.Threshold)
	case FlagCompBasedStats:
		return field.DecodeJSONShortCompBasedStats(dec, &d.CompBasedStats)
	case FlagSubjectFile:
		return field.DecodeJSONSubjectFile(dec, &d.SubjectFile)
	case FlagSubjectLocation:
		return field.DecodeJSONSubjectLocation(dec, &d.SubjectLocation)
	case FlagSeg:
		if tmp, err := field.DecodeJSONSeg(dec); err != nil {
			return err
		} else {
			d.Seg = tmp
			return nil
		}
	case FlagCullingLimit:
		return field.DecodeJSONCullingLimit(dec, &d.CullingLimit)
	case FlagSumStats:
		return field.DecodeJSONSumStats(dec, &d.SumStats)
	case FlagExtensionDropoffPrelimGapped:
		return field.DecodeJSONExtensionDropoffPrelimGapped(dec, &d.ExtensionDropoffPrelimGapped)
	case FlagExtensionDropoffFinalGapped:
		return field.DecodeJSONExtensionDropoffFinalGapped(dec, &d.ExtensionDropoffFinalGapped)
	case FlagGapTrigger:
		return field.DecodeJSONGapTrigger(dec, &d.GapTrigger)
	case FlagNumThreads:
		return field.DecodeJSONThreadCount(dec, &d.NumThreads)
	case FlagUseSmithWatermanTraceback:
		return field.DecodeJSONUseSmithWatermanTraceback(dec, &d.UseSmithWatermanTraceback)
	case FlagNumIterations:
		return field.DecodeJSONNumIterations(dec, &d.NumIterations)
	case FlagOutPSSMFile:
		return field.DecodeJSONOutPSSMFile(dec, &d.OutPSSMFile)
	case FlagOutASCIIPSSMFile:
		return field.DecodeJSONOutASCIIPSSMFile(dec, &d.OutASCIIPSSMFile)
	case FlagSavePSSMAfterLastRound:
		return field.DecodeJSONSavePSSMAfterLastRound(dec, &d.SavePSSMAfterLastRound)
	case FlagSaveEachPSSM:
		return field.DecodeJSONSaveEachPSSM(dec, &d.SaveEachPSSM)
	case FlagPseudocount:
		return field.DecodeJSONPseudocount(dec, &d.Pseudocount)
	case FlagDomainInclusionEThreshold:
		return field.DecodeJSONDomainInclusionEThreshold(dec, &d.DomainInclusionEThreshold)
	case FlagInclusionEThreshold:
		return field.DecodeJSONInclusionEThreshold(dec, &d.InclusionEThreshold)
	case FlagRPSDBFile:
		return field.DecodeJSONRPSDBFile(dec, &d.RPSDBFile)
	case FlagShowDomainHits:
		return field.DecodeJSONShowDomainHits(dec, &d.ShowDomainHits)
	default:
		return nil
	}
}

func (d *DeltaBlast) NKeys() int {
	return 0
}

func (d *DeltaBlast) ToCLI() *exec.Cmd {
	var builder = cli.NewBuilder().AppendFlag(ToolDeltaBlast.String())

	d.baseBlast.ToCLI(builder)
	d.baseListContainer.ToCLI(builder)
	d.bestHit.ToCLI(builder)

	builder.
		Append(d.WordSize).
		Append(d.GapOpen).
		Append(d.GapExtend).
		Append(d.Matrix).
		Append(d.Threshold).
		Append(d.CompBasedStats).
		Append(d.SubjectFile).
		Append(d.SubjectLocation).
		Append(d.Seg).
		Append(d.CullingLimit).
		Append(d.SumStats).
		Append(d.ExtensionDropoffPrelimGapped).
		Append(d.ExtensionDropoffFinalGapped).
		Append(d.GapTrigger).
		Append(d.NumThreads).
		Append(d.UseSmithWatermanTraceback).
		Append(d.NumIterations).
		Append(d.OutPSSMFile).
		Append(d.OutASCIIPSSMFile).
		Append(d.SavePSSMAfterLastRound).
		Append(d.SaveEachPSSM).
		Append(d.Pseudocount).
		Append(d.DomainInclusionEThreshold).
		Append(d.InclusionEThreshold).
		Append(d.RPSDBFile).
		Append(d.ShowDomainHits)

	out := new(exec.Cmd)
	out.Args = *builder

	return out
}
