package FuncStrPoint

import (
	"fmt"
)

type Steps struct {
	On    bool
	Ammo  int
	Power int
}

func (s *Steps) Shot() bool {
	if s.On == false {
		return false
	}
	if s.Ammo > 0 {
		s.Ammo--
		return true
	}
	return true
}

func (s *Steps) RideBike() bool {
	if s.On == false {
		return false
	}
	if s.Power > 0 {
		s.Power--
		return true
	}
	return false
}

func Task6() {
	fmt.Println("В задаче реализована структура и пара методов, \n" +
		"мы назвали условно нашего робота Степ, и нанесли ему урон:")

	Step := Steps{true, 100, 100}
	Step.Shot()
	Step.RideBike()
	fmt.Println("Step:", Step)

	testStruct := Steps{On: true, Ammo: 100, Power: 100}
	for i := 1; i < 55; i++ {
		testStruct.Shot()
	}

	fmt.Println("Step:", testStruct)
}
