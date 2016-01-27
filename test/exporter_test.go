package test

import (
	"github.com/byorty/hardcore/test/models"
	"time"
	"testing"
	"github.com/byorty/hardcore/test/exporters"
)

func TestExporter(t *testing.T) {
	now := time.Now()
	role := models.LoggedUserRole
	user := new(models.User).
		SetId(1).
		SetEmail("user@example.com").
		SetRole(&role).
		SetRegisterDate(now)

	userExporter := exporters.NewUser(user)
	exportable := userExporter.GetExportable()
	for i := 0;i < userExporter.Len();i++ {
		prop := userExporter.Get(i)
		t.Log(prop.GetName(), prop.GetValue(exportable))
		getter := user.Proto().GetByName(prop.GetName()).GetGetter()
		if getter(user) != prop.GetValue(exportable) {
			t.Fail()
		}
	}
}