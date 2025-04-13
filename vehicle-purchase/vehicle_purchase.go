package purchase

import "fmt"

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	var betterChoice string
	if option1 < option2 {
		betterChoice = option1
	} else {
		betterChoice = option2
	}
	return fmt.Sprintf("%s is clearly the better choice.", betterChoice)
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var discountedPrice float64
	if age < 3 {
		discountedPrice = originalPrice * 0.8
	} else if age >= 3 && age < 10 {
		discountedPrice = originalPrice * 0.7
	} else {
		discountedPrice = originalPrice * 0.5
	}
	return discountedPrice

}
