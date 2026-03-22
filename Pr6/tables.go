package main

func getCalculationFactor(effectiveNumber float64, utilizationFactor float64) float64 {
	if effectiveNumber == 1 {
		switch utilizationFactor {
		case 0.1:
			return 8.00
		case 0.15:
			return 5.33
		case 0.2:
			return 4.00
		case 0.3:
			return 2.67
		case 0.4:
			return 2.00
		case 0.5:
			return 1.60
		case 0.6:
			return 1.33
		case 0.7:
			return 1.14
		case 0.8:
			return 1
		}
	} else if effectiveNumber == 2 {
		switch utilizationFactor {
		case 0.1:
			return 6.22
		case 0.15:
			return 4.33
		case 0.2:
			return 3.39
		case 0.3:
			return 2.45
		case 0.4:
			return 1.98
		case 0.5:
			return 1.60
		case 0.6:
			return 1.33
		case 0.7:
			return 1.14
		case 0.8:
			return 1
		}
	} else if effectiveNumber == 3 {
		switch utilizationFactor {
		case 0.1:
			return 4.06
		case 0.15:
			return 2.89
		case 0.2:
			return 2.31
		case 0.3:
			return 1.74
		case 0.4:
			return 1.45
		case 0.5:
			return 1.34
		case 0.6:
			return 1.22
		case 0.7:
			return 1.14
		case 0.8:
			return 1
		}
	} else if effectiveNumber == 4 {
		switch utilizationFactor {
		case 0.1:
			return 3.24
		case 0.15:
			return 2.35
		case 0.2:
			return 1.91
		case 0.3:
			return 1.47
		case 0.4:
			return 1.25
		case 0.5:
			return 1.21
		case 0.6:
			return 1.12
		case 0.7:
			return 1.06
		case 0.8:
			return 1
		}
	} else if effectiveNumber == 5 {
		switch utilizationFactor {
		case 0.1:
			return 2.84
		case 0.15:
			return 2.09
		case 0.2:
			return 1.72
		case 0.3:
			return 1.35
		case 0.4:
			return 1.16
		case 0.5:
			return 1.16
		case 0.6:
			return 1.08
		case 0.7:
			return 1.03
		case 0.8:
			return 1
		}
	} else if effectiveNumber == 6 {
		switch utilizationFactor {
		case 0.1:
			return 2.64
		case 0.15:
			return 1.96
		case 0.2:
			return 1.62
		case 0.3:
			return 1.28
		case 0.4:
			return 1.14
		case 0.5:
			return 1.13
		case 0.6:
			return 1.06
		case 0.7:
			return 1.01
		case 0.8:
			return 1
		}
	} else if effectiveNumber == 7 {
		switch utilizationFactor {
		case 0.1:
			return 2.49
		case 0.15:
			return 1.86
		case 0.2:
			return 1.54
		case 0.3:
			return 1.23
		case 0.4:
			return 1.12
		case 0.5:
			return 1.10
		case 0.6:
			return 1.04
		default:
			return 1
		}
	} else if effectiveNumber == 8 {
		switch utilizationFactor {
		case 0.1:
			return 2.37
		case 0.15:
			return 1.78
		case 0.2:
			return 1.48
		case 0.3:
			return 1.19
		case 0.4:
			return 1.10
		case 0.5:
			return 1.08
		case 0.6:
			return 1.02
		default:
			return 1
		}
	} else if effectiveNumber == 9 {
		switch utilizationFactor {
		case 0.1:
			return 2.27
		case 0.15:
			return 1.71
		case 0.2:
			return 1.43
		case 0.3:
			return 1.16
		case 0.4:
			return 1.09
		case 0.5:
			return 1.07
		case 0.6:
			return 1.01
		default:
			return 1
		}
	} else if effectiveNumber == 10 {
		switch utilizationFactor {
		case 0.1:
			return 2.18
		case 0.15:
			return 1.65
		case 0.2:
			return 1.39
		case 0.3:
			return 1.13
		case 0.4:
			return 1.07
		case 0.5:
			return 1.05
		default:
			return 1
		}
	} else if effectiveNumber == 11 {
		switch utilizationFactor {
		case 0.1:
			return (2.18 + 2.04) / 2
		case 0.15:
			return (1.65 + 1.56) / 2
		case 0.2:
			return (1.39 + 1.32) / 2
		case 0.3:
			return (1.13 + 1.08) / 2
		case 0.4:
			return (1.07 + 1.05) / 2
		case 0.5:
			return (1.05 + 1.03) / 2
		default:
			return 1
		}
	} else if effectiveNumber == 12 {
		switch utilizationFactor {
		case 0.1:
			return 2.04
		case 0.15:
			return 1.56
		case 0.2:
			return 1.32
		case 0.3:
			return 1.08
		case 0.4:
			return 1.05
		case 0.5:
			return 1.03
		default:
			return 1
		}
	} else if effectiveNumber == 13 {
		switch utilizationFactor {
		case 0.1:
			return (2.04 + 1.94) / 2
		case 0.15:
			return (1.56 + 1.49) / 2
		case 0.2:
			return (1.32 + 1.27) / 2
		case 0.3:
			return (1.08 + 1.05) / 2
		case 0.4:
			return (1.05 + 1.02) / 2
		case 0.5:
			return (1.03 + 1) / 2
		default:
			return 1
		}
	} else if effectiveNumber == 14 {
		switch utilizationFactor {
		case 0.1:
			return 1.94
		case 0.15:
			return 1.49
		case 0.2:
			return 1.27
		case 0.3:
			return 1.05
		case 0.4:
			return 1.02
		default:
			return 1
		}
	} else if effectiveNumber == 15 {
		switch utilizationFactor {
		case 0.1:
			return (1.94 + 1.85) / 2
		case 0.15:
			return (1.49 + 1.43) / 2
		case 0.2:
			return (1.27 + 1.23) / 2
		case 0.3:
			return (1.05 + 1.02) / 2
		case 0.4:
			return (1.02 + 1) / 2
		default:
			return 1
		}
	} else if effectiveNumber == 16 {
		switch utilizationFactor {
		case 0.1:
			return 1.85
		case 0.15:
			return 1.43
		case 0.2:
			return 1.23
		case 0.3:
			return 1.02
		default:
			return 1
		}
	} else if effectiveNumber == 17 {
		switch utilizationFactor {
		case 0.1:
			return (1.85 + 1.78) / 2
		case 0.15:
			return (1.43 + 1.39) / 2
		case 0.2:
			return (1.23 + 1.19) / 2
		case 0.3:
			return (1.02 + 1) / 2
		default:
			return 1
		}
	} else if effectiveNumber == 18 {
		switch utilizationFactor {
		case 0.1:
			return 1.78
		case 0.15:
			return 1.39
		case 0.2:
			return 1.19
		default:
			return 1
		}
	} else if effectiveNumber == 19 {
		switch utilizationFactor {
		case 0.1:
			return (1.78 + 1.72) / 2
		case 0.15:
			return (1.39 + 1.35) / 2
		case 0.2:
			return (1.19 + 1.16) / 2
		default:
			return 1
		}
	} else if effectiveNumber == 20 {
		switch utilizationFactor {
		case 0.1:
			return 1.72
		case 0.15:
			return 1.35
		case 0.2:
			return 1.16
		default:
			return 1
		}
	} else if effectiveNumber > 20 && effectiveNumber < 25 {
		switch utilizationFactor {
		case 0.1:
			return (1.72 + 1.60) / 2
		case 0.15:
			return (1.35 + 1.27) / 2
		case 0.2:
			return (1.16 + 1.10) / 2
		default:
			return 1
		}
	} else if effectiveNumber == 25 {
		switch utilizationFactor {
		case 0.1:
			return 1.60
		case 0.15:
			return 1.27
		case 0.2:
			return 1.10
		default:
			return 1
		}
	} else if effectiveNumber > 25 && effectiveNumber < 30 {
		switch utilizationFactor {
		case 0.1:
			return (1.60 + 1.51) / 2
		case 0.15:
			return (1.27 + 1.21) / 2
		case 0.2:
			return (1.10 + 1.05) / 2
		default:
			return 1
		}
	} else if effectiveNumber == 30 {
		switch utilizationFactor {
		case 0.1:
			return 1.51
		case 0.15:
			return 1.21
		case 0.2:
			return 1.05
		default:
			return 1
		}
	} else if effectiveNumber > 30 && effectiveNumber < 35 {
		switch utilizationFactor {
		case 0.1:
			return (1.51 + 1.44) / 2
		case 0.15:
			return (1.21 + 1.16) / 2
		case 0.2:
			return (1.05 + 1) / 2
		default:
			return 1
		}
	} else if effectiveNumber == 35 {
		switch utilizationFactor {
		case 0.1:
			return 1.44
		case 0.15:
			return 1.16
		default:
			return 1
		}
	} else if effectiveNumber > 35 && effectiveNumber < 40 {
		switch utilizationFactor {
		case 0.1:
			return (1.44 + 1.40) / 2
		case 0.15:
			return (1.16 + 1.13) / 2
		default:
			return 1
		}
	} else if effectiveNumber == 40 {
		switch utilizationFactor {
		case 0.1:
			return 1.40
		case 0.15:
			return 1.13
		default:
			return 1
		}
	} else if effectiveNumber > 40 && effectiveNumber < 50 {
		switch utilizationFactor {
		case 0.1:
			return (1.40 + 1.30) / 2
		case 0.15:
			return (1.13 + 1.07) / 2
		default:
			return 1
		}
	} else if effectiveNumber == 50 {
		switch utilizationFactor {
		case 0.1:
			return 1.30
		case 0.15:
			return 1.07
		default:
			return 1
		}
	} else if effectiveNumber > 50 && effectiveNumber < 60 {
		switch utilizationFactor {
		case 0.1:
			return (1.30 + 1.25) / 2
		case 0.15:
			return (1.07 + 1.03) / 2
		default:
			return 1
		}
	} else if effectiveNumber == 60 {
		switch utilizationFactor {
		case 0.1:
			return 1.25
		case 0.15:
			return 1.03
		default:
			return 1
		}
	} else if effectiveNumber > 60 && effectiveNumber < 80 {
		switch utilizationFactor {
		case 0.1:
			return (1.25 + 1.16) / 2
		case 0.15:
			return (1.03 + 1) / 2
		default:
			return 1
		}
	} else if effectiveNumber == 80 {
		switch utilizationFactor {
		case 0.1:
			return 1.16
		default:
			return 1
		}
	} else if effectiveNumber > 80 && effectiveNumber < 100 {
		switch utilizationFactor {
		case 0.1:
			return (1.16 + 1) / 2
		default:
			return 1
		}
	} else {
		return 1
	}
	return 1
}

func getAllCalculationFactor(effectiveNumber float64, utilizationFactor float64) float64 {
	if effectiveNumber == 1 {
		switch utilizationFactor {
		case 0.1:
			return 8.00
		case 0.15:
			return 5.33
		case 0.2:
			return 4.00
		case 0.3:
			return 2.67
		case 0.4:
			return 2.00
		case 0.5:
			return 1.60
		case 0.6:
			return 1.33
		case 0.7:
			return 1.14
		default:
			return 1.14
		}
	} else if effectiveNumber == 2 {
		switch utilizationFactor {
		case 0.1:
			return 5.01
		case 0.15:
			return 3.44
		case 0.2:
			return 2.69
		case 0.3:
			return 1.9
		case 0.4:
			return 1.52
		case 0.5:
			return 1.24
		case 0.6:
			return 1.11
		case 0.7:
			return 1.0
		default:
			return 1.0
		}
	} else if effectiveNumber == 3 {
		switch utilizationFactor {
		case 0.1:
			return 2.4
		case 0.15:
			return 2.17
		case 0.2:
			return 1.8
		case 0.3:
			return 1.42
		case 0.4:
			return 1.23
		case 0.5:
			return 1.14
		case 0.6:
			return 1.08
		case 0.7:
			return 1.0
		default:
			return 1.0
		}
	} else if effectiveNumber == 4 {
		switch utilizationFactor {
		case 0.1:
			return 2.28
		case 0.15:
			return 1.73
		case 0.2:
			return 1.46
		case 0.3:
			return 1.19
		case 0.4:
			return 1.06
		case 0.5:
			return 1.04
		case 0.6:
			return 1.0
		case 0.7:
			return 0.97
		default:
			return 0.97
		}
	} else if effectiveNumber == 5 {
		switch utilizationFactor {
		case 0.1:
			return 1.31
		case 0.15:
			return 1.12
		case 0.2:
			return 1.02
		case 0.3:
			return 1.0
		case 0.4:
			return 0.98
		case 0.5:
			return 0.96
		case 0.6:
			return 0.94
		case 0.7:
			return 0.93
		default:
			return 0.93
		}
	} else if effectiveNumber >= 6 && effectiveNumber <= 8 {
		switch utilizationFactor {
		case 0.1:
			return 1.2
		case 0.15:
			return 1.0
		case 0.2:
			return 0.96
		case 0.3:
			return 0.95
		case 0.4:
			return 0.94
		case 0.5:
			return 0.93
		case 0.6:
			return 0.92
		case 0.7:
			return 0.91
		default:
			return 0.91
		}
	} else if effectiveNumber >= 9 && effectiveNumber <= 10 {
		switch utilizationFactor {
		case 0.1:
			return 1.1
		case 0.15:
			return 0.97
		case 0.2:
			return 0.91
		case 0.3:
			return 0.9
		default:
			return 0.9
		}
	} else if effectiveNumber >= 11 && effectiveNumber <= 25 {
		switch utilizationFactor {
		case 0.1:
			return 0.8
		case 0.15:
			return 0.8
		case 0.2:
			return 0.8
		case 0.3:
			return 0.85
		case 0.4:
			return 0.85
		case 0.5:
			return 0.85
		case 0.6:
			return 0.9
		case 0.7:
			return 0.9
		default:
			return 0.9
		}
	} else if effectiveNumber >= 26 && effectiveNumber <= 50 {
		switch utilizationFactor {
		case 0.1:
			return 0.75
		case 0.15:
			return 0.75
		case 0.2:
			return 0.75
		case 0.3:
			return 0.75
		case 0.4:
			return 0.75
		case 0.5:
			return 0.8
		case 0.6:
			return 0.85
		case 0.7:
			return 0.85
		default:
			return 0.85
		}
	} else {
		switch utilizationFactor {
		case 0.1:
			return 0.65
		case 0.15:
			return 0.65
		case 0.2:
			return 0.65
		case 0.3:
			return 0.7
		case 0.4:
			return 0.7
		case 0.5:
			return 0.75
		case 0.6:
			return 0.8
		case 0.7:
			return 0.8
		default:
			return 0.8
		}
	}
}
