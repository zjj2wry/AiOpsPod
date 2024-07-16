package document

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/zjj2wry/AiOpsPod/config"
	"go.uber.org/zap"
)

// Helper function to create a temporary directory with test files
func createTestFiles(t *testing.T, files map[string]string) string {
	tempDir, err := ioutil.TempDir("", "testfiles")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}

	for name, content := range files {
		path := filepath.Join(tempDir, name)
		if err := ioutil.WriteFile(path, []byte(content), 0644); err != nil {
			t.Fatalf("Failed to write test file %s: %v", name, err)
		}
	}

	return tempDir
}

// TestLocalDocumentSource_FetchDocuments tests the FetchDocuments method
func TestLocalDocumentSource_FetchDocuments(t *testing.T) {
	// Initialize logger
	logger, err := zap.NewDevelopment()
	if err != nil {
		t.Fatalf("Failed to create logger: %v", err)
	}
	defer logger.Sync()

	// Create a temporary directory with test files
	testFiles := map[string]string{
		"doc1.txt": "Content of document 1",
		"doc2.txt": "Content of document 2",
		"doc3.txt": "Content of document 3",
	}

	tempDir := createTestFiles(t, testFiles)
	defer os.RemoveAll(tempDir)

	// Initialize LocalDocumentSource
	source := &LocalDocumentSource{
		LocalDir: config.LocalDir{
			Directory: tempDir,
		},
		Logger: logger,
	}

	// Call FetchDocuments
	documents, err := source.FetchDocuments()
	if err != nil {
		t.Fatalf("FetchDocuments returned an error: %v", err)
	}

	// Check the number of documents
	if len(documents) != len(testFiles) {
		t.Fatalf("Expected %d documents, got %d", len(testFiles), len(documents))
	}

	// Check the contents of the documents
	for _, doc := range documents {
		expectedContent, exists := testFiles[doc.Title]
		if !exists {
			t.Errorf("Unexpected document: %s", doc.Title)
			continue
		}

		if doc.Content != expectedContent {
			t.Errorf("Expected content '%s' for document %s, got '%s'", expectedContent, doc.Title, doc.Content)
		}
	}
}

// TestLocalDocumentSource_FetchDocuments_Error tests that FetchDocuments handles directory errors correctly
func TestLocalDocumentSource_FetchDocuments_Error(t *testing.T) {
	// Initialize logger
	logger, err := zap.NewDevelopment()
	if err != nil {
		t.Fatalf("Failed to create logger: %v", err)
	}
	defer logger.Sync()

	// Initialize LocalDocumentSource with a non-existent directory
	source := &LocalDocumentSource{
		LocalDir: config.LocalDir{
			Directory: "non-existent-dir",
		},
		Logger: logger,
	}

	// Call FetchDocuments
	_, err = source.FetchDocuments()
	if err == nil {
		t.Fatal("Expected an error when fetching documents from a non-existent directory, but got nil")
	}
}
