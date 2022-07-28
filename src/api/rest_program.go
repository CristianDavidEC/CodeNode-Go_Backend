package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"
)

//Post : /save-program
func SaveProgram(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	render.Status(r, http.StatusOK)
	render.JSON(w, r, map[string]interface{}{"message": "Save Program"})
}

//Get : /get-all-programs
func GetAllPrograms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res := map[string]interface{}{"message": "All Programs"}
	_ = json.NewEncoder(w).Encode(res)
}

//Get : /get-program
func GetProgram(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res := map[string]interface{}{"message": "One Program"}
	_ = json.NewEncoder(w).Encode(res)
}

//Post : /run-program
func RunProgram(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//res := map[string]interface{}{"message": "Execute Program"}
	//_ = json.NewEncoder(w).Encode(res)
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, map[string]interface{}{"message": "Save Program", "otra": "otra"})
}
