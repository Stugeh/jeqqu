package main

import (
	"fmt"
	"os"
	"os/exec"
)

func ExecuteQueryCommand(path, query string) (string, error) {
	out, err := exec.Command("jq", query, path).Output()

	if err != nil {
		println("query " + query + " failed for at " + path)
		println(err.Error())
	}

	return string(out), err
}

func main() {
	basePath := "./test-jsons/"
	files, _ := os.ReadDir(basePath)

	for _, file := range files {

		if file.IsDir() {
			continue
		}

		println()
		println(file.Name())

		baseType, err := ExecuteQueryCommand(basePath+file.Name(), ".|type")

		if err != nil {
			continue
		}

		baseType = baseType[1 : len(baseType)-2]

		switch baseType {
		case "object":
			keys, err := ExecuteQueryCommand(basePath+file.Name(), ".|keys")
			if err != nil {
				continue
			}
			fmt.Printf("keys: %v\n", keys)

		case "array":
			keys, err := ExecuteQueryCommand(basePath+file.Name(), ".[0]|keys")
			if err != nil {
				continue
			}
			fmt.Printf("keys: %v\n", keys)
		}

	}

}
