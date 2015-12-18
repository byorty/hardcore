package types

type StraightMappingModelScanner interface {
    One(StraightMappingModel)
    All(StraightMappingModel)
}

type EnumScanner interface {
    One(Named)
//    All(Named)
}