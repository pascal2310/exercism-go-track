// Package weather does the forecast for you.
package weather

// CurrentCondition describes the current condition.
var CurrentCondition string

// CurrentLocation describes the current location.
var CurrentLocation string

// Forecast gives you an forecast based on the city and the condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
