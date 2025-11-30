package main

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
)

func main() {
	fmt.Printf("\n\nКалькулятор Эрланга\n\n")
	fmt.Println("По любым двум параметрам, найти остальные")
	
	fmt.Printf("+%s+\n", strings.Repeat("-", 55))
	fmt.Printf("| %10s | %40s |\n", "Параметр", "Описание")
	fmt.Printf("|%s|\n", strings.Repeat("-", 55))
	fmt.Printf("| %10s | %40s |\n", "a", "Интенсивность предложенного трафика")
	fmt.Printf("|%s|\n", strings.Repeat("-", 55))
	fmt.Printf("| %10s | %40s |\n", "v", "Число каналов")
	fmt.Printf("|%s|\n", strings.Repeat("-", 55))
	fmt.Printf("| %10s | %40s |\n", "π", "Доля потерянных заявок")
	fmt.Printf("|%s|\n", strings.Repeat("-", 55))
	fmt.Printf("| %10s | %40s |\n", "m", "Среднее число занятых каналов")
	fmt.Printf("+%s+\n", strings.Repeat("-", 55))

	fmt.Printf("\n\nВарианты\n")
	
	fmt.Printf("+%s+\n", strings.Repeat("-", 33))
	fmt.Printf("| %5s | %10s | %10s |\n", "№","Известны", "Найти")
	fmt.Printf("|%s|\n", strings.Repeat("-", 33))
	fmt.Printf("| %5s | %10s | %10s |\n", "1", "a, v", "π, m")
	fmt.Printf("|%s|\n", strings.Repeat("-", 33))
	fmt.Printf("| %5s | %10s | %10s |\n", "2", "a, π", "v, m")
	fmt.Printf("|%s|\n", strings.Repeat("-", 33))
	fmt.Printf("| %5s | %10s | %10s |\n", "3", "a, m", "v, π")
	fmt.Printf("|%s|\n", strings.Repeat("-", 33))
	fmt.Printf("| %5s | %10s | %10s |\n", "4", "v, m", "a, π")
	fmt.Printf("|%s|\n", strings.Repeat("-", 33))
	fmt.Printf("| %5s | %10s | %10s |\n", "5", "v, π", "a, m")
	fmt.Printf("|%s|\n", strings.Repeat("-", 33))
	fmt.Printf("| %5s | %10s | %10s |\n", "6", "π, m", "a, v")
	fmt.Printf("+%s+\n", strings.Repeat("-", 33))

	fmt.Print("Введите номер вариант: ")
	var varNum int
	if _ , err := fmt.Scan(&varNum); err != nil{
		slog.Error("невозможно сканировать переменную", slog.String("error", err.Error()))
		os.Exit(1)
	}
	
	fmt.Printf("\n")

	var a float64
	var v int
	var pi float64
	var m float64

	switch varNum{
	case 1:
		fmt.Print("Введите a: ")
		if _ , err := fmt.Scan(&a); err != nil{
			slog.Error("невозможно сканировать переменную", slog.String("error", err.Error()))
			os.Exit(1)
		}
		
		fmt.Print("Введите v: ")
		if _ , err := fmt.Scan(&v); err != nil{
			slog.Error("невозможно сканировать переменную", slog.String("error", err.Error()))
			os.Exit(1)
		}

		if !validateInput(map[string]float64{"a": a, "v": float64(v)}){
			os.Exit(1)
		}
		
		pi, m = CalculateCase1(a, v)

	case 2:
		fmt.Print("Введите a: ")
		if _ , err := fmt.Scan(&a); err != nil{
			slog.Error("невозможно сканировать переменную", slog.String("error", err.Error()))
			os.Exit(1)
		}
		
		fmt.Print("Введите π: ")
		if _ , err := fmt.Scan(&pi); err != nil{
			slog.Error("невозможно сканировать переменную", slog.String("error", err.Error()))
			os.Exit(1)
		}

		if !validateInput(map[string]float64{"a": a, "pi": pi}){
			os.Exit(1)
		}
		
		v, m = CalculateCase2(a, pi)

	case 3:
		fmt.Print("Введите a: ")
		if _ , err := fmt.Scan(&a); err != nil{
			slog.Error("невозможно сканировать переменную", slog.String("error", err.Error()))
			os.Exit(1)
		}
		
		fmt.Print("Введите m: ")
		if _ , err := fmt.Scan(&m); err != nil{
			slog.Error("невозможно сканировать переменную", slog.String("error", err.Error()))
			os.Exit(1)
		}

		if !validateInput(map[string]float64{"a": a, "m": m}){
			os.Exit(1)
		}
		
		v, pi = CalculateCase3(a, m)

	case 4:
		fmt.Print("Введите v: ")
		if _ , err := fmt.Scan(&v); err != nil{
			slog.Error("невозможно сканировать переменную", slog.String("error", err.Error()))
			os.Exit(1)
		}
		
		fmt.Print("Введите m: ")
		if _ , err := fmt.Scan(&m); err != nil{
			slog.Error("невозможно сканировать переменную", slog.String("error", err.Error()))
			os.Exit(1)
		}

		
		if !validateInput(map[string]float64{"v": float64(v), "m": m}){
			os.Exit(1)
		}
		
		a, pi =CalculateCase4(v, m)

	case 5:
		fmt.Print("Введите v: ")
		if _ , err := fmt.Scan(&v); err != nil{
			slog.Error("невозможно сканировать переменную", slog.String("error", err.Error()))
			os.Exit(1)
		}
		
		fmt.Print("Введите π: ")
		if _ , err := fmt.Scan(&pi); err != nil{
			slog.Error("невозможно сканировать переменную", slog.String("error", err.Error()))
			os.Exit(1)
		}

		if !validateInput(map[string]float64{"v": float64(v), "pi": pi}){
			os.Exit(1)
		}
		
		a, m = CalculateCase5(v, pi)

	case 6:
		fmt.Print("Введите π: ")
		if _ , err := fmt.Scan(&pi); err != nil{
			slog.Error("невозможно сканировать переменную", slog.String("error", err.Error()))
			os.Exit(1)
		}
		
		fmt.Print("Введите m: ")
		if _ , err := fmt.Scan(&m); err != nil{
			slog.Error("невозможно сканировать переменную", slog.String("error", err.Error()))
			os.Exit(1)
		}

		if !validateInput(map[string]float64{"pi": pi, "m": m}){
			os.Exit(1)
		}
		
		a, v = CalculateCase6(pi, m)

	default:
		slog.Error("вариант не в диапозоне значений (1-6)", slog.Int("предложенный вариант", varNum))
		os.Exit(1)
	}

	fmt.Printf("+%s+\n", strings.Repeat("-", 25))
	fmt.Printf("| %10s | %10s |\n", "Параметр", "Значение")
	fmt.Printf("|%s|\n", strings.Repeat("-", 25))
	fmt.Printf("| %10s | %10.3f |\n", "a", a)
	fmt.Printf("|%s|\n", strings.Repeat("-", 25))
	fmt.Printf("| %10s | %10d |\n", "v", v)
	fmt.Printf("|%s|\n", strings.Repeat("-", 25))
	fmt.Printf("| %10s | %10.3f |\n", "π", pi)
	fmt.Printf("|%s|\n", strings.Repeat("-", 25))
	fmt.Printf("| %10s | %10.3f |\n", "m", m)
	fmt.Printf("+%s+\n", strings.Repeat("-", 25))
}
