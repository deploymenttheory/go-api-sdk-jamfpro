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
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

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

	var data map[string]interface{}
	if err := json.Unmarshal(byteValue, &data); err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

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

	structs, err := generateRootStructs(schemaData)
	if err != nil {
		return "", fmt.Errorf("failed to generate root structs: %w", err)
	}

	return structs, nil
}

// generateRootStructs generates Go struct definitions for the root-level fields of the OpenAPI schema
func generateRootStructs(schemaData map[string]interface{}) (string, error) {
	var structsBuilder strings.Builder
	structsBuilder.WriteString("package generatedstructs\n\n")

	// Generate the OpenAPI root struct
	openAPIFields := []string{
		"Openapi string `json:\"openapi\"`",
		"Servers []Server `json:\"servers\"`",
		"Security []map[string]interface{} `json:\"security\"`",
		"Paths Paths `json:\"paths\"`",
		"Components Components `json:\"components\"`",
	}

	structsBuilder.WriteString("type OpenAPI struct {\n")
	for _, field := range openAPIFields {
		structsBuilder.WriteString(field + "\n")
	}
	structsBuilder.WriteString("}\n\n")

	// Generate structs for the individual top-level fields
	if err := generateServersStruct(&structsBuilder, schemaData); err != nil {
		return "", err
	}

	if err := generatePathsStruct(&structsBuilder, schemaData); err != nil {
		return "", err
	}

	if err := generateComponentsStruct(&structsBuilder, schemaData); err != nil {
		return "", err
	}

	return structsBuilder.String(), nil
}

// generateServersStruct generates the Go struct for the "servers" field
func generateServersStruct(structsBuilder *strings.Builder, schemaData map[string]interface{}) error {
	servers, ok := schemaData["servers"].([]interface{})
	if !ok {
		return nil
	}

	for _, server := range servers {
		serverMap, ok := server.(map[string]interface{})
		if !ok {
			continue
		}

		serverStruct, err := generateStruct("Server", serverMap)
		if err != nil {
			return err
		}

		structsBuilder.WriteString(serverStruct)
		structsBuilder.WriteString("\n\n")
	}

	return nil
}

// generatePathsStruct generates the Go struct for the "paths" field
func generatePathsStruct(structsBuilder *strings.Builder, schemaData map[string]interface{}) error {
	paths, ok := schemaData["paths"].(map[string]interface{})
	if !ok {
		return nil
	}

	taggedPaths := make(map[string]map[string]interface{})

	for path, pathItem := range paths {
		pathItemMap, ok := pathItem.(map[string]interface{})
		if !ok {
			continue
		}

		for method, operation := range pathItemMap {
			operationMap, ok := operation.(map[string]interface{})
			if !ok {
				continue
			}

			if tags, ok := operationMap["tags"].([]interface{}); ok {
				for _, tag := range tags {
					tagStr, ok := tag.(string)
					if !ok {
						continue
					}

					if _, exists := taggedPaths[tagStr]; !exists {
						taggedPaths[tagStr] = make(map[string]interface{})
					}

					sanitizedPath := sanitizeFieldName(path)
					taggedPaths[tagStr][method+"_"+sanitizedPath] = operationMap
				}
			}
		}
	}

	titleCaser := cases.Title(language.English)

	// Create top-level Paths struct
	structsBuilder.WriteString("type Paths struct {\n")
	tagStructNames := make([]string, 0, len(taggedPaths))

	for tag := range taggedPaths {
		tagStructName := sanitizeFieldName(tag)
		tagStructNames = append(tagStructNames, tagStructName)
	}
	sort.Strings(tagStructNames)

	for _, tagStructName := range tagStructNames {
		structsBuilder.WriteString(fmt.Sprintf("%s %s `json:\"%s\"`\n", titleCaser.String(tagStructName), titleCaser.String(tagStructName), tagStructName))
	}
	structsBuilder.WriteString("}\n\n")

	// Generate structs for each tag
	for _, tag := range tagStructNames {
		paths := taggedPaths[tag]
		tagStructName := sanitizeFieldName(tag)
		structsBuilder.WriteString(fmt.Sprintf("type %s struct {\n", titleCaser.String(tagStructName)))

		methodPaths := make([]string, 0, len(paths))
		for methodPath := range paths {
			methodPaths = append(methodPaths, methodPath)
		}
		sort.Strings(methodPaths)

		for _, methodPath := range methodPaths {
			structName := sanitizeFieldName(methodPath)
			fieldName := titleCaser.String(structName)
			structsBuilder.WriteString(fmt.Sprintf("%s %s `json:\"%s\"`\n", fieldName, structName, methodPath))
		}

		structsBuilder.WriteString("}\n\n")

		for _, methodPath := range methodPaths {
			structName := sanitizeFieldName(methodPath)
			pathItem := paths[methodPath]
			pathItemStruct, err := generateStruct(structName, pathItem.(map[string]interface{}))
			if err != nil {
				return err
			}
			structsBuilder.WriteString(pathItemStruct)
			structsBuilder.WriteString("\n\n")
		}
	}

	return nil
}

// generateComponentsStruct generates the Go struct for the "components" field
func generateComponentsStruct(structsBuilder *strings.Builder, schemaData map[string]interface{}) error {
	components, ok := schemaData["components"].(map[string]interface{})
	if !ok {
		return nil
	}

	componentStruct, err := generateStruct("Components", components)
	if err != nil {
		return err
	}
	structsBuilder.WriteString(componentStruct)
	structsBuilder.WriteString("\n\n")

	for key, value := range components {
		if nestedProps, ok := value.(map[string]interface{}); ok {
			nestedStructDef, err := generateStruct(cases.Title(language.English).String(key), nestedProps)
			if err != nil {
				return err
			}
			structsBuilder.WriteString(nestedStructDef)
			structsBuilder.WriteString("\n\n")
		}
	}

	return nil
}

// generateStruct generates a Go struct definition from a JSON object
func generateStruct(structName string, structData map[string]interface{}) (string, error) {
	var fieldsBuilder strings.Builder
	titleCaser := cases.Title(language.English)

	for key, value := range structData {
		fieldType, err := getFieldType(key, value)
		if err != nil {
			return "", err
		}

		fieldName := sanitizeFieldName(titleCaser.String(key))
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
func getFieldType(fieldName string, value interface{}) (string, error) {
	switch v := value.(type) {
	case map[string]interface{}:
		if _, ok := v["properties"]; ok {
			return cases.Title(language.English).String(fieldName), nil
		}
		return "map[string]interface{}", nil
	case []interface{}:
		if len(v) > 0 {
			elemType, err := getFieldType(fieldName, v[0])
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

// sanitizeFieldName removes invalid characters in struct field names
func sanitizeFieldName(fieldName string) string {
	// Remove invalid characters, specifically '/'
	reg := regexp.MustCompile(`[/]`)
	return reg.ReplaceAllString(fieldName, "")
}
