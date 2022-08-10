package python

import (
	"os"
	"os/exec"
)

/*Write the python file*/
func WritePyFile(code string) string {
	response := ""
	err := os.WriteFile("script.py", []byte(code), 0644)
	if err != nil {
		response = "Error writing file: " + err.Error()
	}
	return response
}

/*Execute the python code and returns the result to the execution */
func RunPyCode(code string) string {
	WritePyFile(code)
	cmd := exec.Command("python", "./script.py")
	out, err := cmd.Output()

	if err != nil {
		return "error to execute python script: " + err.Error()
	}
	return string(out)
}
