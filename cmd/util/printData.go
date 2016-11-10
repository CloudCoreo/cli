package util

import (
	"reflect"

	"encoding/json"
	"fmt"

	"github.com/bndr/gotabulate"
)

// Table struct
type Table struct {
	Header      []string
	Rows        [][]interface{}
	MaxCellSize int
}

// NewTable create new table
func NewTable() *Table {
	return new(Table)
}

// SetHeader set header for table
func (c *Table) SetHeader(header []string) *Table {
	c.Header = header
	return c
}

// SetMaxCellSize set max cell size
func (c *Table) SetMaxCellSize(maxCellSize int) *Table {
	c.MaxCellSize = maxCellSize
	return c
}

// UseObj Serialize obj to table
func (c *Table) UseObj(obj interface{}) *Table {

	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
		t = t.Elem()
	}
	kind := v.Kind()
	switch kind {
	case reflect.Slice:
		for _, objv := range obj.([]interface{}) {
			c.UseObj(objv)
		}
	case reflect.Map:
		var newrow []interface{}
		if len(c.Header) != 0 {
			for _, h := range c.Header {
				newrow = append(newrow, obj.(map[string]interface{})[h])
			}
			c.Rows = append(c.Rows, newrow)
		} else {

			for k, subv := range obj.(map[string]interface{}) {
				c.Header = append(c.Header, k)
				if reflect.TypeOf(subv).Kind() == reflect.Slice {
					c.UseObj(subv)
					return c
				}
				newrow = append(newrow, subv)
			}
			c.Rows = append(c.Rows, newrow)
		}
	case reflect.Struct:
		var newrow []interface{}
		if len(c.Header) != 0 {
			for _, h := range c.Header {
				newrow = append(newrow, v.FieldByName(h).Interface())
			}
		} else {
			for i := 0; i < t.NumField(); i++ {
				newrow = append(newrow, v.Field(i).Interface())
				c.Header = append(c.Header, t.Field(i).Name)
			}
		}

		c.Rows = append(c.Rows, newrow)
	case reflect.String:
		if len(c.Header) != 1 {
			return c
		}
		var newrow []interface{}
		newrow = append(newrow, obj)
		c.Rows = append(c.Rows, newrow)

	case reflect.Int:
		if len(c.Header) != 1 {
			return c
		}
		var newrow []interface{}
		newrow = append(newrow, obj)
		c.Rows = append(c.Rows, newrow)

	case reflect.Float64:
		if len(c.Header) != 1 {
			return c
		}
		var newrow []interface{}
		newrow = append(newrow, obj)
		c.Rows = append(c.Rows, newrow)
	}
	return c
}

// Render table
func (c *Table) Render() string {
	t := gotabulate.Create(c.Rows)
	t.SetHeaders(c.Header)
	t.SetAlign("left")
	// Set the Empty String (optional)
	t.SetEmptyString(" None")
	if c.MaxCellSize != 0 {
		t.SetMaxCellSize(c.MaxCellSize)
	}
	return t.Render("grid")
}

// PrettyPrintJSON pretty print json
func PrettyPrintJSON(obj interface{}) {
	jsonData, err := json.MarshalIndent(obj, "", "\t")
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	fmt.Println(string(jsonData))
}
