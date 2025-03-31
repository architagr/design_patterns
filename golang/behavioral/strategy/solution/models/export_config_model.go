package models

import "strategy_pattern/enums"

type ExportConfig struct {
	Type     enums.ExportType
	FileName string
}
