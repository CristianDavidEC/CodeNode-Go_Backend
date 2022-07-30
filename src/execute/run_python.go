package python

import (
	"fmt"
	"os"
	"os/exec"
)

func WritePyFile(code string) string {
	response := ""
	err := os.WriteFile("script.py", []byte(code), 0644)
	if err != nil {
		response = "Error writing file: " + err.Error()
	}
	return response
}

func RunPyCode(code string) {
	WritePyFile(code)
	cmd := exec.Command("python", "./script.py")
	out, err := cmd.Output()

	if err != nil {
		fmt.Println("error to execute python script: ", err)
		panic(err)
	}
	fmt.Println(string(out))
}
