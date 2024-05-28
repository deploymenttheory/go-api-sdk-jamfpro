package schemaprocessing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/mitchellh/mapstructure"
)

// OpenAPI is the top-level struct for the OAS3 standard
type OpenAPI struct {
	Openapi    string                   `json:"openapi"`
	Servers    []map[string]interface{} `json:"servers"`
	Security   []map[string]interface{} `json:"security"`
	Paths      map[string]interface{}   `json:"paths"`
	Components map[string]interface{}   `json:"components"`
}

// ProcessJSONFile reads a JSON file, decodes it into a Go struct, and returns the result.
func ProcessJSONFile(filePath string, result interface{}) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	return decodeJSONToStruct(byteValue, result)
}

// decodeJSONToStruct reads a JSON byte slice, decodes it into a Go struct, and returns the result.
func decodeJSONToStruct(data []byte, result interface{}) error {
	var jsonData map[string]interface{}
	if err := json.Unmarshal(data, &jsonData); err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	if err := mapstructure.Decode(jsonData, result); err != nil {
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

	return generateStructs(schemaData, "OpenAPI")
}

// generateStructs generates Go struct definitions for the schema
func generateStructs(schemaData map[string]interface{}, rootStructName string) (string, error) {
	var structsBuilder strings.Builder
	structsBuilder.WriteString("package generatedstructs\n\n")

	if err := generateStruct(&structsBuilder, rootStructName, schemaData); err != nil {
		return "", err
	}

	return structsBuilder.String(), nil
}

// generateStruct generates a Go struct definition from a JSON object
func generateStruct(structsBuilder *strings.Builder, parentStructName string, structData map[string]interface{}) error {
	structDef, err := createStructDefinition(parentStructName, structData)
	if err != nil {
		return err
	}
	structsBuilder.WriteString(structDef)
	structsBuilder.WriteString("\n\n")

	for key, value := range structData {
		if subMap, ok := value.(map[string]interface{}); ok {
			subStructName := normalizeToSafeNameGoStruct(parentStructName + key)
			if err := generateStruct(structsBuilder, subStructName, subMap); err != nil {
				return err
			}
		}
	}

	return nil
}

// createStructDefinition creates a Go struct definition string
func createStructDefinition(structName string, structData map[string]interface{}) (string, error) {
	var fieldsBuilder strings.Builder

	// Sort keys for consistent order
	keys := sortKeys(structData)

	for _, key := range keys {
		value := structData[key]
		fieldType, err := inferFieldType(key, value, structName)
		if err != nil {
			return "", err
		}

		fieldName := normalizeToSafeNameGoStruct(key)
		fieldsBuilder.WriteString(fmt.Sprintf("%s %s `json:\"%s\"`\n", fieldName, fieldType, key))
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

// inferFieldType infers the Go type of a field based on its value
func inferFieldType(fieldName string, value interface{}, parentStructName string) (string, error) {
	switch v := value.(type) {
	case map[string]interface{}:
		return normalizeToSafeNameGoStruct(parentStructName + fieldName), nil
	case []interface{}:
		if len(v) > 0 {
			elemType, err := inferFieldType(fieldName, v[0], parentStructName)
			if err != nil {
				return "", err
			}
			return "[]" + elemType, nil
		}
		return "[]interface{}", nil
	case string:
		return "string", nil
	case float64:
		if v == float64(int(v)) {
			return "int", nil
		}
		return "float64", nil
	case bool:
		return "bool", nil
	case nil:
		return "interface{}", nil
	default:
		return "", fmt.Errorf("unknown field type: %T", v)
	}
}

// normalizeToSafeNameGoStruct normalizes field names to be safe for Go struct field names
func normalizeToSafeNameGoStruct(fieldName string) string {
	// Replace invalid characters with a space
	reg := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	fieldName = reg.ReplaceAllString(fieldName, " ")

	// Split the string by spaces and capitalize each part
	parts := strings.Fields(fieldName)
	for i, part := range parts {
		parts[i] = strings.Title(part)
	}

	// Join the parts back together without spaces
	fieldName = strings.Join(parts, "")

	// Ensure the field name starts with an uppercase letter
	if len(fieldName) > 0 && (fieldName[0] < 'A' || fieldName[0] > 'Z') {
		fieldName = "A" + fieldName
	}

	// Ensure the field name ends with a letter
	if len(fieldName) > 0 && (fieldName[len(fieldName)-1] < 'A' || fieldName[len(fieldName)-1] > 'Z') && (fieldName[len(fieldName)-1] < 'a' || fieldName[len(fieldName)-1] > 'z') {
		fieldName = fieldName + "A"
	}

	return fieldName
}

// sortKeys sorts map keys alphabetically and returns them as a slice
func sortKeys(data map[string]interface{}) []string {
	keys := make([]string, 0, len(data))
	for key := range data {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}
