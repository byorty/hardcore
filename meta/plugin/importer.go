package plugin

type Importer struct {
	imports []string
}

func (i *Importer) addImport(newImport string) {
	hasImport := false
	for _, existsImport := range i.imports {
		if existsImport == newImport {
			hasImport = true
			break
		}
	}
	if !hasImport {
		i.imports = append(i.imports, newImport)
	}
}
