package responses

import "encoding/json"

type Organization struct {
	ID           uint
	Name         *string
	AccountLevel string
	Members      int
}

func (o Organization) String() string {
	ju, _ := json.MarshalIndent(o, "", " ")
	return string(ju)
}
