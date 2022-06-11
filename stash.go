package structs

import (
	"encoding/json"
	"fmt"
	"github.com/nooize/go-assist"
	uuid "github.com/satori/go.uuid"
	"net/url"
	"strconv"
)

// FieldKey represent stash key for store and access
type FieldKey string

// Stash type for store and manipulate key/value data
type Stash map[FieldKey]interface{}

// Add store value in stash with access key
func (s Stash) Add(key FieldKey, v interface{}) {
	if s == nil {
		s = make(map[FieldKey]interface{})
	}
	s[key] = v
}

// Has return true if value requested key exists
func (s Stash) Has(key FieldKey) bool {
	if s == nil {
		return false
	}
	_, ok := s[key]
	return ok
}

// NewStashFromStruct create Stash from any structure
func NewStashFromStruct(i interface{}) Stash {
	out := make(map[FieldKey]interface{})
	for key, v := range assist.Struct2Map(i) {
		out[FieldKey(key)] = v
	}
	return out
}

// Get read value from stash
// nil return if key not found
func (s Stash) Get(key FieldKey) interface{} {
	if s == nil {
		return nil
	}
	if i, ok := s[key]; ok {
		return i
	}
	return nil
}

// Delete remove item from stash by key
func (s Stash) Delete(key FieldKey) {
	if s == nil {
		return
	}
	delete(s, key)
}

// GetBool read value as bool from stash by key
// false return if key not found
func (s Stash) GetBool(key FieldKey, def bool) bool {
	v := s.Get(key)
	if v == nil {
		return def
	}
	switch v.(type) {
	case float64:
		return int(v.(float64)) == 1
	case int64:
		return v.(int) == 1
	case string:
		return v.(string) == "true"
	case bool:
		return v.(bool)
	default:
		return def
	}
}

// GetInt read value as int from stash by key
// def parameter return if key not found
func (s Stash) GetInt(key FieldKey, def int) int {
	v := s.Get(key)
	if v == nil {
		return def
	}
	switch v.(type) {
	case float64:
		return int(v.(float64))
	case int:
		return v.(int)
	case int64:
		return v.(int)
	case string:
		iv, err := strconv.ParseInt(v.(string), 10, 64)
		if err != nil {
			return def
		}
		return int(iv)
	default:
		return def
	}
}

// GetStr read value as string from stash by key
// empty string return if key not found
func (s Stash) GetStr(key FieldKey) string {
	if v := s.Get(key); v != nil {
		return fmt.Sprintf("%v", v)
	}
	return ""
}

// GetStruct read value as string from stash by key
// empty string return if key not found
func (s Stash) GetStruct(key FieldKey, i interface{}) (bool, error) {
	v := s.Get(key)
	if v == nil {
		return false, nil
	}
	switch v.(type) {
	case map[string]interface{}:
		return true, assist.Map2Struct(v.(map[string]interface{}), i)
	case string:
		return true, json.Unmarshal([]byte(v.(string)), i)
	}
	return false, nil
}

// GetUrl read value as *url.URL from stash by key
// nil return if key not found
func (s Stash) GetUrl(key FieldKey) *url.URL {
	if v := s.Get(key); v != nil {
		if u, err := url.ParseRequestURI(v.(string)); err == nil {
			return u
		}
		// TODO log not critical error to sameware
	}
	return nil
}

// GetUuid read value as UUID from stash by key
// 00000000-0000-0000-0000-000000000000 return if key not found
func (s Stash) GetUuid(key FieldKey) uuid.UUID {
	return uuid.FromStringOrNil(s.GetStr(key))
}

// MarshalJSON implements the json.Marshaler interface.
func (s Stash) MarshalJSON() ([]byte, error) {
	filtered := make(map[FieldKey]interface{})
	for key, v := range s {
		if len(key) == 0 || key[0] == '_' {
			continue
		}
		filtered[key] = v
	}
	return json.Marshal(filtered)
}

// IStashable interface for represent and operate Stashable structs
type IStashable interface {
	StashAdd(FieldKey, interface{})
	StashGet(FieldKey) (interface{}, bool)
	StashGetUrl(FieldKey) *url.URL
	StashGetStruct(FieldKey, interface{}) (bool, error)
	StashGetUuid(FieldKey) uuid.UUID
	StashGetStr(FieldKey) string
	StashToStruct(interface{}) error
	StashDelete(FieldKey)
}

// Stashable struct for integrate in to other struct
// 	type MyBox struct {
// 		Stashable
//		MyFields int
// 	}
type Stashable struct {
	Stash Stash `json:"Stash,omitempty" pg:"stash,type:jsonb"`
}

// StashAdd store value in stash with access key
func (s *Stashable) StashAdd(key FieldKey, v interface{}) {
	if s.Stash == nil {
		s.Stash = Stash{}
	}
	s.Stash.Add(key, v)
}

// StashHas return true if value requested key exists
func (s *Stashable) StashHas(key FieldKey) bool {
	return s.Stash.Has(key)
}

// StashGet read value from stash
// nil return if key not found
func (s *Stashable) StashGet(key FieldKey) interface{} {
	return s.Stash.Get(key)
}

// StashDelete remove item from stash by key
func (s *Stashable) StashDelete(key FieldKey) {
	s.Stash.Delete(key)
}

// StashGetBool read value as bool from stash by key
// false return if key not found
func (s *Stashable) StashGetBool(key FieldKey, def bool) bool {
	return s.Stash.GetBool(key, def)
}

// StashGetInt read value as int from stash by key
// def parameter return if key not found
func (s *Stashable) StashGetInt(key FieldKey, def int) int {
	return s.Stash.GetInt(key, def)
}

// StashGetStr read value as string from stash by key
// empty string return if key not found
func (s *Stashable) StashGetStr(key FieldKey) string {
	return s.Stash.GetStr(key)
}

// StashGetStruct read value as string from stash by key
// empty string return if key not found
func (s *Stashable) StashGetStruct(key FieldKey, i interface{}) (bool, error) {
	return s.Stash.GetStruct(key, i)
}

// StashGetUrl read value as *url.URL from stash by key
// nil return if key not found
func (s *Stashable) StashGetUrl(key FieldKey) *url.URL {
	return s.Stash.GetUrl(key)
}

// StashGetUuid read value as UUID from stash by key
// 00000000-0000-0000-0000-000000000000 return if key not found
func (s *Stashable) StashGetUuid(key FieldKey) uuid.UUID {
	return s.Stash.GetUuid(key)
}
