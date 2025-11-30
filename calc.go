package main

import (
	"fmt"
	"math"
)

func ErlangB(r int, a float64) float64 {
	if r == 0 {
		return 1.0
	}

	invB := 1.0
	for k := 1; k <= r; k++ {
		invB = 1.0 + (float64(k)/a)*invB
	}
	return 1.0 / invB
}

func FindChannelsByLoss(a float64, pi float64) (int, float64) {
	v := 0
	for {
		piCalc := ErlangB(v, a)
		if piCalc <= pi {
			m := a * (1 - piCalc)
			return v, m
		}
		v++

		if v > 1000000 {
			return -1, 0
		}
	}
}



// FindChannelsByLoadBusy находит количество каналов v для заданной нагрузки a и среднего числа занятых каналов m
func FindChannelsByLoadBusy(a float64, m float64) (int, float64) {
	targetPi := 1 - m/a

	v := 0
	for {
		piCalc := ErlangB(v, a)
		if piCalc <= targetPi {
			return v, piCalc
		}
		v++

		if v > 1000000 {
			return -1, 0
		}
	}
}

func FindLoadByChannelsLoss(v int, pi float64) (float64, float64) {
	low := 0.0
	high := float64(v) * 10.0
	
	// Если вероятность потерь очень мала, увеличиваем верхнюю границу
	if pi < 0.001 {
		high = float64(v) * 100.0
	}
	
	for i := 0; i < 100; i++ {
		mid := (low + high) / 2.0
		piCalc := ErlangB(v, mid)
		
		if math.Abs(piCalc-pi) < 1e-10 {
			m := mid * (1 - pi)
			return mid, m
		}
		
		if piCalc > pi {
			high = mid
		} else {
			low = mid
		}
	}
	
	a := (low + high) / 2.0
	m := a * (1 - pi)
	return a, m
}


// FindLoadByChannelsBusy находит нагрузку a для заданного количества каналов v и среднего числа занятых каналов m
func FindLoadByChannelsBusy(v int, m float64) (float64, float64) {

	low := m
	high := m * 10.0

	for i := 0; i < 100; i++ {
		mid := (low + high) / 2.0
		pi := ErlangB(v, mid)
		mCalc := mid * (1 - pi)

		if math.Abs(mCalc-m) < 1e-10 {
			return mid, pi
		}

		if mCalc < m {
			low = mid
		} else {
			high = mid
		}
	}

	a := (low + high) / 2.0
	pi := ErlangB(v, a)
	return a, pi
}

// FindLoadByLossBusy находит нагрузку a и количество каналов v для заданной вероятности потерь π и среднего числа занятых каналов m
func FindLoadByLossBusy(pi float64, m float64) (float64, int) {
	a := m / (1 - pi)

	v := 0
	for {
		piCalc := ErlangB(v, a)
		if piCalc <= pi {
			return a, v
		}
		v++

		if v > 1000000 {
			return -1, -1
		}
	}
}

// CalculateCase1: Даны a, v -> найти π, m
func CalculateCase1(a float64, v int) (float64, float64) {
	pi := ErlangB(v, a)
	m := a * (1 - pi)
	return pi, m
}

// CalculateCase2: Даны a, π -> найти v, m
func CalculateCase2(a float64, pi float64) (int, float64) {
	return FindChannelsByLoss(a, pi)
}

// CalculateCase3: Даны a, m -> найти v, π
func CalculateCase3(a float64, m float64) (int, float64) {
	return FindChannelsByLoadBusy(a, m)
}

// CalculateCase4: Даны v, m -> найти a, π
func CalculateCase4(v int, m float64) (float64, float64) {
	return FindLoadByChannelsBusy(v, m)
}

// CalculateCase5: Даны v, π -> найти a, m
func CalculateCase5(v int, pi float64) (float64, float64) {
	return FindLoadByChannelsLoss(v, pi)
}

// CalculateCase6: Даны π, m -> найти a, v
func CalculateCase6(pi float64, m float64) (float64, int) {
	return FindLoadByLossBusy(pi, m)
}

func validateInput(params map[string]float64) bool {
	// Проверяем, что все значения неотрицательны
	for key, value := range params {
		if value < 0 {
			fmt.Printf("Ошибка: параметр %s не может быть отрицательным\n", key)
			return false
		}
	}
	
	// Дополнительные проверки для конкретных параметров
	if pi, ok := params["pi"]; ok {
		if pi < 0.0001 || pi > 1 {
			fmt.Printf("Ошибка: вероятность потерь π должна быть в диапазоне [0, 1]\n")
			return false
		}
	}
	
	if a, ok := params["a"]; ok {
		if a < 0.001 || a > 1000000{
			fmt.Printf("Ошибка: a должна лежать в диапозоне [0.001; 1000000]\n")
			return false
		}
	}
	
	if m, ok := params["m"]; ok {
		if m < 0.0001 || m > 999900 {
			fmt.Printf("Ошибка: m должна лежать в диапозоне [0.0001; 999900] \n")
			return false
		}
	}
	
	if v, ok := params["v"]; ok {
		if v < 1 || v > 100000 || math.Floor(v) != v {
			fmt.Printf("Ошибка: v должно быть неотрицательным числом в диапозоне [1; 100000]\n")
			return false
		}
	}
	
	return true
}
