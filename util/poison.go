// https://stackoverflow.com/questions/27744493/can-i-use-gorilla-schema-with-an-sql-nullstring
// https://gist.github.com/carbocation/51b55297702c7d30d3ef
package util

import (
	"database/sql"
	"reflect"

	"github.com/gorilla/schema"
)

var SchemaDecoder = schema.NewDecoder()

// Convertors for sql.Null* types so that they can be
// used with gorilla/schema
func init() {
	SchemaRegisterSQLNulls(SchemaDecoder)
}

func SchemaRegisterSQLNulls(d *schema.Decoder) {
	nullString, nullBool, nullInt64, nullFloat64 := sql.NullString{}, sql.NullBool{}, sql.NullInt64{}, sql.NullFloat64{}

	d.RegisterConverter(nullString, ConvertSQLNullString)
	d.RegisterConverter(nullBool, ConvertSQLNullBool)
	d.RegisterConverter(nullInt64, ConvertSQLNullInt64)
	d.RegisterConverter(nullFloat64, ConvertSQLNullFloat64)
}

func ConvertSQLNullString(value string) reflect.Value {
	v := sql.NullString{}
	if err := v.Scan(value); err != nil {
		return reflect.Value{}
	}

	return reflect.ValueOf(v)
}

func ConvertSQLNullBool(value string) reflect.Value {
	v := sql.NullBool{}
	if err := v.Scan(value); err != nil {
		return reflect.Value{}
	}

	return reflect.ValueOf(v)
}

func ConvertSQLNullInt64(value string) reflect.Value {
	v := sql.NullInt64{}
	if err := v.Scan(value); err != nil {
		return reflect.Value{}
	}

	return reflect.ValueOf(v)
}

func ConvertSQLNullFloat64(value string) reflect.Value {
	v := sql.NullFloat64{}
	if err := v.Scan(value); err != nil {
		return reflect.Value{}
	}

	return reflect.ValueOf(v)
}