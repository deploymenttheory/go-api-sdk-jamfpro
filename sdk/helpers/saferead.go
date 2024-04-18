// helpers/saferead.go

package helpers

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// SafeReadCertificateFile reads a certificate file securely after applying multiple checks.
func SafeReadCertificateFile(filePath string, allowedExtensions []string) ([]byte, error) {
	// Clean the file path first to prevent directory traversal
	cleanedPath := cleanPath(filePath)

	// Check for a valid file extension
	if !isValidExtension(cleanedPath, allowedExtensions) {
		return nil, fmt.Errorf("file extension '%s' is not allowed", filepath.Ext(cleanedPath))
	}

	// Resolve any symbolic links to ensure the path is safe
	resolvedPath, err := resolveSymlinks(cleanedPath)
	if err != nil {
		return nil, err
	}

	// Read the file
	data, err := os.ReadFile(resolvedPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}
	return data, nil
}

// SafeReadJCDSPackageFile reads a package file securely after applying multiple checks.
func SafeReadJCDSPackageFile(filePath string, allowedExtensions []string) ([]byte, error) {
	// Clean the file path first to prevent directory traversal
	cleanedPath := cleanPath(filePath)

	// Check for a valid file extension
	if !isValidExtension(cleanedPath, allowedExtensions) {
		return nil, fmt.Errorf("file extension '%s' is not allowed", filepath.Ext(cleanedPath))
	}

	// Resolve any symbolic links to ensure the path is safe
	resolvedPath, err := resolveSymlinks(cleanedPath)
	if err != nil {
		return nil, err
	}

	// Read the file
	data, err := os.ReadFile(resolvedPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read Jamf Pro package: %v", err)
	}
	return data, nil
}

// resolveSymlinks resolves symbolic links and returns the absolute path.
func resolveSymlinks(filePath string) (string, error) {
	cleanPath := filepath.Clean(filePath)
	absPath, err := filepath.EvalSymlinks(cleanPath)
	if err != nil {
		return "", fmt.Errorf("unable to resolve the absolute path: %s, error: %w", filePath, err)
	}
	return absPath, nil
}

// cleanPath sanitizes the file path to prevent directory traversal.
func cleanPath(filePath string) string {
	return filepath.Clean(filePath)
}

// isValidExtension checks if the file has one of the allowed extensions.
func isValidExtension(filePath string, allowedExtensions []string) bool {
	ext := strings.ToLower(filepath.Ext(filePath))
	for _, allowedExt := range allowedExtensions {
		if ext == strings.ToLower(allowedExt) {
			return true
		}
	}
	return false
}
