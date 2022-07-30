package api

import (
	"encoding/json"
	"net/http"

	python "codenode/packages/src/execute"
	"codenode/packages/src/model"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

//Get : /get-all-programs
func GetAllPrograms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//res := map[string]interface{}{"message": "All Programs"}
	//_ = json.NewEncoder(w).Encode(res)
	render.JSON(w, r, map[string]interface{}{"message": "Save Program"})
}

//Get : /get-program
func GetProgram(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idProgram := chi.URLParam(r, "id")
	w.Write([]byte(idProgram))
	//res := map[string]interface{}{"message": "One Program"}
	//_ = json.NewEncoder(w).Encode(res)
}

//Post : /save-program
func SaveProgram(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	render.Status(r, http.StatusOK)
	render.JSON(w, r, map[string]interface{}{"message": "Save Program"})
}

//Post : /run-program
func RunProgram(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	var codePython model.CodePython
	_ = json.NewDecoder(r.Body).Decode(&codePython)
	python.RunPyCode(codePython.Code)

}
