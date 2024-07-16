package document

// Document represents a standard document structure
type Document struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// DocumentSource interface defines methods to fetch documents
type DocumentSource interface {
	FetchDocuments() ([]Document, error)
}
