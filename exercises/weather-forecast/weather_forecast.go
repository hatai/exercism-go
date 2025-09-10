// Package weather provides a function to forecast the weather condition of a city.
package weather

// CurrentCondition is the current weather condition of the city.
var CurrentCondition string
// CurrentLocation is the current location of the city.
var CurrentLocation string

// Forecast is the function to forecast the weather condition of a city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
