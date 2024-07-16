package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// LLM structure
type LLM struct {
	EmbeddingModel string  `mapstructure:"embeddingModel" yaml:"embeddingModel"`
	OpenAI         *OpenAI `mapstructure:"openai" yaml:"openai"`
}

// OpenAI structure
type OpenAI struct {
	Key string `mapstructure:"key" yaml:"key"`
}

// Vector structure
type Vector struct {
	Weaviate *WEAVIATE `mapstructure:"weaviate" yaml:"weaviate"`
}

// WEAVIATE structure
type WEAVIATE struct {
	Host   string `mapstructure:"host" yaml:"host"`
	Scheme string `mapstructure:"scheme" yaml:"scheme"`
}

// Config structure to hold entire configuration
type Config struct {
	LLM    LLM    `mapstructure:"llm" yaml:"llm"`
	Vector Vector `mapstructure:"vector" yaml:"vector"`
}

func LoadConfig(path string) (*Config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(path)
	viper.AutomaticEnv()

	// Bind environment variables to the config fields
	viper.BindEnv("LLM.EmbeddingModel", "LLM_EMBEDDING_MODEL")
	viper.BindEnv("LLM.OpenAI.Key", "OPENAI_API_KEY")
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
