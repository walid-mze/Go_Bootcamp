package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Person struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Salary    int    `json:"salary"`
	Education string `json:"education"`
}

type PeopleData struct {
	People []Person `json:"people"`
}

func main() {
	fileBytes, err := os.ReadFile("people.json")
	if err != nil {
		log.Fatal("failed to read file")
	}

	var data PeopleData
	err = json.Unmarshal(fileBytes, &data)
	if err != nil {
		log.Fatal("failed to unmarshal json")
	}


	// Calculate average age
	people:=data.People
	var avgAge float64
	for _, person:=range people {
		avgAge += float64(person.Age)
	}
	avgAge = avgAge / float64(len(people))
	fmt.Printf("Average Age: %f\n", avgAge)

	//find youngest age
	minAge := people[0].Age
	for _, person := range people {
		if person.Age < minAge {
			minAge = person.Age
		}
	}
	//collect all youngest people
	var youngestPeople []Person
	for _, person := range people {
		if person.Age==minAge {
			youngestPeople=append(youngestPeople, person)
		}
	}
	fmt.Println("Youngest People:")
	for _, p := range youngestPeople {
		fmt.Printf("- %s (Age: %d)\n", p.Name, p.Age)
	}

	//find oldest age
	maxAge:=people[0].Age
	for _, person:=range people {
		if person.Age>maxAge {
			maxAge=person.Age
		}
	}
	//collect all oldest people
	var oldestPeople []Person
	for _, person:=range people {
		if person.Age==maxAge {
			oldestPeople=append(oldestPeople, person)
		}
	}
	fmt.Println("Oldest People:")
	for _,p:=range oldestPeople {
		fmt.Printf("-%s (Age: %d)\n", p.Name, p.Age)
	}

	//avegare saalary
	var avgsalary float64
	for _, person := range people {
		avgsalary += float64(person.Salary)
	}
	avgsalary = avgsalary / float64(len(people))
	fmt.Printf("Average Salary: %f\n", avgsalary)

	//find highest salary
	maxSalary := people[0].Salary
	for _, person := range people {
		if person.Salary > maxSalary {
			maxSalary = person.Salary
		}
	}
	//collect all people with highest salary
	var highestSalaryPeople []Person
	for _, person := range people {
		if person.Salary == maxSalary {
			highestSalaryPeople = append(highestSalaryPeople, person)
		}
	}
	fmt.Println("People with Highest Salary:")
	for _, p := range highestSalaryPeople {
		fmt.Printf("-%s(Salary: %d)\n", p.Name, p.Salary)
	}

	//find lowest salary
	minSalary := people[0].Salary
	for _, person := range people {
		if person.Salary < minSalary {
			minSalary = person.Salary
		}
	}
	//collect all people with lowest salary
	var lowestSalaryPeople []Person
	for _, person:=range people {
		if person.Salary==minSalary {
			lowestSalaryPeople = append(lowestSalaryPeople, person)
		}
	}
	fmt.Println("People with Lowest Salary:")
	for _, p:=range lowestSalaryPeople {
		fmt.Printf("-%s(Salary: %d)\n", p.Name, p.Salary)
	}

	//counts people by education level
	var res=map[string]int{}
	for _, person:=range people {
		res[person.Education]++
	}
	for education, count := range res {
		fmt.Println(education, count)
	}

}
