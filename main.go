package main

import (
	"encoding/json"
	"flag"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/dynamicpb"
	"google.golang.org/protobuf/types/pluginpb"
)

func main() {

	var flags flag.FlagSet

	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {
		// this enables optional fields to be supported.
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

		// iterate over the files included in the request.
		for _, file := range gen.Files {

			// will limit file generation to files included in the request.
			// if you want to generate for imports this statement can be excluded.
			if !file.Generate {
				continue
			}

			// determine file name GeneratedFilenamePrefix will get the generated file to be aligned with the written proto file.
			// second parameter is the go import path, normally "" or ".".
			generatedFile := gen.NewGeneratedFile("protoslides/"+file.GeneratedFilenamePrefix+".slide", "")

			generatedFile.P("# protoc slides for " + file.GeneratedFilenamePrefix)
			generatedFile.P("\n")

			for _, service := range file.Services {

				// can print to file through the P method.
				generatedFile.P("## Service - " + service.GoName)
				generatedFile.P("\n")

				// iterate over a services methods
				for _, method := range service.Methods {

					// for all the methods in a service print info for the method.
					generatedFile.P("## Method - " + method.GoName)
					generatedFile.P("\n")
					generatedFile.P("FullName: " + method.Desc.FullName())
					// getting comments from a message.
					generatedFile.P("\n")
					generatedFile.P("Comments: " + strings.ReplaceAll(method.Comments.Leading.String(), "//", ""))
					generatedFile.P("\n")
					// you can also get [] LeadingDetached comments & Trailing comments
					// method.Comments.LeadingDetached && method.Comments.Trailing

					generatedFile.P("## Method - "+method.GoName, " Request")
					generatedFile.P("\n")

					generatedFile.P("Type: " + method.Input.Desc.FullName())
					generatedFile.P("\n")

					comments := strings.ReplaceAll(method.Input.Comments.Leading.String(), "//", "")
					if len(comments) > 0 {
						generatedFile.P("Message Comments: ", comments, "\n")
					}

					skellyReq := genSkeletonMessage(method.Input.Desc)
					generatedFile.P("```")
					generatedFile.P(makeCurlJson(skellyReq))
					generatedFile.P("```", "\n")

					generatedFile.P("## Method - "+method.GoName, " Response", "\n")
					generatedFile.P("Type: "+method.Output.Desc.FullName(), "\n")
					comments = strings.ReplaceAll(method.Output.Comments.Leading.String(), "//", "")
					if len(comments) > 0 {
						generatedFile.P("Message Comments: ", comments, "\n")
					}
					skellyRes := genSkeletonMessage(method.Output.Desc)
					generatedFile.P("\n")
					generatedFile.P("```")
					generatedFile.P(makeCurlJson(skellyRes))
					generatedFile.P("```")
					generatedFile.P("\n")
				}
			}

			// show all enums.
			for _, enum := range file.Enums {
				generatedFile.P("## Enum - "+enum.Desc.FullName(), "\n")

				generatedFile.P("Desc", strings.ReplaceAll(enum.Comments.Leading.String(), "//", ""))
				generatedFile.P("\n")
				generatedFile.P("Values: \n")

				for _, ev := range enum.Values {
					generatedFile.P("- ", ev.Desc.FullName(), "\n")
				}
			}
		}

		return nil
	})
}

func genSkeletonMessage(md protoreflect.MessageDescriptor) any {
	defaultMessage := map[string]any{}

	dm := dynamicpb.NewMessage(md)

	for i := 0; i < dm.Descriptor().Fields().Len(); i++ {
		fd := dm.Descriptor().Fields().Get(i)
		name := string(fd.Name())

		switch fd.Kind() {
		case protoreflect.BoolKind:
			defaultMessage[name] = true
		case protoreflect.EnumKind:
			defaultMessage[name] = fd.Enum().Values().Get(0).Name()
		case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Uint32Kind, protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Uint64Kind, protoreflect.Sfixed32Kind, protoreflect.Fixed32Kind:
			defaultMessage[name] = 1
		case protoreflect.StringKind:
			defaultMessage[name] = ""
		case protoreflect.BytesKind:
			defaultMessage[name] = []byte(fd.JSONName())
		case protoreflect.MessageKind:
			// edge case where object contains itself as a field.
			if fd.Message().FullName() == md.FullName() {
				defaultMessage[name] = ""
				continue
			}
			defaultMessage[name] = genSkeletonMessage(fd.Message())
		default:
			return defaultMessage
		}
		if fd.Cardinality() == protoreflect.Repeated {
			if fd.Kind() == protoreflect.StringKind {
				defaultMessage[name] = []string{""}
				continue
			}

			if fd.Kind() != protoreflect.MessageKind {
				defaultMessage[name] = []any{}
				continue
			}

			// it is possible to have this be recursive and include nested fields & messages.
			// but runs risk of endless loop on messages that are recursive.
			defaultMessage[name] = []any{genSkeletonMessage(fd.Message())}
		}
	}

	return defaultMessage
}

func makeCurlJson(in any) string {
	jsonString, _ := json.MarshalIndent(in, "", "  ")
	return string(jsonString)
}
