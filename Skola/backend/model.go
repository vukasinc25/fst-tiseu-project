package main

type User struct {
	id        string
	firstName string
	lastName  string
	jmbg      string
}

type RequestBody struct {
	UserId string `json:"userId"`
}

type Diploma struct {
	AverageGrade    float32   `json:"averageGrade"`
	TotalHighPoints float32   `json:"totalHighPoints"`
	YearFinished    int       `json:"yearFinished"`
	Subjects        []Subject `json:"subjects"`
}

type Subject struct {
	Grades     string  `json:"grades"`
	FinalGrade float32 `json:"finalGrade"`
}
