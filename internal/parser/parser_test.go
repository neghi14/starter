package parser

import (
	"reflect"
	"testing"
)

func TestParser(T *testing.T) {
	type Model struct {
		ID       string `db:"_id"`
		Email    string `db:"email"`
		Name     string `db:"name"`
		Password string `db:"password"`
		Age      int    `db:"age"`
		Phone    int    `db:"phone"`
	}
	parser := New()
	data := Model{
		Email:    "jon@doe.com",
		Name:     "Jon Doe",
		Password: "Pass1234.",
		Age:      24,
	}
	T.Run("Test Parser", func(t *testing.T) {

		res := []*model{
			{
				attributes: nil,
				fieldName:  "ID",
				fieldValue: reflect.ValueOf(data.ID),
				fieldTag:   "_id",
			},
			{
				attributes: nil,
				fieldName:  "Email",
				fieldValue: reflect.ValueOf(data.Email),
				fieldTag:   "email",
			},
			{
				attributes: nil,
				fieldName:  "Name",
				fieldValue: reflect.ValueOf(data.Name),
				fieldTag:   "name",
			},
			{
				attributes: nil,
				fieldName:  "Password",
				fieldValue: reflect.ValueOf(data.Password),
				fieldTag:   "password",
			},
			{
				attributes: nil,
				fieldName:  "Age",
				fieldValue: reflect.ValueOf(data.Age),
				fieldTag:   "age",
			},
			{
				attributes: nil,
				fieldName:  "Phone",
				fieldValue: reflect.ValueOf(data.Phone),
				fieldTag:   "phone",
			},
		}

		mo, err := parser.parse(&data)
		if err != nil {
			t.Error(err.Error())
		}

		if !reflect.ValueOf(mo).IsNil() && reflect.ValueOf(mo) == reflect.ValueOf(res) {
			t.Errorf("Error completing tests, expected type")
		}
		t.Log(mo, res)
	})
}
