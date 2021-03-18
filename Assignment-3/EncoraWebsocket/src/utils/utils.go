package utils

import (
    INC "EncoraWebsocket/src/messages"
	"encoding/xml"
    "fmt"
    "io/ioutil"
    "os"
)

func ParsingXml() interface{} {

	// Open our xmlFile
    xmlFile, err := os.Open("users.xml")
    // if we os.Open returns an error then handle it
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println("Successfully Opened users.xml")
    // defer the closing of our xmlFile so that we can parse it later on
    defer xmlFile.Close()

    // read our opened xmlFile as a byte array.
    byteValue, _ := ioutil.ReadAll(xmlFile)

    // we initialize our Users array
    var users INC.Users
    // we unmarshal our byteArray which contains our
    // xmlFiles content into 'users' which we defined above
    xml.Unmarshal(byteValue, &users)

    return users

   

}

