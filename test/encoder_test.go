package test

import (
	"fmt"
	"github.com/byorty/hardcore/encoder"
	"github.com/byorty/hardcore/test/exporters"
	"github.com/byorty/hardcore/test/models"
	"testing"
	"time"
)

func TestEncoder(t *testing.T) {
	now := time.Now()
	role := models.LoggedAdminRole
	user := new(models.User).
		SetId(1).
		SetEmail("user@example.com").
		SetRole(role).
		SetRegisterDate(now)
	json := encoder.NewJson()
	str := string(json.Encode(exporters.NewUser(user)))
	t.Log(str)
	if fmt.Sprintf(`{"id":1,"email":"user@example.com","role":2,"registerDate":"%s"}`, now.Format("2006-01-02 15:04:05")) != str {
		t.Fail()
	}
}
