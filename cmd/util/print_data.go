// Copyright © 2016 Paul Allen <paul@cloudcoreo.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package util

import (
	"bytes"
	"reflect"

	"encoding/json"
	"fmt"

	"os"

	"io"

	"github.com/CloudCoreo/cli/cmd/content"
	"github.com/bndr/gotabulate"
)

// Table struct
type Table struct {
	HeaderMap   map[string]string
	Header      []string
	Rows        [][]interface{}
	MaxCellSize int
}

// NewTable create new table
func NewTable() *Table {
	return new(Table)
}

// SetHeaderMap set headerMap for table
func (c *Table) SetHeaderMap(headerMap map[string]string) *Table {
	c.HeaderMap = headerMap
	return c
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

	if c.HeaderMap != nil {
		var headers []string
		for _, h := range c.Header {
			headers = append(headers, c.HeaderMap[h])
		}

		t.SetHeaders(headers)
	} else {
		t.SetHeaders(c.Header)
	}

	t.SetAlign("center")
	// Set the Empty String (optional)
	t.SetEmptyString("None")
	if c.MaxCellSize != 0 {
		t.SetMaxCellSize(c.MaxCellSize)
	}
	return t.Render("simple")
}

// PrettyPrintJSON pretty print json
func PrettyPrintJSON(obj interface{}) {
	fmt.Println(PrettyJSON(obj))
}

// PrettyJSON return pretty json
func PrettyJSON(obj interface{}) string {
	buf := new(bytes.Buffer)
	encoder := json.NewEncoder(buf)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "\t")
	err := encoder.Encode(obj)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return ""
	}

	return string(buf.String())
}

//PrintError print error
func PrintError(err error, json bool) {
	if json {
		PrettyPrintJSON(err)
	} else {
		fmt.Fprintf(os.Stderr, err.Error())
	}
}

//PrintResult print result
func PrintResult(out io.Writer, t interface{}, headers []string, headersMap map[string]string, json, verbose bool) {
	if json {
		PrettyPrintJSON(t)
	} else {
		table := NewTable()
		table.SetHeader(headers)
		table.SetHeaderMap(headersMap)

		table.UseObj(t)
		fmt.Fprintln(out, table.Render())
	}

	if verbose {
		fmt.Fprintln(out, content.InfoCommandSuccess)
	}
}
