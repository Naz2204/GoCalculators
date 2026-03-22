package main

import (
	"html/template"
	"math"
	"net/http"
	"net/url"
	"strconv"
)

type PageData struct {
	S_m     float64
	U_nom   float64
	I_k     float64
	T_f     float64
	S_k     float64
	U_cn    float64
	U_k     float64
	S_nom_t float64
	L       float64
	U_k_max float64
	U_vn    float64
	U_nn    float64
	R_sn    float64
	X_sn    float64
	R_s_min float64
	X_s_min float64

	Result map[string]interface{}
}

var calcPage = template.Must(template.ParseFiles("index.html"))

func chooseCable(S_m, U_nom, I_k, T_f float64) [2]float64 {
	j_ek := 1.4
	C_t := 92.0

	I_m := (S_m / 2) / (math.Sqrt(3) * U_nom)
	I_m = math.Round(I_m*100) / 100

	I_m_pa := 2 * I_m
	I_m_pa = math.Round(I_m_pa*100) / 100

	S_ek := I_m / j_ek
	S_ek = math.Round(S_ek*100) / 100

	S := (I_k * math.Sqrt(T_f)) / C_t
	S = math.Round(S*100) / 100

	return [2]float64{S, S_ek}
}

func calculateKZ(S_k, U_cn, U_k, S_nom_t float64) float64 {
	X_c := math.Pow(U_cn, 2) / S_k
	X_c = math.Round(X_c*100) / 100

	X_t := (U_k / 100) * (math.Pow(U_cn, 2) / S_nom_t)
	X_t = math.Round(X_t*100) / 100

	X_sum := X_c + X_t
	X_sum = math.Round(X_sum*100) / 100

	I_p0 := U_cn / (math.Sqrt(3) * X_sum)
	I_p0 = math.Round(I_p0*100) / 100

	return I_p0
}

func calculateKZonStation(L, U_k_max, U_vn, U_nn, S_nom_t, R_sn, X_sn, R_s_min, X_s_min float64) [3][4]float64 {
	R0 := 0.64
	X0 := 0.363

	round := func(val float64, precision int) float64 {
		pow := math.Pow(10, float64(precision))
		return math.Round(val*pow) / pow
	}

	X_t := round((U_k_max*math.Pow(U_vn, 2))/(100*S_nom_t), 2)

	R_h := R_sn
	X_h := round(X_sn+X_t, 2)

	Z_h := round(math.Sqrt(math.Pow(R_h, 2)+math.Pow(X_h, 2)), 2)

	R_h_min := R_s_min
	X_h_min := round(X_s_min+X_t, 2)

	Z_h_min := round(math.Sqrt(math.Pow(R_h_min, 2)+math.Pow(X_h_min, 2)), 2)

	I_h3 := round((U_vn*math.Pow(10, 3))/(math.Sqrt(3)*Z_h), 2)
	I_h2 := round(I_h3*(math.Sqrt(3)/2), 2)

	I_h3_min := round((U_vn*math.Pow(10, 3))/(math.Sqrt(3)*Z_h_min), 2)
	I_h2_min := round(I_h3_min*(math.Sqrt(3)/2), 2)

	k_pr := round(math.Pow(U_nn, 2)/math.Pow(U_vn, 2), 3)

	R_h_n := round(R_h*k_pr, 2)
	X_h_n := round(X_h*k_pr, 2)
	Z_h_n := round(math.Sqrt(math.Pow(R_h_n, 2)+math.Pow(X_h_n, 2)), 2)

	R_h_n_min := round(R_h_min*k_pr, 2)
	X_h_n_min := round(X_h_min*k_pr, 2)
	Z_h_n_min := round(math.Sqrt(math.Pow(R_h_n_min, 2)+math.Pow(X_h_n_min, 2)), 2)

	I_h_n3 := round((U_nn*math.Pow(10, 3))/(math.Sqrt(3)*Z_h_n), 2)
	I_h_n2 := round(I_h_n3*(math.Sqrt(3)/2), 2)

	I_h_n3_min := round((U_nn*math.Pow(10, 3))/(math.Sqrt(3)*Z_h_n_min), 2)
	I_h_n2_min := round(I_h_n3_min*(math.Sqrt(3)/2), 2)

	R_l := round(L*R0, 2)
	X_l := round(L*X0, 2)

	R_sum_n := round(R_l+R_h_n, 2)
	X_sum_n := round(X_l+X_h_n, 2)
	Z_sum_n := round(math.Sqrt(math.Pow(R_sum_n, 2)+math.Pow(X_sum_n, 2)), 2)

	R_sum_n_min := round(R_l+R_h_n_min, 2)
	X_sum_n_min := round(X_l+X_h_n_min, 2)
	Z_sum_n_min := round(math.Sqrt(math.Pow(R_sum_n_min, 2)+math.Pow(X_sum_n_min, 2)), 2)

	I_l_n3 := round((U_nn*math.Pow(10, 3))/(math.Sqrt(3)*Z_sum_n), 2)
	I_l_n2 := round(I_l_n3*(math.Sqrt(3)/2), 2)

	I_l_n3_min := round((U_nn*math.Pow(10, 3))/(math.Sqrt(3)*Z_sum_n_min), 2)
	I_l_n2_min := round(I_l_n3_min*(math.Sqrt(3)/2), 2)

	return [3][4]float64{
		{I_h3, I_h2, I_h3_min, I_h2_min},
		{I_h_n3, I_h_n2, I_h_n3_min, I_h_n2_min},
		{I_l_n3, I_l_n2, I_l_n3_min, I_l_n2_min},
	}
}

func calcHandler(w http.ResponseWriter, r *http.Request) {

	data := PageData{}

	if r.Method == http.MethodGet {

		query := r.URL.Query()

		data.S_m, _ = strconv.ParseFloat(query.Get("S_m"), 64)
		data.U_nom, _ = strconv.ParseFloat(query.Get("U_nom"), 64)
		data.I_k, _ = strconv.ParseFloat(query.Get("I_k"), 64)
		data.T_f, _ = strconv.ParseFloat(query.Get("T_f"), 64)
		data.S_k, _ = strconv.ParseFloat(query.Get("S_k"), 64)
		data.U_cn, _ = strconv.ParseFloat(query.Get("U_cn"), 64)
		data.U_k, _ = strconv.ParseFloat(query.Get("U_k"), 64)
		data.S_nom_t, _ = strconv.ParseFloat(query.Get("S_nom_t"), 64)
		data.L, _ = strconv.ParseFloat(query.Get("L"), 64)
		data.U_k_max, _ = strconv.ParseFloat(query.Get("U_k_max"), 64)
		data.U_vn, _ = strconv.ParseFloat(query.Get("U_vn"), 64)
		data.U_nn, _ = strconv.ParseFloat(query.Get("U_nn"), 64)
		data.R_sn, _ = strconv.ParseFloat(query.Get("R_sn"), 64)
		data.X_sn, _ = strconv.ParseFloat(query.Get("X_sn"), 64)
		data.R_s_min, _ = strconv.ParseFloat(query.Get("R_s_min"), 64)
		data.X_s_min, _ = strconv.ParseFloat(query.Get("X_s_min"), 64)

		cable := chooseCable(data.S_m, data.U_nom, data.I_k, data.T_f)
		kz := calculateKZ(data.S_k, data.U_cn, data.U_k, data.S_nom_t)
		stationKz := calculateKZonStation(data.L, data.U_k_max, data.U_vn, data.U_nn, data.S_nom_t, data.R_sn, data.X_sn, data.R_s_min, data.X_s_min)

		data.Result = map[string]interface{}{
			"cable":     cable,
			"kz":        kz,
			"stationKz": stationKz,
		}

		calcPage.Execute(w, data)
	}

	if r.Method == http.MethodPost {
		values := url.Values{}

		values.Add("S_m", r.FormValue("S_m"))
		values.Add("U_nom", r.FormValue("U_nom"))
		values.Add("I_k", r.FormValue("I_k"))
		values.Add("T_f", r.FormValue("T_f"))
		values.Add("S_k", r.FormValue("S_k"))
		values.Add("U_cn", r.FormValue("U_cn"))
		values.Add("U_k", r.FormValue("U_k"))
		values.Add("S_nom_t", r.FormValue("S_nom_t"))
		values.Add("L", r.FormValue("L"))
		values.Add("U_k_max", r.FormValue("U_k_max"))
		values.Add("U_vn", r.FormValue("U_vn"))
		values.Add("U_nn", r.FormValue("U_nn"))
		values.Add("R_sn", r.FormValue("R_sn"))
		values.Add("X_sn", r.FormValue("X_sn"))
		values.Add("R_s_min", r.FormValue("R_s_min"))
		values.Add("X_s_min", r.FormValue("X_s_min"))

		http.Redirect(w, r, "/?"+values.Encode(), http.StatusSeeOther)
	}
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
