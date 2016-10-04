package plugin

import (
	"github.com/byorty/hardcore/meta/types"
	"os/exec"
	"strings"
)

type Formatter struct {
}

func (f *Formatter) Do(env types.Environment) {
	config := env.GetConfiguration()
	var packeges []string
	for _, container := range config.GetContainers() {
		packeges = append(packeges, container.GetImport())
	}
	if len(packeges) > 0 {
		args := append([]string{"fmt"}, packeges...)
		cmd := exec.Command("go", args...)
		cmd.Run()
		env.GetLogger().Debug("formatter format packages: %s", strings.Join(packeges, ", "))
	}
}
