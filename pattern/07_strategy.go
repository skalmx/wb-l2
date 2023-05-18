package pattern

import (
	"fmt"
)

type SortAlg interface {
	Sort(info []string) []string
}
type Data struct {
	slice []string
	sort  SortAlg
}

func (d *Data) SetSortAlg(s SortAlg) {
	d.sort = s
}

type QuickSort struct {
}

func (s *QuickSort) Sort(info []string) []string {
	fmt.Println("QuickSort")
	// Имитация алгоритма сортировки - вернем тоже самое что получили для краткости
	return info
}

type BubbleSort struct{
}

func (s *BubbleSort) Sort(info []string) []string {
	fmt.Println("BubbleSort")
	// Имитация алгоритма сортировки - вернем тоже самое что получили для краткости
	return info
}

type InsertrionSort struct {
}

func (s *InsertrionSort) Sort(info []string) []string {
	fmt.Println("InsertrionSort")
	// Имитация алгоритма сортировки - вернем тоже самое что получили для краткости
	return info
}

 