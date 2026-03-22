package main

import (
	"html/template"
	"math"
	"net/http"
	"net/url"
	"strconv"
)

type PageData struct {
	Sig1 float64
	Sig2 float64
	P_s  float64
	B    float64

	Result map[string]interface{}
}

var calcPage = template.Must(template.ParseFiles("index.html"))

func calcHandler(w http.ResponseWriter, r *http.Request) {

	data := PageData{}

	if r.Method == http.MethodGet {

		query := r.URL.Query()

		data.Sig1, _ = strconv.ParseFloat(query.Get("sig1"), 64)
		data.Sig2, _ = strconv.ParseFloat(query.Get("sig2"), 64)
		data.P_s, _ = strconv.ParseFloat(query.Get("P_s"), 64)
		data.B, _ = strconv.ParseFloat(query.Get("B"), 64)

		data.Result = calcMargin(data.Sig1, data.Sig2, data.P_s, data.B)

		calcPage.Execute(w, data)
	}

	if r.Method == http.MethodPost {
		values := url.Values{}

		values.Add("sig1", r.FormValue("sig1"))
		values.Add("sig2", r.FormValue("sig2"))
		values.Add("P_s", r.FormValue("P_s"))
		values.Add("B", r.FormValue("B"))

		http.Redirect(w, r, "/?"+values.Encode(), http.StatusSeeOther)
	}
}

func calcMargin(sig1, sig2, P_s, B float64) map[string]interface{} {
	const delta = 0.05
	pMin, pMax := P_s-P_s*delta, P_s+P_s*delta

	deltaW1 := math.Round(integrate(pMin, pMax, sig1, P_s)*100) / 100
	W1 := int(math.Round(P_s * 24 * deltaW1))
	P1 := int(math.Round(float64(W1) * B))
	W2 := int(math.Round(P_s * 24 * (1 - deltaW1)))
	H1 := int(math.Round(float64(W2) * B))

	deltaW2 := math.Round(integrate(pMin, pMax, sig2, P_s)*100) / 100
	W3 := int(math.Round(P_s * 24 * deltaW2))
	P2 := int(math.Round(float64(W3) * B))
	W4 := int(math.Round(P_s * 24 * (1 - deltaW2)))
	H2 := int(math.Round(float64(W4) * B))

	return map[string]interface{}{
		"deltaW1": deltaW1 * 100,
		"P1":      P1,
		"H1":      H1,
		"P1-H1":   P1 - H1,
		"deltaW2": deltaW2 * 100,
		"P2":      P2,
		"H2":      H2,
		"P2-H2":   P2 - H2,
	}
}

func integrate(pStart, pEnd, sig, P_s float64) float64 {
	toIntegrate := func(sig, P_s, p float64) float64 {
		dif := p - P_s
		power := math.Pow(dif, 2) / (2 * sig * sig)
		numerator := math.Exp(power)
		denominator := sig * math.Sqrt(6.28)
		return numerator / denominator
	}

	sum := 0.5*toIntegrate(sig, P_s, pStart) + 0.5*toIntegrate(sig, P_s, pEnd)
	iterations := 1000
	delta := (pEnd - pStart) / float64(iterations)
	for i := 1; i < iterations; i++ {
		sum += toIntegrate(sig, P_s, pStart+float64(i)*delta)
	}
	sum *= delta
	return sum
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
