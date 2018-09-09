package atom

import "encoding/json"

type Category string

func (c *Category) UnmarshalJSON(b []byte) error {
	s := &struct {
		Term string `json:"term"`
	}{}
	err := json.Unmarshal(b, s)
	if err != nil {
		return err
	}
	*c = Category(s.Term)
	return nil
}
