package schemaprocessing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"os"
	"strings"

	"github.com/mitchellh/mapstructure"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// ProcessJSONFile reads a JSON file, decodes it into a Go struct, and returns the result.
// The result is returned as an interface{} and can be further processed as needed.
func ProcessJSONFile(filePath string, result interface{}) error {
	// Read the JSON file
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Read the file contents into a byte slice
	byteValue, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	// Unmarshal the JSON data into a map[string]interface{}
	var data map[string]interface{}
	if err := json.Unmarshal(byteValue, &data); err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	// Decode the map into the provided result struct using mapstructure
	if err := mapstructure.Decode(data, result); err != nil {
		return fmt.Errorf("failed to decode data: %w", err)
	}

	return nil
}

// ParseJSONSchema parses the JSON schema and generates Go struct definitions
func ParseJSONSchema(schema []byte) (string, error) {
	var schemaData map[string]interface{}
	if err := json.Unmarshal(schema, &schemaData); err != nil {
		return "", fmt.Errorf("failed to unmarshal JSON schema: %w", err)
	}

	structs, err := generateStructs("Root", schemaData)
	if err != nil {
		return "", fmt.Errorf("failed to generate structs: %w", err)
	}

	return structs, nil
}

// generateStructs generates Go struct definitions from the JSON schema
func generateStructs(structName string, schemaData map[string]interface{}) (string, error) {
	var structsBuilder strings.Builder

	// Add package name at the top of the file
	structsBuilder.WriteString("package generatedstructs\n\n")

	structDef, err := generateStruct(structName, schemaData)
	if err != nil {
		return "", err
	}
	structsBuilder.WriteString(structDef)
	structsBuilder.WriteString("\n\n")

	for key, value := range schemaData {
		switch value := value.(type) {
		case map[string]interface{}:
			structDef, err := generateStruct(cases.Title(language.English).String(key), value)
			if err != nil {
				return "", err
			}
			structsBuilder.WriteString(structDef)
			structsBuilder.WriteString("\n\n")
		case []interface{}:
			// Handle arrays of objects
			if len(value) > 0 {
				if elem, ok := value[0].(map[string]interface{}); ok {
					structDef, err := generateStruct(cases.Title(language.English).String(key), elem)
					if err != nil {
						return "", err
					}
					structsBuilder.WriteString(structDef)
					structsBuilder.WriteString("\n\n")
				}
			}
		default:
			// Handle primitive types at the top level
			fieldType, err := getFieldType(value)
			if err != nil {
				return "", err
			}
			fieldName := cases.Title(language.English).String(key)
			structsBuilder.WriteString(fmt.Sprintf("%s %s `json:\"%s\" xml:\"%s\"`\n", fieldName, fieldType, key, key))
		}
	}

	return structsBuilder.String(), nil
}

// generateStruct generates a Go struct definition from a JSON object
func generateStruct(structName string, structData map[string]interface{}) (string, error) {
	var fieldsBuilder strings.Builder

	titleCaser := cases.Title(language.English)

	for key, value := range structData {
		fieldType, err := getFieldType(value)
		if err != nil {
			return "", err
		}

		fieldName := titleCaser.String(key)
		fieldsBuilder.WriteString(fmt.Sprintf("%s %s `json:\"%s\" xml:\"%s\"`\n", fieldName, fieldType, key, key))
	}

	structTemplate := `type {{.StructName}} struct {
	{{.Fields}}
}`

	tmpl, err := template.New("struct").Parse(structTemplate)
	if err != nil {
		return "", fmt.Errorf("failed to parse struct template: %w", err)
	}

	structDef := &bytes.Buffer{}
	err = tmpl.Execute(structDef, map[string]string{
		"StructName": structName,
		"Fields":     fieldsBuilder.String(),
	})
	if err != nil {
		return "", fmt.Errorf("failed to execute struct template: %w", err)
	}

	return structDef.String(), nil
}

// getFieldType returns the Go type of a field based on its value
func getFieldType(value interface{}) (string, error) {
	switch v := value.(type) {
	case map[string]interface{}:
		return "struct", nil
	case []interface{}:
		if len(v) > 0 {
			elemType, err := getFieldType(v[0])
			if err != nil {
				return "", err
			}
			return "[]" + elemType, nil
		}
		return "[]interface{}", nil
	case string:
		return "string", nil
	case float64:
		return "float64", nil
	case bool:
		return "bool", nil
	case nil:
		return "interface{}", nil
	default:
		return "", fmt.Errorf("unknown field type: %T", v)
	}
}
