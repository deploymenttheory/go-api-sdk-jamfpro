package main

import (
	"bufio"
	"bytes"
	"fmt"
	"html"
	"os"
	"path/filepath"
	"strings"

	"howett.net/plist"
)

// RemoveEscapedCharacters removes escaped characters from a plist / .mobileconfig file
func RemoveEscapedCharacters() {
	reader := bufio.NewReader(os.Stdin)

	// Prompt for and read the source plist file path
	fmt.Print("Enter the path to the source plist file: ")
	sourceFilePath, err := readInput(reader)
	if err != nil {
		fmt.Println("Error reading source file path:", err)
		return
	}

	// Prompt for and read the destination folder path
	fmt.Print("Enter the destination folder path: ")
	destFolderPath, err := readInput(reader)
	if err != nil {
		fmt.Println("Error reading destination folder path:", err)
		return
	}

	// Read and decode the plist file content
	data, err := decodePlistFile(sourceFilePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Re-encode to XML to ensure correct formatting
	newPlist, err := plist.MarshalIndent(data, plist.XMLFormat, "\t")
	if err != nil {
		fmt.Println("Error marshaling plist:", err)
		return
	}

	// Construct the output file path and write the new plist content
	outputFilePath := filepath.Join(destFolderPath, filepath.Base(sourceFilePath))
	if err := writeFile(outputFilePath, newPlist); err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Printf("Plist file has been reformatted and written successfully to '%s'.\n", outputFilePath)
}

// ReadInput simplifies reading a line of text input
func readInput(reader *bufio.Reader) (string, error) {
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return filepath.Clean(strings.TrimSpace(input)), nil
}

// DecodePlistFile handles reading and unmarshaling a plist file
func decodePlistFile(filePath string) (interface{}, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file '%s': %v", filePath, err)
	}
	defer file.Close()

	// Read the file content
	var buf bytes.Buffer
	if _, err := buf.ReadFrom(file); err != nil {
		return nil, fmt.Errorf("error reading file '%s': %v", filePath, err)
	}
	content := buf.Bytes()

	// Check if content needs unescaping
	if bytes.Contains(content, []byte("&lt;")) {
		content = []byte(UnescapeXML(string(content)))
	}

	decoder := plist.NewDecoder(bytes.NewReader(content))
	var data interface{}
	if err := decoder.Decode(&data); err != nil {
		return nil, fmt.Errorf("error decoding plist: %v", err)
	}
	return data, nil
}

// UnescapeXML takes an XML string with HTML entities and returns the unescaped XML.
func UnescapeXML(encodedXML string) string {
	return html.UnescapeString(encodedXML)
}

// WriteFile handles the file writing process
func writeFile(filePath string, data []byte) error {
	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return fmt.Errorf("error writing file '%s': %v", filePath, err)
	}
	return nil
}
