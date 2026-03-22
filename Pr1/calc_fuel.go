package main

func calcFuel(
	carbon, hydrogen, oxygen, sulfur, nitrogen, wetness, ash float64,
) map[string]interface{} {

	dryMassCoef := calcDryMassCoef(wetness)
	fireMassCoef := calcFireMassCoef(wetness, ash)

	dryValues := calcDryMass(
		carbon, hydrogen, sulfur, nitrogen, oxygen, ash, dryMassCoef,
	)
	fireValues := calcFireMass(
		carbon, hydrogen, sulfur, nitrogen, oxygen, fireMassCoef,
	)

	lowBurnTempK := calcLowBurnTemp(carbon, hydrogen, oxygen, sulfur, wetness) / 1000.0
	lowBurnDryTemp := calcLowBurnDryTemp(lowBurnTempK, wetness, dryMassCoef)
	lowBurnFireTemp := calcLowBurnFireTemp(lowBurnTempK, wetness, fireMassCoef)

	return map[string]interface{}{
		"wetDryCoef":      dryMassCoef,
		"wetFireCoef":     fireMassCoef,
		"dryValues":       dryValues,
		"fireValues":      fireValues,
		"lowBurnTemp":     lowBurnTempK,
		"lowBurnDryTemp":  lowBurnDryTemp,
		"lowBurnFireTemp": lowBurnFireTemp,
	}
}

func calcDryMassCoef(wetness float64) float64 {
	return 100.0 / (100.0 - wetness)
}

func calcFireMassCoef(wetness, ash float64) float64 {
	return 100.0 / (100.0 - wetness - ash)
}

func calcDryMass(
	carbon, hydrogen, sulfur, nitrogen, oxygen, ash, dryMassCoef float64,
) map[string]float64 {
	return map[string]float64{
		"carbon_dry":   carbon * dryMassCoef,
		"hydrogen_dry": hydrogen * dryMassCoef,
		"sulfur_dry":   sulfur * dryMassCoef,
		"nitrogen_dry": nitrogen * dryMassCoef,
		"oxygen_dry":   oxygen * dryMassCoef,
		"ash_dry":      ash * dryMassCoef,
	}
}

func calcFireMass(
	carbon, hydrogen, sulfur, nitrogen, oxygen, fireMassCoef float64,
) map[string]float64 {
	return map[string]float64{
		"carbon_fire":   carbon * fireMassCoef,
		"hydrogen_fire": hydrogen * fireMassCoef,
		"sulfur_fire":   sulfur * fireMassCoef,
		"nitrogen_fire": nitrogen * fireMassCoef,
		"oxygen_fire":   oxygen * fireMassCoef,
	}
}

func calcLowBurnTemp(
	carbon, hydrogen, oxygen, sulfur, wetness float64,
) float64 {
	return 339.0*carbon + 1030.0*hydrogen - 108.8*(oxygen-sulfur) - 25.0*wetness
}

func calcLowBurnDryTemp(lowBurnTemp, wetness, dryMassCoef float64) float64 {
	return (lowBurnTemp + 0.025*wetness) * dryMassCoef
}

func calcLowBurnFireTemp(lowBurnTemp, wetness, fireMassCoef float64) float64 {
	return (lowBurnTemp + 0.025*wetness) * fireMassCoef
}
