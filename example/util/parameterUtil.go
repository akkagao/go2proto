package util

import (
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"

	"go2proto/example/proto"
)

func CreateParameter() *proto.DemoRequest {
	rand.Seed(time.Now().UnixNano())
	demoRequest := &proto.DemoRequest{
		CommonRequest: &proto.CommonRequest{
			Token: "",
			Device: &proto.Device{
				Channel: "999999999",
			},
		},
		Float64Type:      rand.Float64(),
		Float32Type:      rand.Float32(),
		Int32Type:        rand.Int31(),
		Int64Type:        rand.Int63(),
		Uint32Type:       rand.Uint32(),
		Uint64Type:       rand.Uint64(),
		BoolType:         false,
		StringType:       time.Now().Format("2006/01/02 15:04:05"),
		Float64ArrayType: []float64{rand.Float64()},
		Float32ArrayType: []float32{rand.Float32()},
		Int32ArrayType:   []int32{rand.Int31()},
		Int64ArrayType:   []int64{rand.Int63()},
		Uint32ArrayType:  []uint32{rand.Uint32()},
		Uint64ArrayType:  []uint64{rand.Uint64()},
		BoolArrayType:    []bool{true, false},
		StringArrayType:  []string{"1", "2"},
		BytesType:        []byte("hello"),
		StructType:       createStructType(),
		StructTypes: []*proto.StructType{
			createStructType(),
			createStructType(),
		},
	}
	return demoRequest
}

func CreateResult(demoRequest *proto.DemoRequest) *proto.DemoResponse {
	rand.Seed(time.Now().UnixNano())
	demoResponse := &proto.DemoResponse{
		CommonResponse: &proto.CommonResponse{
			RequestId: getReqId(),
			Code:      "200",
			Message:   "success",
		},
		Float64Type:      demoRequest.Float64Type,
		Float32Type:      demoRequest.Float32Type,
		Int32Type:        demoRequest.Int32Type,
		Int64Type:        demoRequest.Int64Type,
		Uint32Type:       demoRequest.Uint32Type,
		Uint64Type:       demoRequest.Uint64Type,
		BoolType:         demoRequest.BoolType,
		StringType:       demoRequest.StringType,
		Float64ArrayType: demoRequest.Float64ArrayType,
		Float32ArrayType: demoRequest.Float32ArrayType,
		Int32ArrayType:   demoRequest.Int32ArrayType,
		Int64ArrayType:   demoRequest.Int64ArrayType,
		Uint32ArrayType:  demoRequest.Uint32ArrayType,
		Uint64ArrayType:  demoRequest.Uint64ArrayType,
		BoolArrayType:    demoRequest.BoolArrayType,
		StringArrayType:  demoRequest.StringArrayType,
		BytesType:        demoRequest.BytesType,
		StructType:       demoRequest.StructType,
		StructTypes:      demoRequest.StructTypes,
	}
	return demoResponse
}

func getReqId() string {
	uuidTmp, err := uuid.NewRandom()
	if err != nil {
		return ""
	}
	return strings.Replace(uuidTmp.String(), "-", "", -1)

}

func createStructType() *proto.StructType {
	structType := &proto.StructType{
		Float64Type:      rand.Float64(),
		Float32Type:      rand.Float32(),
		Int32Type:        rand.Int31(),
		Int64Type:        rand.Int63(),
		Uint32Type:       rand.Uint32(),
		Uint64Type:       rand.Uint64(),
		BoolType:         false,
		StringType:       time.Now().Format("2006/01/02 15:04:05"),
		Float64ArrayType: []float64{rand.Float64()},
		Float32ArrayType: []float32{rand.Float32()},
		Int32ArrayType:   []int32{rand.Int31()},
		Int64ArrayType:   []int64{rand.Int63()},
		Uint32ArrayType:  []uint32{rand.Uint32()},
		Uint64ArrayType:  []uint64{rand.Uint64()},
		BoolArrayType:    []bool{true, false},
		StringArrayType:  []string{"1", "2"},
		BytesType:        []byte("hello"),
	}
	return structType
}
