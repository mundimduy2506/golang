package main

import (
	"fmt"
	"math"
	"strconv"
	"sort"
)

func add(a, b int, text string) string {
	return text + "values: " + fmt.Sprint(a+b)
}

///
func subtract(x, y int, c string) string {
	return c + " this is number " + string(x-y)
}

func swap(ori, des string) (string, string) {
	return des, ori
}

//sumary: return weekdate in english words
func returnWeekDay(date string) {
	switch date {
	case "2":
		fmt.Println("Mon")
		fallthrough
	case "3":
		fmt.Println("Tue")
	case "4":
		fmt.Println("Wed")
	case "5":
		fmt.Println("Thu")
	case "6":
		fmt.Println("Fri")
	case "7":
		fmt.Println("Sat")
	default:
		fmt.Println("Sun")
	}
}

func sortArray(oriArray []int, lastNum string) []int {
	newNum, _ := strconv.Atoi(lastNum)
	oriArray = append(oriArray, newNum)
	for i := 0; i < len(oriArray); i++ {
		for j := i + 1; j < len(oriArray); j++ {
			if oriArray[i] > oriArray[j] {
				var temp = oriArray[i]
				oriArray[i] = oriArray[j]
				oriArray[j] = temp
			}
		}
	}
	return oriArray
}
func mapMyMoney(myName string) {
	fmt.Printf("----Request: find my money holder as %s ------\n", myName)
	var money = map[string]int{
		"Duy": 126,
		"Nhi": 155,
	}
	if currentMoney, ok := money[myName]; ok {
	//if money[myName] != 0 {
		fmt.Printf("%s's current asset is %d \n", myName, currentMoney)
	} else {
		fmt.Println("Cannot find your money")
	}
}
func explorePointer(){
	a,b := 10, 100;
	fmt.Printf("a = %d", a)
	fmt.Println()
	fmt.Printf("b = %d", b)
	fmt.Println()
	var aPointer = &a
	fmt.Printf("before: aPointer = %d", *aPointer)
	fmt.Println()
	*aPointer = 15;
	fmt.Printf("after: aPointer = %d", *aPointer)
}
func main() {
	fmt.Printf("Now you have %g problems.", math.Sqrt(4)) //float
	fmt.Println()
	fmt.Println(math.Pi)
	fmt.Printf("2 + 3 = %s", add(2, 3, "result of"))
	fmt.Println()
	fmt.Println(swap("hello", "world"))
	var originArr = []int{10, 50, 20, 30}
	fmt.Printf("Manual sorting: %v \n",sortArray(originArr, "60"))//manual sorting
	sort.Ints(originArr)
	fmt.Printf("Package sort  : %v \n",originArr)//sort package
	returnWeekDay("5")
	mapMyMoney("Alex")
	mapMyMoney("Duy")
	//use pointer - similarly to C and C++
	explorePointer()

}
