package ollama

import (
	"bytes"
	"encoding/json"
	"errors"
	"os/exec"
	"sort"
	"strings"

	"github.com/nathfavour/ghscaffold/internal/config"
)

// OllamaModelInfo holds info about an installed Ollama model
type OllamaModelInfo struct {
	Name string `json:"name"`
	Size int64  `json:"size"` // in bytes
}

// SelectOllamaModel chooses the model to use based on config and installed models
func SelectOllamaModel(cfg config.Config) (string, error) {
	if cfg.OllamaModel != "" {
		return cfg.OllamaModel, nil
	}
	models, err := listInstalledOllamaModels()
	if err != nil || len(models) == 0 {
		return "", errors.New("no ollama models found")
	}
	sort.Slice(models, func(i, j int) bool {
		if cfg.OllamaModelSelection == "largest" {
			return models[i].Size > models[j].Size
		}
		return models[i].Size < models[j].Size
	})
	return models[0].Name, nil
}

// listInstalledOllamaModels gets installed models using 'ollama list --json'
func listInstalledOllamaModels() ([]OllamaModelInfo, error) {
	cmd := exec.Command("ollama", "list", "--json")
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	// The output may be a JSON array or newline-delimited JSON objects
	var models []OllamaModelInfo
	dec := json.NewDecoder(bytes.NewReader(out))
	for dec.More() {
		var m OllamaModelInfo
		if err := dec.Decode(&m); err != nil {
			// Try to parse as array if newline-delimited fails
			if strings.Contains(string(out), "[") {
				if err := json.Unmarshal(out, &models); err != nil {
					return nil, err
				}
				return models, nil
			}
			return nil, err
		}
		models = append(models, m)
	}
	return models, nil
}

func ScoreReadmeWithOllama(projectDetails, readme string) (int, error) {
	// ...connect to Ollama, send prompt, parse response...
	return 0, nil // placeholder
}
