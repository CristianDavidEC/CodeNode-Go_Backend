package api

import (
	"encoding/json"
	"net/http"

	python "codenode/packages/src/execute"
	"codenode/packages/src/model"
	"codenode/packages/src/repository"

	"github.com/go-chi/chi/v5"
)

/*Get : /get-all-programs*/
func GetAllPrograms(w http.ResponseWriter, r *http.Request) {
	res, err := repository.GetAllProgramsDb()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

/*Get : /get-program*/
func GetProgram(w http.ResponseWriter, r *http.Request) {
	idProgram := chi.URLParam(r, "id")
	res, err := repository.GetProgramDb(idProgram)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

/*Post : /save-program*/
func SaveProgram(w http.ResponseWriter, r *http.Request) {
	var program model.Program
	_ = json.NewDecoder(r.Body).Decode(&program)
	newProgram, err := json.Marshal(program)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = repository.SaveProgramDb(newProgram)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

/*Post : /run-program*/
func RunProgram(w http.ResponseWriter, r *http.Request) {
	var codePython model.CodePython
	_ = json.NewDecoder(r.Body).Decode(&codePython)
	result := python.RunPyCode(codePython.Code)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(model.CodePython{Code: result})
}
