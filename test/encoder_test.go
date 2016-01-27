package test

import (
	"testing"
	"github.com/byorty/hardcore/encoder"
	"github.com/byorty/hardcore/test/models"
	"time"
	"github.com/byorty/hardcore/test/exporters"
)

func TestEncoder(t *testing.T) {
	role := models.LoggedAdminRole
	user := new(models.User).
		SetId(1).
		SetEmail("user@example.com").
		SetRole(&role).
		SetRegisterDate(time.Now())
	json := encoder.NewJson()
	str := string(json.Encode(exporters.NewUser(user)))
	t.Log(str)
	t.Fail()
}
