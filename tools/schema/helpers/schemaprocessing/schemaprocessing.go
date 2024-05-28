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

// OpenAPI is the top-level struct for the OAS3 standard
type OpenAPI struct {
	Openapi    string                   `json:"openapi"`
	Servers    []Server                 `json:"servers"`
	Security   []map[string]interface{} `json:"security"`
	Paths      Paths                    `json:"paths"`
	Components Components               `json:"components"`
}

// Server represents the Server object in OAS3
type Server struct {
	URL         string                 `json:"url"`
	Description string                 `json:"description"`
	Variables   map[string]interface{} `json:"variables"`
}

// Paths represents the Paths object in OAS3
type Paths struct {
	PathItems map[string]PathItem
}

// PathItem represents a single path item in the Paths object
type PathItem struct {
	Get    map[string]interface{} `json:"get"`
	Post   map[string]interface{} `json:"post"`
	Put    map[string]interface{} `json:"put"`
	Delete map[string]interface{} `json:"delete"`
	Patch  map[string]interface{} `json:"patch"`
}

// Components represents the Components object in OAS3
type Components struct {
	ComponentItems map[string]interface{}
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

	return generateStructs(schemaData)
}

// generateStructs generates Go struct definitions for the schema
func generateStructs(schemaData map[string]interface{}) (string, error) {
	var structsBuilder strings.Builder
	structsBuilder.WriteString("package generatedstructs\n\n")

	if err := generateOpenAPIStruct(&structsBuilder); err != nil {
		return "", err
	}

	if err := generateServersStruct(&structsBuilder, schemaData); err != nil {
		return "", err
	}

	if err := generateSecurityStruct(&structsBuilder, schemaData); err != nil {
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

// generateSecurityStruct generates Go struct definitions for the security section of the schema
func generateSecurityStruct(structsBuilder *strings.Builder, schemaData map[string]interface{}) error {
	security, ok := schemaData["security"].([]interface{})
	if !ok {
		return nil
	}

	for i, sec := range security {
		secMap, ok := sec.(map[string]interface{})
		if !ok {
			continue
		}

		if err := appendStruct(structsBuilder, fmt.Sprintf("Security%d", i+1), secMap); err != nil {
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

	pathsStruct := Paths{PathItems: make(map[string]PathItem)}

	for path, pathItem := range paths {
		pathItemMap, ok := pathItem.(map[string]interface{})
		if !ok {
			continue
		}

		sanitizedPath := sanitizeFieldName(path)
		pathStruct := PathItem{}

		if getItem, ok := pathItemMap["get"]; ok {
			pathStruct.Get = getItem.(map[string]interface{})
		}
		if postItem, ok := pathItemMap["post"]; ok {
			pathStruct.Post = postItem.(map[string]interface{})
		}
		if putItem, ok := pathItemMap["put"]; ok {
			pathStruct.Put = putItem.(map[string]interface{})
		}
		if deleteItem, ok := pathItemMap["delete"]; ok {
			pathStruct.Delete = deleteItem.(map[string]interface{})
		}

		pathsStruct.PathItems[sanitizedPath] = pathStruct

		if err := appendStruct(structsBuilder, sanitizedPath, pathItemMap); err != nil {
			return err
		}
	}

	structsBuilder.WriteString("type Paths struct {\n")
	for path := range pathsStruct.PathItems {
		fieldName := sanitizeFieldName(path)
		structsBuilder.WriteString(fmt.Sprintf("%s PathItem `json:\"%s\"`\n", fieldName, path))
	}
	structsBuilder.WriteString("}\n\n")

	return nil
}

// generateComponentsStruct generates Go struct definitions for the components section of the schema
func generateComponentsStruct(structsBuilder *strings.Builder, schemaData map[string]interface{}) error {
	components, ok := schemaData["components"].(map[string]interface{})
	if !ok {
		return nil
	}

	componentsStruct := Components{ComponentItems: make(map[string]interface{})}

	for componentName, componentItem := range components {
		componentItemMap, ok := componentItem.(map[string]interface{})
		if !ok {
			continue
		}

		sanitizedComponentName := sanitizeFieldName(componentName)
		componentsStruct.ComponentItems[sanitizedComponentName] = componentItemMap

		if err := appendStruct(structsBuilder, sanitizedComponentName, componentItemMap); err != nil {
			return err
		}
	}

	structsBuilder.WriteString("type Components struct {\n")
	for componentName := range componentsStruct.ComponentItems {
		fieldName := sanitizeFieldName(componentName)
		structsBuilder.WriteString(fmt.Sprintf("%s map[string]interface{} `json:\"%s\"`\n", fieldName, componentName))
	}
	structsBuilder.WriteString("}\n\n")

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
	reg := regexp.MustCompile(`[/{}-]`)
	return reg.ReplaceAllString(fieldName, "_")
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
