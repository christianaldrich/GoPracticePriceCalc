package conversion

import (
	"errors"
	"fmt"
	"strconv"
)

func StringsToFloat(strings []string) ([]float64, error) {

	var floats []float64

	for _, stringVal := range strings {
		floatval, err := strconv.ParseFloat(stringVal, 64)

		if err != nil {
			fmt.Println(err)
			return nil, errors.New("failed to convert string to float")
		}

		floats = append(floats, floatval)

	}
	return floats, nil
}
