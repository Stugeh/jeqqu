package main

import (
	"jeqqu/src/utils"
	"os"
)

// Executes jq query commands
//
// path: path to input file
//
// query: jq query

func main() {
	basePath := "./test-jsons/stringified_fields.json"
	bytes, _ := os.ReadFile(basePath)
	bytes = utils.UnStringifyJsonStrings(bytes)

	path, err := utils.NewTempFile(bytes)

	if err != nil {
		println(err.Error())
		panic("Failed to create initial temp file.")

	}

	out, err := utils.ExecuteQueryCommand(path, ".")

	if err != nil {
		println("Failed to pretty print.")
	}

	println(string(out))

	// baseType, err := ExecuteQueryCommand(basePath+file.Name(), ".|type")
	//
	// if err != nil {
	// 	continue
	// }
	//
	// baseType = baseType[1 : len(baseType)-2]
	//
	// switch baseType {
	// case "object":
	// 	keys, err := ExecuteQueryCommand(basePath+file.Name(), ".|keys")
	// 	if err != nil {
	// 		continue
	// 	}
	// 	fmt.Printf("keys: %v\n", keys)
	//
	// case "array":
	// 	keys, err := ExecuteQueryCommand(basePath+file.Name(), ".[0]|keys")
	// 	if err != nil {
	// 		continue
	// 	}
	// 	fmt.Printf("keys: %v\n", keys)
	// }

}
