package models

import (
	"database/sql/driver"
	"encoding/json"
)

// EntryCounters

func (pc *EntryCounters) Scan(val interface{}) error {
	switch v := val.(type) {
	case []byte:
		json.Unmarshal(v, &pc)
		return nil
	case string:
		json.Unmarshal([]byte(v), &pc)
		return nil
	default:
		return nil
	}
}

func (pc EntryCounters) Value() (driver.Value, error) {
	return json.Marshal(pc)
}

// EntryAuthor

func (pc *EntryAuthor) Scan(val interface{}) error {
	switch v := val.(type) {
	case []byte:
		json.Unmarshal(v, &pc)
		return nil
	case string:
		json.Unmarshal([]byte(v), &pc)
		return nil
	default:
		return nil
	}
}

func (pc EntryAuthor) Value() (driver.Value, error) {
	return json.Marshal(pc)
}

// EntrySubsite

func (pc *EntrySubsite) Scan(val interface{}) error {
	switch v := val.(type) {
	case []byte:
		json.Unmarshal(v, &pc)
		return nil
	case string:
		json.Unmarshal([]byte(v), &pc)
		return nil
	default:
		return nil
	}
}

func (pc EntrySubsite) Value() (driver.Value, error) {
	return json.Marshal(pc)
}

// SubsiteSubscribers

func (pc *SubsiteSubscribers) Scan(val interface{}) error {
	switch v := val.(type) {
	case []byte:
		json.Unmarshal(v, &pc)
		return nil
	case string:
		json.Unmarshal([]byte(v), &pc)
		return nil
	default:
		return nil
	}
}

func (pc SubsiteSubscribers) Value() (driver.Value, error) {
	return json.Marshal(pc)
}

func (pc *SubsiteSubscriptions) Scan(val interface{}) error {
	switch v := val.(type) {
	case []byte:
		json.Unmarshal(v, &pc)
		return nil
	case string:
		json.Unmarshal([]byte(v), &pc)
		return nil
	default:
		return nil
	}
}

func (pc SubsiteSubscriptions) Value() (driver.Value, error) {
	return json.Marshal(pc)
}

func (pc *SubsiteRules) Scan(val interface{}) error {
	switch v := val.(type) {
	case []byte:
		json.Unmarshal(v, &pc)
		return nil
	case string:
		json.Unmarshal([]byte(v), &pc)
		return nil
	default:
		return nil
	}
}

func (pc SubsiteRules) Value() (driver.Value, error) {
	return json.Marshal(pc)
}

func (pc *SubscribersAvatar) Scan(val interface{}) error {
	switch v := val.(type) {
	case []byte:
		json.Unmarshal(v, &pc)
		return nil
	case string:
		json.Unmarshal([]byte(v), &pc)
		return nil
	default:
		return nil
	}
}

func (pc SubscribersAvatar) Value() (driver.Value, error) {
	return json.Marshal(pc)
}

// Bool allows 0/1 to also become boolean.
type JsonBool bool

func (bit *JsonBool) UnmarshalJSON(b []byte) error {
	txt := string(b)
	*bit = JsonBool(txt == "1" || txt == "true")
	return nil
}
