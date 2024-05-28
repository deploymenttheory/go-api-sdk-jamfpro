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

	return generateStructs(schemaData)
}

// generateStructs generates Go struct definitions for the schema
func generateStructs(schemaData map[string]interface{}) (string, error) {
	var structsBuilder strings.Builder
	structsBuilder.WriteString("package generatedstructs\n\n")

	if err := generateOpenAPIStruct(&structsBuilder); err != nil {
		return "", err
	}

	if err := generateTopLevelStructs(&structsBuilder, schemaData); err != nil {
		return "", err
	}

	return structsBuilder.String(), nil
}

// generateOpenAPIStruct generates a Go struct definition for the OpenAPI object
func generateOpenAPIStruct(structsBuilder *strings.Builder) error {
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
	return nil
}

// generateTopLevelStructs generates Go struct definitions for the top-level sections of the schema
func generateTopLevelStructs(structsBuilder *strings.Builder, schemaData map[string]interface{}) error {
	if err := generateServersStruct(structsBuilder, schemaData); err != nil {
		return err
	}

	if err := generatePathsStruct(structsBuilder, schemaData); err != nil {
		return err
	}

	return generateComponentsStruct(structsBuilder, schemaData)
}

// generateServersStruct generates Go struct definitions for the servers section of the schema
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

		if err := appendStruct(structsBuilder, "Server", serverMap); err != nil {
			return err
		}
	}

	return nil
}

// generatePathsStruct generates Go struct definitions for the paths section of the schema
func generatePathsStruct(structsBuilder *strings.Builder, schemaData map[string]interface{}) error {
	paths, ok := schemaData["paths"].(map[string]interface{})
	if !ok {
		return nil
	}

	taggedPaths := extractTaggedPaths(paths)
	titleCaser := cases.Title(language.English)

	structsBuilder.WriteString("type Paths struct {\n")
	tagStructNames := make([]string, 0, len(taggedPaths))
	for tag := range taggedPaths {
		tagStructNames = append(tagStructNames, tag)
	}
	sort.Strings(tagStructNames)

	for _, tagStructName := range tagStructNames {
		structsBuilder.WriteString(fmt.Sprintf("%s %s `json:\"%s\"`\n", titleCaser.String(tagStructName), titleCaser.String(tagStructName), tagStructName))
	}
	structsBuilder.WriteString("}\n\n")

	for _, tag := range tagStructNames {
		if err := generateTaggedPathsStruct(structsBuilder, tag, taggedPaths[tag], titleCaser); err != nil {
			return err
		}
	}

	return nil
}

// extractTaggedPaths extracts paths tagged with a specific tag
func extractTaggedPaths(paths map[string]interface{}) map[string]map[string]interface{} {
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

	return taggedPaths
}

// generateTaggedPathsStruct generates Go struct definitions for tagged paths
func generateTaggedPathsStruct(structsBuilder *strings.Builder, tag string, paths map[string]interface{}, titleCaser cases.Caser) error {
	tagStructName := sanitizeFieldName(tag)
	structsBuilder.WriteString(fmt.Sprintf("type %s struct {\n", titleCaser.String(tagStructName)))

	methodPaths := sortKeys(paths)

	for _, methodPath := range methodPaths {
		structName := sanitizeFieldName(methodPath)
		fieldName := titleCaser.String(structName)
		structsBuilder.WriteString(fmt.Sprintf("%s %s `json:\"%s\"`\n", fieldName, structName, methodPath))
	}

	structsBuilder.WriteString("}\n\n")

	for _, methodPath := range methodPaths {
		structName := sanitizeFieldName(methodPath)
		pathItem := paths[methodPath]
		if err := appendStruct(structsBuilder, structName, pathItem.(map[string]interface{})); err != nil {
			return err
		}
	}

	return nil
}

// generateComponentsStruct generates Go struct definitions for the components section of the schema
func generateComponentsStruct(structsBuilder *strings.Builder, schemaData map[string]interface{}) error {
	components, ok := schemaData["components"].(map[string]interface{})
	if !ok {
		return nil
	}

	if err := appendStruct(structsBuilder, "Components", components); err != nil {
		return err
	}

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

func appendStruct(structsBuilder *strings.Builder, structName string, structData map[string]interface{}) error {
	structDef, err := generateStruct(structName, structData)
	if err != nil {
		return err
	}
	structsBuilder.WriteString(structDef)
	structsBuilder.WriteString("\n\n")
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
	case
		string:
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

// sortKeys sorts map keys alphabetically and returns them as a slice
func sortKeys(data map[string]interface{}) []string {
	keys := make([]string, 0, len(data))
	for key := range data {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}
