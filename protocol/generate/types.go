package main

var types = map[string]spec{
	"bool":     ext("encodeBool", "decodeBool"),
	"byte":     ext("encodeByte", "decodeByte"),
	"uint8":    ui(1),
	"int8":     i(1),
	"uint16":   ui(2),
	"int16":    i(2),
	"uint32":   ui(4),
	"int32":    i(4),
	"uint64":   ui(8),
	"int64":    i(8),
	"varint":   nil, //t(parseVarInt, 4, true),
	"varLong":  nil, //t(parseVarInt, 8, true),
	"float32":  nil, //t(parseFloat, 4, true),
	"float64":  nil, //t(parseFloat, 8, true),
	"string":   ext("encodeString", "decodeString"),
	"Buffer":   ext("encodeBuffer", "decodeBuffer"),
	"Position": ext("encodePosition", "decodePosition"),
	"MetaData": ext("encodeMetaData", "decodeMetaData"),
}
