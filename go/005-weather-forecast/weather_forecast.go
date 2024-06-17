// Package weather provides weather forecast for a given city.
package weather

// CurrentCondition variables to store the current weather condition.
var CurrentCondition string

// CurrentLocation variables to store the current weather location.
var CurrentLocation string

// Forecast returns the current weather condition for a given city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
