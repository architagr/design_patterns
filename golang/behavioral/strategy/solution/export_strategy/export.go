package exportstrategy

import (
	"errors"
	"fmt"
	"strategy_pattern/enums"
	"strategy_pattern/models"
)

var (
	Err_NotSupportedExportType = errors.New("Not Supported export Type")
)

// Strategy interface
type IUserExport interface {
	Export(data []models.UserModel) (string, error)
}

func UserExportBuilder(config *models.ExportConfig) (IUserExport, error) {
	switch config.Type {
	case enums.ExportType_CSV:
		return &CsvExportStrategy{}, nil
	case enums.ExportType_Json:
		return &JSONExportStrategy{}, nil
	case enums.ExportType_Xml:
		return &XMLExportStrategy{}, nil
	default:
		return nil, fmt.Errorf("%w, %d", Err_NotSupportedExportType, config.Type)
	}
}
