package types

type Configuration interface {
    SetContainers([]Container)
    GetContainers() []Container
}
