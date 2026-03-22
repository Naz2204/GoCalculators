package main

import (
	"fmt"
	"html/template"
	"math"
	"net/http"
	"net/url"
	"strconv"
)

type PageData struct {
	InputType               string
	LowBurnTemp             float64
	AshPart                 float64
	AshMass                 float64
	BurnableEjection        float64
	CleaningEfficiency      float64
	SolidEmissionWithSulfur float64
	Mass                    float64

	Result map[string]interface{}
}

var calcPage = template.Must(template.ParseFiles("index.html"))

func pageHandler(w http.ResponseWriter, r *http.Request) {

	data := PageData{}

	if r.Method == http.MethodGet {
		query := r.URL.Query()

		data.InputType = query.Get("fuel_list")
		fmt.Println(data.InputType)
		data.AshPart, _ = strconv.ParseFloat(query.Get("ash_part"), 64)
		data.BurnableEjection, _ = strconv.ParseFloat(query.Get("burnable_ejection"), 64)
		data.CleaningEfficiency, _ = strconv.ParseFloat(query.Get("cleaning_efficiency"), 64)
		data.AshMass, _ = strconv.ParseFloat(query.Get("ash_mass"), 64)
		data.Mass, _ = strconv.ParseFloat(query.Get("mass"), 64)
		data.LowBurnTemp, _ = strconv.ParseFloat(query.Get("low_burn_temp"), 64)
		data.SolidEmissionWithSulfur, _ = strconv.ParseFloat(query.Get("solid_emission_with_sulfur"), 64)

		if data.InputType == "user" {
			data.Result = calculateUserInput(data.LowBurnTemp, data.AshPart, data.AshMass, data.BurnableEjection, data.CleaningEfficiency,
				data.SolidEmissionWithSulfur, data.Mass)
		}
		if data.InputType == "coal" {
			data.Result = calculateCoal(data.Mass)
		}
		if data.InputType == "mazut" {
			data.Result = calculateMazut(data.Mass)
		}
		if data.InputType == "gas" {
			data.Result = calculateGas(data.Mass)
		}

		calcPage.Execute(w, data)
		return
	}

	if r.Method == http.MethodPost {
		values := url.Values{}

		values.Add("fuel_list", r.FormValue("fuel_list"))
		values.Add("ash_part", r.FormValue("ash_part"))
		values.Add("burnable_ejection", r.FormValue("burnable_ejection"))
		values.Add("cleaning_efficiency", r.FormValue("cleaning_efficiency"))
		values.Add("ash_mass", r.FormValue("ash_mass"))
		values.Add("mass", r.FormValue("mass"))
		values.Add("low_burn_temp", r.FormValue("low_burn_temp"))
		values.Add("solid_emission_with_sulfur", r.FormValue("solid_emission_with_sulfur"))

		http.Redirect(w, r, "/?"+values.Encode(), http.StatusSeeOther)
	}

}

func calculateCoal(mass float64) map[string]interface{} {
	lowBurnTemp := 20.47
	ashPart := 0.8
	ashMass := 25.2
	burnableEjection := 1.5
	cleaningEfficiency := 0.985
	solidEmissionWithSulfur := 0.0

	emission := calculateEmission(lowBurnTemp, ashPart, ashMass, burnableEjection, cleaningEfficiency, solidEmissionWithSulfur)
	ejection := calculateEjection(emission, lowBurnTemp, mass)

	return map[string]interface{}{
		"emission": emission,
		"ejection": ejection,
	}
}

func calculateMazut(mass float64) map[string]interface{} {
	lowBurnTemp := 39.48
	ashPart := 1.0
	ashMass := 0.15
	burnableEjection := 0.0
	cleaningEfficiency := 0.985
	solidEmissionWithSulfur := 0.0

	emission := calculateEmission(lowBurnTemp, ashPart, ashMass, burnableEjection, cleaningEfficiency, solidEmissionWithSulfur)
	ejection := calculateEjection(emission, lowBurnTemp, mass)

	return map[string]interface{}{
		"emission": emission,
		"ejection": ejection,
	}
}

func calculateGas(mass float64) map[string]interface{} {
	lowBurnTemp := 33.08
	ashPart := 0.0
	ashMass := 0.0
	burnableEjection := 0.0
	cleaningEfficiency := 0.985
	solidEmissionWithSulfur := 0.0

	emission := calculateEmission(lowBurnTemp, ashPart, ashMass, burnableEjection, cleaningEfficiency, solidEmissionWithSulfur)
	ejection := calculateEjection(emission, lowBurnTemp, mass)

	return map[string]interface{}{
		"emission": emission,
		"ejection": ejection,
	}
}

func calculateUserInput(lowBurnTemp, ashPart, ashMass, burnableEjection, cleaningEfficiency, solidEmissionWithSulfur, mass float64) map[string]interface{} {
	emission := calculateEmission(lowBurnTemp, ashPart, ashMass, burnableEjection, cleaningEfficiency, solidEmissionWithSulfur)
	ejection := calculateEjection(emission, lowBurnTemp, mass)

	return map[string]interface{}{
		"emission": emission,
		"ejection": ejection,
	}
}

func calculateEmission(lowBurnTemp, ashPart, ashMass, burnableEjection, cleaningEfficiency, solidEmissionWithSulfur float64) float64 {
	return (math.Pow(10, 6)/lowBurnTemp)*ashPart*(ashMass/(100-burnableEjection))*(1-cleaningEfficiency) + solidEmissionWithSulfur
}

func calculateEjection(emission, lowBurnTemp, mass float64) float64 {
	return math.Pow(10, -6) * emission * lowBurnTemp * mass
}

func main() {
	http.HandleFunc("/", pageHandler)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", nil)
}
