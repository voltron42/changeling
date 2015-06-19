package xmltogo_test

import (
	"./"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"testing"
)

func Test(t *testing.T) {
	bytes, err := ioutil.ReadFile("./data.xml")
	if err != nil {
		panic(err)
	}
	a := A{}
	err = xml.Unmarshal(bytes, &a)
	if err != nil {
		panic(err)
	}
	for _, item := range a.List {
		fmt.Printf("item: %v\n", item.String())
	}
}

var marshallerOfA = xmltogo.InterfaceMarshaller{
	ChildMap: map[string]func() interface{}{
		"b": func() interface{} { return B{} },
		"c": func() interface{} { return C{} },
		"d": func() interface{} { return D{} },
	},
}

type A struct {
	List []Inner
}

func (a *A) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	//var end *xml.EndElement
	return marshallerOfA.MarshalChildren(d, start, func(item interface{}) error {
		innerItem, ok := item.(Inner)
		if ok {
			a.List = append(a.List, innerItem)
		}
		return nil
	})
}

type Inner interface {
	String() string
}

type B struct {
	Data string `xml:",chardata"`
}

func (b *B) String() string {
	return "b:" + b.Data
}

type C struct {
	Data string `xml:",chardata"`
}

func (c *C) String() string {
	return "c:" + c.Data
}

type D struct {
	Data string `xml:",chardata"`
}

func (d *D) String() string {
	return "d:" + d.Data
}
