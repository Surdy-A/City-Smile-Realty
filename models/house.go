package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type House struct {
	PropertyName string    `json:"property_name"`
	City         string    `json:"city"`
	Address      string    `json:"address"`
	Bedroom      int64     `json:"bedroom"`
	Bathroom     int64     `json:"bathroom"`
	Washroom     int64     `json:"washroom"`
	CarSpace     int64     `json:"car_space"`
	RoomArea     string    `json:"room_area"`
	Price        float64   `json:"price"`
	Availability string    `json:"availability"`
	Status       string    `json:"status"`
	Images       []string  `json:"images"`
	Description  string    `json:"description"`
	Agent        string    `json:"agent"`
	Features     Features  `json:"features"`
	Video        string    `json:"video"`
	Category     string    `json:"category"`
	CreatedDate  time.Time `json:"created_date"`
	UpdatedDate  time.Time `json:"uodated_date"`
	ID           string    `json:"id"`
	PropertyType string    `json:"property_type"`
	State        string    `json:"state"`
	PostalCode   string    `json:"postal_code"`
}

type Features struct {
	Features []string `json:"features" form:"features[]"`
}

func (h House) Value() (driver.Value, error) {
	return json.Marshal(h)
}

func (h *House) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &h)
}
