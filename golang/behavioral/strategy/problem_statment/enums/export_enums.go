package enums

type ExportType int

const (
	ExportType_CSV     = iota
	ExportType_Json    = iota
	ExportType_Xml     = iota
	ExportType_WebHook = iota
)
