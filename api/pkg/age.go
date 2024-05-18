package pkg

import "time"

func GetAgeUser(dateBirth string) (int, error) {
	currentDate := time.Now()

	date, err := time.Parse("2006-01-02", dateBirth)
	if err != nil {
		return 0, err
	}

	years := currentDate.Year() - date.Year()

	return years, nil
}

func GetAgeCategoryId(age int) int {
	if age > 0 && age < 6 {
		return 3
	} else if age >= 6 && age < 12 {
		return 2
	} else if age >= 12 && age < 16 {
		return 1
	} else if age >= 16 && age < 18 {
		return 4
	}

	return 5
}
