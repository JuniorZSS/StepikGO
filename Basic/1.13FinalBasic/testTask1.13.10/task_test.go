package testTask1_13_10

import (
	"testing"
)

func TestTask(t *testing.T) {
	tests := []struct {
		a        int
		b        int
		excepted int
	}{
		{a: 100, b: 500, excepted: 497},
		{a: 1, b: 10, excepted: 7},
		{a: 1, b: 4, excepted: -1},
		{a: 1, b: 4, excepted: -1},
	}
	for _, val := range tests {
		if task10(val.a, val.b) != val.excepted {
			t.Errorf("Не найдено значение для %d, %d. Ожидалось%d", val.a, val.b, val.excepted)
		}
	}
}
