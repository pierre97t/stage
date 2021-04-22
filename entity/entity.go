package entity

import "time"

// Entity is the interface use to do CRUD operations on application's entities
type Entity interface {
	New() Entity
	GenerateSTDParam(sessionID *string, checkRelConsist func(map[string]string) error) error
	UpdateSTDParam(sessionID *string, checkRelConsist func(map[string]string) error) error
	UpdateViewParam(sessionID *string, checkRelConsist func(map[string]string) error) error
	Init(sessionID *string) error
	Validate(sessionID *string, checkRelConsist func(map[string]string) error) error
	Interface() interface{}
	List() interface{}
	String() string
	// Getters
	GetID() *string
	GetName() *string
	GetType() *string
	GetIsDeleted() *bool
	GetCreatedDate() *time.Time
	GetUpdatedDate() *time.Time
	GetLastViewDate() *time.Time
	GetCreatedByID() *string
	GetUpdatedByID() *string
	GetLastViewByID() *string
	GetOwnerID() *string
	GetCAS() uint64
	GetExpiryTime() *time.Time
	GetExpiry() time.Duration
	// Setter
	SetID(value *string)
	SetIsDeleted(value *bool)
	SetCAS(value uint64)
	SetExpiryTime(value time.Time)
	SetExpiry(value *time.Duration)
	SetCreatedDate(*time.Time)
	SetUpdatedDate(*time.Time)
	SetLastViewDate(*time.Time)
	SetCreatedByID(*string)
	SetUpdatedByID(*string)
	SetLastViewByID(*string)
	SetType(*string)
}

// BulkEntity is used to process entities trough bulk operations
type BulkEntity struct {
	ID     string         `json:"id,omitempty"`
	Data   interface{}    `json:"data,omitempty"`
	CAS    uint64         `json:"cas,omitempty"`
	Expiry *time.Duration `json:"expiry,omitempty"`
	Err    error          `json:"error,omitempty"`
}

// SearchEntity contains search result for a specific entity
type SearchEntity struct {
	Err   *error      `json:"error,omitempty"`
	Score float64     `json:"score,omitempty"`
	Data  interface{} `json:"data,omitempty"`
}
