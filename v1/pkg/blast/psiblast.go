package blast

import (
	"os/exec"

	"github.com/francoispqt/gojay"
	. "github.com/veupathdb/lib-go-blast/v1/pkg/blast/consts"
	"github.com/veupathdb/lib-go-blast/v1/pkg/blast/field"
	"lib-go-blast/v1/pkg/cli"
)

type PSIBlast struct {
	baseBlast
	ipgListContainer
	bestHit

	WordSize                     field.WordSize
	GapOpen                      field.GapOpen
	GapExtend                    field.GapExtend
	Matrix                       field.Matrix
	Threshold                    field.Threshold
	CompBasedStats               field.FullCompBasedStats
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
	InMSAFile                    field.InMSAFile
	MSAMasterIDX                 field.MSAMasterIDX
	IgnoreMSAMaster              field.IgnoreMSAMaster
	InPSSMFile                   field.InPSSMFile
	Pseudocount                  field.Pseudocount
	InclusionEThreshold          field.InclusionEThreshold
	PhiPatternFile               field.PhiPatternFile
}

func (*PSIBlast) Tool() Tool {
	return ToolPSIBlast
}

func (p *PSIBlast) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	if v, e := p.baseBlast.UnmarshalJSONObject(dec, key); v || e != nil {
		return e
	}
	if v, e := p.ipgListContainer.UnmarshalJSONObject(dec, key); v || e != nil {
		return e
	}
	if v, e := p.bestHit.UnmarshalJSONObject(dec, key); v || e != nil {
		return e
	}

	switch key {
	case FlagWordSize:
		return field.DecodeJSONWordSize(dec, &p.WordSize)
	case FlagGapOpen:
		return field.DecodeJSONGapOpen(dec, &p.GapOpen)
	case FlagGapExtend:
		return field.DecodeJSONGapExtend(dec, &p.GapExtend)
	case FlagMatrix:
		return field.DecodeJSONMatrix(dec, &p.Matrix)
	case FlagThreshold:
		return field.DecodeJSONThreshold(dec, &p.Threshold)
	case FlagCompBasedStats:
		return field.DecodeJSONFullCompBasedStats(dec, &p.CompBasedStats)
	case FlagSubjectFile:
		return field.DecodeJSONSubjectFile(dec, &p.SubjectFile)
	case FlagSubjectLocation:
		return field.DecodeJSONSubjectLocation(dec, &p.SubjectLocation)
	case FlagSeg:
		if v, e := field.DecodeJSONSeg(dec); e != nil {
			return e
		} else {
			p.Seg = v
			return nil
		}
	case FlagCullingLimit:
		return field.DecodeJSONCullingLimit(dec, &p.CullingLimit)
	case FlagSumStats:
		return field.DecodeJSONSumStats(dec, &p.SumStats)
	case FlagExtensionDropoffPrelimGapped:
		return field.DecodeJSONExtensionDropoffPrelimGapped(dec, &p.ExtensionDropoffPrelimGapped)
	case FlagExtensionDropoffFinalGapped:
		return field.DecodeJSONExtensionDropoffFinalGapped(dec, &p.ExtensionDropoffFinalGapped)
	case FlagGapTrigger:
		return field.DecodeJSONGapTrigger(dec, &p.GapTrigger)
	case FlagNumThreads:
		return field.DecodeJSONThreadCount(dec, &p.NumThreads)
	case FlagUseSmithWatermanTraceback:
		return field.DecodeJSONUseSmithWatermanTraceback(dec, &p.UseSmithWatermanTraceback)
	case FlagNumIterations:
		return field.DecodeJSONNumIterations(dec, &p.NumIterations)
	case FlagOutPSSMFile:
		return field.DecodeJSONOutPSSMFile(dec, &p.OutPSSMFile)
	case FlagOutASCIIPSSMFile:
		return field.DecodeJSONOutASCIIPSSMFile(dec, &p.OutASCIIPSSMFile)
	case FlagSavePSSMAfterLastRound:
		return field.DecodeJSONSavePSSMAfterLastRound(dec, &p.SavePSSMAfterLastRound)
	case FlagSaveEachPSSM:
		return field.DecodeJSONSaveEachPSSM(dec, &p.SaveEachPSSM)
	case FlagInMSAFile:
		return field.DecodeJSONInMSAFile(dec, &p.InMSAFile)
	case FlagMSAMasterIDX:
		return field.DecodeJSONMSAMasterIDX(dec, &p.MSAMasterIDX)
	case FlagIgnoreMSAMaster:
		return field.DecodeJSONIgnoreMSAMaster(dec, &p.IgnoreMSAMaster)
	case FlagInPSSMFile:
		return field.DecodeJSONInPSSMFile(dec, &p.InPSSMFile)
	case FlagPseudocount:
		return field.DecodeJSONPseudocount(dec, &p.Pseudocount)
	case FlagInclusionEThreshold:
		return field.DecodeJSONInclusionEThreshold(dec, &p.InclusionEThreshold)
	case FlagPhiPatternFile:
		return field.DecodeJSONPhiPatternFile(dec, &p.PhiPatternFile)
	default:
		return nil
	}
}

func (p *PSIBlast) NKeys() int {
	return 0
}

func (p *PSIBlast) ToCLI() *exec.Cmd {
	b := cli.NewBuilder().AppendFlag(ToolPSIBlast.String())

	p.baseBlast.ToCLI(b)
	p.ipgListContainer.ToCLI(b)
	p.bestHit.ToCLI(b)

	b.Append(p.WordSize).
		Append(p.GapOpen).
		Append(p.GapExtend).
		Append(p.Matrix).
		Append(p.Threshold).
		Append(p.CompBasedStats).
		Append(p.SubjectFile).
		Append(p.SubjectLocation).
		Append(p.Seg).
		Append(p.CullingLimit).
		Append(p.SumStats).
		Append(p.ExtensionDropoffPrelimGapped).
		Append(p.ExtensionDropoffFinalGapped).
		Append(p.GapTrigger).
		Append(p.NumThreads).
		Append(p.UseSmithWatermanTraceback).
		Append(p.NumIterations).
		Append(p.OutPSSMFile).
		Append(p.OutASCIIPSSMFile).
		Append(p.SavePSSMAfterLastRound).
		Append(p.SaveEachPSSM).
		Append(p.InMSAFile).
		Append(p.MSAMasterIDX).
		Append(p.IgnoreMSAMaster).
		Append(p.InPSSMFile).
		Append(p.Pseudocount).
		Append(p.InclusionEThreshold).
		Append(p.PhiPatternFile)

	out := new(exec.Cmd)
	out.Args = *b

	return out
}
