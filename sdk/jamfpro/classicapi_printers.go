// classicapi_printers.go
// Jamf Pro Classic Api - Printers
// api reference: https://developer.jamf.com/jamf-pro/reference/printers
// Jamf Pro Classic Api requires the structs to support an XML data structure.

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriPrinters = "/JSSResource/printers"

// List

// ResponsePrintersList represents the response for a list of printers.
type ResponsePrintersList struct {
	Size    int                `xml:"size"`
	Printer []PrintersListItem `xml:"printer"`
}

type PrintersListItem struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

type ResponsePrinterCreateAndUpdate struct {
	ID int `xml:"id"`
}

// Resource

// ResourcePrinter represents the detailed structure of a single printer.
type ResourcePrinter struct {
	ID          int    `xml:"id"`
	Name        string `xml:"name"`
	Category    string `xml:"category"`
	URI         string `xml:"uri"`
	CUPSName    string `xml:"CUPS_name"`
	Location    string `xml:"location"`
	Model       string `xml:"model"`
	Info        string `xml:"info"`
	Notes       string `xml:"notes"`
	MakeDefault bool   `xml:"make_default"`
	UseGeneric  bool   `xml:"use_generic"`
	PPD         string `xml:"ppd"`
	PPDPath     string `xml:"ppd_path"`
	PPDContents string `xml:"ppd_contents"`
}

// CRUD

// GetPrinters retrieves a serialized list of printers.
func (c *Client) GetPrinters() (*ResponsePrintersList, error) {
	endpoint := uriPrinters

	var printers ResponsePrintersList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &printers, c.HTTP.Logger)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "printers", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &printers, nil
}

// GetPrinterByID fetches a specific printer by its ID.
func (c *Client) GetPrinterByID(id int) (*ResourcePrinter, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriPrinters, id)

	var printer ResourcePrinter
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &printer, c.HTTP.Logger)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "printer", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &printer, nil
}

// GetPrinterByName fetches a specific printer by its name.
func (c *Client) GetPrinterByName(name string) (*ResourcePrinter, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriPrinters, name)

	var printer ResourcePrinter
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &printer, c.HTTP.Logger)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByName, "printer", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &printer, nil
}

// CreatePrinters creates a new printer on the Jamf Pro server.
func (c *Client) CreatePrinter(printer *ResourcePrinter) (*ResponsePrinterCreateAndUpdate, error) {
	endpoint := fmt.Sprintf("%s/id/0", uriPrinters)

	// Wrap the printer with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"printer"`
		*ResourcePrinter
	}{
		ResourcePrinter: printer,
	}

	var responsePrinter ResponsePrinterCreateAndUpdate
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &responsePrinter, c.HTTP.Logger)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "printer", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responsePrinter, nil
}

// UpdatePrinterByID updates a printer by its ID.
func (c *Client) UpdatePrinterByID(id int, printer *ResourcePrinter) (*ResponsePrinterCreateAndUpdate, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriPrinters, id)

	// Wrap the printer with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"printer"`
		*ResourcePrinter
	}{
		ResourcePrinter: printer,
	}

	var responsePrinter ResponsePrinterCreateAndUpdate
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &responsePrinter, c.HTTP.Logger)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByID, "printer", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responsePrinter, nil
}

// UpdatePrinterByName updates a printer by its name.
func (c *Client) UpdatePrinterByName(name string, printer *ResourcePrinter) (*ResponsePrinterCreateAndUpdate, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriPrinters, name)

	// Wrap the printer with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"printer"`
		*ResourcePrinter
	}{
		ResourcePrinter: printer,
	}

	var responsePrinter ResponsePrinterCreateAndUpdate
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &responsePrinter, c.HTTP.Logger)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByName, "printer", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responsePrinter, nil
}

// DeletePrinterByID deletes a printer by its ID.
func (c *Client) DeletePrinterByID(id int) error {
	endpoint := fmt.Sprintf("%s/id/%d", uriPrinters, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil, c.HTTP.Logger)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByID, "printer", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeletePrinterByName deletes a printer by its name.
func (c *Client) DeletePrinterByName(name string) error {
	endpoint := fmt.Sprintf("%s/name/%s", uriPrinters, name)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil, c.HTTP.Logger)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByName, "printer", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
