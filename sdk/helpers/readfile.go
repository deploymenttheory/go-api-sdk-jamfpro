// helpers/readfile.go
package helpers

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
)

// Base64EncodeCertificate reads a certificate file and returns its content as a base64-encoded string.
func Base64EncodeCertificate(certPath string) (string, error) {
	allowedExtensions := []string{".crt", ".pem", ".cer"} // Define allowed certificate file extensions

	// Use the secure file reading helper
	data, err := SafeReadCertificateFile(certPath, allowedExtensions)
	if err != nil {
		return "", fmt.Errorf("failed to read certificate file securely: %v", err)
	}

	// Base64 encode the file's content
	encoded := base64.StdEncoding.EncodeToString(data)
	return encoded, nil
}

// ReadJCDSPackageTypes returns a reader and size for a package file securely after applying multiple checks.
func ReadJCDSPackageTypes(filePath string) (io.Reader, int64, error) {
	allowedExtensions := []string{".pkg", ".dmg", ".zip"} // Define allowed package file extensions

	// Use the secure file reading helper
	data, err := SafeReadJCDSPackageFile(filePath, allowedExtensions)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to read package file securely: %v", err)
	}

	size := int64(len(data))
	reader := bytes.NewReader(data)

	return reader, size, nil
}
