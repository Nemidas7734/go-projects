package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func convertLength(value float64, fromUnit, toUnit string) float64 {

	// Conversion rates for length units (all units converted to millimeters)
	var conversionRates = map[string]map[string]float64{
		"millimeter": {
			"centimeter": 0.1,
			"meter":      0.001,
			"kilometer":  0.000001,
			"inch":       0.0393701,
			"foot":       0.00328084,
			"yard":       0.00109361,
			"mile":       6.2137e-7,
		},
		"centimeter": {
			"millimeter": 10,
			"meter":      0.01,
			"kilometer":  0.00001,
			"inch":       0.393701,
			"foot":       0.0328084,
			"yard":       0.0109361,
			"mile":       6.2137e-6,
		},
		"meter": {
			"millimeter": 1000,
			"centimeter": 100,
			"kilometer":  0.001,
			"inch":       39.3701,
			"foot":       3.28084,
			"yard":       1.09361,
			"mile":       0.000621371,
		},
		"kilometer": {
			"millimeter": 1000000,
			"centimeter": 100000,
			"meter":      1000,
			"inch":       39370.1,
			"foot":       3280.84,
			"yard":       1093.61,
			"mile":       0.621371,
		},
		"inch": {
			"millimeter": 25.4,
			"centimeter": 2.54,
			"meter":      0.0254,
			"kilometer":  2.54e-5,
			"foot":       0.0833333,
			"yard":       0.0277778,
			"mile":       1.5783e-5,
		},
		"foot": {
			"millimeter": 304.8,
			"centimeter": 30.48,
			"meter":      0.3048,
			"kilometer":  0.0003048,
			"inch":       12,
			"yard":       0.333333,
			"mile":       0.000189394,
		},
		"yard": {
			"millimeter": 914.4,
			"centimeter": 91.44,
			"meter":      0.9144,
			"kilometer":  0.0009144,
			"inch":       36,
			"foot":       3,
			"mile":       0.000568182,
		},
		"mile": {
			"millimeter": 1609344,
			"centimeter": 160934,
			"meter":      1609.34,
			"kilometer":  1.60934,
			"inch":       63360,
			"foot":       5280,
			"yard":       1760,
		},
	}

	return value * conversionRates[fromUnit][toUnit]
}

func convertWeight(value float64, fromUnit, toUnit string) float64 {

	// Conversion rates for weight units (all units converted to grams)
	var weightConversions = map[string]map[string]float64{
		"milligram": {
			"gram":     0.001,
			"kilogram": 1e-6,
			"ounce":    3.527396194e-5,
			"pound":    2.204622621e-6,
		},
		"gram": {
			"milligram": 1000,
			"kilogram":  0.001,
			"ounce":     0.03527396,
			"pound":     0.00220462,
		},
		"kilogram": {
			"milligram": 1000000,
			"gram":      1000,
			"ounce":     35.27396,
			"pound":     2.20462,
		},
		"ounce": {
			"milligram": 28349.5,
			"gram":      28.3495,
			"kilogram":  0.0283495,
			"pound":     0.0625,
		},
		"pound": {
			"milligram": 453592,
			"gram":      453.592,
			"kilogram":  0.453592,
			"ounce":     16,
		},
	}

	return value * weightConversions[fromUnit][toUnit]
}


func convertTemperature(value float64, fromUnit, toUnit string) float64 {

	var tempValue float64

	// Convert to Celsius
	switch fromUnit {
	case "celsius":
		tempValue = value
	case "fahrenheit":
		tempValue = (value - 32) * 5 / 9
	case "kelvin":
		tempValue = value - 273.15
	}

	// Convert from Celsius to the target unit
	switch toUnit {
	case "celsius":
		return tempValue
	case "fahrenheit":
		return (tempValue * 9 / 5) + 32
	case "kelvin":
		return tempValue + 273.15
	}

	return tempValue
}


func convertHandler(w http.ResponseWriter, r *http.Request) {
	// Enable CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		return
	}

	var data struct {
		Value   float64 `json:"value"`
		FromUnit string  `json:"fromUnit"`
		ToUnit   string  `json:"toUnit"`
		UnitType string  `json:"unitType"` 
	}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var convertedValue float64

	switch data.UnitType {
	case "length":
		convertedValue = convertLength(data.Value, data.FromUnit, data.ToUnit)
	case "weight":
		convertedValue = convertWeight(data.Value, data.FromUnit, data.ToUnit)
	case "temperature":
		convertedValue = convertTemperature(data.Value, data.FromUnit, data.ToUnit)
	default:
		http.Error(w, "Invalid unit type", http.StatusBadRequest)
		return
	}

	response := struct {
		Result float64 `json:"result"`
	}{Result: convertedValue}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {

	http.HandleFunc("/convert", convertHandler)

	fmt.Println("Server running on port 8090...")
	http.ListenAndServe(":8090", nil)
}
