// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by unknown module path version unknown version DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
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

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}