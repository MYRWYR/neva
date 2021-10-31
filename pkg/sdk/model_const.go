/*
 * API Title
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package sdk

type Const struct {
	Type string `json:"type,omitempty"`

	Value int `json:"value,omitempty"`
}

// AssertConstRequired checks if the required fields are not zero-ed
func AssertConstRequired(obj Const) error {
	return nil
}

// AssertRecurseConstRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of Const (e.g. [][]Const), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseConstRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aConst, ok := obj.(Const)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertConstRequired(aConst)
	})
}
