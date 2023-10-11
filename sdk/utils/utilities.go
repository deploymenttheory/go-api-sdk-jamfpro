// utilities.go
// For utility/helper functions to support from the main package
package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
	"path/filepath"
)

func Base64EncodeCertificate(certPath string) (string, error) {
	// Read the certificate file
	data, err := os.ReadFile(certPath)
	if err != nil {
		return "", fmt.Errorf("failed to read certificate file: %v", err)
	}

	// Base64 encode the file's content
	encoded := base64.StdEncoding.EncodeToString(data)

	return encoded, nil
}

// GetImageContentType determines the content type based on file extension
func GetImageContentType(filePath string) string {
	ext := filepath.Ext(filePath)
	switch ext {
	case ".png":
		return "image/png"
	case ".jpg", ".jpeg":
		return "image/jpeg"
	default:
		return "application/octet-stream"
	}
}

// Base64Encode encodes the provided data into a base64 string and provides details about the encoding process.
func Base64Encode(data []byte) (string, error) {
	// Check if the provided data is empty
	if len(data) == 0 {
		return "", fmt.Errorf("no data provided for encoding")
	}

	// Encode the data to base64
	encoded := base64.StdEncoding.EncodeToString(data)

	// Check if the encoding process was successful
	if encoded == "" {
		return "", fmt.Errorf("failed to encode the data")
	}

	// Return the encoded data
	return encoded, nil
}

// UnmarshalJSONData unmarshals binary data into the given output structure.
func UnmarshalJSONData(data []byte, out interface{}) error {
	return json.Unmarshal(data, out)
}

// DumpRequestToFile dumps the given request to a specified file.
func DumpRequestToFile(req *http.Request, filename string) error {
	// Dump the request details
	dump, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		return fmt.Errorf("error dumping request: %v", err)
	}

	// Open a file for writing
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("cannot create file: %v", err)
	}
	defer file.Close()

	// Write the dumped request to the file
	_, err = file.WriteString(string(dump))
	if err != nil {
		return fmt.Errorf("cannot write to file: %v", err)
	}

	return nil
}

/*
// Print request headers for troubleshooting
func PrintRequestHeaders(req *http.Request) {
	fmt.Println("Request Headers:")
	for name, values := range req.Header {
		// Each value is a slice of strings since headers can be repeated.
		for _, value := range values {
			fmt.Printf("%s: %s\n", name, value)
		}
	}
}
*/
// PrettyPrintStruct prints the structure in a pretty format
func PrettyPrintStruct(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return err.Error()
	}

	var out map[string]interface{}
	if err := json.Unmarshal(b, &out); err != nil {
		return err.Error()
	}

	pretty, err := json.MarshalIndent(out, "", "  ")
	if err != nil {
		return err.Error()
	}
	return string(pretty)
}
