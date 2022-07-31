package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	database "codenode/packages/src/dataBase"
	python "codenode/packages/src/execute"
	"codenode/packages/src/model"

	"github.com/go-chi/chi/v5"
)

//Get : /get-all-programs
func GetAllPrograms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res, err := database.AllProgramsDb()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res.Json)
}

//Get : /get-program
func GetProgram(w http.ResponseWriter, r *http.Request) {
	idProgram := chi.URLParam(r, "id")
	res := map[string]interface{}{"message": string(idProgram)}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

//Post : /save-program
func SaveProgram(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Body)
	idProgram := chi.URLParam(r, "id")
	res := map[string]interface{}{"message": "Save Program"}
	var program model.Program
	_ = json.NewDecoder(r.Body).Decode(&program)

	w.Write([]byte(idProgram))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)

}

//Post : /run-program
func RunProgram(w http.ResponseWriter, r *http.Request) {
	var codePython model.CodePython
	_ = json.NewDecoder(r.Body).Decode(&codePython)
	result := python.RunPyCode(codePython.Code)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(model.CodePython{Code: result})
}
