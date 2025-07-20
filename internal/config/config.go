package config

import (
	"os"
	"path/filepath"
)

type Config struct {
	LimitOutput int  `json:"limit_output"`
	UseRegex    bool `json:"use_regex"`
	UseOllama   bool `json:"use_ollama"`
	// ...other configs...
	// OllamaModel: If set, use this model name for Ollama API requests. Takes precedence over automatic selection.
	OllamaModel string `json:"ollama_model"`
	// OllamaModelSelection: Controls automatic model selection. "smallest" (default) or "largest". Ignored if OllamaModel is set.
	OllamaModelSelection string `json:"ollama_model_selection"`
}

func DefaultConfig() Config {
	return Config{
		LimitOutput:          5,
		UseRegex:             true,
		UseOllama:            false,
		OllamaModel:          "",
		OllamaModelSelection: "smallest",
	}
}

func LoadConfig() (Config, error) {
	home, _ := os.UserHomeDir()
	filepath.Join(home, ".ghscaffold", "configs.json")
	// ...check existence, create default if missing...
	// ...load and validate config...
	return DefaultConfig(), nil // placeholder
}
