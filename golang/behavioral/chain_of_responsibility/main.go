package main

import "fmt"

type exportRequest struct {
	requestid     int
	currentStatus string
	email         string
	data          []any
	filePaths     map[string]string
}

type exportStepHandler interface {
	execute(req *exportRequest) error
	setNext(next exportStepHandler)
}

type exportFetchUsers struct {
	next exportStepHandler
}

func (h *exportFetchUsers) execute(req *exportRequest) error {
	fmt.Printf("RequestId(%d): Users fetched from DB\n", req.requestid)
	if h.next != nil {
		return h.next.execute(req)
	}
	return nil
}

func (h *exportFetchUsers) setNext(next exportStepHandler) {
	h.next = next
}

type exportParseIntoFile struct {
	next exportStepHandler
}

func (h *exportParseIntoFile) execute(req *exportRequest) error {
	fmt.Printf("RequestId(%d): Users parsed into file\n", req.requestid)
	if h.next != nil {
		return h.next.execute(req)
	}
	return nil
}

func (h *exportParseIntoFile) setNext(next exportStepHandler) {
	h.next = next
}

type exportZipFiles struct {
	next exportStepHandler
}

func (h *exportZipFiles) execute(req *exportRequest) error {
	fmt.Printf("RequestId(%d): file zipped\n", req.requestid)
	if h.next != nil {
		return h.next.execute(req)
	}
	return nil
}

func (h *exportZipFiles) setNext(next exportStepHandler) {
	h.next = next
}

type exportEmailZip struct {
	next exportStepHandler
}

func (h *exportEmailZip) execute(req *exportRequest) error {
	fmt.Printf("RequestId(%d): Zip emailed\n", req.requestid)
	if h.next != nil {
		return h.next.execute(req)
	}
	return nil
}

func (h *exportEmailZip) setNext(next exportStepHandler) {
	h.next = next
}
func InitExportData() exportStepHandler {
	emailZipHandler := &exportEmailZip{}

	zipHandler := &exportZipFiles{}
	zipHandler.setNext(emailZipHandler)

	parseFileHandler := &exportParseIntoFile{}
	parseFileHandler.setNext(zipHandler)

	fetchUserHandler := &exportFetchUsers{}
	fetchUserHandler.setNext(parseFileHandler)

	return fetchUserHandler
}

func main() {
	exportDataHandler := InitExportData()
	exportDataHandler.execute(&exportRequest{
		requestid: 1,
	})
}
