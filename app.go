package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Person struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Job       string `json:"job"`
	Gender    string `json:"gender"`
	City      string `json:"city"`
	Salary    int    `json:"salary"`
	Birthdate string `json:"datetime"`
}

func main() {
	jsonFile, err := os.Open("person.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened person.json")
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var persons []Person

	json.Unmarshal(byteValue, &persons)
	// fmt.Println(GroupPeopleByCity(persons))
	// fmt.Println(GroupPeopleByJob(persons))
	// fmt.Println(Top5JobsByNumer(persons))
	// fmt.Println(TopJobByNumerInEachCity(persons))
	// fmt.Println(AverageSalaryByJob(persons))
	// fmt.Println(SumPersonCity(persons))
	// fmt.Println(SalaryEachJob(persons))
	// fmt.Println(FiveCitiesHasTopAverageSalary(persons))
	fmt.Println(AverageAgePerJob(persons))
}

func GroupPeopleByCity(p []Person) (result map[string][]Person) {
	result = make(map[string][]Person)
	for _, person := range p {
		result[person.City] = append(result[person.City], person)
	}
	return result
}
func GroupPeopleByJob(p []Person) (result map[string]int) {
	result = make(map[string]int)
	for _, person := range p {
		result[person.Job]++
	}
	return result
}
func Top5JobsByNumer(p []Person) (result []string) {
	top5 := GroupPeopleByJob(p)
	type kv struct {
		Key   string
		Value int
	}

	var ss []kv
	for k, v := range top5 {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	for i := 0; i < 5; i++ {
		result = append(result, ss[i].Key)
	}
	return result
}
func TopJobByNumerInEachCity(p []Person) (result map[string]string) {
	PersonofCity := GroupPeopleByCity(p)
	result = make(map[string]string)
	for k, v := range PersonofCity {
		result[k] = Top5JobsByNumer(v)[0]
	}
	return result
}

func AverageSalaryByJob(p []Person) (result map[string]int) {
	result = make(map[string]int)
	sumSalary := SalaryEachJob(p)
	aPeople := GroupPeopleByJob(p)
	for k := range sumSalary {
		result[k] = sumSalary[k] / aPeople[k]
	}
	return result
}

// go
func FiveCitiesHasTopAverageSalary(p []Person) (result []string) {
	result1 := make(map[string]int)
	grPP := GroupPeopleByCity(p)
	for k, v := range grPP {
		result1[k] = SumSalary(v) / len(v)
	}
	type kv struct {
		Key   string
		Value int
	}
	var ss []kv
	for k, v := range result1 {
		ss = append(ss, kv{k, v})
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})
	for i := 0; i < 5; i++ {
		result = append(result, ss[i].Key)
	}
	return result
}
func FiveCitiesHasTopSalaryForDeveloper(p []Person) (result []string) {
	mDeveloper := make(map[string]int)
	ggRP := GroupPeopleByCity(p)
	for k, v := range ggRP {
		mDeveloper[k] = SumDeveloperSalary(v) / len(v)
	}
	type kv struct {
		Key   string
		Value int
	}
	var ss []kv
	for k, v := range mDeveloper {
		ss = append(ss, kv{k, v})
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})
	for i := 0; i < 5; i++ {
		result = append(result, ss[i].Key)
	}
	return result
}

func AverageAgePerCity() {

}
func SalaryEachJob(p []Person) (result map[string]int) {
	result = make(map[string]int)
	for _, person := range p {
		result[person.Job] += person.Salary
	}
	return result
}
func SumPersonCity(p []Person) (result map[string]int) {
	result = make(map[string]int)
	for _, person := range p {
		result[person.City]++
	}
	return result
}
func SumSalary(p []Person) int {
	sum := 0
	for _, person := range p {
		sum += person.Salary
	}
	return sum
}

func SumDeveloperSalary(p []Person) int {
	sum := 0
	job := "developer"
	for _, person := range p {
		if person.Job == job {
			sum += person.Salary
		}
	}
	return sum
}

func SumAgeOfJob(p []Person) (result map[string]int) {
	result = make(map[string]int)
	for _, person := range p{
		result[person.Job]+= CalculateAge(person.Birthdate)
	}
	return result
}
func CalculateAge(birthDate string) (result int) {
	now := time.Now()
	ny := now.Year()
	nm := int(now.Month())
	nd := now.Day()
	birthDateSplit := strings.Split(birthDate, "-")
	tmp := make([]int, 0)
	for _, value := range birthDateSplit {
		if x, err := strconv.Atoi(value); err == nil {
			tmp = append(tmp, x)
		} else {
			fmt.Println(err)
		}
	}
	var age int
	if tmp[1] > nm {
		age = ny - tmp[0]
	}
	if tmp[1] == nm {
		if tmp[2] >= nd {
			age = ny - tmp[0]
		} else {
			age = ny - tmp[0] - 1
		}
	} else if tmp[1] < nm {
		age = ny - tmp[0] - 1
	}
	return age
}

func AverageAgePerJob(p []Person) (result map[string]int){
	result = make(map[string]int)
	sum := SumAgeOfJob(p)
	number := SumPersonCity(p)
	for key := range number {
		result[key] = sum[key] / number[key]
	}
	return result
}
