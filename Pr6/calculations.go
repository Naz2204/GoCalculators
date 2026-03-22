package main

import (
	"math"
)

func calcAllNominalCapacity(number int, nominalCapacity float64) float64 {
	return float64(number) * nominalCapacity
}

func calculatedCurrent(allNominalCapacity, loadVoltage, loadPowerFactor, nominalEfficiency float64) float64 {
	return allNominalCapacity / (math.Sqrt(3) * loadVoltage * loadPowerFactor * nominalEfficiency)
}

func calcGroupUtilizationFactor(generalNominalCapacity []float64, utilizationFactor []float64) float64 {
	numerator := 0.0
	denominator := 0.0
	for i := 0; i < len(utilizationFactor); i++ {
		numerator += generalNominalCapacity[i] * utilizationFactor[i]
		denominator += generalNominalCapacity[i]
	}
	if denominator == 0 {
		return 0
	}
	return numerator / denominator
}

func calcGroupEffectiveNumber(number []int, nominalCapacity []float64) float64 {
	numerator := 0.0
	denominator := 0.0
	for i := 0; i < len(number); i++ {
		numerator += float64(number[i]) * nominalCapacity[i]
		denominator += float64(number[i]) * math.Pow(nominalCapacity[i], 2)
	}
	if denominator == 0 {
		return 0
	}
	return math.Ceil(math.Pow(numerator, 2) / denominator)
}

func calcEstimatedActiveLoad(generalNominalCapacity []float64, utilizationFactor []float64, calculationFactor float64) float64 {
	secondOperand := 0.0
	for i := 0; i < len(generalNominalCapacity); i++ {
		secondOperand += generalNominalCapacity[i] * utilizationFactor[i]
	}
	return calculationFactor * secondOperand
}

func calcEstimatedReactiveLoad(effectiveNumber float64, generalNominalCapacity []float64, utilizationFactor []float64, calculationFactor float64, reactivePowerFactor []float64) float64 {
	secondOperand := 0.0
	for i := 0; i < len(generalNominalCapacity); i++ {
		secondOperand += generalNominalCapacity[i] * utilizationFactor[i] * reactivePowerFactor[i]
	}
	if effectiveNumber <= 10 {
		return 1.1 * secondOperand
	}
	return secondOperand
}

func calcFullPower(activeLoad, reactiveLoad float64) float64 {
	underRoot := math.Pow(activeLoad, 2) + math.Pow(reactiveLoad, 2)
	return math.Sqrt(underRoot)
}

func calcGroupCurrent(activeLoad, loadVoltage float64) float64 {
	return activeLoad / loadVoltage
}

func calcWorkshopUtilizationFactor(allWorkshopNominalCapacity, allWorkshopAverageActiveLoad float64) float64 {
	if allWorkshopNominalCapacity == 0 {
		return 0
	}
	return allWorkshopAverageActiveLoad / allWorkshopNominalCapacity
}

func calcWorkshopEffectiveNumber(allWorkshopNominalCapacity, allSquaredWorkshopNominalCapacity float64) float64 {
	if allSquaredWorkshopNominalCapacity == 0 {
		return 0
	}
	return math.Ceil(math.Pow(allWorkshopNominalCapacity, 2) / allSquaredWorkshopNominalCapacity)
}

func calcEstimatedActiveTyreLoad(calculationFactor, allWorkshopAverageActiveLoad float64) float64 {
	return calculationFactor * allWorkshopAverageActiveLoad
}

func calcEstimatedReactiveTyreLoad(calculationFactor, allWorkshopAverageActiveLoadTg float64) float64 {
	return calculationFactor * allWorkshopAverageActiveLoadTg
}

func calcTyreFullPower(estimatedActiveTyreLoad, estimatedReactiveTyreLoad float64) float64 {
	underRoot := math.Pow(estimatedActiveTyreLoad, 2) + math.Pow(estimatedReactiveTyreLoad, 2)
	return math.Sqrt(underRoot)
}

func calcTyreCurrent(estimatedActiveTyreLoad, loadTireVoltage float64) float64 {
	return estimatedActiveTyreLoad / loadTireVoltage
}
