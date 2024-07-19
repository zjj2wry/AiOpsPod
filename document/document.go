package document

import (
	"github.com/tmc/langchaingo/schema"
)

// DocumentSource interface defines methods to fetch documents
type DocumentSource interface {
	FetchDocuments() ([]schema.Document, error)
}
