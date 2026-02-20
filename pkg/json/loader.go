package json

import (
	"encoding/json"
	"fmt"
	"os"
)

func Load[T any](path string) (T, error) {
	var result T

	data, err := os.ReadFile(path)
	if err != nil {
		return result, fmt.Errorf("read file: %w", err)
	}

	if err := json.Unmarshal(data, &result); err != nil {
		return result, fmt.Errorf("unmarshal json: %w", err)
	}

	return result, nil
}
