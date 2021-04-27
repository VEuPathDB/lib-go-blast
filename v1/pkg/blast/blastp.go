package blast

import (
	"os/exec"

	"github.com/francoispqt/gojay"
	. "github.com/veupathdb/lib-go-blast/v1/pkg/blast/consts"
	"github.com/veupathdb/lib-go-blast/v1/pkg/blast/field"
	"lib-go-blast/v1/pkg/cli"
)

type BlastP struct {
	baseBlast
	ipgListContainer
	bestHit

	Task                         field.BlastPTask
	WordSize                     field.WordSize
	GapOpen                      field.GapOpen
	GapExtend                    field.GapExtend
	Matrix                       field.Matrix
	Threshold                    field.Threshold
	CompBasedStats               field.FullCompBasedStats
	SubjectFile                  field.SubjectFile
	SubjectLocation              field.SubjectLocation
	Seg                          field.Seg
	DBSoftMask                   field.DBSoftMask
	DBHardMask                   field.DBHardMask
	CullingLimit                 field.CullingLimit
	ExtensionDropoffPrelimGapped field.ExtensionDropoffPrelimGapped
	ExtensionDropoffFinalGapped  field.ExtensionDropoffFinalGapped
	UngappedAlignmentsOnly       field.UngappedAlignmentsOnly
	NumThreads                   field.ThreadCount
	UseSmithWatermanTraceback    field.UseSmithWatermanTraceback
}

func (b *BlastP) Tool() Tool {
	return ToolBlastP
}

func (b *BlastP) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	if val, err := b.baseBlast.UnmarshalJSONObject(dec, key); err != nil || val {
		return err
	}

	if val, err := b.ipgListContainer.UnmarshalJSONObject(dec, key); err != nil || val {
		return err
	}

	if val, err := b.bestHit.UnmarshalJSONObject(dec, key); err != nil || val {
		return err
	}

	switch key {
	case FlagTask:
		return field.DecodeJSONBlastPTask(dec, &b.Task)
	case FlagWordSize:
		return field.DecodeJSONWordSize(dec, &b.WordSize)
	case FlagGapOpen:
		return field.DecodeJSONGapOpen(dec, &b.GapOpen)
	case FlagGapExtend:
		return field.DecodeJSONGapExtend(dec, &b.GapExtend)
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
		if val, err := field.DecodeJSONSeg(dec); err != nil {
			return err
		} else {
			b.Seg = val
			return nil
		}
	case FlagDBSoftMask:
		return field.DecodeJSONDBSoftMask(dec, &b.DBSoftMask)
	case FlagDBHardMask:
		return field.DecodeJSONDBHardMask(dec, &b.DBHardMask)
	case FlagCullingLimit:
		return field.DecodeJSONCullingLimit(dec, &b.CullingLimit)
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

func (b *BlastP) NKeys() int {
	return 0
}

func (b *BlastP) ToCLI() *exec.Cmd {
	var builder = cli.NewBuilder().AppendFlag(ToolBlastP.String())

	b.baseBlast.ToCLI(builder)
	b.ipgListContainer.ToCLI(builder)
	b.bestHit.ToCLI(builder)

	builder.
		Append(b.Task).
		Append(b.WordSize).
		Append(b.GapOpen).
		Append(b.GapExtend).
		Append(b.Matrix).
		Append(b.Threshold).
		Append(b.CompBasedStats).
		Append(b.SubjectFile).
		Append(b.SubjectLocation).
		Append(b.Seg).
		Append(b.DBSoftMask).
		Append(b.DBHardMask).
		Append(b.CullingLimit).
		Append(b.ExtensionDropoffPrelimGapped).
		Append(b.ExtensionDropoffFinalGapped).
		Append(b.UngappedAlignmentsOnly).
		Append(b.NumThreads).
		Append(b.UseSmithWatermanTraceback)

	out := new(exec.Cmd)
	out.Args = *builder

	return out
}
