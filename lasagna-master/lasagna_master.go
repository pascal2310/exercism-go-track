package lasagna

import (
	"strings"
)

func PreparationTime(layers []string, prepTime int) int {
	if prepTime == 0 {
		prepTime = 2
	}
	return len(layers) * prepTime
}

func Quantities(ingredients []string) (noodles int, sauce float64) {
	for i := range ingredients {
		ingredient := ingredients[i]

		if strings.EqualFold(ingredient, "sauce") {
			sauce += 0.2
		}
		if strings.EqualFold(ingredient, "noodles") {
			noodles += 50
		}
	}
	return

}

func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	calculatedQuantities := []float64{}
	for i := range quantities {
		perPerson := quantities[i] / 2
		calculatedQuantities = append(calculatedQuantities, perPerson*float64(portions))
	}
	return calculatedQuantities
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
