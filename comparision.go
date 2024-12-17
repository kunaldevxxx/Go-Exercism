package purchase

func NeedsLicense(kind string) bool {
    if kind == "car" || kind == "truck" {
        return true
    } else {
        return false
    }
}
func ChooseVehicle(option1, option2 string) string {
    if option1 < option2 {
        return option1 + " is clearly the better choice."
    } else {
        return option2 + " is clearly the better choice."
    }
}

func CalculateResellPrice(originalPrice, age float64) float64 {
    if age < 3 {
        return 0.8 * originalPrice
    } else if age >= 10 {
        return 0.5 * originalPrice
    } else if age >= 3 && age < 10 {
        return 0.7 * originalPrice
    }
    return 0 
}