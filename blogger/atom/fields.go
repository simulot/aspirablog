package atom

import (
	"encoding/json"
)

type Field struct {
	Content string `json:"$t"`
	Type    string `json:"type"`
	Version string `json:"version"`
}

type ID string

func (f *ID) UnmarshalJSON(b []byte) error {
	s := &Field{}
	err := json.Unmarshal(b, s)
	if err != nil {
		return err
	}
	*f = ID(s.Content)
	return nil
}

type Author string

func (f *Author) UnmarshalJSON(b []byte) error {
	s := &struct {
		Name Field `json:name`
	}{}
	err := json.Unmarshal(b, s)
	if err != nil {
		return err
	}
	*f = Author(s.Name.Content)
	return nil
}
