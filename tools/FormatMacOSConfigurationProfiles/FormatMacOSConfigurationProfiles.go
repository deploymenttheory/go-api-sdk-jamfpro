package main

import (
	"fmt"
	"html"
	"os"
	"path/filepath"
	"strings"

	"howett.net/plist"
)

// FormatMacOSConfigurationProfiles walks through the specified source directory,
// processes each .mobileconfig file by formatting its XML content, and saves the
// formatted files to the specified destination directory with an added suffix.
// This function prompts the user to input the source directory, destination directory,
// and file suffix for the processed files.
func main() {
	var sourceDir, destDir, suffix string

	fmt.Print("Enter the source directory where your existing configuration profiles exist: \n")
	fmt.Scan(&sourceDir)

	fmt.Print("Enter the destination directory where you want to save the formatted configuration profiles: \n")
	fmt.Scan(&destDir)

	fmt.Print("Enter a file suffix if you wish to add one: \n")
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

	// Read the file content
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Check if data is empty
	if len(data) == 0 {
		fmt.Println("Empty file, skipping:", path)
		return
	}

	// Decode HTML entities
	decodedData := html.UnescapeString(string(data))

	// Unmarshal the plist content
	var plistData interface{}
	_, err = plist.Unmarshal([]byte(decodedData), &plistData)
	if err != nil {
		fmt.Println("Error unmarshalling plist:", err)
		return
	}

	// Call to reformatPlistData
	reformattedData, err := reformatPlistData(plistData)
	if err != nil {
		fmt.Println("Error in reformatting plist data:", err)
		return
	}

	// Marshal back to bytes with indentation
	prettyPlist, err := plist.MarshalIndent(reformattedData, plist.XMLFormat, "\t")
	if err != nil {
		fmt.Println("Error marshalling plist:", err)
		return
	}

	// Define the new file path
	filename := filepath.Base(path)
	newPath := filepath.Join(destDir, strings.TrimSuffix(filename, ".mobileconfig")+suffix+".mobileconfig")

	// Write the formatted data to the new file
	err = os.WriteFile(newPath, prettyPlist, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println("Processed and saved:", newPath)
}

// debugLog prints debug messages if debugging is enabled.
// This can be controlled by a debug flag in production code.
func debugLog(message string) {
	fmt.Println("DEBUG:", message)
}

// reformatPlistData takes the original plist data and reformats it
// into the desired structure. It handles key-value pairs, dict, and array elements.
func reformatPlistData(originalData interface{}) (interface{}, error) {
	debugLog("Reformatting plist data")

	// Check if the original data is a dictionary
	if dict, ok := originalData.(map[string]interface{}); ok {
		newDict := make(map[string]interface{})
		for key, value := range dict {
			switch element := value.(type) {
			case []interface{}:
				// Process array elements
				newArray := make([]interface{}, 0)
				for _, item := range element {
					if itemDict, ok := item.(map[string]interface{}); ok {
						// Process each dictionary in the array
						newItemDict := processDict(itemDict)
						newArray = append(newArray, newItemDict)
					} else {
						newArray = append(newArray, item)
					}
				}
				newDict[key] = newArray
			case map[string]interface{}:
				// Process nested dictionary
				newDict[key] = processDict(element)
			default:
				// Handle other types
				newDict[key] = element
			}
		}
		return newDict, nil
	}

	return nil, fmt.Errorf("unexpected data format")
}

// processDict applies specific transformations to a dictionary element.
func processDict(dict map[string]interface{}) map[string]interface{} {
	newDict := make(map[string]interface{})

	for k, v := range dict {
		// Process 'PayloadContent' key specifically
		if k == "PayloadContent" {
			newPayloadContent := processPayloadContent(v)
			newDict[k] = newPayloadContent
		} else {
			// Copy other keys and values as is
			newDict[k] = v
		}
	}

	return newDict
}

// processPayloadContent processes the 'PayloadContent' section of the plist.
func processPayloadContent(value interface{}) interface{} {
	if array, ok := value.([]interface{}); ok {
		newArray := make([]interface{}, 0)
		for _, item := range array {
			if itemDict, ok := item.(map[string]interface{}); ok {
				// Process each dictionary in the array
				newItemDict := processPayloadItem(itemDict)
				newArray = append(newArray, newItemDict)
			}
		}
		return newArray
	}
	return value
}

// processPayloadItem processes each item in the 'PayloadContent' array.
func processPayloadItem(dict map[string]interface{}) map[string]interface{} {
	newItemDict := make(map[string]interface{})

	// Add necessary keys with default values if they are missing
	necessaryKeys := []string{"PayloadDescription", "PayloadDisplayName", "PayloadEnabled", "PayloadIdentifier", "PayloadOrganization", "PayloadType", "PayloadUUID", "PayloadVersion"}
	for _, key := range necessaryKeys {
		if _, exists := dict[key]; !exists {
			switch key {
			case "PayloadDescription", "PayloadDisplayName", "PayloadIdentifier", "PayloadOrganization", "PayloadType", "PayloadUUID":
				newItemDict[key] = ""
			case "PayloadEnabled":
				newItemDict[key] = true // Assuming default is true, adjust as needed
			case "PayloadVersion":
				newItemDict[key] = 1 // Assuming default version is 1, adjust as needed
			}
		}
	}

	// Copy existing keys and values
	for k, v := range dict {
		newItemDict[k] = v
	}

	return newItemDict
}
