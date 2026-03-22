package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
)

type PageData struct {
	TotalNumber          int       `json:"total_number"`
	TotalNominalCapacity float64   `json:"total_nominal_capacity"`
	AverageActiveLoad    float64   `json:"average_active_load"`
	AverageReactiveLoad  float64   `json:"average_reactive_load"`
	TotalSquaredPower    float64   `json:"total_squared_power"`
	LoadVoltage          float64   `json:"load_voltage"`
	Name                 []string  `json:"name"`
	NominalEfficiency    []float64 `json:"nominal_efficiency"`
	LoadCowerCoef        []float64 `json:"load_cower_coef"`
	Number               []int     `json:"number"`
	NominalCapacity      []float64 `json:"nominal_capacity"`
	UtilizationFactor    []float64 `json:"utilization_factor"`
	ReactivePowerFactor  []float64 `json:"reactive_power_factor"`
}

type CalculationResults struct {
	GroupUtilizationFactor    float64
	EffectiveNumber           float64
	EPCalculationFactor       float64
	EstimatedActiveLoad       float64
	EstimatedReactiveLoad     float64
	GroupFullPower            float64
	GroupCurrent              float64
	WorkshopUtilizationFactor float64
	WorkshopEffectiveNumber   float64
	WorkshopCalculationFactor float64
	EstimatedActiveTyreLoad   float64
	EstimatedReactiveTyreLoad float64
	WorkshopFullPower         float64
	WorkshopCurrent           float64
}

func roundToTwo(val float64) float64 {
	return math.Round(val*100) / 100
}

func roundToOne(val float64) float64 {
	return math.Round(val*10) / 10
}

func calculatePowerLoad(data PageData) CalculationResults {

	var generalNominalCapacity []float64
	var generalCalculatedCurrent []float64

	fmt.Println(data)

	for i := 0; i < len(data.Number); i++ {
		capacity := calcAllNominalCapacity(data.Number[i], data.NominalCapacity[i])
		generalNominalCapacity = append(generalNominalCapacity, capacity)

		curr := calculatedCurrent(capacity, data.LoadVoltage, data.LoadCowerCoef[i], data.NominalEfficiency[i])
		generalCalculatedCurrent = append(generalCalculatedCurrent, curr)
	}

	groupUtilizationFactor := roundToTwo(calcGroupUtilizationFactor(generalNominalCapacity, data.UtilizationFactor))
	effectiveNumber := roundToTwo(calcGroupEffectiveNumber(data.Number, data.NominalCapacity))

	var epCalculationFactor float64
	if groupUtilizationFactor == 0.15 {
		epCalculationFactor = getCalculationFactor(effectiveNumber, groupUtilizationFactor)
	} else {
		epCalculationFactor = getCalculationFactor(effectiveNumber, roundToOne(groupUtilizationFactor))
	}

	estimatedActiveLoad := roundToTwo(calcEstimatedActiveLoad(generalNominalCapacity, data.UtilizationFactor, epCalculationFactor))
	estimatedReactiveLoad := roundToTwo(calcEstimatedReactiveLoad(effectiveNumber, generalNominalCapacity, data.UtilizationFactor, epCalculationFactor, data.ReactivePowerFactor))
	groupFullPower := roundToTwo(calcFullPower(estimatedActiveLoad, estimatedReactiveLoad))
	groupCurrent := roundToTwo(calcGroupCurrent(estimatedActiveLoad, data.LoadVoltage))

	workshopUtilizationFactor := roundToTwo(calcWorkshopUtilizationFactor(data.TotalNominalCapacity, data.AverageActiveLoad))
	workshopEffectiveNumber := roundToTwo(calcWorkshopEffectiveNumber(data.TotalNominalCapacity, data.TotalSquaredPower))

	var workshopCalculationFactor float64
	if workshopUtilizationFactor == 0.15 {
		workshopCalculationFactor = getAllCalculationFactor(workshopEffectiveNumber, workshopUtilizationFactor)
	} else {
		workshopCalculationFactor = getAllCalculationFactor(workshopEffectiveNumber, roundToOne(workshopUtilizationFactor))
	}

	estimatedActiveTyreLoad := roundToTwo(calcEstimatedActiveTyreLoad(workshopCalculationFactor, data.AverageActiveLoad))
	estimatedReactiveTyreLoad := roundToTwo(calcEstimatedReactiveTyreLoad(workshopCalculationFactor, data.AverageReactiveLoad))
	workshopFullPower := roundToTwo(calcTyreFullPower(estimatedActiveTyreLoad, estimatedReactiveTyreLoad))
	workshopCurrent := roundToTwo(calcTyreCurrent(estimatedActiveTyreLoad, data.LoadVoltage))

	return CalculationResults{
		GroupUtilizationFactor:    groupUtilizationFactor,
		EffectiveNumber:           effectiveNumber,
		EPCalculationFactor:       epCalculationFactor,
		EstimatedActiveLoad:       estimatedActiveLoad,
		EstimatedReactiveLoad:     estimatedReactiveLoad,
		GroupFullPower:            groupFullPower,
		GroupCurrent:              groupCurrent,
		WorkshopUtilizationFactor: workshopUtilizationFactor,
		WorkshopEffectiveNumber:   workshopEffectiveNumber,
		WorkshopCalculationFactor: workshopCalculationFactor,
		EstimatedActiveTyreLoad:   estimatedActiveTyreLoad,
		EstimatedReactiveTyreLoad: estimatedReactiveTyreLoad,
		WorkshopFullPower:         workshopFullPower,
		WorkshopCurrent:           workshopCurrent,
	}
}

func handler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var data PageData
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(calculatePowerLoad(data))
}

func main() {

	http.HandleFunc("/calculate", handler)
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.ListenAndServe(":8080", nil)
}
