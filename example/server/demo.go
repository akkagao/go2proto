package server

import (
	"go2protoexample/server/common"
)

/*
UserDemo 用户服务
*/
type Demo interface {
	CreateDemo(demoRequest DemoRequest) DemoResponse                            // 创建DEMO
	// DemoStream_stream(userRequest DemoRequestStream) DemoResponseStream         // stream 调用
	// DemoPingPang_pingpang(userRequest DemoRequestPingPang) DemoResponsePingPang // pingpang 调用
}

type DemoRequest struct {
	CommonRequest    common.CommonRequest // 请求公共参数
	Float64Type      float64              // float64 参数测试
	Float32Type      float32              // float32 参数测试
	Int32Type        int32                // int32 参数测试
	Int64Type        int64                // int64 参数测试
	Uint32Type       uint32               // uint32 参数测试
	Uint64Type       uint64               // uint64 参数测试
	BoolType         bool                 // bool 参数测试
	StringType       string               // string 参数测试
	Float64ArrayType []float64            // []float64 参数测试
	Float32ArrayType []float32            // []float32 参数测试
	Int32ArrayType   []int32              // []int32 参数测试
	Int64ArrayType   []int64              // []int64 参数测试
	Uint32ArrayType  []uint32             // []uint32 参数测试
	Uint64ArrayType  []uint64             // []uint64 参数测试
	BoolArrayType    []bool               // []bool 参数测试
	StringArrayType  []string             // []string 参数测试
	BytesType        []byte               // []byte 类型测试
	StructType       StructType           // 结构体类型的测试
	StructTypes      []StructType         // 结构体数组类型测试
}

type DemoRequestPingPang struct {
	CommonRequest    common.CommonRequest // 请求公共参数
	Float64Type      float64              // float64 参数测试
	Float32Type      float32              // float32 参数测试
	Int32Type        int32                // int32 参数测试
	Int64Type        int64                // int64 参数测试
	Uint32Type       uint32               // uint32 参数测试
	Uint64Type       uint64               // uint64 参数测试
	BoolType         bool                 // bool 参数测试
	StringType       string               // string 参数测试
	Float64ArrayType []float64            // []float64 参数测试
	Float32ArrayType []float32            // []float32 参数测试
	Int32ArrayType   []int32              // []int32 参数测试
	Int64ArrayType   []int64              // []int64 参数测试
	Uint32ArrayType  []uint32             // []uint32 参数测试
	Uint64ArrayType  []uint64             // []uint64 参数测试
	BoolArrayType    []bool               // []bool 参数测试
	StringArrayType  []string             // []string 参数测试
	BytesType        []byte               // []byte 类型测试
	StructType       StructType           // 结构体类型的测试
	StructTypes      []StructType         // 结构体数组类型测试
}

type DemoRequestStream struct {
	CommonRequest    common.CommonRequest // 请求公共参数
	Float64Type      float64              // float64 参数测试
	Float32Type      float32              // float32 参数测试
	Int32Type        int32                // int32 参数测试
	Int64Type        int64                // int64 参数测试
	Uint32Type       uint32               // uint32 参数测试
	Uint64Type       uint64               // uint64 参数测试
	BoolType         bool                 // bool 参数测试
	StringType       string               // string 参数测试
	Float64ArrayType []float64            // []float64 参数测试
	Float32ArrayType []float32            // []float32 参数测试
	Int32ArrayType   []int32              // []int32 参数测试
	Int64ArrayType   []int64              // []int64 参数测试
	Uint32ArrayType  []uint32             // []uint32 参数测试
	Uint64ArrayType  []uint64             // []uint64 参数测试
	BoolArrayType    []bool               // []bool 参数测试
	StringArrayType  []string             // []string 参数测试
	BytesType        []byte               // []byte 类型测试
	StructType       StructType           // 结构体类型的测试
	StructTypes      []StructType         // 结构体数组类型测试
}

type StructType struct {
	Float64Type      float64   // float64 参数测试
	Float32Type      float32   // float32 参数测试
	Int32Type        int32     // int32 参数测试
	Int64Type        int64     // int64 参数测试
	Uint32Type       uint32    // uint32 参数测试
	Uint64Type       uint64    // uint64 参数测试
	BoolType         bool      // bool 参数测试
	StringType       string    // string 参数测试
	Float64ArrayType []float64 // []float64 参数测试
	Float32ArrayType []float32 // []float32 参数测试
	Int32ArrayType   []int32   // []int32 参数测试
	Int64ArrayType   []int64   // []int64 参数测试
	Uint32ArrayType  []uint32  // []uint32 参数测试
	Uint64ArrayType  []uint64  // []uint64 参数测试
	BoolArrayType    []bool    // []bool 参数测试
	StringArrayType  []string  // []string 参数测试
	BytesType        []byte    // []byte 类型测试
}

type DemoResponse struct {
	CommonResponse   common.CommonResponse // 返回公共参数
	Float64Type      float64               // float64 参数测试
	Float32Type      float32               // float32 参数测试
	Int32Type        int32                 // int32 参数测试
	Int64Type        int64                 // int64 参数测试
	Uint32Type       uint32                // uint32 参数测试
	Uint64Type       uint64                // uint64 参数测试
	BoolType         bool                  // bool 参数测试
	StringType       string                // string 参数测试
	Float64ArrayType []float64             // []float64 参数测试
	Float32ArrayType []float32             // []float32 参数测试
	Int32ArrayType   []int32               // []int32 参数测试
	Int64ArrayType   []int64               // []int64 参数测试
	Uint32ArrayType  []uint32              // []uint32 参数测试
	Uint64ArrayType  []uint64              // []uint64 参数测试
	BoolArrayType    []bool                // []bool 参数测试
	StringArrayType  []string              // []string 参数测试
	BytesType        []byte                // []byte 类型测试
	StructType       StructType            // 结构体类型的测试
	StructTypes      []StructType          // 结构体数组类型测试
}

type DemoResponseStream struct {
	CommonResponse   common.CommonResponse // 返回公共参数
	Float64Type      float64               // float64 参数测试
	Float32Type      float32               // float32 参数测试
	Int32Type        int32                 // int32 参数测试
	Int64Type        int64                 // int64 参数测试
	Uint32Type       uint32                // uint32 参数测试
	Uint64Type       uint64                // uint64 参数测试
	BoolType         bool                  // bool 参数测试
	StringType       string                // string 参数测试
	Float64ArrayType []float64             // []float64 参数测试
	Float32ArrayType []float32             // []float32 参数测试
	Int32ArrayType   []int32               // []int32 参数测试
	Int64ArrayType   []int64               // []int64 参数测试
	Uint32ArrayType  []uint32              // []uint32 参数测试
	Uint64ArrayType  []uint64              // []uint64 参数测试
	BoolArrayType    []bool                // []bool 参数测试
	StringArrayType  []string              // []string 参数测试
	BytesType        []byte                // []byte 类型测试
	StructType       StructType            // 结构体类型的测试
	StructTypes      []StructType          // 结构体数组类型测试
}

type DemoResponsePingPang struct {
	CommonResponse   common.CommonResponse // 返回公共参数
	Float64Type      float64               // float64 参数测试
	Float32Type      float32               // float32 参数测试
	Int32Type        int32                 // int32 参数测试
	Int64Type        int64                 // int64 参数测试
	Uint32Type       uint32                // uint32 参数测试
	Uint64Type       uint64                // uint64 参数测试
	BoolType         bool                  // bool 参数测试
	StringType       string                // string 参数测试
	Float64ArrayType []float64             // []float64 参数测试
	Float32ArrayType []float32             // []float32 参数测试
	Int32ArrayType   []int32               // []int32 参数测试
	Int64ArrayType   []int64               // []int64 参数测试
	Uint32ArrayType  []uint32              // []uint32 参数测试
	Uint64ArrayType  []uint64              // []uint64 参数测试
	BoolArrayType    []bool                // []bool 参数测试
	StringArrayType  []string              // []string 参数测试
	BytesType        []byte                // []byte 类型测试
	StructType       StructType            // 结构体类型的测试
	StructTypes      []StructType          // 结构体数组类型测试
}
