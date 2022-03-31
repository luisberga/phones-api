package models

type PhoneGroup struct {
	CompanyID       uint64 `json:"company_id,omitempty"`
	AvailablePhones string `json:"available_phones,omitempty"`
}
