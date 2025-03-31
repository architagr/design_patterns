package exportstrategy

import (
	"errors"
	"strategy_pattern/models"
	"strings"
)

// CSV Export: Concrete strategy
type CsvExportStrategy struct {
}

func (e *CsvExportStrategy) Export(data []models.UserModel) (string, error) {
	if len(data) == 0 {
		return "", errors.New("No data")
	}
	sb := strings.Builder{}
	sb.WriteString(data[0].GetHeader())
	for _, u := range data {
		sb.WriteString(u.GetRow())
	}
	return sb.String(), nil
}
