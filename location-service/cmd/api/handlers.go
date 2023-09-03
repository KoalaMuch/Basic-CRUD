package main

import (
	"net/http"
	"basic-crud/data"
	"github.com/go-chi/chi/v5"
)

type NewLocationPayload struct {
	Name string `json:"name"`
	Description string `json:"description"`
}

func (app *Config) AddLocation(w http.ResponseWriter, r *http.Request) {
	var requestPayload NewLocationPayload
	_ = app.readJSON(w, r, &requestPayload)

	location := data.Location{
		Name: requestPayload.Name,
		Description: requestPayload.Description,
	}

	err := app.Models.Location.Insert(location)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	resp := jsonResponse{
		Error: false,
		Message: "Location added successfully",
	}
	app.writeJSON(w, http.StatusAccepted, resp)
}

func (app *Config) GetLocations(w http.ResponseWriter, r *http.Request) {
	locations, err := app.Models.Location.FindAll()
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	resp := jsonResponse{
		Error: false,
		Message: "Locations retrieved successfully",
		Data: locations,
	}
	app.writeJSON(w, http.StatusAccepted, resp)
}

func (app *Config) GetLocation(w http.ResponseWriter, r *http.Request) {

	ID := chi.URLParam(r, "id")

	location, err := app.Models.Location.FindById(ID)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	resp := jsonResponse{
		Error: false,
		Message: "Location retrieved successfully",
		Data: location,
	}
	app.writeJSON(w, http.StatusAccepted, resp)
}

func (app *Config) UpdateLocation(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		ID string `json:"id"`
		Name string `json:"name"`
		Description string `json:"description"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err!= nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	location := data.Location{
		ID: requestPayload.ID,
		Name: requestPayload.Name,
		Description: requestPayload.Description,
	}

	err = app.Models.Location.Update(location)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	resp := jsonResponse{
		Error: false,
		Message: "Location updated successfully",
	}
	app.writeJSON(w, http.StatusAccepted, resp)
}

func (app *Config) DeleteLocation(w http.ResponseWriter, r *http.Request) {
	
	ID := chi.URLParam(r, "id")

	err := app.Models.Location.Delete(ID)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	resp := jsonResponse{
		Error: false,
		Message: "Location deleted successfully",
	}
	app.writeJSON(w, http.StatusAccepted, resp)
}