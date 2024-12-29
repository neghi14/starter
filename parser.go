package starter

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"sync"

	"go.mongodb.org/mongo-driver/v2/bson"
)

var parser *sync.Once
var parser_instance *Parser

// Supported Tags
const (
	// tagDefault name of model variable or list of names
	tagDefault = "db"

	// tagAttributes name of model attributes or list of attributes
	tagAttributes = "attr"

	// mongoIDField is the _id field for mongodb
	mongoIDField = "_id"
)

type model struct {
	attributes []string
	fieldName  string
	fieldValue reflect.Value
	fieldTag   string
}

type E struct {
	Key   string
	Value interface{}
	Type  ParserValueType
}

type M []E

type Parser struct{}

func New() *Parser {
	parser.Do(func() {
		parser_instance = &Parser{}
	})

	return parser_instance
}

func (m *Parser) ParseToKeyValue(d interface{}) (M, error) {
	var res M
	data, err := m.parse(d)
	if err != nil {
		return nil, err
	}
	for _, dd := range data {

		var val interface{}
		switch dd.fieldValue.Kind() {
		case reflect.Int:
			val = int(dd.fieldValue.Int())
		case reflect.Int8:
			val = int8(dd.fieldValue.Int())
		case reflect.Int16:
			val = int16(dd.fieldValue.Int())
		case reflect.Int32:
			val = int32(dd.fieldValue.Int())
		case reflect.Int64:
			val = int64(dd.fieldValue.Int())
		case reflect.String:
			val = dd.fieldValue.String()
		case reflect.Float32:
			val = float32(dd.fieldValue.Float())
		case reflect.Float64:
			val = float64(dd.fieldValue.Float())
		default:
			return nil, errors.New("unsupported field type")
		}

		if !dd.fieldValue.IsZero() {
			res = append(res, E{Key: dd.fieldTag, Value: val})
		}

	}
	return res, nil
}
func (m *Parser) ParseKeyOnly(d interface{}) (M, error) {
	var res M
	data, err := m.parse(d)
	if err != nil {
		return nil, err
	}
	for _, dd := range data {
		var valType ParserValueType
		switch dd.fieldValue.Kind() {
		case reflect.Int:
			valType = Num
		case reflect.Float32:
			valType = Num
		case reflect.Float64:
			valType = Num
		case reflect.String:
			valType = Text
		default:
			return nil, errors.New("unsupported field type")
		}
		res = append(res, E{Key: dd.fieldTag, Type: valType})
	}
	return res, nil
}
func (m *Parser) ParseToStruct(obj interface{}, data M) error {

	mo, err := m.parse(obj)
	if err != nil {
		return err
	}
	for _, mod := range mo {
		if mod.fieldValue.IsValid() && mod.fieldValue.CanSet() {
			for _, re := range data {
				if re.Key == mod.fieldTag {
					mod.fieldValue.Set(reflect.ValueOf(re.Value))
				}
			}
		}
	}

	return nil
}

func (m *Parser) parse(d interface{}) ([]*model, error) {
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
		m.fieldName = fType.Name
		m.fieldValue = v.Field(i)

		if def, ok := fType.Tag.Lookup(tagDefault); ok {
			m.fieldTag = def
		}
		if def, ok := fType.Tag.Lookup(tagAttributes); ok {
			attr := strings.Split(def, ",")
			m.attributes = append(m.attributes, attr...)
		}
		res = append(res, m)
	}

	return res, nil
}

func (m *Parser) ConvertToBson(data M) (bson.D, error) {
	var res bson.D

	for _, d := range data {
		if d.Key == mongoIDField && d.Key != "" {
			id, err := bson.ObjectIDFromHex(d.Value.(string))
			if err != nil {
				return nil, err
			}
			d.Value = id
		}
		res = append(res, bson.E{Key: d.Key, Value: d.Value})
	}

	return res, nil
}

func (m *Parser) ConvertFromBson(data bson.D) (M, error) {
	var res M

	for _, d := range data {
		if d.Key == mongoIDField {
			id, ok := d.Value.(bson.ObjectID)
			if !ok {
				return nil, errors.New("error validating primitive ID")
			}
			d.Value = id.Hex()
		}
		res = append(res, E{Key: d.Key, Value: d.Value})
	}
	return res, nil
}
