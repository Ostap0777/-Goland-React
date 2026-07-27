package makes

import (
	"errors"
	"strings"
	"time"
)

type Make struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	CreatedAt time.Time `json:"createAt"`
}

type Input struct {
		Name string `json:"name"`
}

func (in *Input) Validate() error {
in.Name = strings.TrimSpace(in.Name)
if in.Name == "" {
	return errors.New("name is required")
}
return nil
}