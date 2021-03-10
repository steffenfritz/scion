// Package api provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}
