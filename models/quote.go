package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type QuoteResponse struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Response string             `bson:"response"`
}

type Shipper struct {
	RegisteredNumber string `json:"registered_number"`
	Token            string `json:"token"`
	PlatformCode     string `json:"platform_code"`
}

type Address struct {
	ZipCode string `json:"zipcode,omitempty"`
}

type Recipient struct {
	Type             int     `json:"type"`
	RegisteredNumber string  `json:"registered_number,omitempty"`
	StateInscription string  `json:"state_inscription,omitempty"`
	Country          string  `json:"country"`
	Zipcode          uint32  `json:"zipcode"`
	Address          Address `json:"address,omitempty"`
}

type Volume struct {
	Amount        int         `json:"amount"`
	Price         int         `json:"price"`
	AmountVolumes int         `json:"amount_volumes,omitempty"`
	Category      interface{} `json:"category"`
	SKU           string      `json:"sku,omitempty"`
	Tag           string      `json:"tag,omitempty"`
	Description   string      `json:"description,omitempty"`
	Height        float64     `json:"height"`
	Width         float64     `json:"width"`
	Length        float64     `json:"length"`
	UnitaryPrice  float64     `json:"unitary_price"`
	UnitaryWeight float64     `json:"unitary_weight"`
	Consolidate   bool        `json:"consolidate,omitempty"`
	Overlaid      bool        `json:"overlaid,omitempty"`
	Rotate        bool        `json:"rotate,omitempty"`
}

type Dispatcher struct {
	RegisteredNumber string   `json:"registered_number"`
	Zipcode          uint32   `json:"zipcode"`
	TotalPrice       float64  `json:"total_price,omitempty"`
	Volumes          []Volume `json:"volumes"`
}

type Returns struct {
	Composition  bool `json:"composition,omitempty"`
	Volumes      bool `json:"volumes,omitempty"`
	AppliedRules bool `json:"applied_rules,omitempty"`
}

type QuoteApiRequest struct {
	Shipper        Shipper      `json:"shipper"`
	Recipient      Recipient    `json:"recipient"`
	Dispatchers    []Dispatcher `json:"dispatchers"`
	Channel        string       `json:"channel,omitempty"`
	Filter         int          `json:"filter,omitempty"`
	Limit          int          `json:"limit,omitempty"`
	Identification string       `json:"identification,omitempty"`
	Reverse        bool         `json:"reverse,omitempty"`
	SimulationType []int        `json:"simulation_type"`
	Returns        Returns      `json:"returns,omitempty"`
}

type CRequest struct {
	Recipient Recipient `json:"recipient"`
	Volumes   []Volume  `json:"volumes"`
}

type CarrierResponse struct {
	Carrier []CarrierDetail `json:"carrier"`
}

type CarrierDetail struct {
	Name     string  `json:"name"`
	Service  string  `json:"service"`
	Deadline int     `json:"deadline"`
	Price    float64 `json:"price"`
}
