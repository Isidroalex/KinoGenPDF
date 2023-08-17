package transport

import (
	"testing"
)

func TestShowPuppyCount(t *testing.T) {
	// Создайте тестовые данные
	puppies := []*Puppy{
		{Dog: Dog{Sex: "male"}},
		{Dog: Dog{Sex: "female"}},
		{Dog: Dog{Sex: "male"}},
	}

	litter := Litter{Puppy: puppies}

	// Тестируем метод для разных вариантов аргументов
	testCases := []struct {
		sex           string
		expectedCount int
	}{
		{sex: "male", expectedCount: 2},
		{sex: "female", expectedCount: 1},
		{sex: "", expectedCount: 3},
	}

	for _, tc := range testCases {
		t.Run(tc.sex, func(t *testing.T) {
			actualCount := litter.ShowPuppyCount(tc.sex)
			if actualCount != tc.expectedCount {
				t.Errorf("For sex '%s', expected count: %d, but got: %d", tc.sex, tc.expectedCount, actualCount)
			}
		})
	}
}
