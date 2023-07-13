package openapi_spec_code_diffs

import (
	"fmt"
	"github.com/pb33f/libopenapi"
	v3high "github.com/pb33f/libopenapi/datamodel/high/v3"
	"github.com/rs/zerolog/log"
	"os"
	"strings"
)

// oasModel Represents OpenAPI Spec file and libopenapi document model for the same
type oasModel struct {
	specFile string
	docModel *libopenapi.DocumentModel[v3high.Document]
}

// NewOASModel Returns a new oasModel
func NewOASModel(specFile string) *oasModel {
	return &oasModel{specFile: specFile, docModel: nil}
}

// LoadSpecModel Loads a spec model from the specified specFile
func (o *oasModel) LoadSpecModel() error {
	specFile, err := os.ReadFile(o.specFile)
	if err != nil {
		log.Error().Msg("Error reading spec file: " + err.Error())
		return err
	}

	document, err := libopenapi.NewDocument(specFile)
	if err != nil {
		log.Error().Msg("Cannot create document: " + err.Error())
		return err
	}

	docModel, errors := document.BuildV3Model()
	if len(errors) > 0 {
		for i := range errors {
			log.Error().Msg(errors[i].Error())
		}
		log.Error().Msg(fmt.Sprintf("cannot create v3 model from document: %d errors reported", len(errors)))
	}
	o.docModel = docModel
	return nil
}

// GetPathOps Gets paths defined in OpenAPI Specs file
func (o *oasModel) GetPathOps() []string {
	var ops []string
	for pathName, pathItem := range o.docModel.Model.Paths.PathItems {
		for op := range pathItem.GetOperations() {
			endpoint := strings.ToUpper(op) + " " + pathName
			ops = append(ops, endpoint)
		}
	}

	return ops
}
