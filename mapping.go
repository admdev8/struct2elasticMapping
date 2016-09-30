package struct2elasticMapping

type AnalyzerType string

const (
	AnalyzerStandard   AnalyzerType = "standard"
	AnalyzerWhitespace              = "whitespace"
	AnalyzerSimple                  = "simple"
	AnalyzerArabic                  = "arabic"
	AnalyzerArmenian                = "armenian"
	AnalyzerBasque                  = "basque"
	AnalyzerBrazilian               = "brazilian"
	AnalyzerBulgarian               = "bulgarian"
	AnalyzerCatalan                 = "catalan"
	AnalyzerCjk                     = "cjk"
	AnalyzerCzech                   = "czech"
	AnalyzerDanish                  = "danish"
	AnalyzerDutch                   = "dutch"
	AnalyzerEnglish                 = "english"
	AnalyzerFinnish                 = "finnish"
	AnalyzerFrench                  = "french"
	AnalyzerGalician                = "galician"
	AnalyzerGerman                  = "german"
	AnalyzerGreek                   = "greek"
	AnalyzerHindi                   = "hindi"
	AnalyzerHungarian               = "hungarian"
	AnalyzerIndonesian              = "indonesian"
	AnalyzerIrish                   = "irish"
	AnalyzerItalian                 = "italian"
	AnalyzerLatvian                 = "latvian"
	AnalyzerLithuanian              = "lithuanian"
	AnalyzerNorwegian               = "norwegian"
	AnalyzerPersian                 = "persian"
	AnalyzerPortuguese              = "portuguese"
	AnalyzerRomanian                = "romanian"
	AnalyzerRussian                 = "russian"
	AnalyzerSorani                  = "sorani"
	AnalyzerSpanish                 = "spanish"
	AnalyzerSwedish                 = "swedish"
	AnalyzerTurkish                 = "turkish"
	AnalyzerThai                    = "thai"
)

type FieldType string

const (
	FieldTypeAttachment FieldType = "attachment"
	FieldTypeBinary               = "binary"
	FieldTypeBoolean              = "boolean"
	FieldTypeByte                 = "byte"
	FieldTypeCompletion           = "completion"
	FieldTypeDate                 = "date"
	FieldTypeDouble               = "double"
	FieldTypeFloat                = "float"
	FieldTypeGeoPoint             = "geo_point"
	FieldTypeGeoShape             = "geo_shape"
	FieldTypeInteger              = "integer"
	FieldTypeIpAddress            = "ip" // IPv4 Address
	FieldTypeLong                 = "long"
	FieldTypeMurmur3              = "murmur3"
	FieldTypeObject               = "object"
	FieldTypeShort                = "short"
	FieldTypeString               = "string"
	FieldTypeTokenCount           = "token_count"
)

type IndexType string

const (
	IndexFullText      IndexType = "analyzed"
	IndexExact                   = "not_analyzed"
	IndexNotSearchable           = "no"
)

type Property struct {
	Type       FieldType           `json:"type,omitempty"`
	Index      IndexType           `json:"index,omitempty"`
	Analyzer   AnalyzerType        `json:"analyzer,omitempty"`
	Properties map[string]Property `json:"properties,omitempty"`
}

type DynamicType string

const (
	DynamicTypeTrue   DynamicType = "true"   // Add new fields dynamically—the default
	DynamicTypeFalse              = "false"  // Ignore new fields
	DynamicTypeStrict             = "strict" // Throw an exception if an unknown field is encountered
)

type Properties map[string]Property

type Mapping struct {
	Properties          Properties   `json:"properties,omitempty"`
	Analyzer            AnalyzerType `json:"analyzer,omitempty"`
	Dynamic             DynamicType  `json:"dynamic,omitempty"`
	DynamicDate_formats string       `json:"dynamic_date_formats,omitempty"`
	DynamicTemplates    string       `json:"dynamic_templates,omitempty"`
}
