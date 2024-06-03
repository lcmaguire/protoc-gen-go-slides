package main

import (
	"flag"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

func main() {

	var flags flag.FlagSet

	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {
		// this enables optional fields to be supported.
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

		// create a new generated file.
		generatedFile := gen.NewGeneratedFile("service.slide", "") // Second paramater is the go import path, normally "" or ".".

		// can print to file through the P method.
		generatedFile.P("// ")

		// iterate over the files included in the request.
		for _, file := range gen.Files {
			// iterate over the services within a file
			for _, service := range file.Services {
				// iterate over a services methods
				for _, method := range service.Methods {
					// with your method do something.
					generatedFile.P(method.Comments.Leading.String())
					//method.Comments
				}
			}
		}

		return nil
	})
}
