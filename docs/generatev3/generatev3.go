package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type FunctionDoc struct {
	FunctionName string
	Method       string
	Path         string
	Description  string
}

func main() {
	startPath, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current working directory: %v\n", err)
		return
	}

	sdkPath, err := findSDKDirectory(startPath)
	if err != nil {
		fmt.Printf("Error finding sdk directory: %v\n", err)
		return
	}

	exportPath := "export"

	// Create the export directory if it doesn't exist
	if _, err := os.Stat(exportPath); os.IsNotExist(err) {
		if err := os.MkdirAll(exportPath, 0755); err != nil {
			fmt.Printf("Error creating export directory: %v\n", err)
			return
		}
	}

	// Process the SDK directory files
	if err := filepath.Walk(sdkPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
			processFile(path, exportPath)
		}
		return nil
	}); err != nil {
		fmt.Printf("Error walking through files: %v\n", err)
	}

	fmt.Println("Processing complete. Markdown files generated in the 'export' folder.")
}

func findSDKDirectory(startPath string) (string, error) {
	for {
		fmt.Printf("Searching for 'sdk' directory in: %s\n", startPath)
		sdkPath := filepath.Join(startPath, "sdk")
		if _, err := os.Stat(sdkPath); err == nil {
			fmt.Printf("Found 'sdk' directory in: %s\n", sdkPath)
			return sdkPath, nil
		}

		if filepath.Base(startPath) == "go-api-sdk-jamfpro" {
			return "", fmt.Errorf("reached the root of the repository without finding sdk directory")
		}

		parent := filepath.Dir(startPath)
		if parent == startPath {
			return "", fmt.Errorf("sdk directory not found")
		}

		startPath = parent
	}
}

func processFile(filePath, exportPath string) {
	fmt.Printf("Processing file: %s\n", filePath)
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var currentDoc *FunctionDoc
	var docs []*FunctionDoc
	inCommentBlock := false

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if strings.HasPrefix(line, "/*") {
			inCommentBlock = true
		}

		if inCommentBlock {
			if strings.HasPrefix(line, "Function:") {
				if currentDoc != nil {
					docs = append(docs, currentDoc)
				}
				currentDoc = &FunctionDoc{}
				currentDoc.FunctionName = strings.TrimSpace(strings.TrimPrefix(line, "Function:"))
			} else if strings.HasPrefix(line, "Method:") {
				if currentDoc != nil {
					currentDoc.Method = strings.TrimSpace(strings.TrimPrefix(line, "Method:"))
				}
			} else if strings.HasPrefix(line, "Path:") {
				if currentDoc != nil {
					currentDoc.Path = strings.TrimSpace(strings.TrimPrefix(line, "Path:"))
				}
			} else if strings.HasPrefix(line, "Description:") {
				if currentDoc != nil {
					currentDoc.Description = strings.TrimSpace(strings.TrimPrefix(line, "Description:"))
				}
			}
		}

		if strings.HasSuffix(line, "*/") {
			inCommentBlock = false
			if currentDoc != nil {
				docs = append(docs, currentDoc)
				currentDoc = nil
			}
		}
	}
	if currentDoc != nil {
		docs = append(docs, currentDoc)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	if len(docs) > 0 {
		fmt.Printf("Found matching documentation in file: %s\n", filePath)
		groupedDocs := groupDocsByPath(docs)
		for endpoint, funcs := range groupedDocs {
			writeMarkdown(exportPath, endpoint, funcs)
		}
	}
}

func groupDocsByPath(docs []*FunctionDoc) map[string][]*FunctionDoc {
	grouped := make(map[string][]*FunctionDoc)
	for _, doc := range docs {
		endpoint := getEndpoint(doc.Path)
		grouped[endpoint] = append(grouped[endpoint], doc)
	}
	return grouped
}

func getEndpoint(path string) string {
	re := regexp.MustCompile(`\/JSSResource\/([^\/]+)`)
	matches := re.FindStringSubmatch(path)
	if len(matches) > 1 {
		return matches[1]
	}
	return "unknown"
}

func writeMarkdown(exportPath, endpoint string, docs []*FunctionDoc) {
	filename := filepath.Join(exportPath, fmt.Sprintf("%s.md", endpoint))
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Error creating markdown file: %v\n", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	writer.WriteString(fmt.Sprintf("# %s API Documentation\n\n", strings.Title(endpoint)))

	// Calculate the maximum length of the path column
	maxPathLength := 0
	for _, doc := range docs {
		if len(doc.Path) > maxPathLength {
			maxPathLength = len(doc.Path)
		}
	}

	writer.WriteString(fmt.Sprintf("| Function | Method | Path%s | Description |\n", strings.Repeat(" ", maxPathLength-4)))
	writer.WriteString(fmt.Sprintf("|----------|--------|-%s|-------------|\n", strings.Repeat("-", maxPathLength+2)))

	for _, doc := range docs {
		writer.WriteString(fmt.Sprintf("| %s | %s | %s%s | %s |\n",
			doc.FunctionName,
			doc.Method,
			doc.Path,
			strings.Repeat(" ", maxPathLength-len(doc.Path)),
			doc.Description))
	}
}
