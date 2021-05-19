package xorm

import (
	"reflect"
	"strings"
	"time"

	"github.com/xormplus/xorm/schemas"
)

var (
	Bit       = "BIT"
	TinyInt   = "TINYINT"
	SmallInt  = "SMALLINT"
	MediumInt = "MEDIUMINT"
	Int       = "INT"
	Integer   = "INTEGER"
	BigInt    = "BIGINT"

	Enum = "ENUM"
	Set  = "SET"

	Char             = "CHAR"
	Varchar          = "VARCHAR"
	NChar            = "NCHAR"
	NVarchar         = "NVARCHAR"
	TinyText         = "TINYTEXT"
	Text             = "TEXT"
	NText            = "NTEXT"
	Clob             = "CLOB"
	MediumText       = "MEDIUMTEXT"
	LongText         = "LONGTEXT"
	Uuid             = "UUID"
	UniqueIdentifier = "UNIQUEIDENTIFIER"
	SysName          = "SYSNAME"

	Date          = "DATE"
	DateTime      = "DATETIME"
	SmallDateTime = "SMALLDATETIME"
	Time          = "TIME"
	TimeStamp     = "TIMESTAMP"
	TimeStampz    = "TIMESTAMPZ"
	Year          = "YEAR"

	Decimal    = "DECIMAL"
	Numeric    = "NUMERIC"
	Money      = "MONEY"
	SmallMoney = "SMALLMONEY"

	Real   = "REAL"
	Float  = "FLOAT"
	Double = "DOUBLE"

	Binary     = "BINARY"
	VarBinary  = "VARBINARY"
	TinyBlob   = "TINYBLOB"
	Blob       = "BLOB"
	MediumBlob = "MEDIUMBLOB"
	LongBlob   = "LONGBLOB"
	Bytea      = "BYTEA"

	Bool = "BOOL"

	Serial    = "SERIAL"
	BigSerial = "BIGSERIAL"
)

// SQLType2GoType default sql type change to go type
func SQLType2GoType(typ string) reflect.Type {
	switch strings.ToUpper(typ) {
	case Bit, TinyInt, SmallInt, MediumInt, Int, Integer, Serial:
		return reflect.TypeOf(1)
	case BigInt, BigSerial:
		return reflect.TypeOf(int64(1))
	case Float, Real:
		return reflect.TypeOf(float64(1))
	case Double:
		return reflect.TypeOf(float64(1))
	case Char, NChar, Varchar, NVarchar, TinyText, Text, NText, MediumText, LongText, Enum, Set, Uuid, Clob, SysName:
		return reflect.TypeOf("")
	case TinyBlob, Blob, LongBlob, Bytea, Binary, MediumBlob, VarBinary, UniqueIdentifier:
		return reflect.TypeOf([]byte{})
	case Bool:
		return reflect.TypeOf(true)
	case DateTime, Date, Time, TimeStamp, TimeStampz, SmallDateTime, Year:
		return reflect.TypeOf(time.Time{})
	case Decimal, Numeric, Money, SmallMoney:
		return reflect.TypeOf(float64(1))
	default:
		return reflect.TypeOf("")
	}
}

func TypeTransform(col *schemas.Column) string {
	typ := SQLType2GoType(col.SQLType.Name).String()

	// 这三个字段无视数据库的数据类型 全都用 time
	if col.Name == "created_at" || col.Name == "updated_at" || col.Name == "deleted_at" {
		return "time.Time"
	}

	switch typ {
	case "[]uint8":
		return "[]byte"
	case "int":
		return "int64"
	case "time.Time":
		return "time.Time"
	default:
		return typ
	}
}
