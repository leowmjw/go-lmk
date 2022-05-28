// This file is generated by github.com/mikekonan/go-oas3. DO NOT EDIT.

package rippling

import (
	"encoding/json"
	"fmt"
)

type error struct {
	Code    *int    `json:"code"`
	Message *string `json:"message"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (body *Error) UnmarshalJSON(data []byte) error {
	var value error
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	if value.Message == nil {
		return fmt.Errorf("message is required")
	}

	body.Message = *value.Message

	if value.Code == nil {
		return fmt.Errorf("code is required")
	}

	body.Code = *value.Code

	return nil
}
func (body Error) Validate() error {
	return nil
}

type pet struct {
	ID   *int    `json:"id"`
	Name *string `json:"name"`
	Tag  string  `json:"tag"`
}

type Pet struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Tag  string `json:"tag"`
}

func (body *Pet) UnmarshalJSON(data []byte) error {
	var value pet
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	body.Tag = value.Tag

	if value.ID == nil {
		return fmt.Errorf("id is required")
	}

	body.ID = *value.ID

	if value.Name == nil {
		return fmt.Errorf("name is required")
	}

	body.Name = *value.Name

	return nil
}
func (body Pet) Validate() error {
	return nil
}

type Pets []Pet
