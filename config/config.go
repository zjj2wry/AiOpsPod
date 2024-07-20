package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// LLM structure
type LLM struct {
	EmbeddingModel string  `yaml:"embeddingModel"`
	OpenAI         *OpenAI `yaml:"openai"`
}

type DocumentConfig struct {
	FeishuConfig *FeishuConfig `yaml:"feishuConfig"`
	LocalDir     *LocalDir     `yaml:"localDir"`
}

type LocalDir struct {
	Directory string ` yaml:"directory"`
}

// FeishuConfig holds the configuration for Feishu
type FeishuConfig struct {
}

// OpenAI structure
type OpenAI struct {
	Key   string `yaml:"key"`
	Model string `yaml:"model"`
}

// Vector structure
type Vector struct {
	Weaviate *WEAVIATE `yaml:"weaviate"`
	// scoreThreshold specifies the minimum similarity score required for a result to be included in the query results.
	//
	// In vector similarity searches, the score represents the degree of relevance between the query vector and the result vectors.
	// A higher scoreThreshold ensures that only results with high similarity to the query vector are returned, which can improve the relevance of the results.
	// Conversely, a lower scoreThreshold may include more results with lower relevance, which might be useful for applications requiring broader searches.
	//
	// Setting the threshold too high might exclude relevant results, while setting it too low might include irrelevant ones.
	// The ideal threshold value depends on the specific use case, the quality of the data, and the desired balance between precision and recall.
	//
	// Example:
	//   For a search engine with high precision requirements, a threshold of 0.8 or higher might be appropriate.
	//   For exploratory search or broad recommendation systems, a threshold of 0.5 might be used to include more diverse results.
	ScoreThreshold float32 `yaml:"scoreThreshold"`
}

// WEAVIATE structure
type WEAVIATE struct {
	Host   string `yaml:"host"`
	Scheme string `yaml:"scheme"`
}

type Prometheus struct {
	Address string `yaml:"address"`
}

// Config structure to hold entire configuration
type Config struct {
	LLM        LLM             `yaml:"llm"`
	Vector     Vector          `yaml:"vector"`
	Document   *DocumentConfig `yaml:"document"`
	Prometheus *Prometheus     `yaml:"prometheus"`
}

func LoadConfig(path string) (*Config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(path)
	viper.AutomaticEnv()

	// Bind environment variables to the config fields
	viper.BindEnv("LLM.EmbeddingModel", "LLM_EMBEDDING_MODEL")
	viper.BindEnv("LLM.OpenAI.Key", "OPENAI_API_KEY")
	viper.BindEnv("LLM.OpenAI.Model", "OPENAI_MODEL")
	viper.BindEnv("Vector.Weaviate.Host", "WEAVIATE_HOST")
	viper.BindEnv("Vector.Weaviate.Scheme", "WEAVIATE_SCHEME")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("unable to decode config into struct: %w", err)
	}

	return &config, nil
}
