package main

import (
	"html/template"
	"math"
	"net/http"
	"net/url"
	"strconv"
)

type PageData struct {
	W_os   float64
	T_w_os float64
	Z_p_a  float64
	Z_p_p  float64

	Result map[string]interface{}
}

var calcPage = template.Must(template.ParseFiles("index.html"))

func calcHandler(w http.ResponseWriter, r *http.Request) {

	data := PageData{}

	if r.Method == http.MethodGet {

		query := r.URL.Query()

		data.W_os, _ = strconv.ParseFloat(query.Get("W_os"), 64)
		data.T_w_os, _ = strconv.ParseFloat(query.Get("T_w_os"), 64)
		data.Z_p_a, _ = strconv.ParseFloat(query.Get("Z_p_a"), 64)
		data.Z_p_p, _ = strconv.ParseFloat(query.Get("Z_p_p"), 64)

		W_ds := calcTask1(data.W_os, data.T_w_os)
		M_Z_p := calcTask2(data.Z_p_a, data.Z_p_p)
		Comp := results(data.W_os, W_ds)

		data.Result = map[string]interface{}{
			"W_ds":  W_ds,
			"W_os":  data.W_os,
			"M_Z_p": M_Z_p,
			"Comp":  Comp,
		}

		calcPage.Execute(w, data)
	}

	if r.Method == http.MethodPost {
		values := url.Values{}

		values.Add("W_os", r.FormValue("W_os"))
		values.Add("T_w_os", r.FormValue("T_w_os"))
		values.Add("Z_p_a", r.FormValue("Z_p_a"))
		values.Add("Z_p_p", r.FormValue("Z_p_p"))

		http.Redirect(w, r, "/?"+values.Encode(), http.StatusSeeOther)
	}
}

func calcTask1(w_oc, t_w_oc float64) float64 {
	const k_p_max = 43.0
	const w_sv = 0.02

	k_aoc := math.Round(w_oc*t_w_oc/8760*1e5) / 1e5
	k_poc := math.Round(1.2*k_p_max/8760*1e5) / 1e5

	w_dk := math.Round(2*w_oc*(k_aoc+k_poc)*1e5) / 1e5
	w_ds := w_sv + w_dk

	return w_ds
}

func calcTask2(Z_p_a, Z_p_p float64) int {
	const omega = 0.01
	const t_v = 0.045
	const k_p = 0.004
	const P_m = 5120.0
	const T_m = 6451.0

	M_W_n_a := int(math.Ceil(omega * t_v * P_m * T_m))
	M_W_n_p := int(math.Ceil(k_p * P_m * T_m))
	M_Z_p := int(math.Ceil(Z_p_a*float64(M_W_n_a) + Z_p_p*float64(M_W_n_p)))

	return M_Z_p
}

func results(w_os, w_ds float64) string {
	var compare string
	if w_os == w_ds {
		compare = "Надійність одноколової і двоколової систем рівна"
	} else if w_os > w_ds {
		compare = "Надійність двоколової системи вища ніж одноколової"
	} else {
		compare = "Надійність одноколової системи вища ніж двоколової"
	}
	return compare
}

func main() {
	http.HandleFunc("/", calcHandler)

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static")),
		),
	)

	http.ListenAndServe(":8080", nil)
}
