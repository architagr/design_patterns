package exportstrategy

import (
	"encoding/xml"
	"errors"
	"strategy_pattern/models"
)

// XML Export: Concrete Strategy
type XMLExportStrategy struct {
}

func (e *XMLExportStrategy) Export(data []models.UserModel) (string, error) {
	if len(data) == 0 {
		return "", errors.New("No data")
	}
	xmlData, err := xml.Marshal(data)
	return string(xmlData), err
}
