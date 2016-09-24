package importer

import "github.com/byorty/hardcore/types"

type PropertyImpl struct {
	kind types.ImportablePropertyKind
	protoKind types.ProtoKind
	importer  types.Importer
}

func NewProperty(kind types.ImportablePropertyKind, protoKind types.ProtoKind) PropertyImpl {
	return PropertyImpl{
		kind: kind,
		protoKind: protoKind,
	}
}

func (p PropertyImpl) GetKind() types.ImportablePropertyKind {
	return p.kind
}

func (p PropertyImpl) GetProtoKind() types.ProtoKind {
	return p.protoKind
}

func (p PropertyImpl) GetImporter() types.Importer {
	return p.importer
}
