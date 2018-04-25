// https://stackoverflow.com/questions/27744493/can-i-use-gorilla-schema-with-an-sql-nullstring
// https://gist.github.com/carbocation/51b55297702c7d30d3ef
package util

import (
	"log"
	"reflect"

	"github.com/gorilla/schema"
	"github.com/guregu/null"
)

var SchemaDecoder = schema.NewDecoder()

// Convertors for sql.Null* types so that they can be
// used with gorilla/schema
func init() {
	SchemaRegisterSQLNulls(SchemaDecoder)
}

func SchemaRegisterSQLNulls(d *schema.Decoder) {
	// nullString, nullBool, nullInt64, nullFloat64, nullTime := null.String{}, null.Bool{}, null.Int{}, null.Float{}, null.Time{}
	nullString, nullBool, nullInt64, nullFloat64 := null.String{}, null.Bool{}, null.Int{}, null.Float{}

	d.RegisterConverter(nullString, ConvertSQLNullString)
	d.RegisterConverter(nullBool, ConvertSQLNullBool)
	d.RegisterConverter(nullInt64, ConvertSQLNullInt64)
	d.RegisterConverter(nullFloat64, ConvertSQLNullFloat64)
	// d.RegisterConverter(nullTime, ConvertSQLTime)
}

func ConvertSQLNullString(value string) reflect.Value {
	v := null.String{}
	if err := v.Scan(value); err != nil {
		return reflect.Value{}
	}

	return reflect.ValueOf(v)
}

func ConvertSQLNullBool(value string) reflect.Value {
	v := null.Bool{}
	if err := v.Scan(value); err != nil {
		return reflect.Value{}
	}

	return reflect.ValueOf(v)
}

func ConvertSQLNullInt64(value string) reflect.Value {
	v := null.Int{}
	if err := v.Scan(value); err != nil {
		return reflect.Value{}
	}

	return reflect.ValueOf(v)
}

func ConvertSQLNullFloat64(value string) reflect.Value {
	v := null.Float{}
	if err := v.Scan(value); err != nil {
		return reflect.Value{}
	}

	return reflect.ValueOf(v)
}

func ConvertSQLTime(value string) reflect.Value {
	v := null.Time{}
	if err := v.Scan(value); err != nil {
		log.Println("nil")
		log.Println(err)
		return reflect.Value{}
	}

	return reflect.ValueOf(v)
}
