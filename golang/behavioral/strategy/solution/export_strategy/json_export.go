package exportstrategy

import (
	"encoding/json"
	"errors"
	"strategy_pattern/models"
)

// JSON Export: Concrete Strategy
type JSONExportStrategy struct {
}

func (e *JSONExportStrategy) Export(data []models.UserModel) (string, error) {
	if len(data) == 0 {
		return "", errors.New("No data")
	}
	jsonData, err := json.Marshal(data)
	return string(jsonData), err
}
