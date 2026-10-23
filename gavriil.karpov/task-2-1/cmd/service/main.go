package main

import (
	"errors"
	"fmt"
	"os"
)

const (
	maximumTemperature = 30
	minimumTemperature = 15
	lessOrEqual        = "<="
	greaterOrEqual     = ">="
)

var (
	ErrInvalidSign = errors.New("invalid sign")
)

type DepartmentTemperatureRange struct {
	lowestTemperature  int
	highestTemperature int
}

func New() *DepartmentTemperatureRange {
	return &DepartmentTemperatureRange{lowestTemperature: minimumTemperature,
		highestTemperature: maximumTemperature,
	}
}

func (t *DepartmentTemperatureRange) ChangeTemperatures(inputTemperature int,
	sign string) error {
	switch sign {
	case lessOrEqual:
		if inputTemperature < t.highestTemperature {
			t.highestTemperature = inputTemperature
		}
	case greaterOrEqual:
		if inputTemperature > t.lowestTemperature {
			t.lowestTemperature = inputTemperature
		}
	default:
		return fmt.Errorf("Invalid sign %s: %w", sign, ErrInvalidSign)
	}
	return nil
}

func (t *DepartmentTemperatureRange) calculateOptimalTemperature() int {
	if t.highestTemperature < t.lowestTemperature {
		return -1
	}
	return t.lowestTemperature
}

func main() {
	var departmentsNumber int
	if _, err := fmt.Scan(&departmentsNumber); err != nil {
		fmt.Fprintln(os.Stderr, "Error: invalid departments number:", err)
		return
	}
	if departmentsNumber < 1 {
		fmt.Fprintln(os.Stderr, "Error: departments number must be positive")
		return
	}
	for i := 0; i < departmentsNumber; i++ {
		var headcount int
		if _, err := fmt.Scan(&headcount); err != nil {
			fmt.Fprintln(os.Stderr, "Error: invalid headcount:", err)
			return
		}
		if headcount < 1 {
			fmt.Fprintln(os.Stderr, "Error: headcount must be positive")
			return
		}
		temperatureRange := New()
		for j := 0; j < headcount; j++ {
			var sign string
			var temperature int
			if _, err := fmt.Scan(&sign, &temperature); err != nil {
				fmt.Fprintln(os.Stderr, "Error: invalid Input:", err)
				return
			}
			if err := temperatureRange.ChangeTemperatures(temperature, sign); err != nil {
				fmt.Fprintln(os.Stderr, "Error:", err)
				return
			}
			fmt.Println(temperatureRange.calculateOptimalTemperature())
		}
	}
}
