package xmltogo

import (
	"encoding/xml"
	"errors"
	"fmt"
	"reflect"
)

type InterfaceMarshaller struct {
	ChildMap map[string]func() interface{}
}

func (i *InterfaceMarshaller) MarshalChildren(d *xml.Decoder, start xml.StartElement, appender func(item interface{}) error) error {
	for true {
		t, err := d.Token()
		if err != nil {
			return err
		}
		if reflect.DeepEqual(t, start.End()) {
			break
		}
		if reflect.DeepEqual(reflect.TypeOf(t), reflect.TypeOf(start)) {
			bookend, ok := t.(xml.StartElement)
			if ok {
				itemType, ok := i.ChildMap[bookend.Name.Local]
				if !ok {
					return errors.New(fmt.Sprintf("Token name not recognized: %v", bookend.Name.Local))
				}
				item := itemType()
				fmt.Printf("type: %v\n", reflect.TypeOf(item).Name())
				fmt.Printf("kind: %v\n", reflect.TypeOf(item).Kind())
				fmt.Printf("ptr kind: %v\n", reflect.ValueOf(&item).Kind())
				fmt.Printf("start: %V\n", bookend)
				err := d.DecodeElement(&item, &bookend)
				if err != nil {
					return err
				}
				fmt.Printf("item: %v\n", item)
				err = appender(&item)
				if err != nil {
					return err
				}
			}
		}
	}
	fmt.Println("no errors")
	return nil
}
