package common

type Container struct {
    Package string `xml:"package,attr"`
    ShortPackage string
    Import string
    Path string
}
