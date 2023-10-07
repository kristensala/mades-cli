package commands

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"

	"github.com/tiaguinho/gosoap"
)

const (
    Ecp = "ecp"
    Edx = "edx"
)

type Mades struct {
    Ecp string
    Edx string
}

type SendMessageResponseResult struct {
    MessageId string `xml:"messageID"`
}

type SendMessageResponse struct {
    SendMessageResponseResult string `xml:"SendMessageResponseResult"`
}

var (
    sendMessageResponse SendMessageResponse
)

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
    wsdlUrl := m.getUrl(command)

    httpClient := &http.Client{}
    soapClient, err := gosoap.SoapClient(wsdlUrl, httpClient)
    if err != nil {
        log.Fatalf("Soap client error: %s", err)
        return
    }

    params := gosoap.Params{
        "message": gosoap.Params{
            "content": "message here", //todo: has to be byte array
            "baMessageID": "", // optional
        },
        "conversationID": "",
    }

    res, err := soapClient.Call("SendMessage", params)
    if err != nil {
        log.Fatalf("Could not set the message. Error: %s", err)
        return
    }

    res.Unmarshal(&sendMessageResponse)
    result := SendMessageResponseResult{}

    err = xml.Unmarshal([]byte(sendMessageResponse.SendMessageResponseResult), &result)
    if err != nil {
        log.Fatalf("Could not unmarchal response. Error: %s", err)
        return
    }

    fmt.Println(result)
}

