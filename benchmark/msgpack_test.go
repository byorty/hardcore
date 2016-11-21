package benchmark

import (
	"bytes"
	"github.com/byorty/hardcore/encoder"
	"github.com/byorty/hardcore/test/exporters"
	"github.com/byorty/hardcore/test/models"
	msgpack "github.com/msgpack/msgpack-go"
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
	msgNow   = time.Now()
	msgUser1 = new(models.User).
			SetId(1).
			SetEmail("user@example.com")
	msgUser2 = map[string]interface{}{
		"id":    1,
		"email": "user@example.com",
	}
)

func BenchmarkMsgpack(b *testing.B) {
	buf := new(bytes.Buffer)
	for i := 0; i < b.N; i++ {
		msgpack.Pack(buf, msgUser2)
	}
}

func BenchmarkMsgpackExporter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		encoder.NewMsgpack().Encode(exporters.NewUser(msgUser1))
	}
}
