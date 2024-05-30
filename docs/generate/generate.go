package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type OAS3Schema struct {
	Paths map[string]map[string]PathDetails `json:"paths"`
}

type PathDetails struct {
	Tags        []string `json:"tags"`
	Summary     string   `json:"summary"`
	Description string   `json:"description"`
}

func extractPaths(oas3Schema OAS3Schema) map[string][]map[string]string {
	groupedPaths := make(map[string][]map[string]string)

	for path, methods := range oas3Schema.Paths {
		for method, details := range methods {
			apiTag := "default"
			if len(details.Tags) > 0 {
				apiTag = details.Tags[0]
			}
			pathInfo := map[string]string{
				"path":        path,
				"method":      strings.ToUpper(method),
				"summary":     details.Summary,
				"description": details.Description,
			}
			groupedPaths[apiTag] = append(groupedPaths[apiTag], pathInfo)
		}
	}

	return groupedPaths
}

func generateMarkdownTable(groupedPaths map[string][]map[string]string) map[string]string {
	markdownTables := make(map[string]string)

	for api, paths := range groupedPaths {
		var markdown []string
		markdown = append(markdown, fmt.Sprintf("## %s API\n", api))
		markdown = append(markdown, "| Path | Method | Summary | Description | go-api-sdk-jamfpro coverage |")
		markdown = append(markdown, "|------|--------|---------|-------------|-----------------------------|")

		for _, pathInfo := range paths {
			row := fmt.Sprintf("| %s | %s | %s | %s |  |", pathInfo["path"], pathInfo["method"], pathInfo["summary"], pathInfo["description"])
			markdown = append(markdown, row)
		}
		markdownTables[api] = strings.Join(markdown, "\n")
	}

	return markdownTables
}

func saveToMarkdown(markdownTables map[string]string, outputDir string) error {
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		if err := os.Mkdir(outputDir, 0755); err != nil {
			return fmt.Errorf("failed to create output directory: %v", err)
		}
	}

	for api, markdown := range markdownTables {
		filename := fmt.Sprintf("%s/%s.md", outputDir, strings.ReplaceAll(api, "/", "_"))
		if err := os.WriteFile(filename, []byte(markdown), 0644); err != nil {
			return fmt.Errorf("failed to write markdown file: %v", err)
		}
		fmt.Printf("Markdown file saved to %s\n", filename)
	}

	return nil
}

func main() {
	// Load the OAS3 schema
	schemaURL := "https://lbgsandbox.jamfcloud.com/api/schema/"
	resp, err := http.Get(schemaURL)
	if err != nil {
		fmt.Printf("failed to get schema: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("failed to read response body: %v\n", err)
		return
	}

	var oas3Schema OAS3Schema
	if err := json.Unmarshal(body, &oas3Schema); err != nil {
		fmt.Printf("failed to unmarshal JSON: %v\n", err)
		return
	}

	// Extract and group paths
	groupedPaths := extractPaths(oas3Schema)

	// Generate Markdown tables
	markdownTables := generateMarkdownTable(groupedPaths)

	// Save to separate GitHub-based Markdown files for each API tag
	outputDir := "api_paths"
	if err := saveToMarkdown(markdownTables, outputDir); err != nil {
		fmt.Printf("failed to save markdown files: %v\n", err)
	}
}
