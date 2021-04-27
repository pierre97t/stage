package model

import (
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"
)

const (
	createdDateField uint8 = iota
	updatedDateField
	lastViewDateField
	createdByIDField
	updatedByIDField
	lastViewByIDField
	ownerIDField
)

const (
	// TypeUser coressponds to the type used to identify a record in database
	TypeUser     = "user"
	TypeCustomer = "customer"
	TypeCompany  = "company"
	TypeContact  = "contact"
	TypeLead     = "lead"
	TypeProspect = "prospect"
	TypeAccount  = "account"
)

// Std contains standard properties
type Std struct {
	ID         *string        `json:"id,omitempty" validate:"uuid4"`
	Name       *string        `json:"name,omitempty"`
	Type       *string        `json:"type,omitempty" validate:"alpha"`
	IsDeleted  *bool          `json:"isDeleted,omitempty"`
	Created    *minStd        `json:"created,omitempty"`
	Updated    *minStd        `json:"updated,omitempty"`
	LastView   *minStd        `json:"lastview,omitempty"`
	Owner      *minStd        `json:"owner,omitempty"`
	CAS        *uint64        `json:"cas,omitempty"`
	ExpiryTime *time.Time     `json:"expiryTime,omitempty"`
	Expiry     *time.Duration `json:"expiry,omitempty"`
}

type minStd struct {
	ID   *string    `json:"id,omitempty" validate:"uuid4"`
	Name *string    `json:"name,omitempty"`
	Type *string    `json:"type,omitempty" validate:"alpha"`
	Date *time.Time `json:"date,omitempty"`
}

// GetID returns minStd's ID as string
func (m *minStd) GetID() string {
	if m.ID != nil {
		return *m.ID
	}
	return ""
}

// GetName returns minStd's name as string
func (m *minStd) GetName() string {
	if m.Name != nil {
		return *m.Name
	}
	return ""
}

// GetType returns minStd's type as string
func (m *minStd) GetType() string {
	if m.Type != nil {
		return *m.Type
	}
	return ""
}

// GetType returns minStd's type as string
func (m *minStd) GetDate() time.Time {
	if m.Date != nil {
		return *m.Date
	}
	return time.Time{}
}

// GenerateID generates a random UUID V4
func GenerateID() string {
	return uuid.NewV4().String()
}

// GenerateSTDParam sets record's created date and created by
func (s *Std) GenerateSTDParam(sessionID *string, checkRelConsist func(map[string]string) error) error {
	// Generats and sets entity's ID
	id := GenerateID()
	s.ID = &id

	// Replace nil dates with current one
	s.setNilDates([]uint8{createdDateField})
	// IDs check if setlled IDs exist in database and replace null ones with provided sessionID parameter
	return s.checkIDsExistAndSetNilOnes([]uint8{createdByIDField, ownerIDField}, sessionID, checkRelConsist)

}

// UpdateSTDParam sets record's updated date and updated by
func (s *Std) UpdateSTDParam(sessionID *string, checkRelConsist func(map[string]string) error) error {
	// Dates
	s.setNilDates([]uint8{updatedDateField})
	// Updated By ID
	return s.checkIDsExistAndSetNilOnes([]uint8{updatedByIDField}, sessionID, checkRelConsist)
}

// UpdateViewParam sets record's last view date and last view by
func (s *Std) UpdateViewParam(sessionID *string, checkRelConsist func(map[string]string) error) error {
	// Dates
	t := time.Now()
	if s.Updated == nil {
		s.Updated = &minStd{}
	}
	s.Updated.Date = &t
	// Last View By ID
	return s.checkIDsExistAndSetNilOnes([]uint8{lastViewByIDField}, sessionID, checkRelConsist)
}

// GetID returns database's record ID
func (s *Std) GetID() *string {
	return s.ID
}

// GetType returns record's type
func (s *Std) GetType() *string {
	return s.Type
}

// GetName returns record's name
func (s *Std) GetName() *string {
	return s.Name
}

// GetIsDeleted returns record's IsDeleted
func (s *Std) GetIsDeleted() *bool {
	return s.IsDeleted
}

// GetCreatedDate returns record's CreatedDate
func (s *Std) GetCreatedDate() *time.Time {
	if s.Created != nil {
		return s.Created.Date
	}
	return nil
}

// GetUpdatedDate returns record's UpdatedDate
func (s *Std) GetUpdatedDate() *time.Time {
	if s.Updated != nil {
		return s.Updated.Date
	}
	return nil
}

// GetLastViewDate returns record's LastViewDate
func (s *Std) GetLastViewDate() *time.Time {
	if s.LastView != nil {
		return s.LastView.Date
	}
	return nil
}

// GetCreatedByID returns record's CreatedByID
func (s *Std) GetCreatedByID() *string {
	if s.Created != nil {
		return s.Created.ID
	}
	return nil
}

// GetUpdatedByID returns record's UpdatedByID
func (s *Std) GetUpdatedByID() *string {
	if s.Updated != nil {
		return s.Updated.ID
	}
	return nil
}

// GetLastViewByID returns record's LastViewByID
func (s *Std) GetLastViewByID() *string {
	if s.LastView != nil {
		return s.LastView.ID
	}
	return nil
}

// GetOwnerID returns record's OwnerID
func (s *Std) GetOwnerID() *string {
	if s.Owner != nil {
		return s.Owner.ID
	}
	return nil
}

// GetCAS returns record's CAS
func (s *Std) GetCAS() uint64 {
	if s.CAS == nil {
		return 0
	}
	return *s.CAS
}

// GetExpiryTime returns record's expiry date
func (s *Std) GetExpiryTime() *time.Time {
	return s.ExpiryTime
}

// GetExpiry returns record's expiry duration
func (s *Std) GetExpiry() time.Duration {
	if s.Expiry == nil {
		return 0
	}
	return *s.Expiry
}

// SetID sets record's ID
func (s *Std) SetID(value *string) {
	s.ID = value
}

// SetIsDeleted sets IsDeleted entitie's value
func (s *Std) SetIsDeleted(value *bool) {
	s.IsDeleted = value
}

// SetCAS sets record's CAS
func (s *Std) SetCAS(value uint64) {
	if value > 0 {
		s.CAS = &value
	} else {
		s.CAS = nil
	}
}

// SetExpiryTime sets record's expiry date
func (s *Std) SetExpiryTime(value time.Time) {
	if !value.IsZero() {
		s.ExpiryTime = &value
	}
}

// SetExpiry sets record's expiry duration
func (s *Std) SetExpiry(value *time.Duration) {
	if value != nil && *value > 0 {
		s.Expiry = value
	}
}

// SetCreatedDate sets record's CreatedDate
func (s *Std) SetCreatedDate(value *time.Time) {
	if s.Created == nil {
		s.Created = &minStd{}
	}
	s.Created.Date = value
}

// SetUpdatedDate sets record's UpdatedDate
func (s *Std) SetUpdatedDate(value *time.Time) {
	if s.Updated == nil {
		s.Updated = &minStd{}
	}
	s.Updated.Date = value
}

// SetLastViewDate sets record's LastViewDate
func (s *Std) SetLastViewDate(value *time.Time) {
	if s.LastView == nil {
		s.LastView = &minStd{}
	}
	s.LastView.Date = value
}

// SetCreatedByID sets record's CreatedByID
func (s *Std) SetCreatedByID(value *string) {
	if s.Created == nil {
		s.Created = &minStd{}
	}
	s.Created.ID = value
}

// SetUpdatedByID sets record's UpdatedByID
func (s *Std) SetUpdatedByID(value *string) {
	if s.Updated == nil {
		s.Updated = &minStd{}
	}
	s.Updated.ID = value
}

// SetLastViewByID sets record's LastViewByID
func (s *Std) SetLastViewByID(value *string) {
	if s.LastView == nil {
		s.LastView = &minStd{}
	}
	s.LastView.ID = value
}

// SetType sets record's type
func (s *Std) SetType(value *string) {
	s.Type = value
}

// SetNilDates will set provided date ptr if they are nill
func (s *Std) setNilDates(dates []uint8) {
	t := time.Now()
	for _, d := range dates {
		switch d {
		case createdDateField:
			if s.Created == nil {
				s.Created = &minStd{}
				s.Created.Date = &t
			} else if s.Created.Date == nil {
				s.Created.Date = &t
			} else if s.Created.Date.IsZero() {
				s.Created.Date = nil
			}
		case updatedDateField:
			if s.Updated == nil {
				s.Updated = &minStd{}
				s.Updated.Date = &t
			} else if s.Updated.Date == nil {
				s.Updated.Date = &t
			} else if s.Updated.Date.IsZero() {
				s.Updated.Date = nil
			}
		case lastViewDateField:
			if s.LastView == nil {
				s.LastView = &minStd{}
				s.LastView.Date = &t
			} else if s.LastView.Date == nil {
				s.LastView.Date = &t
			} else if s.LastView.Date.IsZero() {
				s.LastView.Date = nil
			}
		}
	}
}

// checkIDsExistAndSetNilOnes if provided IDs aren't nill their existance will be checked else they will be settled regarding default ID
func (s *Std) checkIDsExistAndSetNilOnes(ids []uint8, defaultID *string, checkRelConsist func(map[string]string) error) error {
	uniqIDs := make(map[string]string)
	for _, id := range ids {
		switch id {
		case createdByIDField:
			if s.Created == nil {
				s.Created = &minStd{}
				s.Created.ID = defaultID
			} else if s.Created.ID == nil {
				s.Created.ID = defaultID
			} else if _, valid := uniqIDs[*s.Created.ID]; *s.Created.ID != "" && !valid {
				uniqIDs[*s.Created.ID] = TypeUser
			}
		case updatedByIDField:
			if s.Updated == nil {
				s.Updated = &minStd{}
				s.Updated.ID = defaultID
			} else if s.Updated.ID == nil {
				s.Updated.ID = defaultID
			} else if _, valid := uniqIDs[*s.Updated.ID]; *s.Updated.ID != "" && !valid {
				uniqIDs[*s.Updated.ID] = TypeUser
			}
		case lastViewByIDField:
			if s.LastView == nil {
				s.LastView = &minStd{}
				s.LastView.ID = defaultID
			} else if s.LastView.ID == nil {
				s.LastView.ID = defaultID
			} else if _, valid := uniqIDs[*s.LastView.ID]; *s.LastView.ID != "" && !valid {
				uniqIDs[*s.LastView.ID] = TypeUser
			}
		case ownerIDField:
			if s.Owner == nil {
				s.Owner = &minStd{}
				s.Owner.ID = defaultID
			} else if s.Owner.ID == nil {
				s.Owner.ID = defaultID
			} else if _, valid := uniqIDs[*s.Owner.ID]; *s.Owner.ID != "" && !valid {
				uniqIDs[*s.Owner.ID] = TypeUser
			}
		}
	}

	if len(uniqIDs) > 0 {
		if err := checkRelConsist(uniqIDs); err != nil {
			return err
		}
	}

	return nil
}

// removeNilAndDuplicatedStrPtr remove nil and duplicates from a slice of *string
/*func removeNilAndDuplicatedStrPtr(arr []*string) map[string]string {
	occured := map[string]bool{}
	result := make(map[string]string)
	for e := range arr {

		// check if already the mapped
		// variable is set to true or not
		if arr[e] != nil && occured[*arr[e]] != true {
			occured[*arr[e]] = true

			// Append to result slice.
			result[*arr[e]] = TypeUser
		}
	}

	return result
}*/

func handleJSONErr(tpe string, data []byte, err error) error {
	if err != nil {
		return errors.New("Unable to parse request body into " + tpe + " structure")
	}
	return err
}
