package xmltogo_test

import (
	//"../clouseau/reckon"
	"./"
	"encoding/xml"
	"fmt"
	//"io/ioutil"
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	/*
		bytes, err := ioutil.ReadFile("./data.xml")
		if err != nil {
			panic(err)
		}
	*/
	bytes := []byte("<a><b value=\"1\"/><c value=\"2\"/><d value=\"3\"/></a>")
	a := A{}
	a2 := A2{}
	err := xml.Unmarshal(bytes, &a)
	if err != nil {
		panic(err)
	}
	err = xml.Unmarshal(bytes, &a2)
	if err != nil {
		panic(err)
	}
	list := append([]Inner{}, a2.B, a2.C, a2.D)

	fmt.Printf("%v, %v\n", len(a.List), a.List)
	fmt.Printf("%v\n", list)
	//out := fmt.Sprintf("%v", a.List)
	//reckon.That(out).Is.EqualTo("[]")
}

var marshallerOfA = xmltogo.InterfaceMarshaller{
	ChildMap: map[string]func() interface{}{
		"b": func() interface{} { return B{} },
		"c": func() interface{} { return C{} },
		"d": func() interface{} { return D{} },
	},
}

type A2 struct {
	B B `xml:"b"`
	C C `xml:"c"`
	D D `xml:"d"`
}

type A struct {
	List []Inner
}

func (a *A) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	//var end *xml.EndElement
	return marshallerOfA.MarshalChildren(d, start, func(item interface{}) error {
		ptr := reflect.ValueOf(item)
		for ptr.Kind() == reflect.Ptr || ptr.Kind() == reflect.Interface {
			ptr = ptr.Elem()
		}
		fmt.Printf("kind: %v\n", ptr.Kind())
		fmt.Printf("type: %v\n", ptr.Type().Name())
		innerItem, ok := ptr.Interface().(Inner)
		if ok {
			fmt.Printf("child: %v\n", innerItem.Str())
			fmt.Println("this is ok")
			a.List = append(a.List, innerItem)
		}
		return nil
	})
}

type Inner interface {
	Str() string
}

type B struct {
	Data string `xml:"value,attr"`
}

func (b B) Str() string {
	return "b:" + b.Data
}

type C struct {
	Data string `xml:"value,attr"`
}

func (c C) Str() string {
	return "c:" + c.Data
}

type D struct {
	Data string `xml:"value,attr"`
}

func (d D) Str() string {
	return "d:" + d.Data
}
