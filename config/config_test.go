package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	// Create a temporary config file for testing
	tempDir := t.TempDir()
	tempFile := filepath.Join(tempDir, "config.yaml")
	configContent := `
llm:
  embedding_model: "bert"
  openai:
    key: "test-openai-key"
vector:
  weaviate:
    host: "test-host"
    scheme: "https"
document:
  localDir:
    directory: sop
`
	if err := os.WriteFile(tempFile, []byte(configContent), 0644); err != nil {
		t.Fatalf("Failed to write temp config file: %v", err)
	}

	// Setup environment variables
	os.Setenv("LLM_EMBEDDING_MODEL", "env-bert")
	os.Setenv("OPENAI_API_KEY", "env-openai-key")
	os.Setenv("WEAVIATE_HOST", "env-weaviate-host")
	os.Setenv("WEAVIATE_SCHEME", "env-https")
	defer func() {
		os.Unsetenv("LLM_EMBEDDING_MODEL")
		os.Unsetenv("OPENAI_API_KEY")
		os.Unsetenv("WEAVIATE_HOST")
		os.Unsetenv("WEAVIATE_SCHEME")
	}()

	// Load config
	config, err := LoadConfig(tempDir)
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	// Test assertions
	assert.Equal(t, "env-bert", config.LLM.EmbeddingModel)
	assert.NotNil(t, config.LLM.OpenAI)
	assert.Equal(t, "env-openai-key", config.LLM.OpenAI.Key)

	assert.NotNil(t, config.Vector.Weaviate)
	assert.Equal(t, "env-weaviate-host", config.Vector.Weaviate.Host)
	assert.Equal(t, "env-https", config.Vector.Weaviate.Scheme)
	assert.Equal(t, "sop", config.Document.LocalDir.Directory)
}
