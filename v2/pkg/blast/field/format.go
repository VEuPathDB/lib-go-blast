package field

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/bval"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
)

func DecodeJSONOutFormat(dec *gojay.Decoder, val *Format) error {
	return dec.Object(val)
}

type Format struct {
	Type      FormatType
	Delimiter byte
	Fields    FormatFieldSlice
}

func (f *Format) IsDefault() bool {
	return f.Type == FormatTypePairwise &&
		f.Delimiter == 0 &&
		len(f.Fields) == 0
}

func (f *Format) ArgString() string {
	if f.Delimiter == 0 && len(f.Fields) == 0 {
		return f.Type.String()
	}

	out := strings.Builder{}
	out.Grow(32)
	out.Reset()

	out.WriteString(f.Type.String())

	if f.Delimiter != 0 {
		out.WriteByte(' ')
		out.WriteString("delim=")
		out.WriteByte(f.Delimiter)
	}

	for _, f := range f.Fields {
		out.WriteByte(' ')
		out.WriteString(f.String())
	}

	return out.String()
}

func (f *Format) Flag() string {
	return consts.FlagOutFormat
}

func (f *Format) FlagString() string {
	return f.Flag() + "=" + f.ArgString()
}

const (
	errIllegalDelim = "Custom delimiters are only allowed for report types 6, 7, 10, and 17."
	errIllegalField = "Field selection is only allowed for report types 6, 7, 10, and 17."
	errInvalidType  = "Invalid %s type \"%d\"."
	errInvalidField = "Invalid report field \"%s\"."
)

func (f *Format) Validate(em bval.ValidationBuilder) {
	if f.Delimiter != 0 {
		if f.Type != 6 && f.Type != 7 && f.Type != 10 && f.Type != 17 {
			_ = em.AppendError(f.Flag(), errIllegalDelim)
		}
	}

	if len(f.Fields) > 0 {
		if f.Type != 6 && f.Type != 7 && f.Type != 10 && f.Type != 17 {
			_ = em.AppendError(f.Flag(), errIllegalField)
		}
	}

	if !f.Type.IsValid() {
		_ = em.AppendError(f.Flag(), fmt.Sprintf(errInvalidType, f.Flag(), f.Type))
	}

	if badFields := f.Fields.invalids(); len(badFields) > 0 {
		for i := range badFields {
			_ = em.AppendError(f.Flag(), fmt.Sprintf(errInvalidField, badFields[i]))
		}
	}
}

func (f *Format) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	switch key {
	case consts.KeyType:
		return dec.Uint8((*uint8)(&f.Type))
	case consts.KeyDelimiter:
		var tmp string
		if err := dec.String(&tmp); err != nil {
			return err
		} else {
			f.Delimiter = tmp[0]
			return nil
		}
	case consts.KeyFields:
		return dec.Array(&f.Fields)
	default:
		return nil
	}
}

func (f *Format) NKeys() int {
	return 0
}

// ////////////////////////////////////////////////////////////////////////////////////////////// //

type FormatType uint8

const (
	FormatTypePairwise FormatType = iota
	FormatTypeQueryAnchoredShowingIdentities
	FormatTypeQueryAnchoredNoIdentities
	FormatTypeFlatQueryAnchoredShowingIdentities
	FormatTypeFlatQueryAnchoredNoIdentities
	FormatTypeBlastXML
	FormatTypeTabular
	FormatTypeTabularWithCommentLines
	FormatTypeSeqAlignTextASN1
	FormatTypeSeqAlignBinaryASN1
	FormatTypeCommaSeparatedValues
	FormatTypeBlastArchiveASN1
	FormatTypeSeqAlignJSON
	FormatTypeMultipleFileBlastJSON
	FormatTypeMultipleFileBlastXML2
	FormatTypeSingleFileBlastJSON
	FormatTypeSingleFileBlastXML2
	FormatTypeSequenceAlignmentMap
	FormatTypeOrganismReport
)

func (f FormatType) IsDefault() bool {
	return f == FormatTypePairwise
}

func (f FormatType) String() string {
	return strconv.Itoa(int(f))
}

func (f FormatType) IsValid() bool {
	switch f {
	case FormatTypePairwise:
		return true
	case FormatTypeQueryAnchoredShowingIdentities:
		return true
	case FormatTypeQueryAnchoredNoIdentities:
		return true
	case FormatTypeFlatQueryAnchoredShowingIdentities:
		return true
	case FormatTypeFlatQueryAnchoredNoIdentities:
		return true
	case FormatTypeBlastXML:
		return true
	case FormatTypeTabular:
		return true
	case FormatTypeTabularWithCommentLines:
		return true
	case FormatTypeSeqAlignTextASN1:
		return true
	case FormatTypeSeqAlignBinaryASN1:
		return true
	case FormatTypeCommaSeparatedValues:
		return true
	case FormatTypeBlastArchiveASN1:
		return true
	case FormatTypeSeqAlignJSON:
		return true
	case FormatTypeMultipleFileBlastJSON:
		return true
	case FormatTypeMultipleFileBlastXML2:
		return true
	case FormatTypeSingleFileBlastJSON:
		return true
	case FormatTypeSingleFileBlastXML2:
		return true
	case FormatTypeSequenceAlignmentMap:
		return true
	case FormatTypeOrganismReport:
		return true
	default:
		return false
	}
}

// ////////////////////////////////////////////////////////////////////////////////////////////// //

type FormatField string

const (
	FormatFieldQuerySeqID                         FormatField = "qseqid"
	FormatFieldQueryGI                            FormatField = "qgi"
	FormatFieldQueryAccession                     FormatField = "qacc"
	FormatFieldQueryAccessionVersion              FormatField = "qaccver"
	FormatFieldQuerySequenceLength                FormatField = "qlen"
	FormatFieldSubjectSeqID                       FormatField = "sseqid"
	FormatFieldAllSubjectSeqIDs                   FormatField = "sallseqid"
	FormatFieldSubjectGI                          FormatField = "sgi"
	FormatFieldAllSubjectGIs                      FormatField = "sallgi"
	FormatFieldSubjectAccession                   FormatField = "sacc"
	FormatFieldSubjectAccessionVersion            FormatField = "saccver"
	FormatFieldAllSubjectAccessions               FormatField = "sallacc"
	FormatFieldSubjectSequenceLength              FormatField = "slen"
	FormatFieldStartOfAlignmentInQuery            FormatField = "qstart"
	FormatFieldEndOfAlignmentInQuery              FormatField = "qend"
	FormatFieldStartOfAlignmentInSubject          FormatField = "sstart"
	FormatFieldEndOfAlignmentInSubject            FormatField = "send"
	FormatFieldAlignedPartOfQuerySequence         FormatField = "qseq"
	FormatFieldAlignedPartOfSubjectSequence       FormatField = "sseq"
	FormatFieldExpectValue                        FormatField = "evalue"
	FormatFieldBitScore                           FormatField = "bitscore"
	FormatFieldRawScore                           FormatField = "score"
	FormatFieldAlignmentLength                    FormatField = "length"
	FormatFieldPercentageOfIdenticalMatches       FormatField = "pident"
	FormatFieldNumberOfIdenticalMatches           FormatField = "nident"
	FormatFieldNumberOfMismatches                 FormatField = "mismatch"
	FormatFieldNumberOfPositiveScoringMatches     FormatField = "positive"
	FormatFieldNumberOfGapOpenings                FormatField = "gapopen"
	FormatFieldTotalNumberOfGaps                  FormatField = "gaps"
	FormatFieldPercentageOfPositiveScoringMatches FormatField = "ppos"
	FormatFieldQueryAndSubjectFrames              FormatField = "frames"
	FormatFieldQueryFrame                         FormatField = "qframe"
	FormatFieldSubjectFrame                       FormatField = "sframe"
	FormatFieldBlastTracebackOperations           FormatField = "btop"
	FormatFieldSubjectTaxonomyID                  FormatField = "staxid"
	FormatFieldSubjectScientificName              FormatField = "ssciname"
	FormatFieldSubjectCommonName                  FormatField = "scomname"
	FormatFieldSubjectBlastName                   FormatField = "sblastname"
	FormatFieldSubjectSuperKingdom                FormatField = "sskingdom"
	FormatFieldUniqueSubjectTaxonomyIDs           FormatField = "staxids"
	FormatFieldUniqueSubjectScientificNames       FormatField = "sscinames"
	FormatFieldUniqueSubjectCommonNames           FormatField = "scomnames"
	FormatFieldUniqueSubjectBlastNames            FormatField = "sblastnames"
	FormatFieldUniqueSubjectSuperKingdoms         FormatField = "sskingdoms"
	FormatFieldSubjectTitle                       FormatField = "stitle"
	FormatFieldAllSubjectTitles                   FormatField = "salltitles"
	FormatFieldSubjectStrand                      FormatField = "sstrand"
	FormatFieldQueryCoveragePerSubject            FormatField = "qcovs"
	FormatFieldQueryCoveragePerHSP                FormatField = "qcovhsp"
	FormatFieldQueryCoveragePerUniqueSubject      FormatField = "qcovus"
	FormatFieldIncludeSequenceData                FormatField = "SQ"
	FormatFieldSubjectAsReferenceSeq              FormatField = "SR"
	FormatFieldStandardFields                     FormatField = "std"
)

func (f FormatField) IsDefault() bool {
	return f == FormatFieldStandardFields
}

func (f FormatField) String() string {
	return string(f)
}

var formatFieldMap = map[FormatField]bool {
	FormatFieldQuerySeqID: true,
	FormatFieldQueryGI: true,
	FormatFieldQueryAccession: true,
	FormatFieldQueryAccessionVersion: true,
	FormatFieldQuerySequenceLength: true,
	FormatFieldSubjectSeqID: true,
	FormatFieldAllSubjectSeqIDs: true,
	FormatFieldSubjectGI: true,
	FormatFieldAllSubjectGIs: true,
	FormatFieldSubjectAccession: true,
	FormatFieldSubjectAccessionVersion: true,
	FormatFieldAllSubjectAccessions: true,
	FormatFieldSubjectSequenceLength: true,
	FormatFieldStartOfAlignmentInQuery: true,
	FormatFieldEndOfAlignmentInQuery: true,
	FormatFieldStartOfAlignmentInSubject: true,
	FormatFieldEndOfAlignmentInSubject: true,
	FormatFieldAlignedPartOfQuerySequence: true,
	FormatFieldAlignedPartOfSubjectSequence: true,
	FormatFieldExpectValue: true,
	FormatFieldBitScore: true,
	FormatFieldRawScore: true,
	FormatFieldAlignmentLength: true,
	FormatFieldPercentageOfIdenticalMatches: true,
	FormatFieldNumberOfIdenticalMatches: true,
	FormatFieldNumberOfMismatches: true,
	FormatFieldNumberOfPositiveScoringMatches: true,
	FormatFieldNumberOfGapOpenings: true,
	FormatFieldTotalNumberOfGaps: true,
	FormatFieldPercentageOfPositiveScoringMatches: true,
	FormatFieldQueryAndSubjectFrames: true,
	FormatFieldQueryFrame: true,
	FormatFieldSubjectFrame: true,
	FormatFieldBlastTracebackOperations: true,
	FormatFieldSubjectTaxonomyID: true,
	FormatFieldSubjectScientificName: true,
	FormatFieldSubjectCommonName: true,
	FormatFieldSubjectBlastName: true,
	FormatFieldSubjectSuperKingdom: true,
	FormatFieldUniqueSubjectTaxonomyIDs: true,
	FormatFieldUniqueSubjectScientificNames: true,
	FormatFieldUniqueSubjectCommonNames: true,
	FormatFieldUniqueSubjectBlastNames: true,
	FormatFieldUniqueSubjectSuperKingdoms: true,
	FormatFieldSubjectTitle: true,
	FormatFieldAllSubjectTitles: true,
	FormatFieldSubjectStrand: true,
	FormatFieldQueryCoveragePerSubject: true,
	FormatFieldQueryCoveragePerHSP: true,
	FormatFieldQueryCoveragePerUniqueSubject: true,
	FormatFieldIncludeSequenceData: true,
	FormatFieldSubjectAsReferenceSeq: true,
	FormatFieldStandardFields: true,
}

// ////////////////////////////////////////////////////////////////////////////////////////////// //

func (f FormatField) IsValid() bool {
	return formatFieldMap[f]
}

type FormatFieldSlice []FormatField

func (f *FormatFieldSlice) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var tmp string
	if err := dec.String(&tmp); err != nil {
		return err
	}

	*f = append(*f, FormatField(tmp))

	return nil
}

func (f FormatFieldSlice) invalids() []string {
	out := make([]string, 0, 3)

	for i := range f {
		if !f[i].IsValid() {
			out = append(out, string(f[i]))
		}
	}

	return out
}
