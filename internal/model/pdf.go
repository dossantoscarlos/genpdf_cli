package model

type Pdf struct {
	InputValueName     string
	OutputValueName    string
	CompressionLevel   string
	PdfVersion         string
	PageRange          []string
	ShouldExtractPages bool
	Password           string
	Encrypt            string
	Metadata           map[string]string
}
