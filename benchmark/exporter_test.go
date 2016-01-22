package benchmark

import (
	"testing"
	"github.com/byorty/hardcore/test/models"
	"time"
	"bytes"
	"github.com/byorty/hardcore/test/exporters"
	"github.com/byorty/hardcore/types"
	"fmt"
	"encoding/json"
	"strconv"
)

type User2 struct {
	Id int64 `json:"id"`
	Email string `json:"email"`
	Role int `json:"role"`
	RegisterDate time.Time `json:"registerDate"`
}

var (
	now = time.Now()
	role = models.LoggedUserRole
	user1 = new(models.User).
		SetId(1).
		SetEmail("user@example.com").
		SetRole(&role).
		SetRegisterDate(time.Now())
	user2 = &User2{1, "user@example.com", role.GetId(), now}
)

func BenchmarkJson(b *testing.B) {
	for i := 0;i < b.N;i++ {
		json.Marshal(user2)
	}
}

func BenchmarkExporter(b *testing.B) {
	for i := 0;i < b.N;i++ {
		buf := new(bytes.Buffer)
		exp := exporters.NewUser().Export(user1)
		expLen := exp.Len()
		lastIndex := expLen - 1

		buf.WriteByte('{')

		for i := 0;i < expLen;i++ {
			prop := exp.Get(i)

			buf.WriteByte('"')
			buf.WriteString(prop.GetName())
			buf.WriteByte('"')
			buf.WriteByte(':')
			switch value := prop.GetValue().(type) {
			case string:
				buf.WriteByte('"')
				buf.WriteString(value)
				buf.WriteByte('"')
			case time.Time:
				buf.WriteByte('"')
				buf.WriteString(value.Format("2006-01-02T15:04:05.999999-07:00"))
				buf.WriteByte('"')
			case types.IntEnum:
				buf.WriteString(strconv.Itoa(value.GetId()))
			case int64:
				buf.WriteString(strconv.FormatInt(value, 10))
			default:
				buf.WriteString(fmt.Sprintf("%v", value))
			}

			if i < lastIndex {
				buf.WriteByte(',')
			}
		}

		buf.WriteByte('}')
		buf.Bytes()
	}
}
