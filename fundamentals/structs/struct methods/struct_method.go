package main

import "fmt"

func main() {

	s1 := Student{"ak", []int{70,90,80,85}, 21}

	// fmt.Println(s1.age)

	s1.setage(6)

	// fmt.Println(s1.age)

	averagegrade := s1.getAverageGrad()

	fmt.Println(averagegrade)
}

type Student struct {
	name   string
	grades []int
	age    int
}

func (s *Student) setage(age int) {
	s.age = age
}

func (s *Student) getAverageGrad() float32 {
	sum := 0
	for _, v := range s.grades {
		sum += v
	}
	return float32(sum) / float32(len(s.grades))
}
