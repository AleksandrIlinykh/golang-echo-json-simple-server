package models

import "encoding/json"

type Magazine struct {
	Id      int     `json:"id"`
	Title   string  `json:"title"`
	Company string  `json:"company"`
	Price   float64 `json:"price"`
	Month   int     `json:"month"`
	Year    int     `json:"year"`
}

type Document struct {
	Id      int             `json:"id"`
	Name    string          `json:"name"`
	Version string          `json:"version"`
	DocJSON json.RawMessage `json:"doc_json"`
	// CreatedAt time.Time       `json:"created_at"`		//Неиспользуемые поля
	// CreatedBy uuid.UUID       `json:"created_by"`
	//UpdatedAt time.Time       `json:"updated_at"`
	//UpdatedBy uuid.UUID       `json:"updated_by"`
}
