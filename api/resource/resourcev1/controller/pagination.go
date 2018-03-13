package controller

import (
	"encoding/json"
	"reflect"
)

type GenericPaginatedResourcesHandler struct {
	resourceType reflect.Type
}

func NewRCPaginatedResources(resource interface{}) GenericPaginatedResourcesHandler {
	return GenericPaginatedResourcesHandler{
		resourceType: reflect.TypeOf(resource),
	}
}

func (pr GenericPaginatedResourcesHandler) Resources(bytes []byte, curURL string) ([]interface{}, string, error) {
	var paginatedResources = struct {
		NextUrl        string          `json:"next_url"`
		ResourcesBytes json.RawMessage `json:"resources"`
	}{}

	err := json.Unmarshal(bytes, &paginatedResources)

	slicePtr := reflect.New(reflect.SliceOf(pr.resourceType))
	err = json.Unmarshal([]byte(paginatedResources.ResourcesBytes), slicePtr.Interface())
	slice := reflect.Indirect(slicePtr)

	contents := make([]interface{}, 0, slice.Len())
	for i := 0; i < slice.Len(); i++ {
		contents = append(contents, slice.Index(i).Interface())
	}

	return contents, paginatedResources.NextUrl, err
}
