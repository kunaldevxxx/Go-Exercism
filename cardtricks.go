package cards

func FavoriteCards() []int {
    s := []int{2, 6, 9}
    return s
}

func GetItem(slice []int, index int) int {
    if index < 0 || index >= len(slice) {
        return -1
    }
    return slice[index]
}

func SetItem(slice []int, index, value int) []int {
    if index < 0 || index >= len(slice) {
        return append(slice, value)
    }
    slice[index] = value
    return slice
}

func PrependItems(slice []int, values ...int) []int {
    return append(values, slice...)
}

func RemoveItem(slice []int, index int) []int {
    if index < 0 || index >= len(slice) {
        return slice
    }
    return append(slice[:index], slice[index+1:]...)
}