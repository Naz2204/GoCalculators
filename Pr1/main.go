package main

import (
	"html/template"
	"net/http"
	"net/url"
	"strconv"
)

type FuelPageData struct {
	Carbon   float64
	Hydrogen float64
	Sulfur   float64
	Nitrogen float64
	Oxygen   float64
	Wetness  float64
	Ash      float64

	Result map[string]interface{}
}

type OilPageData struct {
	Carbon          float64
	Hydrogen        float64
	Sulfur          float64
	Vanadium        float64
	Oxygen          float64
	Wetness         float64
	Ash             float64
	LowBurnFireTemp float64

	Result map[string]interface{}
}

var fuel_tmpl = template.Must(template.ParseFiles("pages/fuel.html"))
var oil_tmpl = template.Must(template.ParseFiles("pages/fuel_oil.html"))

func fuelHandler(w http.ResponseWriter, r *http.Request) {

	data := FuelPageData{}

	if r.Method == http.MethodGet {

		query := r.URL.Query()

		data.Carbon, _ = strconv.ParseFloat(query.Get("carbon"), 64)
		data.Hydrogen, _ = strconv.ParseFloat(query.Get("hydrogen"), 64)
		data.Sulfur, _ = strconv.ParseFloat(query.Get("sulfur"), 64)
		data.Nitrogen, _ = strconv.ParseFloat(query.Get("nitrogen"), 64)
		data.Oxygen, _ = strconv.ParseFloat(query.Get("oxygen"), 64)
		data.Wetness, _ = strconv.ParseFloat(query.Get("wetness"), 64)
		data.Ash, _ = strconv.ParseFloat(query.Get("ash"), 64)

		if query.Get("carbon") != "" {
			data.Result = calcFuel(
				data.Carbon,
				data.Hydrogen,
				data.Oxygen,
				data.Sulfur,
				data.Nitrogen,
				data.Wetness,
				data.Ash,
			)
		}

		fuel_tmpl.Execute(w, data)
		return
	}

	if r.Method == http.MethodPost {

		values := url.Values{}

		values.Add("carbon", r.FormValue("carbon"))
		values.Add("hydrogen", r.FormValue("hydrogen"))
		values.Add("sulfur", r.FormValue("sulfur"))
		values.Add("nitrogen", r.FormValue("nitrogen"))
		values.Add("oxygen", r.FormValue("oxygen"))
		values.Add("wetness", r.FormValue("wetness"))
		values.Add("ash", r.FormValue("ash"))

		http.Redirect(w, r, "/fuel?"+values.Encode(), http.StatusSeeOther)
	}
}

func oilHandler(w http.ResponseWriter, r *http.Request) {

	data := OilPageData{}

	if r.Method == http.MethodGet {

		query := r.URL.Query()

		data.Carbon, _ = strconv.ParseFloat(query.Get("carbon"), 64)
		data.Hydrogen, _ = strconv.ParseFloat(query.Get("hydrogen"), 64)
		data.Sulfur, _ = strconv.ParseFloat(query.Get("sulfur"), 64)
		data.Vanadium, _ = strconv.ParseFloat(query.Get("vanadium"), 64)
		data.Oxygen, _ = strconv.ParseFloat(query.Get("oxygen"), 64)
		data.Wetness, _ = strconv.ParseFloat(query.Get("wetness"), 64)
		data.Ash, _ = strconv.ParseFloat(query.Get("ash"), 64)
		data.LowBurnFireTemp, _ = strconv.ParseFloat(query.Get("low_burn_fire_temp"), 64)

		if query.Get("carbon") != "" {
			data.Result = calcFuelOil(
				data.Carbon,
				data.Hydrogen,
				data.Sulfur,
				data.Oxygen,
				data.Ash,
				data.Wetness,
				data.Vanadium,
			)
		}

		oil_tmpl.Execute(w, data)
		return
	}

	if r.Method == http.MethodPost {

		values := url.Values{}

		values.Add("carbon", r.FormValue("carbon"))
		values.Add("hydrogen", r.FormValue("hydrogen"))
		values.Add("sulfur", r.FormValue("sulfur"))
		values.Add("vanadium", r.FormValue("vanadium"))
		values.Add("oxygen", r.FormValue("oxygen"))
		values.Add("wetness", r.FormValue("wetness"))
		values.Add("ash", r.FormValue("ash"))
		values.Add("low_burn_fire_temp", r.FormValue("low_burn_fire_temp"))

		http.Redirect(w, r, "/fuel_oil?"+values.Encode(), http.StatusSeeOther)
	}
}

func main() {

	http.HandleFunc("/fuel", fuelHandler)
	http.HandleFunc("/fuel_oil", oilHandler)

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static")),
		),
	)

	http.ListenAndServe(":8080", nil)
}
