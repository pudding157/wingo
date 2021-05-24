package utils

import (
	"fmt"
	"reflect"
)

type Enum int

const (
	PHONE_NUMBER Enum = iota
	EMAIL
)

var enumStrings = []string{
	"phone-number",
	"email",
}

func EnumFromKey(key string) (*Enum, error) {

	var enum Enum
	for index, enumKey := range enumStrings {
		if key == enumKey {
			enum = Enum(index)
			return &enum, nil
		}
	}
	return nil, fmt.Errorf("%T  : invalid document type '%s'", enum, key)
}

func EnumFromIndex(i int) (*Enum, error) {
	var enum Enum

	if i >= len(enumStrings) {
		return nil, fmt.Errorf("%T  : defined index is out of range", enum)
	}

	return EnumFromKey(enumStrings[i])

}

func ParseEnum(e Enum, key string) Enum {
	return e.ParseKey(&key)
}

func (d Enum) Index() int {
	return int(d)
}

func (d Enum) String() string {
	return enumStrings[d.Index()]
}

func (d Enum) SetIndex(i int64) {

	v := reflect.ValueOf(d)
	v.SetInt(i)
}

func (d Enum) ParseKey(key *string) Enum {

	enum, err := EnumFromIndex(0)
	if err != nil {
		panic(fmt.Errorf("%T : no enum define", enum))
	}
	if key != nil {
		e, err := EnumFromKey(*key)
		if err != nil {
			return *enum
		}
		enum = e
	}
	return *enum
}
