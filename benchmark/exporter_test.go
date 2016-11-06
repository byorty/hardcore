package benchmark

import (
	"encoding/json"
	"github.com/byorty/hardcore/encoder"
	"github.com/byorty/hardcore/test/exporters"
	"github.com/byorty/hardcore/test/models"
	"testing"
	"time"
)

type User2 struct {
	Id           int64     `json:"id"`
	Email        string    `json:"email"`
	Role         int       `json:"role"`
	RegisterDate time.Time `json:"registerDate"`
}

var (
	now   = time.Now()
	role  = models.LoggedUserRole
	user1 = new(models.User).
		SetId(1).
		SetEmail("user@example.com").
		SetRole(role).
		SetRegisterDate(time.Now())
	user2 = &User2{1, "user@example.com", role.GetId(), now}
)

func BenchmarkJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.Marshal(user2)
	}
}

func BenchmarkExporter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		encoder.NewJson().One(exporters.NewUser(user1))
	}
}
