package blast

import (
	"os/exec"

	"github.com/francoispqt/gojay"
	. "github.com/veupathdb/lib-go-blast/v1/pkg/blast/consts"
	"github.com/veupathdb/lib-go-blast/v1/pkg/blast/field"
	"github.com/veupathdb/lib-go-blast/v1/pkg/cli"
)

type BlastN struct {
	baseBlast
	baseListContainer
	bestHit

	Strand                       field.Strand
	Task                         field.BlastNTask
	WordSize                     field.WordSize
	GapOpen                      field.GapOpen
	GapExtend                    field.GapExtend
	Penalty                      field.Penalty
	Reward                       field.Reward
	UseIndex                     field.UseIndex
	IndexName                    field.IndexName
	SubjectFile                  field.SubjectFile
	SubjectLocation              field.SubjectLocation
	Dust                         field.Dust
	FilteringDB                  field.FilteringDB
	WindowMaskerTaxID            field.WindowMaskerTaxID
	WindowMaskerDB               field.WindowMaskerDB
	DBSoftMask                   field.DBSoftMask
	DBHardMask                   field.DBHardMask
	PercentIdentity              field.PercentIdentity
	CullingLimit                 field.CullingLimit
	TemplateType                 field.TemplateType
	TemplateLength               field.TemplateLength
	SumStats                     field.SumStats
	ExtensionDropoffPrelimGapped field.ExtensionDropoffPrelimGapped
	ExtensionDropoffFinalGapped  field.ExtensionDropoffFinalGapped
	NonGreedy                    field.NonGreedy
	MinRawGappedScore            field.MinRawGappedScore
	UngappedAlignmentsOnly       field.UngappedAlignmentsOnly
	OffDiagonalRange             field.OffDiagonalRange
	NumThreads                   field.ThreadCount
}

func (b *BlastN) Tool() Tool {
	return ToolBlastN
}

func (b *BlastN) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	if val, err := b.baseBlast.UnmarshalJSONObject(dec, key); err != nil || val {
		return nil
	}

	if val, err := b.baseListContainer.UnmarshalJSONObject(dec, key); err != nil || val {
		return err
	}

	if val, err := b.bestHit.UnmarshalJSONObject(dec, key); err != nil || val {
		return err
	}

	switch key {
	case FlagStrand:
		return field.DecodeJSONStrand(dec, &b.Strand)
	case FlagTask:
		return field.DecodeJSONBlastNTask(dec, &b.Task)
	case FlagWordSize:
		return field.DecodeJSONWordSize(dec, &b.WordSize)
	case FlagGapOpen:
		return field.DecodeJSONGapOpen(dec, &b.GapOpen)
	case FlagGapExtend:
		return field.DecodeJSONGapExtend(dec, &b.GapExtend)
	case FlagPenalty:
		return field.DecodeJSONPenalty(dec, &b.Penalty)
	case FlagReward:
		return field.DecodeJSONReward(dec, &b.Reward)
	case FlagUseIndex:
		return field.DecodeJSONUseIndex(dec, &b.UseIndex)
	case FlagIndexName:
		return field.DecodeJSONIndexName(dec, &b.IndexName)
	case FlagSubjectFile:
		return field.DecodeJSONSubjectFile(dec, &b.SubjectFile)
	case FlagSubjectLocation:
		return field.DecodeJSONSubjectLocation(dec, &b.SubjectLocation)
	case FlagDust:
		if tmp, err := field.DecodeJSONDust(dec); err != nil {
			return err
		} else {
			b.Dust = tmp
			return nil
		}
	case FlagFilteringDB:
		return field.DecodeJSONFilteringDB(dec, &b.FilteringDB)
	case FlagWindowMaskerTaxID:
		return field.DecodeJSONWindowMaskerTaxID(dec, &b.WindowMaskerTaxID)
	case FlagWindowMaskerDB:
		return field.DecodeJSONWindowMaskerDB(dec, &b.WindowMaskerDB)
	case FlagDBSoftMask:
		return field.DecodeJSONDBSoftMask(dec, &b.DBSoftMask)
	case FlagDBHardMask:
		return field.DecodeJSONDBHardMask(dec, &b.DBHardMask)
	case FlagPercentIdentity:
		return field.DecodeJSONPercentIdentity(dec, &b.PercentIdentity)
	case FlagCullingLimit:
		return field.DecodeJSONCullingLimit(dec, &b.CullingLimit)
	case FlagTemplateType:
		return field.DecodeJSONTemplateType(dec, &b.TemplateType)
	case FlagTemplateLength:
		return field.DecodeJSONTemplateLength(dec, &b.TemplateLength)
	case FlagSumStats:
		return field.DecodeJSONSumStats(dec, &b.SumStats)
	case FlagExtensionDropoffPrelimGapped:
		return field.DecodeJSONExtensionDropoffPrelimGapped(dec, &b.ExtensionDropoffPrelimGapped)
	case FlagExtensionDropoffFinalGapped:
		return field.DecodeJSONExtensionDropoffFinalGapped(dec, &b.ExtensionDropoffFinalGapped)
	case FlagNonGreedy:
		return field.DecodeJSONNonGreedy(dec, &b.NonGreedy)
	case FlagMinRawGappedScore:
		return field.DecodeJSONMinRawGappedScore(dec, &b.MinRawGappedScore)
	case FlagUngappedAlignmentsOnly:
		return field.DecodeJSONUngappedAlignmentsOnly(dec, &b.UngappedAlignmentsOnly)
	case FlagOffDiagonalRange:
		return field.DecodeJSONOffDiagonalRange(dec, &b.OffDiagonalRange)
	case FlagNumThreads:
		return field.DecodeJSONThreadCount(dec, &b.NumThreads)
	default:
		return nil
	}
}

func (b *BlastN) NKeys() int {
	return 0
}

func (b *BlastN) ToCLI() *exec.Cmd {
	var builder = cli.NewBuilder().AppendFlag(ToolBlastN.String())

	b.baseBlast.ToCLI(builder)
	b.baseListContainer.ToCLI(builder)
	b.bestHit.ToCLI(builder)

	builder.
		Append(b.Strand).
		Append(b.Task).
		Append(b.WordSize).
		Append(b.GapOpen).
		Append(b.GapExtend).
		Append(b.Penalty).
		Append(b.Reward).
		Append(b.UseIndex).
		Append(b.IndexName).
		Append(b.SubjectFile).
		Append(b.SubjectLocation).
		Append(b.Dust).
		Append(b.FilteringDB).
		Append(b.WindowMaskerTaxID).
		Append(b.WindowMaskerDB).
		Append(b.DBSoftMask).
		Append(b.DBHardMask).
		Append(b.PercentIdentity).
		Append(b.CullingLimit).
		Append(b.TemplateType).
		Append(b.TemplateLength).
		Append(b.SumStats).
		Append(b.ExtensionDropoffPrelimGapped).
		Append(b.ExtensionDropoffFinalGapped).
		Append(b.NonGreedy).
		Append(b.MinRawGappedScore).
		Append(b.UngappedAlignmentsOnly).
		Append(b.OffDiagonalRange).
		Append(b.NumThreads)

	out := new(exec.Cmd)
	out.Args = *builder

	return out
}
