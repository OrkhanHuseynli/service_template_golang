package models

type SimpleRequest struct {
	Product string `json:"product"`
}

type SimpleResponse struct {
	Message string `json:"message"`
	//don't output this field
	Author string `json:"-"`
	// don't output this field if the  value is empty
	Date string `json:"date, omitempty"`
	// convert output to a string and rename "id"
	Id int `json:"id,string"`
}

type SimpleDataItem struct {
	Name string `json:"name"`
}