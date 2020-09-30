package main

import (
	"encoding/hex"
	"github.com/safarifone/Iso8583/Iso8583"
	"log"
)

func main() {
	// Creates a new Iso8583 message
	msg := Iso8583.NewIso8583()

	//Set various fields as required
	msg.SetFieldValue(4,"4000")
	msg.SetFieldValue(70,"677")

	//You can even set a sub field on field 127 (this conforms to the Postilion switch standard)
	msg.SetSubFieldValue(127,2,"999999999")
	msg.SetSubFieldValue(127,6,"p1")
	msg.SetSubFieldValue(127,34,"98")

	//You can get the raw bytes of the message
	data := msg.ToMsg()

	//Print out the message in the default formatted message style
	log.Printf("\n%s",msg.ToString(""))

	//retrieve values from the message
	log.Printf("Data stored is %v\n", msg.GetFieldValue(4))

	//View the hex dump of the message also
	log.Println("hex dump data is ")
	log.Printf("\n%v",hex.Dump(data))
}
