package model

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"go.mongodb.org/mongo-driver/v2/bson"
)

// Supported Tags
const (
	// tagMongo name of mongo variable or list of names
	tagMongo = "mongo"

	// tagDefault name of model variable or list of names
	tagDefault = "model"

	// tagAttributes name of model attributes or list of attributes
	tagAttributes = "attr"
)

type model struct {
	attributes   []string
	fieldName    string
	fieldValue   reflect.Value
	mongoField   string
	defaultField string
}

type Model struct{}

func New() *Model {
	return &Model{}
}

func (m *Model) Extract(d interface{}) ([]byte, error) {
	_, err := m.parse(d)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (m *Model) ExtractForMongo(d interface{}) (bson.D, error) {
	var res bson.D
	data, err := m.parse(d)
	if err != nil {
		return nil, err
	}
	for _, dd := range data {
		var field string
		if dd.mongoField == "" {
			field = dd.defaultField
		} else {
			field = dd.mongoField
		}
		res = append(res, bson.E{Key: field, Value: dd.fieldValue.String()})
	}
	return res, nil
}

func (m *Model) parse(d interface{}) ([]*model, error) {
	var res []*model
	v := reflect.ValueOf(d)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		buf := make([]byte, 2)
		buf = fmt.Appendf(buf, "Invalid type recieved, expected type %s, got type %s", reflect.Struct, v.Kind())
		return nil, errors.New(string(buf))
	}

	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fType := t.Field(i)
		m := &model{}
		m.fieldName = v.Type().Field(i).Name
		m.fieldValue = v.Field(i)

		if def, ok := fType.Tag.Lookup(tagDefault); ok {
			m.defaultField = def
		}
		if def, ok := fType.Tag.Lookup(tagMongo); ok {
			m.mongoField = def
		}
		if def, ok := fType.Tag.Lookup(tagAttributes); ok {
			attr := strings.Split(def, ",")
			m.attributes = append(m.attributes, attr...)
		}
		res = append(res, m)
	}

	return res, nil
}
