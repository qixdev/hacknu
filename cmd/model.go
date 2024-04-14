package main

type Student struct {
	Name      string             `json:"name"`
	Goal      string             `json:"goal"`
	Materials map[string]Subject `json:"materials"`
}

type Teacher struct {
	Students []Student `json:"students"`
}

type Parent struct {
	Child Student  `json:"child"`
	Goals []string `json:"parents_goals"`
}

type Subject struct {
	Grade  float64          `json:"grade"`
	Topics map[string]Topic `json:"topics"`
}

type Topic struct {
	NumberOfQuestions int     `json:"numberOfQuestions"`
	RightAnswer       int     `json:"rightAnswer"`
	AverageGrade      float64 `json:"averageGrade"`
}
type Prompt struct {
	Answer string   `json:"prompt"`
	Files  []string `json:"files"`
}

type Python struct {
	Files   []string `json:"files"`
	Message string   `json:"message"`
}
type Data struct {
	Materials map[string]Subject `json:"materials"`
}
type Prediction struct {
	Subjects map[string]float64 `json:"subjects"`
}
type Exam struct {
	Results        []int `json:"results"`
	Attempt        int   `json:"attempt"`
	QuestionNumber int   `json:"numberOfQuestions"`
	MaxRes         int   `json:"max"`
}
