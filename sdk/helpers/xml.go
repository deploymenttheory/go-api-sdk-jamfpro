// helpers/xml.go

// For utility/helper functions to support the jamf pro package
package helpers

import (
	"strings"
)

// ConvertToXMLSafeString replaces disallowed XML characters in a string with their corresponding XML entity references.
// This function is useful for preparing a string to be safely included in an XML document.
func ConvertToXMLSafeString(s string) string {
	// Define a map of disallowed characters and their XML entity equivalents.
	replacements := map[string]string{
		"&":  "&amp;",
		"<":  "&lt;",
		">":  "&gt;",
		"'":  "&apos;",
		"\"": "&quot;",
	}

	for key, val := range replacements {
		s = strings.ReplaceAll(s, key, val)
	}

	return s
}

// ConvertFromXMLSafeString reverses the process of ConvertToXMLSafeString.
// It replaces XML entity references in a string back to their original characters.
// This is useful when reading XML data that contains entity references and converting them back to normal characters.
func ConvertFromXMLSafeString(s string) string {
	replacements := map[string]string{
		"&amp;":  "&",
		"&lt;":   "<",
		"&gt;":   ">",
		"&apos;": "'",
		"&quot;": "\"",
	}

	for key, val := range replacements {
		s = strings.ReplaceAll(s, key, val)
	}

	return s
}

// EnsureXMLSafeString checks if a string contains disallowed XML characters.
// If it does, it converts the string to an XML-safe format using ConvertToXMLSafeString.
// This function is useful for ensuring that strings are safe for inclusion in XML documents.
func EnsureXMLSafeString(s string) string {
	disallowedChars := []string{"&", "<", ">", "'", "\""}

	for _, char := range disallowedChars {
		if strings.Contains(s, char) {
			return ConvertToXMLSafeString(s)
		}
	}

	return s
}
