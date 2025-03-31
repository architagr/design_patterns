package models

import "strategy/enums"

type ExportConfig struct {
	Type     enums.ExportType
	FileName string
}
