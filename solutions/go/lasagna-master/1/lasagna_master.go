package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, averageTime int) int{
    if averageTime == 0 {
		averageTime = 2
	}
    total := len(layers)*averageTime
    return total
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	var noodles int
	var sauce float64

	for _, layer := range layers {
		if layer == "noodles" {
			noodles += 50
		}
		if layer == "sauce" {
			sauce += 0.2
		}
	}

	return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendIngredients, myIngredients []string) {
	secret := friendIngredients[len(friendIngredients)-1]
	myIngredients[len(myIngredients)-1] = secret
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amounts []float64, portions int) []float64 {
    // Calculate scaling factor (original recipe is for 2 portions)
    factor := float64(portions) / 2.0

    // Create a new slice to avoid modifying the original
    scaled := make([]float64, len(amounts))

    // Scale each amount
    for i, v := range amounts {
        scaled[i] = v * factor
    }

    return scaled
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
