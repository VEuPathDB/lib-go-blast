package blast

import (
	"os/exec"

	"github.com/francoispqt/gojay"
	. "lib-go-blast/v1/pkg/blast/consts"
	"lib-go-blast/v1/pkg/blast/field"
	"lib-go-blast/v1/pkg/cli"
)

type BlastX struct {
	baseBlast
	ipgListContainer
	bestHit

	Strand                       field.Strand
	QueryGenCode                 field.QueryGenCode
	Task                         field.BlastXTask
	WordSize                     field.WordSize
	GapOpen                      field.GapOpen
	GapExtend                    field.GapExtend
	MaxIntronLength              field.MaxIntronLength
	Matrix                       field.Matrix
	Threshold                    field.Threshold
	CompBasedStats               field.FullCompBasedStats
	SubjectFile                  field.SubjectFile
	SubjectLocation              field.SubjectLocation
	Seg                          field.Seg
	DBSoftMask                   field.DBSoftMask
	DBHardMask                   field.DBHardMask
	CullingLimit                 field.CullingLimit
	SumStats                     field.SumStats
	ExtensionDropoffPrelimGapped field.ExtensionDropoffPrelimGapped
	ExtensionDropoffFinalGapped  field.ExtensionDropoffFinalGapped
	UngappedAlignmentsOnly       field.UngappedAlignmentsOnly
	NumThreads                   field.ThreadCount
	UseSmithWatermanTraceback    field.UseSmithWatermanTraceback
}

func (b *BlastX) Tool() Tool {
	return ToolBlastX
}

func (b *BlastX) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	if tmp, err := b.baseBlast.UnmarshalJSONObject(dec, key); err != nil || tmp {
		return err
	}

	if tmp, err := b.ipgListContainer.UnmarshalJSONObject(dec, key); err != nil || tmp {
		return err
	}

	if tmp, err := b.bestHit.UnmarshalJSONObject(dec, key); err != nil || tmp {
		return err
	}

	switch key {
	case FlagStrand:
		return field.DecodeJSONStrand(dec, &b.Strand)
	case FlagQueryGenCode:
		return field.DecodeJSONQueryGenCode(dec, &b.QueryGenCode)
	case FlagTask:
		return field.DecodeJSONBlastXTask(dec, &b.Task)
	case FlagWordSize:
		return field.DecodeJSONWordSize(dec, &b.WordSize)
	case FlagGapOpen:
		return field.DecodeJSONGapOpen(dec, &b.GapOpen)
	case FlagGapExtend:
		return field.DecodeJSONGapExtend(dec, &b.GapExtend)
	case FlagMaxIntronLength:
		return field.DecodeJSONMaxIntronLength(dec, &b.MaxIntronLength)
	case FlagMatrix:
		return field.DecodeJSONMatrix(dec, &b.Matrix)
	case FlagThreshold:
		return field.DecodeJSONThreshold(dec, &b.Threshold)
	case FlagCompBasedStats:
		return field.DecodeJSONFullCompBasedStats(dec, &b.CompBasedStats)
	case FlagSubjectFile:
		return field.DecodeJSONSubjectFile(dec, &b.SubjectFile)
	case FlagSubjectLocation:
		return field.DecodeJSONSubjectLocation(dec, &b.SubjectLocation)
	case FlagSeg:
		if tmp, err := field.DecodeJSONSeg(dec); err != nil {
			return err
		} else {
			b.Seg = tmp
			return nil
		}
	case FlagDBSoftMask:
		return field.DecodeJSONDBSoftMask(dec, &b.DBSoftMask)
	case FlagDBHardMask:
		return field.DecodeJSONDBHardMask(dec, &b.DBHardMask)
	case FlagCullingLimit:
		return field.DecodeJSONCullingLimit(dec, &b.CullingLimit)
	case FlagSumStats:
		return field.DecodeJSONSumStats(dec, &b.SumStats)
	case FlagExtensionDropoffPrelimGapped:
		return field.DecodeJSONExtensionDropoffPrelimGapped(dec, &b.ExtensionDropoffPrelimGapped)
	case FlagExtensionDropoffFinalGapped:
		return field.DecodeJSONExtensionDropoffFinalGapped(dec, &b.ExtensionDropoffFinalGapped)
	case FlagUngappedAlignmentsOnly:
		return field.DecodeJSONUngappedAlignmentsOnly(dec, &b.UngappedAlignmentsOnly)
	case FlagNumThreads:
		return field.DecodeJSONThreadCount(dec, &b.NumThreads)
	case FlagUseSmithWatermanTraceback:
		return field.DecodeJSONUseSmithWatermanTraceback(dec, &b.UseSmithWatermanTraceback)
	default:
		return nil
	}
}

func (b *BlastX) NKeys() int {
	return 0
}

func (b *BlastX) ToCLI() *exec.Cmd {
	var builder = cli.NewBuilder().AppendFlag(ToolBlastX.String())

	b.baseBlast.ToCLI(builder)
	b.ipgListContainer.ToCLI(builder)
	b.bestHit.ToCLI(builder)

	builder.
		Append(b.Strand).
		Append(b.QueryGenCode).
		Append(b.Task).
		Append(b.WordSize).
		Append(b.GapOpen).
		Append(b.GapExtend).
		Append(b.MaxIntronLength).
		Append(b.Matrix).
		Append(b.Threshold).
		Append(b.CompBasedStats).
		Append(b.SubjectFile).
		Append(b.SubjectLocation).
		Append(b.Seg).
		Append(b.DBSoftMask).
		Append(b.DBHardMask).
		Append(b.CullingLimit).
		Append(b.SumStats).
		Append(b.ExtensionDropoffPrelimGapped).
		Append(b.ExtensionDropoffFinalGapped).
		Append(b.UngappedAlignmentsOnly).
		Append(b.NumThreads).
		Append(b.UseSmithWatermanTraceback)

	out := new(exec.Cmd)
	out.Args = *builder

	return out
}
