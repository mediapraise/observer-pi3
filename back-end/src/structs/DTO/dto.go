package DTO

import (
	"errors"
	"observer-go/src/structs/model"
	"reflect"
)

var modelToDTOMap = map[reflect.Type]reflect.Type{
	reflect.TypeOf((*model.User)(nil)).Elem():           reflect.TypeOf(UserDTO{}),
	reflect.TypeOf((*model.HistoryPayment)(nil)).Elem(): reflect.TypeOf(HistoryPaymentDTO{}),
	reflect.TypeOf((*model.Event)(nil)).Elem():         reflect.TypeOf(EventDTO{}),
	reflect.TypeOf((*model.Company)(nil)).Elem():       reflect.TypeOf(CompanyDTO{}),
	reflect.TypeOf((*model.Registration)(nil)).Elem(): reflect.TypeOf(RegistrationDTO{}),
	// Add other model-to-DTO mappings here
}

func ParseFromModel(model interface{}) (interface{}, error) {
	modelValue := reflect.ValueOf(model)

	// Ensure the model is a pointer to a struct
	if modelValue.Kind() != reflect.Ptr || modelValue.Elem().Kind() != reflect.Struct {
		return nil, errors.New("model must be a pointer to a struct")
	}

	modelType := modelValue.Elem().Type()

	// Find the corresponding DTO type
	dtoType, ok := modelToDTOMap[modelType]
	if !ok {
		return nil, errors.New("no corresponding DTO type found for the given model")
	}

	// Create a new instance of the DTO
	dtoValue := reflect.New(dtoType).Elem()

	// Populate the DTO fields with values from the model
	for i := 0; i < modelValue.Elem().NumField(); i++ {
		modelField := modelType.Field(i)
		modelFieldValue := modelValue.Elem().Field(i)

		// Find the corresponding field in the DTO
		dtoField, ok := dtoType.FieldByName(modelField.Name)
		if !ok {
			continue // Skip if the field does not exist in the DTO
		}

		// Ensure the field types are assignable
		dtoFieldValue := dtoValue.FieldByName(dtoField.Name)
		if dtoFieldValue.IsValid() && dtoFieldValue.CanSet() && modelFieldValue.Type().AssignableTo(dtoFieldValue.Type()) {
			dtoFieldValue.Set(modelFieldValue)
		}
	}

	// Return the populated DTO as an interface{}
	return dtoValue.Interface(), nil
}
