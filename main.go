package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path"
	"strings"
	"text/template"
	"unicode"
)

const streamFlag = "_stream"
const pingpang = "_pingpang"

var filePath string

// var dir string
var target string

func init() {
	flag.StringVar(&filePath, "f", "", "source file path")
	// flag.StringVar(&dir, "d", "", "source file dir path")
	flag.StringVar(&target, "t", "proto", "proto file target path")
	flag.Usage = usage
}

func main() {
	flag.Parse()
	if filePath == "" {
		flag.Usage()
		return
	}
	start(filePath)
}

func start(interfacePath string) {
	// Create the AST by parsing src.
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, interfacePath, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	microService := MicroService{PackageName: target}
	imports := []string{}
	messages := []Message{}
	ast.Inspect(f, func(node ast.Node) bool {
		if node == nil {
			return true
		}
		if importSpec, ok := node.(*ast.ImportSpec); ok {
			if importSpec.Path.Value == "\"time\"" {
				microService.ImportTime = true
			} else {
				importpath := importSpec.Path.Value
				importpath = fmt.Sprintf("%v%v.proto", target, importpath[strings.LastIndex(importpath, "/"):len(importpath)-1])
				imports = append(imports, importpath)
			}
		}

		if typeSpecNode, ok := node.(*ast.TypeSpec); ok {
			// 处理接口
			if interfaceNode, f := typeSpecNode.Type.(*ast.InterfaceType); f {
				fmt.Println("接口名称：", typeSpecNode.Name.Name)
				serviceFunctions := interfaceParser(interfaceNode)
				service := Service{}
				service.Name = typeSpecNode.Name.Name
				service.PackageName = genPackageName(typeSpecNode.Name.Name)
				service.ServiceFunctions = serviceFunctions
				microService.Service = service
				// spew.Dump(serviceFunctions)
			}
			// 处理结构体
			if structNode, f := typeSpecNode.Type.(*ast.StructType); f {
				structName := typeSpecNode.Name.Name
				log.Println("struct名称：", structName)
				messageFields := structParser(structName, structNode)
				message := Message{}
				message.Name = structName
				message.MessageFields = messageFields
				messages = append(messages, message)
				// spew.Dump(messageFields)
			}

		}
		return true
	})
	microService.Messages = messages
	microService.Imports = imports

	targetFileName := strings.Replace(path.Base(interfacePath), ".go", ".proto", -1)
	saveToFile(microService, fmt.Sprintf("%v/%v", target, targetFileName))
}

func genPackageName(s string) string {
	if unicode.IsUpper([]rune(s)[0]) {
		return strings.ToLower(string(s[0])) + string(s[1:])
	}
	return s
}

/*
proto 文件模板
*/
func saveToFile(microService MicroService, protoPath string) {
	t, err := template.New("protoTemplate").Parse(protoTemplate)
	if err != nil {
		log.Panic(err)
	}
	os.MkdirAll(path.Dir(protoPath), 0755)
	os.Remove(protoPath)
	f, err := os.OpenFile(protoPath, os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		log.Panic(err)
	}

	err = t.Execute(f, microService)
	if err != nil {
		log.Panic(err)
	}
}

/*
解析结构体
*/
func structParser(structName string, structNode *ast.StructType) []MessageField {
	messageFields := []MessageField{}
	for i, field := range structNode.Fields.List {
		messageField := MessageField{}
		messageField.Index = i + 1
		// 如果是类似name,address string 这样的定义则报错
		if len(field.Names) != 1 {
			log.Fatalf("struct %v error,the field can't define like 'name,address string'", structName)
		}
		messageField.FieldName = field.Names[0].Name

		// 基本类型处理
		if fieldType, ok := field.Type.(*ast.Ident); ok {
			messageField.FieldType = getProtoType(fieldType.Name)
		}

		// map类型处理
		if fieldType, ok := field.Type.(*ast.MapType); ok {
			key, value := "", ""
			if keyType, ok := fieldType.Key.(*ast.Ident); ok {
				key = getProtoType(keyType.Name)
			}
			if valueType, ok := fieldType.Value.(*ast.Ident); ok {
				value = getProtoType(valueType.Name)
			}
			messageField.FieldType = fmt.Sprintf("map<%v,%v>", key, value)
		}

		// 处理引用类型
		if fieldType, ok := field.Type.(*ast.SelectorExpr); ok {
			if p, ok := fieldType.X.(*ast.Ident); ok && p.Name == "time" {
				messageField.FieldType = "google.protobuf.Timestamp"
			} else {
				messageField.FieldType = fieldType.Sel.Name
			}
		}
		// 处理参数是数组的情况
		if fieldType, ok := field.Type.(*ast.ArrayType); ok {
			if fieldTypeElt, ok := fieldType.Elt.(*ast.Ident); ok {
				// byte 数组特殊处理 转换成bytes
				if fieldTypeElt.Name == "byte" {
					messageField.FieldType = "bytes"
				} else {
					messageField.FieldType = "repeated " + getProtoType(fieldTypeElt.Name)
				}
			}
		}
		// 获取注释
		if field.Comment != nil {
			messageField.Comment = field.Comment.List[0].Text
		}
		messageFields = append(messageFields, messageField)
	}
	return messageFields
}

/*
解析接口代码
*/
func interfaceParser(interfaceNode *ast.InterfaceType) []ServiceFunction {
	serviceFunctions := []ServiceFunction{}
	// 解析方法列表
	for _, function := range interfaceNode.Methods.List {

		serviceFunction := ServiceFunction{}
		if function.Comment != nil {
			serviceFunction.Comment = function.Comment.List[0].Text
		}
		// 获取方法名称
		if len(function.Names) != 1 {
			log.Fatal("parser function error")
		}

		functionName := function.Names[0].Name
		if strings.HasSuffix(functionName, streamFlag) {
			serviceFunction.Name = strings.Replace(functionName, streamFlag, "", -1)
			serviceFunction.Stream = true
		} else if strings.HasSuffix(functionName, pingpang) {
			serviceFunction.Name = strings.Replace(functionName, pingpang, "", -1)
			serviceFunction.PingPong = true
		} else {
			serviceFunction.Name = functionName
		}

		// 解析方法
		if funcBody, ok := function.Type.(*ast.FuncType); ok {
			// 解析参数列表
			for i, param := range funcBody.Params.List {
				// 获取参数名称
				for _, paramName := range param.Names {
					log.Printf("function:%v index:%v paramName:%v ", functionName, i+1, paramName.Name)
				}
				// 获取参数类型
				if paramType, ok := param.Type.(*ast.Ident); ok {
					serviceFunction.ParamType = paramType.Name
					log.Printf("function:%v index:%v paramType:%v ", functionName, i+1, paramType.Name)
				}
			}
			// 解析返回值
			for i, result := range funcBody.Results.List {
				// 获取参数名称
				for _, resultName := range result.Names {
					log.Printf("function:%v index:%v resultName:%v ", functionName, i+1, resultName.Name)
				}
				// 获取参数类型
				if resultType, ok := result.Type.(*ast.Ident); ok {
					serviceFunction.ResultType = resultType.Name
					log.Printf("function:%v index:%v resultType:%v ", functionName, i+1, resultType.Name)
				}
			}
		}
		serviceFunctions = append(serviceFunctions, serviceFunction)
	}
	return serviceFunctions
}

/*
go 类型转换成proto类型
*/
func getProtoType(fieldType string) string {
	switch fieldType {
	case "float64":
		return "double";
	case "float32":
		return "float"
	case "int32", "int":
		return "int32"
	case "int64":
		return "int64"
	case "uint32":
		return "uint32"
	case "uint64":
		return "uint64"
	case "bool":
		return "bool"
	case "string":
		return "string"
	case "int8", "int16":
		log.Fatal("int8, int16 is nonsupport type")
		return fieldType
	default:
		// 如果类型是引用其他结构体，则直接返回名称
		return fieldType
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, `go2proto version: go2proto/1.0.0
Usage: go2proto [-f] [-t]

Options:
`)
	flag.PrintDefaults()
}
