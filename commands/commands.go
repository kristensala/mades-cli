package commands

import "fmt"

const (
    Ecp = "ecp"
    Edx = "edx"
)

type Mades struct {
    Ecp string
    Edx string
}

func (m Mades) getUrl(command string) string {
    switch command {
    case Edx:
        return m.Edx
    case Ecp:
        return m.Ecp
    default:
        panic("Could not find url")
    }
}

func (m Mades) SendMessage(command string, message string) {
    //wsdlUrl := w.getUrl(command)

    fmt.Println(message)

}

