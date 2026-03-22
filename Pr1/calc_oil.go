package main

// calcWorkMass computes the work mass for each element based on fire amounts
// and a conversion coefficient.
func calcWorkMass(carbonFire, hydrogenFire, sulfurFire, oxygenFire, fireToWorkMassCoef float64) map[string]float64 {
	workMass := make(map[string]float64)
	workMass["carbon_work"] = carbonFire * fireToWorkMassCoef
	workMass["hydrogen_work"] = hydrogenFire * fireToWorkMassCoef
	workMass["sulfur_work"] = sulfurFire * fireToWorkMassCoef
	workMass["oxygen_work"] = oxygenFire * fireToWorkMassCoef
	return workMass
}

// calcFireToWorkMassCoef calculates the coefficient that converts fire mass to work mass.
func calcFireToWorkMassCoef(wetness, ash float64) float64 {
	return (100 - wetness - ash) / 100
}

// calcBurnWorkTemp calculates the temperature of the burning work.
func calcBurnWorkTemp(wetness, ashWork float64) float64 {
	return ((100 - wetness - ashWork) / 100) - 0.025*wetness
}

// calcFuelOil aggregates all intermediate calculations into a single result map.
func calcFuelOil(carbon, hydrogen, sulfur, oxygen, ash, wetness, vanadium float64) map[string]interface{} {
	ashWork := ash * ((100 - wetness) / 100)
	vanadiumWork := vanadium * ((100 - wetness) / 100)

	fireToWorkMassCoef := calcFireToWorkMassCoef(wetness, ashWork)
	workMass := calcWorkMass(carbon, hydrogen, sulfur, oxygen, fireToWorkMassCoef)

	burnWorkTemp := calcBurnWorkTemp(wetness, ashWork)

	return map[string]interface{}{
		"ashWork":      ashWork,
		"vanadiumWork": vanadiumWork,
		"carbonWork":   workMass["carbon_work"],
		"hydrogenWork": workMass["hydrogen_work"],
		"sulfurWork":   workMass["sulfur_work"],
		"oxygenWork":   workMass["oxygen_work"],
		"burnWorkTemp": burnWorkTemp,
	}
}
