package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Opportunities struct {
	gorm.Model
	Role 		string
	Company 	string
	Location 	string
	Remote 		bool
	Link 		string
	Salary 		int64
}


type OpportunitiesResponse struct {
	ID 			uint  		`json:"id"`
	CreatedAt 	time.Time 	`json:"createdAt"`
	UpdatedAt 	time.Time 	`json:"updatedAt"`
	DeletedAt 	time.Time 	`json:"deletedAt"` // passando o omitempty o campo deletedAt é omitido da estrutura do json quando vier null
	Role 		string 		`json:"role"`
	Company 	string 		`json:"company"`
	Location 	string	 	`json:"location"`
	Remote 		bool 		`json:"remote"`
	Link 		string	 	`json:"link"`
	Salary 		int64 		`json:"salary"`
}