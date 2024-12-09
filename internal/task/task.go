package task

import "github.com/guregu/null"

type Task struct {
	ID			int			`json:"id"`
	Title 		string		`json:"task"`
	Description	null.String `json:"description"`
	Status		string		`json:"To do"`
	AddedDate 	string		`json:"added_date"`
	UpdatedDate null.String `json:"updated_date"`
	Tags		null.String	`json:"tags"`
}