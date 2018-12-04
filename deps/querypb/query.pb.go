// Code generated by protoc-gen-go. DO NOT EDIT.
// source: query.proto
package querypb

import "strconv"

// EnumName is a helper function to simplify printing protocol buffer enums
// by name.  Given an enum map and a value, it returns a useful string.
func EnumName(m map[int32]string, v int32) string {
	s, ok := m[v]
	if ok {
		return s
	}
	return strconv.Itoa(int(v))
}

// Flags sent from the MySQL C API
type MySqlFlag int32

const (
	MySqlFlag_EMPTY                 MySqlFlag = 0
	MySqlFlag_NOT_NULL_FLAG         MySqlFlag = 1
	MySqlFlag_PRI_KEY_FLAG          MySqlFlag = 2
	MySqlFlag_UNIQUE_KEY_FLAG       MySqlFlag = 4
	MySqlFlag_MULTIPLE_KEY_FLAG     MySqlFlag = 8
	MySqlFlag_BLOB_FLAG             MySqlFlag = 16
	MySqlFlag_UNSIGNED_FLAG         MySqlFlag = 32
	MySqlFlag_ZEROFILL_FLAG         MySqlFlag = 64
	MySqlFlag_BINARY_FLAG           MySqlFlag = 128
	MySqlFlag_ENUM_FLAG             MySqlFlag = 256
	MySqlFlag_AUTO_INCREMENT_FLAG   MySqlFlag = 512
	MySqlFlag_TIMESTAMP_FLAG        MySqlFlag = 1024
	MySqlFlag_SET_FLAG              MySqlFlag = 2048
	MySqlFlag_NO_DEFAULT_VALUE_FLAG MySqlFlag = 4096
	MySqlFlag_ON_UPDATE_NOW_FLAG    MySqlFlag = 8192
	MySqlFlag_NUM_FLAG              MySqlFlag = 32768
	MySqlFlag_PART_KEY_FLAG         MySqlFlag = 16384
	MySqlFlag_GROUP_FLAG            MySqlFlag = 32768
	MySqlFlag_UNIQUE_FLAG           MySqlFlag = 65536
	MySqlFlag_BINCMP_FLAG           MySqlFlag = 131072
)

var MySqlFlag_name = map[int32]string{
	0:     "EMPTY",
	1:     "NOT_NULL_FLAG",
	2:     "PRI_KEY_FLAG",
	4:     "UNIQUE_KEY_FLAG",
	8:     "MULTIPLE_KEY_FLAG",
	16:    "BLOB_FLAG",
	32:    "UNSIGNED_FLAG",
	64:    "ZEROFILL_FLAG",
	128:   "BINARY_FLAG",
	256:   "ENUM_FLAG",
	512:   "AUTO_INCREMENT_FLAG",
	1024:  "TIMESTAMP_FLAG",
	2048:  "SET_FLAG",
	4096:  "NO_DEFAULT_VALUE_FLAG",
	8192:  "ON_UPDATE_NOW_FLAG",
	32768: "NUM_FLAG",
	16384: "PART_KEY_FLAG",
	// Duplicate value: 32768: "GROUP_FLAG",
	65536:  "UNIQUE_FLAG",
	131072: "BINCMP_FLAG",
}
var MySqlFlag_value = map[string]int32{
	"EMPTY":                 0,
	"NOT_NULL_FLAG":         1,
	"PRI_KEY_FLAG":          2,
	"UNIQUE_KEY_FLAG":       4,
	"MULTIPLE_KEY_FLAG":     8,
	"BLOB_FLAG":             16,
	"UNSIGNED_FLAG":         32,
	"ZEROFILL_FLAG":         64,
	"BINARY_FLAG":           128,
	"ENUM_FLAG":             256,
	"AUTO_INCREMENT_FLAG":   512,
	"TIMESTAMP_FLAG":        1024,
	"SET_FLAG":              2048,
	"NO_DEFAULT_VALUE_FLAG": 4096,
	"ON_UPDATE_NOW_FLAG":    8192,
	"NUM_FLAG":              32768,
	"PART_KEY_FLAG":         16384,
	"GROUP_FLAG":            32768,
	"UNIQUE_FLAG":           65536,
	"BINCMP_FLAG":           131072,
}

func (x MySqlFlag) String() string {
	return EnumName(MySqlFlag_name, int32(x))
}

// Flag allows us to qualify types by their common properties.
type Flag int32

const (
	Flag_NONE       Flag = 0
	Flag_ISINTEGRAL Flag = 256
	Flag_ISUNSIGNED Flag = 512
	Flag_ISFLOAT    Flag = 1024
	Flag_ISQUOTED   Flag = 2048
	Flag_ISTEXT     Flag = 4096
	Flag_ISBINARY   Flag = 8192
)

var Flag_name = map[int32]string{
	0:    "NONE",
	256:  "ISINTEGRAL",
	512:  "ISUNSIGNED",
	1024: "ISFLOAT",
	2048: "ISQUOTED",
	4096: "ISTEXT",
	8192: "ISBINARY",
}
var Flag_value = map[string]int32{
	"NONE":       0,
	"ISINTEGRAL": 256,
	"ISUNSIGNED": 512,
	"ISFLOAT":    1024,
	"ISQUOTED":   2048,
	"ISTEXT":     4096,
	"ISBINARY":   8192,
}

func (x Flag) String() string {
	return EnumName(Flag_name, int32(x))
}

// Type defines the various supported data types in bind vars
// and query results.
type Type int32

const (
	// NULL_TYPE specifies a NULL type.
	Type_NULL_TYPE Type = 0
	// INT8 specifies a TINYINT type.
	// Properties: 1, IsNumber.
	Type_INT8 Type = 257
	// UINT8 specifies a TINYINT UNSIGNED type.
	// Properties: 2, IsNumber, IsUnsigned.
	Type_UINT8 Type = 770
	// INT16 specifies a SMALLINT type.
	// Properties: 3, IsNumber.
	Type_INT16 Type = 259
	// UINT16 specifies a SMALLINT UNSIGNED type.
	// Properties: 4, IsNumber, IsUnsigned.
	Type_UINT16 Type = 772
	// INT24 specifies a MEDIUMINT type.
	// Properties: 5, IsNumber.
	Type_INT24 Type = 261
	// UINT24 specifies a MEDIUMINT UNSIGNED type.
	// Properties: 6, IsNumber, IsUnsigned.
	Type_UINT24 Type = 774
	// INT32 specifies a INTEGER type.
	// Properties: 7, IsNumber.
	Type_INT32 Type = 263
	// UINT32 specifies a INTEGER UNSIGNED type.
	// Properties: 8, IsNumber, IsUnsigned.
	Type_UINT32 Type = 776
	// INT64 specifies a BIGINT type.
	// Properties: 9, IsNumber.
	Type_INT64 Type = 265
	// UINT64 specifies a BIGINT UNSIGNED type.
	// Properties: 10, IsNumber, IsUnsigned.
	Type_UINT64 Type = 778
	// FLOAT32 specifies a FLOAT type.
	// Properties: 11, IsFloat.
	Type_FLOAT32 Type = 1035
	// FLOAT64 specifies a DOUBLE or REAL type.
	// Properties: 12, IsFloat.
	Type_FLOAT64 Type = 1036
	// TIMESTAMP specifies a TIMESTAMP type.
	// Properties: 13, IsQuoted.
	Type_TIMESTAMP Type = 2061
	// DATE specifies a DATE type.
	// Properties: 14, IsQuoted.
	Type_DATE Type = 2062
	// TIME specifies a TIME type.
	// Properties: 15, IsQuoted.
	Type_TIME Type = 2063
	// DATETIME specifies a DATETIME type.
	// Properties: 16, IsQuoted.
	Type_DATETIME Type = 2064
	// YEAR specifies a YEAR type.
	// Properties: 17, IsNumber, IsUnsigned.
	Type_YEAR Type = 785
	// DECIMAL specifies a DECIMAL or NUMERIC type.
	// Properties: 18, None.
	Type_DECIMAL Type = 18
	// TEXT specifies a TEXT type.
	// Properties: 19, IsQuoted, IsText.
	Type_TEXT Type = 6163
	// BLOB specifies a BLOB type.
	// Properties: 20, IsQuoted, IsBinary.
	Type_BLOB Type = 10260
	// VARCHAR specifies a VARCHAR type.
	// Properties: 21, IsQuoted, IsText.
	Type_VARCHAR Type = 6165
	// VARBINARY specifies a VARBINARY type.
	// Properties: 22, IsQuoted, IsBinary.
	Type_VARBINARY Type = 10262
	// CHAR specifies a CHAR type.
	// Properties: 23, IsQuoted, IsText.
	Type_CHAR Type = 6167
	// BINARY specifies a BINARY type.
	// Properties: 24, IsQuoted, IsBinary.
	Type_BINARY Type = 10264
	// BIT specifies a BIT type.
	// Properties: 25, IsQuoted.
	Type_BIT Type = 2073
	// ENUM specifies an ENUM type.
	// Properties: 26, IsQuoted.
	Type_ENUM Type = 2074
	// SET specifies a SET type.
	// Properties: 27, IsQuoted.
	Type_SET Type = 2075
	// TUPLE specifies a a tuple. This cannot
	// be returned in a QueryResult, but it can
	// be sent as a bind var.
	// Properties: 28, None.
	Type_TUPLE Type = 28
	// GEOMETRY specifies a GEOMETRY type.
	// Properties: 29, IsQuoted.
	Type_GEOMETRY Type = 2077
	// JSON specifies a JSON type.
	// Properties: 30, IsQuoted.
	Type_JSON Type = 2078
	// EXPRESSION specifies a SQL expression.
	// This type is for internal use only.
	// Properties: 31, None.
	Type_EXPRESSION Type = 31
)

var Type_name = map[int32]string{
	0:     "NULL_TYPE",
	257:   "INT8",
	770:   "UINT8",
	259:   "INT16",
	772:   "UINT16",
	261:   "INT24",
	774:   "UINT24",
	263:   "INT32",
	776:   "UINT32",
	265:   "INT64",
	778:   "UINT64",
	1035:  "FLOAT32",
	1036:  "FLOAT64",
	2061:  "TIMESTAMP",
	2062:  "DATE",
	2063:  "TIME",
	2064:  "DATETIME",
	785:   "YEAR",
	18:    "DECIMAL",
	6163:  "TEXT",
	10260: "BLOB",
	6165:  "VARCHAR",
	10262: "VARBINARY",
	6167:  "CHAR",
	10264: "BINARY",
	2073:  "BIT",
	2074:  "ENUM",
	2075:  "SET",
	28:    "TUPLE",
	2077:  "GEOMETRY",
	2078:  "JSON",
	31:    "EXPRESSION",
}
var Type_value = map[string]int32{
	"NULL_TYPE":  0,
	"INT8":       257,
	"UINT8":      770,
	"INT16":      259,
	"UINT16":     772,
	"INT24":      261,
	"UINT24":     774,
	"INT32":      263,
	"UINT32":     776,
	"INT64":      265,
	"UINT64":     778,
	"FLOAT32":    1035,
	"FLOAT64":    1036,
	"TIMESTAMP":  2061,
	"DATE":       2062,
	"TIME":       2063,
	"DATETIME":   2064,
	"YEAR":       785,
	"DECIMAL":    18,
	"TEXT":       6163,
	"BLOB":       10260,
	"VARCHAR":    6165,
	"VARBINARY":  10262,
	"CHAR":       6167,
	"BINARY":     10264,
	"BIT":        2073,
	"ENUM":       2074,
	"SET":        2075,
	"TUPLE":      28,
	"GEOMETRY":   2077,
	"JSON":       2078,
	"EXPRESSION": 31,
}

func (x Type) String() string {
	return EnumName(Type_name, int32(x))
}

// TransactionState represents the state of a distributed transaction.
type TransactionState int32

const (
	TransactionState_UNKNOWN  TransactionState = 0
	TransactionState_PREPARE  TransactionState = 1
	TransactionState_COMMIT   TransactionState = 2
	TransactionState_ROLLBACK TransactionState = 3
)

var TransactionState_name = map[int32]string{
	0: "UNKNOWN",
	1: "PREPARE",
	2: "COMMIT",
	3: "ROLLBACK",
}
var TransactionState_value = map[string]int32{
	"UNKNOWN":  0,
	"PREPARE":  1,
	"COMMIT":   2,
	"ROLLBACK": 3,
}

func (x TransactionState) String() string {
	return EnumName(TransactionState_name, int32(x))
}

type ExecuteOptions_IncludedFields int32

const (
	ExecuteOptions_TYPE_AND_NAME ExecuteOptions_IncludedFields = 0
	ExecuteOptions_TYPE_ONLY     ExecuteOptions_IncludedFields = 1
	ExecuteOptions_ALL           ExecuteOptions_IncludedFields = 2
)

var ExecuteOptions_IncludedFields_name = map[int32]string{
	0: "TYPE_AND_NAME",
	1: "TYPE_ONLY",
	2: "ALL",
}
var ExecuteOptions_IncludedFields_value = map[string]int32{
	"TYPE_AND_NAME": 0,
	"TYPE_ONLY":     1,
	"ALL":           2,
}

func (x ExecuteOptions_IncludedFields) String() string {
	return EnumName(ExecuteOptions_IncludedFields_name, int32(x))
}

type ExecuteOptions_Workload int32

const (
	ExecuteOptions_UNSPECIFIED ExecuteOptions_Workload = 0
	ExecuteOptions_OLTP        ExecuteOptions_Workload = 1
	ExecuteOptions_OLAP        ExecuteOptions_Workload = 2
	ExecuteOptions_DBA         ExecuteOptions_Workload = 3
)

var ExecuteOptions_Workload_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "OLTP",
	2: "OLAP",
	3: "DBA",
}
var ExecuteOptions_Workload_value = map[string]int32{
	"UNSPECIFIED": 0,
	"OLTP":        1,
	"OLAP":        2,
	"DBA":         3,
}

func (x ExecuteOptions_Workload) String() string {
	return EnumName(ExecuteOptions_Workload_name, int32(x))
}

type ExecuteOptions_TransactionIsolation int32

const (
	ExecuteOptions_DEFAULT          ExecuteOptions_TransactionIsolation = 0
	ExecuteOptions_REPEATABLE_READ  ExecuteOptions_TransactionIsolation = 1
	ExecuteOptions_READ_COMMITTED   ExecuteOptions_TransactionIsolation = 2
	ExecuteOptions_READ_UNCOMMITTED ExecuteOptions_TransactionIsolation = 3
	ExecuteOptions_SERIALIZABLE     ExecuteOptions_TransactionIsolation = 4
)

var ExecuteOptions_TransactionIsolation_name = map[int32]string{
	0: "DEFAULT",
	1: "REPEATABLE_READ",
	2: "READ_COMMITTED",
	3: "READ_UNCOMMITTED",
	4: "SERIALIZABLE",
}
var ExecuteOptions_TransactionIsolation_value = map[string]int32{
	"DEFAULT":          0,
	"REPEATABLE_READ":  1,
	"READ_COMMITTED":   2,
	"READ_UNCOMMITTED": 3,
	"SERIALIZABLE":     4,
}

func (x ExecuteOptions_TransactionIsolation) String() string {
	return EnumName(ExecuteOptions_TransactionIsolation_name, int32(x))
}

// The category of one statement.
type StreamEvent_Statement_Category int32

const (
	StreamEvent_Statement_Error StreamEvent_Statement_Category = 0
	StreamEvent_Statement_DML   StreamEvent_Statement_Category = 1
	StreamEvent_Statement_DDL   StreamEvent_Statement_Category = 2
)

var StreamEvent_Statement_Category_name = map[int32]string{
	0: "Error",
	1: "DML",
	2: "DDL",
}
var StreamEvent_Statement_Category_value = map[string]int32{
	"Error": 0,
	"DML":   1,
	"DDL":   2,
}

func (x StreamEvent_Statement_Category) String() string {
	return EnumName(StreamEvent_Statement_Category_name, int32(x))
}

type SplitQueryRequest_Algorithm int32

const (
	SplitQueryRequest_EQUAL_SPLITS SplitQueryRequest_Algorithm = 0
	SplitQueryRequest_FULL_SCAN    SplitQueryRequest_Algorithm = 1
)

var SplitQueryRequest_Algorithm_name = map[int32]string{
	0: "EQUAL_SPLITS",
	1: "FULL_SCAN",
}
var SplitQueryRequest_Algorithm_value = map[string]int32{
	"EQUAL_SPLITS": 0,
	"FULL_SCAN":    1,
}

func (x SplitQueryRequest_Algorithm) String() string {
	return EnumName(SplitQueryRequest_Algorithm_name, int32(x))
}

// Value represents a typed value.
type Value struct {
	Type  Type   `protobuf:"varint,1,opt,name=type,enum=query.Type" json:"type,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *Value) Reset()         { *m = Value{} }
func (m *Value) String() string { return "TODO" }

func (m *Value) GetType() Type {
	if m != nil {
		return m.Type
	}
	return Type_NULL_TYPE
}

func (m *Value) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

// BindVariable represents a single bind variable in a Query.
type BindVariable struct {
	Type  Type   `protobuf:"varint,1,opt,name=type,enum=query.Type" json:"type,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	// values are set if type is TUPLE.
	Values []*Value `protobuf:"bytes,3,rep,name=values" json:"values,omitempty"`
}

func (m *BindVariable) Reset()         { *m = BindVariable{} }
func (m *BindVariable) String() string { return "TODO" }

func (m *BindVariable) GetType() Type {
	if m != nil {
		return m.Type
	}
	return Type_NULL_TYPE
}

func (m *BindVariable) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *BindVariable) GetValues() []*Value {
	if m != nil {
		return m.Values
	}
	return nil
}

// BoundQuery is a query with its bind variables
type BoundQuery struct {
	// sql is the SQL query to execute
	Sql string `protobuf:"bytes,1,opt,name=sql" json:"sql,omitempty"`
	// bind_variables is a map of all bind variables to expand in the query.
	// nil values are not allowed. Use NULL_TYPE to express a NULL value.
	BindVariables map[string]*BindVariable `protobuf:"bytes,2,rep,name=bind_variables,json=bindVariables" json:"bind_variables,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *BoundQuery) Reset()         { *m = BoundQuery{} }
func (m *BoundQuery) String() string { return "TODO" }

func (m *BoundQuery) GetSql() string {
	if m != nil {
		return m.Sql
	}
	return ""
}

func (m *BoundQuery) GetBindVariables() map[string]*BindVariable {
	if m != nil {
		return m.BindVariables
	}
	return nil
}
