package plugin

type PackageImporter struct {
	imports []string
}

func (p *PackageImporter) addImport(newImport string) {
	hasImport := false
	for _, existsImport := range p.imports {
		if existsImport == newImport {
			hasImport = true
			break
		}
	}
	if !hasImport {
		p.imports = append(p.imports, newImport)
	}
}
