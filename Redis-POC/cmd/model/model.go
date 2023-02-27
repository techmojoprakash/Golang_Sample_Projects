package model

type Student struct {
	Id         string  `json:"id"`
	Name       string  `json:"name"`
	Gpa        float64 `json:"gpa"`
	IsEligible bool    `json:"isEligible"`
}
