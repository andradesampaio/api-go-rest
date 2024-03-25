package controllers

import (
	"api-rest/internal/database"
	"api-rest/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)


func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	var personalidades []models.Personalidade
	database.DB.Find(&personalidades)
    json.NewEncoder(w).Encode(personalidades)
}

func RetornaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	personalidadeID := vars["id"]
	var personalidade models.Personalidade
    database.DB.First(&personalidade, personalidadeID)
    json.NewEncoder(w).Encode(personalidade)
}

func CriaUmaNovaPersonalidade(w http.ResponseWriter, r *http.Request) {
    var novaPersonalidade models.Personalidade
    json.NewDecoder(r.Body).Decode(&novaPersonalidade)
    database.DB.Create(&novaPersonalidade)
    json.NewEncoder(w).Encode(novaPersonalidade)
}

func DeletaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	personalidadeID := vars["id"]
	database.DB.Delete(&models.Personalidade{}, personalidadeID)
	fmt.Fprintf(w, "Personalidade deletada com sucesso")
}

func EditaPersonalidade(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    var personalidade models.Personalidade

    database.DB.First(&personalidade, id)

    json.NewDecoder(r.Body).Decode(&personalidade)
    database.DB.Save(&personalidade)
	
    json.NewEncoder(w).Encode(personalidade)
}