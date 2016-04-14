package models


import(

	"time"
)

type TodoList struct {

		ID                     	int64        `json:"id"`
		UserID                 	int64        `json:"userId"`
		Item                   	string     	 `json:"name"`
		Description             string   	 `sql:"size:1024" json:"description"`
		StartDate              	time.Time	 `json:"startdate"`
		EndDate                	time.Time 	 `json:"enddate"`
		Enabled                 bool   		 `json:"enabled"`
		CreatedAt               time.Time 	 `json:"createdAt"`
		UpdatedAt               time.Time	 `json:"updatedAt"`
}

type Users struct {

		ID                     	int64    	 `json:"id"`
		Login                  	string   	 `sql:"size:1024" json:"login"` 
		Email                  	string   	 `sql:"size:1024" json:"email"` 
		Passwords              	string   	 `sql:"size:1024" json:"password"` 
		Fullname                string   	 `sql:"size:1024" json:"fullname"`
		Description             string   	 `sql:"size:1024" json:"description"`
		Enabled                 bool   		 `json:"enabled"`
		Verified                bool     	 `json:"verified"`
		Role					int64		 `json:"role"`
		CreatedAt               time.Time 	 `json:"createdAt" `
		UpdatedAt               time.Time    `json:"updatedAt" ` 
}






