package main

var types = map[string]spec{
	"bool":     ext("EncodeBool", "DecodeBool"),
	"byte":     nil, //ext("encodeByte", "decodeByte"),
	"uint8":    nil, //ui(1),
	"int8":     ext("EncodeInt8", "DecodeInt8"),
	"uint16":   ui(2),
	"int16":    i(2),
	"uint32":   ui(4),
	"int32":    i(4),
	"uint64":   ui(8),
	"int64":    i(8),
	"VarInt":   ext("EncodeVarInt", "DecodeVarInt"),
	"varLong":  nil, //t(parseVarInt, 8, true),
	"float32":  nil, //t(parseFloat, 4, true),
	"float64":  nil, //t(parseFloat, 8, true),
	"string":   ext("EncodeString", "DecodeString"),
	"Buffer":   nil, //ext("encodeBuffer", "decodeBuffer"),
	"Position": nil, //ext("encodePosition", "decodePosition"),
	"MetaData": nil, //ext("encodeMetaData", "decodeMetaData"),
}
