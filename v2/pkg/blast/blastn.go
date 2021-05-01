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

type BlastN struct {
	traits.QueryConfig
	traits.WithListContainer
	traits.WithBestHit

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
	if val, err := b.QueryConfig.UnmarshalJSONObject(dec, key); err != nil || val {
		return nil
	}

	if val, err := b.WithListContainer.UnmarshalJSONObject(dec, key); err != nil || val {
		return err
	}

	if val, err := b.WithBestHit.UnmarshalJSONObject(dec, key); err != nil || val {
		return err
	}

	switch key {
	case consts.FlagStrand:
		return field.DecodeJSONStrand(dec, &b.Strand)
	case consts.FlagTask:
		return field.DecodeJSONBlastNTask(dec, &b.Task)
	case consts.FlagWordSize:
		return field.DecodeJSONWordSize(dec, &b.WordSize)
	case consts.FlagGapOpen:
		return field.DecodeJSONGapOpen(dec, &b.GapOpen)
	case consts.FlagGapExtend:
		return field.DecodeJSONGapExtend(dec, &b.GapExtend)
	case consts.FlagPenalty:
		return field.DecodeJSONPenalty(dec, &b.Penalty)
	case consts.FlagReward:
		return field.DecodeJSONReward(dec, &b.Reward)
	case consts.FlagUseIndex:
		return field.DecodeJSONUseIndex(dec, &b.UseIndex)
	case consts.FlagIndexName:
		return field.DecodeJSONIndexName(dec, &b.IndexName)
	case consts.FlagSubjectFile:
		return field.DecodeJSONSubjectFile(dec, &b.SubjectFile)
	case consts.FlagSubjectLocation:
		return field.DecodeJSONSubjectLocation(dec, &b.SubjectLocation)
	case consts.FlagDust:
		if tmp, err := field.DecodeJSONDust(dec); err != nil {
			return err
		} else {
			b.Dust = tmp
			return nil
		}
	case consts.FlagFilteringDB:
		return field.DecodeJSONFilteringDB(dec, &b.FilteringDB)
	case consts.FlagWindowMaskerTaxID:
		return field.DecodeJSONWindowMaskerTaxID(dec, &b.WindowMaskerTaxID)
	case consts.FlagWindowMaskerDB:
		return field.DecodeJSONWindowMaskerDB(dec, &b.WindowMaskerDB)
	case consts.FlagDBSoftMask:
		return field.DecodeJSONDBSoftMask(dec, &b.DBSoftMask)
	case consts.FlagDBHardMask:
		return field.DecodeJSONDBHardMask(dec, &b.DBHardMask)
	case consts.FlagPercentIdentity:
		return field.DecodeJSONPercentIdentity(dec, &b.PercentIdentity)
	case consts.FlagCullingLimit:
		return field.DecodeJSONCullingLimit(dec, &b.CullingLimit)
	case consts.FlagTemplateType:
		return field.DecodeJSONTemplateType(dec, &b.TemplateType)
	case consts.FlagTemplateLength:
		return field.DecodeJSONTemplateLength(dec, &b.TemplateLength)
	case consts.FlagSumStats:
		return field.DecodeJSONSumStats(dec, &b.SumStats)
	case consts.FlagExtensionDropoffPrelimGapped:
		return field.DecodeJSONExtensionDropoffPrelimGapped(dec, &b.ExtensionDropoffPrelimGapped)
	case consts.FlagExtensionDropoffFinalGapped:
		return field.DecodeJSONExtensionDropoffFinalGapped(dec, &b.ExtensionDropoffFinalGapped)
	case consts.FlagNonGreedy:
		return field.DecodeJSONNonGreedy(dec, &b.NonGreedy)
	case consts.FlagMinRawGappedScore:
		return field.DecodeJSONMinRawGappedScore(dec, &b.MinRawGappedScore)
	case consts.FlagUngappedAlignmentsOnly:
		return field.DecodeJSONUngappedAlignmentsOnly(dec, &b.UngappedAlignmentsOnly)
	case consts.FlagOffDiagonalRange:
		return field.DecodeJSONOffDiagonalRange(dec, &b.OffDiagonalRange)
	case consts.FlagNumThreads:
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

	b.QueryConfig.ToCLI(builder)
	b.WithListContainer.ToCLI(builder)
	b.WithBestHit.ToCLI(builder)

	builder.
		Append(&b.Strand).
		Append(&b.Task).
		Append(&b.WordSize).
		Append(&b.GapOpen).
		Append(&b.GapExtend).
		Append(&b.Penalty).
		Append(&b.Reward).
		Append(&b.UseIndex).
		Append(&b.IndexName).
		Append(&b.SubjectFile).
		Append(&b.SubjectLocation).
		Append(b.Dust).
		Append(&b.FilteringDB).
		Append(&b.WindowMaskerTaxID).
		Append(&b.WindowMaskerDB).
		Append(&b.DBSoftMask).
		Append(&b.DBHardMask).
		Append(&b.PercentIdentity).
		Append(&b.CullingLimit).
		Append(&b.TemplateType).
		Append(&b.TemplateLength).
		Append(&b.SumStats).
		Append(&b.ExtensionDropoffPrelimGapped).
		Append(&b.ExtensionDropoffFinalGapped).
		Append(&b.NonGreedy).
		Append(&b.MinRawGappedScore).
		Append(&b.UngappedAlignmentsOnly).
		Append(&b.OffDiagonalRange).
		Append(&b.NumThreads)

	out := new(exec.Cmd)
	out.Args = *builder

	return out
}

func (b *BlastN) Validate() bval.ValidationError {
	errors := make(bval.ValidationBuilder)

	b.QueryConfig.Validate(errors)
	b.WithListContainer.Validate(errors)

	_ = errors.Validate(&b.Strand, &b.Task, &b.Penalty).
		IncompatibleWith(&b.DBFile, &b.SubjectFile, &b.SubjectLocation).
		U32GTEQ(&b.WordSize, 4).
		IncompatibleWith(
			&b.SubjectFile,
			&b.GIList,
			&b.NegativeGIList,
			&b.SequenceIDList,
			&b.NegativeSequenceIDList,
			b.TaxIDs,
			b.NegativeTaxIDs,
			&b.TaxIDList,
			&b.NegativeTaxIDList,
			&b.DBSoftMask,
			&b.DBHardMask,
		).
		IncompatibleWith(
			&b.SubjectLocation,
			&b.GIList,
			&b.NegativeGIList,
			&b.SequenceIDList,
			&b.NegativeSequenceIDList,
			b.TaxIDs,
			b.NegativeTaxIDs,
			&b.TaxIDList,
			&b.NegativeTaxIDList,
			&b.DBSoftMask,
			&b.DBHardMask,
			&b.Remote,
		).
		Incompatible(&b.DBSoftMask, &b.DBHardMask).
		Validate(&b.PercentIdentity).
		IncompatibleWith(&b.CullingLimit, &b.BestHitOverhang, &b.BestHitScoreEdge).
		Validate(
			&b.BestHitOverhang,
			&b.BestHitScoreEdge,
			&b.TemplateType,
			&b.TemplateLength,
		).
		Requires(&b.TemplateLength, &b.TemplateType).
		IncompatibleWith(
			&b.Remote,
			&b.GIList,
			&b.SequenceIDList,
			&b.TaxIDs,
			&b.TaxIDList,
			&b.NegativeGIList,
			&b.NegativeSequenceIDList,
			&b.NegativeTaxIDs,
			&b.NegativeTaxIDList,
			&b.NumThreads,
		)

	if len(errors) == 0 {
		return nil
	}

	return bval.ValidationError(errors)
}

