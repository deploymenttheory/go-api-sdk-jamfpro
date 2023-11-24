package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// FormatMacOSConfigurationProfiles walks through the specified source directory,
// processes each .mobileconfig file by formatting its XML content, and saves the
// formatted files to the specified destination directory with an added suffix.
// This function prompts the user to input the source directory, destination directory,
// and file suffix for the processed files.

func main() {
	var sourceDir, destDir, suffix string

	fmt.Print("Enter the source directory where your existing configuration profiles exist: ")
	fmt.Scan(&sourceDir)

	fmt.Print("Enter the destination directory where you want to save the formatted configuration profiles: ")
	fmt.Scan(&destDir)

	fmt.Print("Enter a file suffix if you wish to add one: ")
	fmt.Scan(&suffix)

	err := filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if strings.HasSuffix(path, ".mobileconfig") {
			processMacOSConfigurationProfileFile(path, destDir, suffix)
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error processing files:", err)
		return
	}

	fmt.Println("Processing completed.")
}

// processMacOSConfigurationProfileFile takes the path of a .mobileconfig file, reads its content,
// and formats the XML structure. The formatted XML content is then saved to a new file in the
// specified destination directory, optionally appending a suffix to the file name.
// This function is called for each .mobileconfig file found in the source directory by
// FormatMacOSConfigurationProfiles.
//
// Args:
// path (string): The file path of the .mobileconfig file to be processed.
// destDir (string): The destination directory where the formatted file will be saved.
// suffix (string): An optional suffix to be added to the file name of the formatted file.
//
// The function reads the XML content from the specified file, unmarshals it to ensure
// proper XML structure, and then marshals it back with proper indentation. The result
// is written to a new file in the destination directory with the appropriate suffix.
func processMacOSConfigurationProfileFile(path, destDir, suffix string) {
	fmt.Println("Processing:", path)
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var prettyXML string
	err = xml.Unmarshal(data, &prettyXML)
	if err != nil {
		fmt.Println("Error unmarshalling XML:", err)
		return
	}

	output, err := xml.MarshalIndent(prettyXML, "", "    ")
	if err != nil {
		fmt.Println("Error marshalling XML:", err)
		return
	}

	filename := filepath.Base(path)
	newPath := filepath.Join(destDir, strings.TrimSuffix(filename, ".mobileconfig")+suffix+".mobileconfig")

	err = os.WriteFile(newPath, output, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println("Processed and saved:", newPath)
}
