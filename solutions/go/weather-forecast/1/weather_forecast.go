// Package weather provides tools for forecasting weather conditions for a given location.
package weather


var (
    // CurrentCondition represents the current weather condition.
	CurrentCondition string
    // CurrentLocation represents the current location for which the weather is reported.
	CurrentLocation  string
)
// Forecast sets the current location and weather condition and returns a formatted forecast string.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
