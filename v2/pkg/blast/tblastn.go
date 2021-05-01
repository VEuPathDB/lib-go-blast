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

type TBlastN struct {
	traits.QueryConfig
	traits.WithListContainer
	traits.WithBestHit

	Task                         field.TBlastNTask
	WordSize                     field.WordSize
	GapOpen                      field.GapOpen
	GapExtend                    field.GapExtend
	DBGenCode                    field.DBGenCode
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
	InPSSMFile                   field.InPSSMFile
}

func (*TBlastN) Tool() Tool {
	return ToolTBlastN
}

func (t *TBlastN) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	if v, e := t.QueryConfig.UnmarshalJSONObject(dec, key); v || e != nil {
		return e
	}
	if v, e := t.WithListContainer.UnmarshalJSONObject(dec, key); v || e != nil {
		return e
	}
	if v, e := t.WithBestHit.UnmarshalJSONObject(dec, key); v || e != nil {
		return e
	}

	switch key {
	case consts.FlagTask:
		return field.DecodeJSONTBlastNTask(dec, &t.Task)
	case consts.FlagWordSize:
		return field.DecodeJSONWordSize(dec, &t.WordSize)
	case consts.FlagGapOpen:
		return field.DecodeJSONGapOpen(dec, &t.GapOpen)
	case consts.FlagGapExtend:
		return field.DecodeJSONGapExtend(dec, &t.GapExtend)
	case consts.FlagDBGenCode:
		return field.DecodeJSONDBGenCode(dec, &t.DBGenCode)
	case consts.FlagMaxIntronLength:
		return field.DecodeJSONMaxIntronLength(dec, &t.MaxIntronLength)
	case consts.FlagMatrix:
		return field.DecodeJSONMatrix(dec, &t.Matrix)
	case consts.FlagThreshold:
		return field.DecodeJSONThreshold(dec, &t.Threshold)
	case consts.FlagCompBasedStats:
		return field.DecodeJSONFullCompBasedStats(dec, &t.CompBasedStats)
	case consts.FlagSubjectFile:
		return field.DecodeJSONSubjectFile(dec, &t.SubjectFile)
	case consts.FlagSubjectLocation:
		return field.DecodeJSONSubjectLocation(dec, &t.SubjectLocation)
	case consts.FlagSeg:
		if v, e := field.DecodeJSONSeg(dec); e != nil {
			return e
		} else {
			t.Seg = v
			return nil
		}
	case consts.FlagDBSoftMask:
		return field.DecodeJSONDBSoftMask(dec, &t.DBSoftMask)
	case consts.FlagDBHardMask:
		return field.DecodeJSONDBHardMask(dec, &t.DBHardMask)
	case consts.FlagCullingLimit:
		return field.DecodeJSONCullingLimit(dec, &t.CullingLimit)
	case consts.FlagSumStats:
		return field.DecodeJSONSumStats(dec, &t.SumStats)
	case consts.FlagExtensionDropoffPrelimGapped:
		return field.DecodeJSONExtensionDropoffPrelimGapped(dec, &t.ExtensionDropoffPrelimGapped)
	case consts.FlagExtensionDropoffFinalGapped:
		return field.DecodeJSONExtensionDropoffFinalGapped(dec, &t.ExtensionDropoffFinalGapped)
	case consts.FlagUngappedAlignmentsOnly:
		return field.DecodeJSONUngappedAlignmentsOnly(dec, &t.UngappedAlignmentsOnly)
	case consts.FlagNumThreads:
		return field.DecodeJSONThreadCount(dec, &t.NumThreads)
	case consts.FlagUseSmithWatermanTraceback:
		return field.DecodeJSONUseSmithWatermanTraceback(dec, &t.UseSmithWatermanTraceback)
	case consts.FlagInPSSMFile:
		return field.DecodeJSONInPSSMFile(dec, &t.InPSSMFile)
	default:
		return nil
	}
}

func (t *TBlastN) NKeys() int {
	return 0
}

func (t *TBlastN) ToCLI() *exec.Cmd {
	b := cli.NewBuilder().AppendFlag(ToolTBlastN.String())

	t.QueryConfig.ToCLI(b)
	t.WithListContainer.ToCLI(b)
	t.WithBestHit.ToCLI(b)

	b.Append(&t.Task).
		Append(&t.WordSize).
		Append(&t.GapOpen).
		Append(&t.GapExtend).
		Append(&t.DBGenCode).
		Append(&t.MaxIntronLength).
		Append(&t.Matrix).
		Append(&t.Threshold).
		Append(&t.CompBasedStats).
		Append(&t.SubjectFile).
		Append(&t.SubjectLocation).
		Append(t.Seg).
		Append(&t.DBSoftMask).
		Append(&t.DBHardMask).
		Append(&t.CullingLimit).
		Append(&t.SumStats).
		Append(&t.ExtensionDropoffPrelimGapped).
		Append(&t.ExtensionDropoffFinalGapped).
		Append(&t.UngappedAlignmentsOnly).
		Append(&t.NumThreads).
		Append(&t.UseSmithWatermanTraceback).
		Append(&t.InPSSMFile)

	out := new(exec.Cmd)
	out.Args = *b

	return out
}

func (t *TBlastN) Validate() bval.ValidationError {
	em := make(bval.ValidationBuilder)

	em.Validate(&t.QueryConfig, &t.WithListContainer, &t.Task).
		U32GTEQ(&t.WordSize, 2).
		Incompatible(&t.QueryFile, &t.InPSSMFile).
		Incompatible(&t.QueryLocation, &t.InPSSMFile).
		IncompatibleWith(&t.DBFile, &t.SubjectFile, &t.SubjectLocation).
		Validate(&t.DBGenCode, &t.Threshold, &t.CompBasedStats).
		IncompatibleWith(
			&t.SubjectFile,
			&t.GIList,
			&t.NegativeGIList,
			&t.SequenceIDList,
			&t.NegativeSequenceIDList,
			t.TaxIDs,
			t.NegativeTaxIDs,
			&t.TaxIDList,
			&t.NegativeTaxIDList,
			&t.DBSoftMask,
			&t.DBHardMask,
		).
		IncompatibleWith(
			&t.SubjectLocation,
			&t.GIList,
			&t.NegativeGIList,
			&t.SequenceIDList,
			&t.NegativeSequenceIDList,
			t.TaxIDs,
			t.NegativeTaxIDs,
			&t.TaxIDList,
			&t.NegativeTaxIDList,
			&t.DBSoftMask,
			&t.DBHardMask,
			&t.Remote,
		)

	if t.Format.Type == field.FormatTypeSequenceAlignmentMap {
		em.AppendError(t.Format.Flag(), "Format type 17 is not supported with tblastn")
	}

	em.Incompatible(&t.DBSoftMask, &t.DBHardMask).
		IncompatibleWith(&t.CullingLimit, &t.BestHitOverhang, &t.BestHitScoreEdge).
		IncompatibleWith(
			&t.Remote,
			&t.GIList,
			&t.SequenceIDList,
			&t.TaxIDs,
			&t.TaxIDList,
			&t.NegativeGIList,
			&t.NegativeSequenceIDList,
			&t.NegativeTaxIDs,
			&t.NegativeTaxIDList,
			&t.NumThreads,
			&t.InPSSMFile,
		)

	if len(em) == 0 {
		return nil
	}

	return bval.ValidationError(em)
}

