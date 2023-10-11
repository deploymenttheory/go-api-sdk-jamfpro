// printers.go
// Jamf Pro Classic Api
// Classic API requires the structs to support both XML and JSON.

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriAPIPrinters = "/JSSResource/printers"

type ResponsePrinter struct {
	ID          int    `json:"id,omitempty" xml:"id,omitempty"`
	Name        string `json:"name" xml:"name"`
	Category    string `json:"category" xml:"category"`
	URI         string `json:"uri" xml:"uri"`
	CUPSName    string `json:"CUPS_name" xml:"CUPS_name"`
	Location    string `json:"location" xml:"location"`
	Model       string `json:"model" xml:"model"`
	Info        string `json:"info,omitempty" xml:"info,omitempty"`
	Notes       string `json:"notes,omitempty" xml:"notes,omitempty"`
	MakeDefault bool   `json:"make_default" xml:"make_default"`
	UseGeneric  bool   `json:"use_generic" xml:"use_generic"`
	PPD         string `json:"ppd" xml:"ppd"`
	PPDPath     string `json:"ppd_path" xml:"ppd_path"`
	PPDContents string `json:"ppd_contents" xml:"ppd_contents"`
}

type ResponsePrintersList struct {
	Size     int             `json:"size" xml:"size"`
	Printers []PrinterDetail `json:"printer" xml:"printer"`
}

type PrinterDetail struct {
	ID   int    `json:"id" xml:"id"`
	Name string `json:"name" xml:"name"`
}

// GetPrinterByID gets a printer by its id
func (c *Client) GetPrinterByID(id int) (*ResponsePrinter, error) {
	url := fmt.Sprintf("%s/id/%d", uriAPIPrinters, id)

	var printer ResponsePrinter
	if err := c.DoRequest("GET", url, nil, nil, &printer); err != nil {
		return nil, fmt.Errorf("failed to execute request: %v", err)
	}

	return &printer, nil
}

// GetPrinterByName gets a printer by its name
func (c *Client) GetPrinterByName(name string) (*ResponsePrinter, error) {
	url := fmt.Sprintf("%s/name/%s", uriAPIPrinters, name)

	var printer ResponsePrinter
	if err := c.DoRequest("GET", url, nil, nil, &printer); err != nil {
		return nil, fmt.Errorf("failed to execute request: %v", err)
	}

	return &printer, nil
}

// GetPrinters gets a list of all printers
func (c *Client) GetPrinters() ([]ResponsePrintersList, error) {
	url := uriAPIPrinters

	var printersList []ResponsePrintersList
	if err := c.DoRequest("GET", url, nil, nil, &printersList); err != nil {
		return nil, fmt.Errorf("failed to execute request: %v", err)
	}

	return printersList, nil
}

// CreatePrinter creates a new Jamf Pro Printer.
func (c *Client) CreatePrinter(printer *ResponsePrinter) (*ResponsePrinter, error) {
	url := fmt.Sprintf("%s/id/0", uriAPIPrinters) // ID 0 is typically used for creation in many APIs

	reqBody := &struct {
		XMLName xml.Name `xml:"printer"`
		*ResponsePrinter
	}{
		ResponsePrinter: printer,
	}

	var responsePrinter ResponsePrinter
	if err := c.DoRequest("POST", url, reqBody, nil, &responsePrinter); err != nil {
		return nil, fmt.Errorf("failed to create printer: %v", err)
	}

	return &responsePrinter, nil
}

// UpdatePrinterByID updates an existing Jamf Pro Printer by ID.
func (c *Client) UpdatePrinterByID(id int, printer *ResponsePrinter) (*ResponsePrinter, error) {
	url := fmt.Sprintf("%s/id/%d", uriAPIPrinters, id)

	reqBody := &struct {
		XMLName xml.Name `xml:"printer"`
		*ResponsePrinter
	}{
		ResponsePrinter: printer,
	}

	var responsePrinter ResponsePrinter
	if err := c.DoRequest("PUT", url, reqBody, nil, &responsePrinter); err != nil {
		return nil, fmt.Errorf("failed to update printer by ID: %v", err)
	}

	return &responsePrinter, nil
}

// UpdatePrinterByName updates an existing Jamf Pro Printer by Name.
func (c *Client) UpdatePrinterByName(name string, printer *ResponsePrinter) (*ResponsePrinter, error) {
	url := fmt.Sprintf("%s/name/%s", uriAPIPrinters, name)

	reqBody := &struct {
		XMLName xml.Name `xml:"printer"`
		*ResponsePrinter
	}{
		ResponsePrinter: printer,
	}

	var responsePrinter ResponsePrinter
	if err := c.DoRequest("PUT", url, reqBody, nil, &responsePrinter); err != nil {
		return nil, fmt.Errorf("failed to update printer by Name: %v", err)
	}

	return &responsePrinter, nil
}

// DeletePrinterByID deletes an existing Jamf Pro Printer by ID.
func (c *Client) DeletePrinterByID(id int) error {
	url := fmt.Sprintf("%s/id/%d", uriAPIPrinters, id)

	if err := c.DoRequest("DELETE", url, nil, nil, nil); err != nil {
		return fmt.Errorf("failed to delete printer by ID: %v", err)
	}

	return nil
}

// DeletePrinterByName deletes an existing Jamf Pro Printer by Name.
func (c *Client) DeletePrinterByName(name string) error {
	url := fmt.Sprintf("%s/name/%s", uriAPIPrinters, name)

	if err := c.DoRequest("DELETE", url, nil, nil, nil); err != nil {
		return fmt.Errorf("failed to delete printer by Name: %v", err)
	}

	return nil
}
