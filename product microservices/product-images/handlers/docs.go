package handlers

import (
	"os"
)

// NOTE: Types defined here are purely for documentation Purposes
// these types are not used by handlers

// Generic Error message returned as string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// No content is returned by this API endpoint
// swagger:response noContentResponse
type noContentResponseWrapper struct {
}

// FileContent is returned by this API endpoint
// swagger:response fileResponse
type fileResponseWrapper struct {
	file *os.File
}

//swagger:parameters SaveImage GetImage
type productImageParamsWrapper struct {
	// id for which the operation relates
	// in: path
	// required: true
	ID string `json:"id"`
	// filename for image in basepath
	// in: path
	//required: true
	Filename string `json:"filename"`
}
