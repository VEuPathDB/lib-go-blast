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

type TBlastX struct {
	traits.QueryConfig
	traits.WithListContainer
	traits.WithBestHit

	Strand          field.Strand
	QueryGenCode    field.QueryGenCode
	WordSize        field.WordSize
	MaxIntronLength field.MaxIntronLength
	Matrix          field.Matrix
	Threshold       field.Threshold
	DBGenCode       field.DBGenCode
	SubjectFile     field.SubjectFile
	SubjectLocation field.SubjectLocation
	Seg             field.Seg
	DBSoftMask      field.DBSoftMask
	DBHardMask      field.DBHardMask
	CullingLimit    field.CullingLimit
	SumStats        field.SumStats
	NumThreads      field.ThreadCount
}

func (*TBlastX) Tool() Tool {
	return ToolTBlastX
}

func (t *TBlastX) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
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
	case consts.FlagStrand:
		return field.DecodeJSONStrand(dec, &t.Strand)
	case consts.FlagQueryGenCode:
		return field.DecodeJSONQueryGenCode(dec, &t.QueryGenCode)
	case consts.FlagWordSize:
		return field.DecodeJSONWordSize(dec, &t.WordSize)
	case consts.FlagMaxIntronLength:
		return field.DecodeJSONMaxIntronLength(dec, &t.MaxIntronLength)
	case consts.FlagMatrix:
		return field.DecodeJSONMatrix(dec, &t.Matrix)
	case consts.FlagThreshold:
		return field.DecodeJSONThreshold(dec, &t.Threshold)
	case consts.FlagDBGenCode:
		return field.DecodeJSONDBGenCode(dec, &t.DBGenCode)
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
	case consts.FlagNumThreads:
		return field.DecodeJSONThreadCount(dec, &t.NumThreads)
	default:
		return nil
	}
}

func (t *TBlastX) NKeys() int {
	return 0
}

func (t *TBlastX) ToCLI() *exec.Cmd {
	b := cli.NewBuilder().AppendFlag(ToolTBlastX.String())

	t.QueryConfig.ToCLI(b)
	t.WithListContainer.ToCLI(b)
	t.WithBestHit.ToCLI(b)

	b.Append(&t.Strand).
		Append(&t.QueryGenCode).
		Append(&t.WordSize).
		Append(&t.MaxIntronLength).
		Append(&t.Matrix).
		Append(&t.Threshold).
		Append(&t.DBGenCode).
		Append(&t.SubjectFile).
		Append(&t.SubjectLocation).
		Append(t.Seg).
		Append(&t.DBSoftMask).
		Append(&t.DBHardMask).
		Append(&t.CullingLimit).
		Append(&t.SumStats).
		Append(&t.NumThreads)

	out := new(exec.Cmd)
	out.Args = *b

	return out
}

func (t *TBlastX) Validate() bval.ValidationError {
	e := make(bval.ValidationBuilder)

	e.Validate(&t.QueryConfig, &t.WithListContainer, &t.Strand, &t.QueryGenCode).
		IncompatibleWith(&t.DBFile, &t.SubjectFile, &t.SubjectLocation).
		U32GTEQ(&t.WordSize, 2).
		Validate(&t.Threshold, &t.DBGenCode).
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
		e.AppendError(t.Format.Flag(), "Format type 17 is not supported with tblastx")
	}

	e.Incompatible(&t.DBSoftMask, &t.DBHardMask).
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
		)

	if len(e) == 0 {
		return nil
	}

	return bval.ValidationError(e)
}

