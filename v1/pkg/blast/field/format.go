package field

import (
	"strconv"
	"strings"

	"github.com/francoispqt/gojay"
	"lib-go-blast/v1/pkg/blast/consts"
)

func DecodeJSONOutFormat(dec *gojay.Decoder, val *Format) error {
	return dec.DecodeObject(val)
}

type Format struct {
	Type      FormatType
	Delimiter byte
	Fields    FormatFieldSlice
}

func (f Format) IsDefault() bool {
	return f.Type == FormatTypePairwise &&
		f.Delimiter == 0 &&
		len(f.Fields) == 0
}

func (f Format) ArgString() string {
	if f.Delimiter == 0 && len(f.Fields) == 0 {
		return f.Type.String()
	}

	out := strings.Builder{}
	out.Grow(32)
	out.Reset()

	out.WriteByte('\'')
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

	out.WriteByte('\'')

	return out.String()
}

func (f Format) Flag() string {
	return consts.FlagOutFormat
}

func (f Format) FlagString() string {
	return f.Flag() + "=" + f.ArgString()
}

func (f Format) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	switch key {
	case consts.KeyType:
		return dec.DecodeUint8((*uint8)(&f.Type))
	case consts.KeyDelimiter:
		var tmp string
		if err := dec.DecodeString(&tmp); err != nil {
			return err
		} else {
			f.Delimiter = tmp[0]
			return nil
		}
	case consts.KeyFields:
		return dec.DecodeArray(&f.Fields)
	default:
		return nil
	}
}

func (f Format) NKeys() int {
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
	FormatTypeSeqalignTextASN1
	FormatTypeSeqalignBinaryASN1
	FormatTypeCommaSeparatedValues
	FormatTypeBLASTArchiveASN1
	FormatTypeSeqalignJSON
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

type FormatFieldSlice []FormatField

func (f *FormatFieldSlice) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var tmp string
	if err := dec.DecodeString(&tmp); err != nil {
		return err
	}

	*f = append(*f, FormatField(tmp))

	return nil
}
