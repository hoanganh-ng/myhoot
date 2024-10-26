package helpers

type Comparable[T any] interface {
	Is(T) bool
}

// AddIfNotExist add a comparable value into a slice of comparable, return the addition
func AddIfNotExist[T Comparable[any]](family []T, needle T) ([]T, bool) {
	isExisted := false
	for i := 0; i < len(family); i++ {
		if family[i].Is(needle) {
			isExisted = true
			break
		}
	}
	if isExisted {
		return family, false
	}
	family = append(family, needle)
	return family, true
}
