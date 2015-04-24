package meta

import (
	"testing"
)

func TestBuilder1(t *testing.T) {
	RegisterEnums(
		&SomeStatus{},
//		&SomeType{},
	)
	RegisterControllers(
		&SomeController{},
	)
	Build()
}
