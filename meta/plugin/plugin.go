package plugin

import "github.com/byorty/hardcore/meta"

type Plugin interface {
	Do(*meta.Configuration)
}
