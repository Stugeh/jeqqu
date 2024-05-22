package utils

import "os/exec"

func ExecuteQueryCommand(path, query string) ([]byte, error) {
	out, err := exec.Command("jq", query, path).Output()

	if err != nil {
		println("query " + query + " failed for at " + path)
		panic(err.Error())
	}

	return out, err
}

func PrettyPrintLatest(bytes []byte) {
	// _, err := exec.Command(jq", "'.'" ).Output()
	//
	// if err != nil {
	// 	println("failed PrettyPrintJson")
	// 	println(err.Error())
	// 	return
	// }
	// println()
	// // println(string(out))
	// println()
}
