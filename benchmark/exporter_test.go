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
	"unicode/utf8"
)

type User2 struct {
	Id int64 `json:"id"`
	Email string `json:"email"`
	Role int `json:"role"`
	RegisterDate time.Time `json:"registerDate"`
}

type Marshaller struct {}

func (m Marshaller) MarshalInt(i int) string {
	return strconv.FormatInt(int64(i), 10)
}

func (m Marshaller) MarshalInt64(i int64) string {
	return strconv.FormatInt(i, 10)
}

func (m Marshaller) MarshalTime(t time.Time) string {
	return t.Format("2006-01-02T15:04:05.999999-07:00")
}

type UserMarshaller struct {
	Marshaller
	user *models.User
}

func NewUserMarshaller(user *models.User) json.Marshaler {
	m := new(UserMarshaller)
	m.user = user
	return m
}

func (u UserMarshaller) MarshalJSON() ([]byte, error) {
	j := `{"id":` + u.MarshalInt64(u.user.GetId()) +
	`,"email":"` + u.user.GetEmail() +
	`","role":` + u.MarshalInt(u.user.GetRoleId()) +
	`,"registerDate":"` + u.MarshalTime(u.user.GetRegisterDate()) + `"}`
	return []byte(j), nil
}

type Encoder struct {}

func NewEncoder() *Encoder {
	return new(Encoder)
}

func (e *Encoder) Encode(exp types.Exporter) []byte {
	expLen := exp.Len()
	lastIndex := expLen - 1
	buf := new(bytes.Buffer)
	buf.Write(startBrace)

	for i := 0;i < expLen;i++ {
		prop := exp.Get(i)

		buf.Write(quotes)
		buf.WriteString(prop.GetName())
		buf.Write(quotes)
		buf.Write(colon)
		switch value := prop.GetValue().(type) {
		case string:
			buf.Write(quotes)
			start := 0
			for i := 0; i < len(value); {
				if b := value[i]; b < utf8.RuneSelf {
					if 0x20 <= b && b != '\\' && b != '"' && b != '<' && b != '>' && b != '&' {
						i++
						continue
					}
					if start < i {
						buf.WriteString(value[start:i])
					}
					switch b {
					case '\\', '"':
						buf.WriteByte('\\')
						buf.WriteByte(b)
					case '\n':
						buf.WriteByte('\\')
						buf.WriteByte('n')
					case '\r':
						buf.WriteByte('\\')
						buf.WriteByte('r')
					case '\t':
						buf.WriteByte('\\')
						buf.WriteByte('t')
					default:
						// This encodes bytes < 0x20 except for \n and \r,
						// as well as <, > and &. The latter are escaped because they
						// can lead to security holes when user-controlled strings
						// are rendered into JSON and served to some browsers.
						buf.WriteString(`\u00`)
						buf.WriteByte(hex[b>>4])
						buf.WriteByte(hex[b&0xF])
					}
					i++
					start = i
					continue
				}
				c, size := utf8.DecodeRuneInString(value[i:])
				if c == utf8.RuneError && size == 1 {
					if start < i {
						buf.WriteString(value[start:i])
					}
					buf.WriteString(`\ufffd`)
					i += size
					start = i
					continue
				}
				// U+2028 is LINE SEPARATOR.
				// U+2029 is PARAGRAPH SEPARATOR.
				// They are both technically valid characters in JSON strings,
				// but don't work in JSONP, which has to be evaluated as JavaScript,
				// and can lead to security holes there. It is valid JSON to
				// escape them, so we do so unconditionally.
				// See http://timelessrepo.com/json-isnt-a-javascript-subset for discussion.
				if c == '\u2028' || c == '\u2029' {
					if start < i {
						buf.WriteString(value[start:i])
					}
					buf.WriteString(`\u202`)
					buf.WriteByte(hex[c&0xF])
					i += size
					start = i
					continue
				}
				i += size
			}
			if start < len(value) {
				buf.WriteString(value[start:])
			}
			buf.Write(quotes)
		case time.Time:
			buf.Write(quotes)
			buf.WriteString(value.Format("2006-01-02T15:04:05.999999-07:00"))
			buf.Write(quotes)
		case types.IntEnum:
			buf.WriteString(strconv.Itoa(value.GetId()))
		case int64:
			buf.WriteString(strconv.FormatInt(value, 10))
		default:
			buf.WriteString(fmt.Sprintf("%v", value))
		}

		if i < lastIndex {
			buf.Write(comma)
		}
	}

	buf.Write(endBrace)
	return buf.Bytes()
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

var (
	startBrace = []byte("{")
	endBrace = []byte("}")
	quotes = []byte("\"")
	colon = []byte(":")
	comma = []byte(",")
	hex = "0123456789abcdef"
)

func BenchmarkExporter(b *testing.B) {
	for i := 0;i < b.N;i++ {
		NewEncoder().Encode(exporters.NewUser(user1))
	}
}

func BenchmarkMarshaller(b *testing.B) {
	for i := 0;i < b.N;i++ {
		json.Marshal(NewUserMarshaller(user1))
	}
}
