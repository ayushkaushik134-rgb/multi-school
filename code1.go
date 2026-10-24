package main

import "fmt"

type School struct {
	ID   int
	Name string
}

type Class struct {
	ID       int
	Name     string
	SchoolID int
}

type Section struct {
	ID      int
	Name    string
	ClassID int
}
type Student1 struct {
	ID        int
	Name      string
	SectionID int
}
type Teacher struct {
	ID       int
	Name     string
	SchoolID int
}

// create
func Create[T any](items []T, item T) []T {
	return append(items, item)
}

// read
func Read[T any](items []T) []T {
	return items
}

// update
func Update[T any](items []T, index int, item T) []T {
	if index >= 0 && index < len(items) {
		items[index] = item
	} else {
		fmt.Println("invalid index")
	}
	return items
}
func Delete[T any](items []T, index int) []T {
	if index >= 0 && index < len(items) {
		return append(items[:index], items[index+1:]...)
	}
	fmt.Println("invalid index")
	return items

}
func Display[T any](title string, items []T) {
	fmt.Println("\n--", title, "--")
	for _, item := range items {
		fmt.Printf("%+v/n", item)
	}
}
func main() {
	schools := []School{
		{ID: 1, Name: "sv international"},
		{ID: 2, Name: "abc school"},
	}
	//create
	schools = Create(schools, School{ID: 3, Name: "xyz school"})
//read
	Display("schools -read", Read(schools))
//update
	schools = Update(schools, 0, School{ID: 1, Name: "svinternational schl"})
	Display("school -updated", schools)
//delete
	schools = Delete(schools, 1)
	Display("school - deleted", schools)

	classes := []Class{
		{ID: 1, Name: "10th Grade", SchoolID: 123},
		{ID: 2, Name: "11th Grade", SchoolID: 234},
	}
	//create
	classes = Create(classes, Class{ID: 3, Name: "12th grade", SchoolID: 333})
	
	//read
	Display("class read", Read(classes))
	//update
	classes = Update(classes, 0, Class{ID: 2, Name: "11th class", SchoolID: 343})
	Display("class update", classes)
//delete
	classes = Delete(classes, 1)
	Display("class deleted ", classes)
//section
	section := []Section{
		{ID: 1, Name: "Rose", ClassID: 12},
		{ID: 2, Name: "Lotus", ClassID: 13},
	}
	//create
	section = Create(section, Section{ID: 3, Name: "Lily", ClassID: 14})
	//read
	Display("section read", Read(section))
	//update
	section = Update(section, 0, Section{ID: 2, Name: "Lottus", ClassID: 13})
	Display("section update", section)
	//delete
	section = Delete(section, 1)
	Display("section deleted", section)
//student
	student1 := []Student1{
		{ID: 1, Name: "ayush", SectionID: 12},
		{ID: 3, Name: "sohan ", SectionID: 34},
	}
	//create
	student1 = Create(student1, Student1{ID: 2, Name: "rohan", SectionID: 23})
	//read
	Display("student read", Read(student1))
	//update
	student1 = Update(student1, 0, Student1{ID: 3, Name: "sohaan ", SectionID: 34})
	Display("student update", student1)
	//delete
	student1 = Delete(student1, 1)
	Display("student deleted ", student1)
// teacher
	teacher := []Teacher{
		{ID: 1, Name: "sonika singh", SchoolID: 23},
		{ID: 2, Name: "johny", SchoolID: 45},
	}
	//create
	teacher = Create(teacher, Teacher{ID: 6, Name: "den", SchoolID: 98})
	//read
	Display("teacher read", Read(teacher))
	//update
	teacher = Update(teacher, 0, Teacher{ID: 2, Name: "johnny", SchoolID: 45})
	Display("teacher update", teacher)
	//delete
	teacher = Delete(teacher, 1)
	Display("teacher deleted", teacher)
}
