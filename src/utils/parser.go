package utils

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
)

// Takes in json and returns a tree of keys
func GetKeyTree(jsonBytes []byte) (map[string]interface{}, error) {
	var jsonData interface{}
	if err := json.Unmarshal(jsonBytes, &jsonData); err != nil {
		return nil, err
	}

	return buildKeyTree(jsonData), nil
}

// buildKeyTree recursively builds a tree of keys from the JSON data
func buildKeyTree(data interface{}) map[string]interface{} {
	keyTree := make(map[string]interface{})

	switch v := data.(type) {
	case map[string]interface{}:
		for k, val := range v {
			keyTree[k] = buildKeyTree(val)
		}
	case []interface{}:
		// For arrays, we'll build a key tree for the first element (assuming uniform structure)
		if len(v) > 0 {
			keyTree["[]"] = buildKeyTree(v[0])
		}
	default:
		return nil
	}

	return keyTree
}

// Makes any embedded json strings into valid json
func UnStringifyJsonStrings(input []byte) []byte {
	// Regular expression to find JSON strings
	jsonStringPattern := `"(.*?)":\s*"((?:\\.|[^"\\])*)"`
	re := regexp.MustCompile(jsonStringPattern)

	deStringifiedJson := re.ReplaceAllFunc(input, func(match []byte) []byte {
		// Extract the key and value from the match
		matches := re.FindSubmatch(match)
		if len(matches) < 3 {
			return match
		}

		key := matches[1]
		jsonString := matches[2]

		// Unquote the JSON string to handle escaped characters
		unquotedJSON, err := strconv.Unquote(`"` + string(jsonString) + `"`)
		if err != nil {
			return match
		}

		var parsedContent map[string]interface{}
		err = json.Unmarshal([]byte(unquotedJSON), &parsedContent)
		if err != nil {
			return match
		}

		parsedJSON, err := json.Marshal(parsedContent)
		if err != nil {
			return match
		}

		return []byte(fmt.Sprintf(`"%s":%s`, key, parsedJSON))
	})

	return deStringifiedJson
}
