package types

type ModelScanner interface {
    One(Model)
    All(Model)
}

type EnumScanner interface {
    One(Named)
//    All(Named)
}