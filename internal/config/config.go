package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type InputTarget struct {
	Name  string `json:"name"`
	URL   string `json:"url"`
	Owner string `json:"owner"`
}

func LoadTargetsFromFile(filePath string) ([]InputTarget, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filePath, err)
	}

	var targets []InputTarget
	if err := json.Unmarshal(data, &targets); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON from file %s: %w", filePath, err)
	}

	return targets, nil
}

func SaveTargetsToFile(filePath string, targets []InputTarget) error {
	data, err := json.MarshalIndent(targets, "", " ")
	if err != nil {
		return fmt.Errorf("impossible de lire le fichier %s: %w", filePath, err)
	}

	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return fmt.Errorf("impossible d'écrire dans le fichier %s: %w", filePath, err)
	}
	return nil
}
