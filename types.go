package hardcore

const (
	EOL = '\n'
)

type Enum interface {
	GetId() int
	GetName() string
}
