package models

type Student struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

var Students = make(map[int]Student)

func init() {
	Students[1] = Student{Id: 1, Name: "Rohit sharma", Age: 20, Email: "rohit@example.com"}
	Students[2] = Student{Id: 2, Name: "Virat Kohli", Age: 22, Email: "virat@example.com"}
	Students[3] = Student{Id: 3, Name: "Prithviraj Awatade", Age: 19, Email: "prithviraj@example.com"}
}
