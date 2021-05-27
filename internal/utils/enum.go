package utils

import (
	"fmt"
	"reflect"
)

type Enum int

const (
	PHONE_NUMBER Enum = 0 // otp type
	EMAIL        Enum = 1 // otp type

	MEMBER Enum = 0 // user status
	VIP    Enum = 1 // user status
)

var otpTypeStrings = []string{
	"phone-number", // otp type
	"email",        // otp type

}

var userStatusString = []string{
	"member", // otp type
	"vip",    // otp type
}

func GetEnumArray(arrName string) []string {
	switch n := arrName; n {
	case "userStatus":
		return userStatusString
	case "otp":
		return otpTypeStrings
	default:
		return nil
	}
}
func EnumFromKey(key string, eString []string) (*Enum, error) {

	var enum Enum
	for index, enumKey := range eString {
		if key == enumKey {
			enum = Enum(index)
			return &enum, nil
		}
	}
	return nil, fmt.Errorf("%T  : invalid document type '%s'", enum, key)
}

func EnumFromIndex(i int, eString []string) (*Enum, error) {
	var enum Enum

	if i >= len(eString) {
		return nil, fmt.Errorf("%T  : defined index is out of range", enum)
	}

	return EnumFromKey(eString[i], eString)

}

func ParseEnum(e Enum, key string, eString []string) Enum {
	return e.ParseKey(&key, eString)
}

func (d Enum) Index() int {
	return int(d)
}

func (d Enum) String(eString []string) string {
	return eString[d.Index()]
}

func (d Enum) SetIndex(i int64) {

	v := reflect.ValueOf(d)
	v.SetInt(i)
}

func (d Enum) ParseKey(key *string, eString []string) Enum {

	enum, err := EnumFromIndex(0, eString)
	if err != nil {
		panic(fmt.Errorf("%T : no enum define", enum))
	}
	if key != nil {
		e, err := EnumFromKey(*key, eString)
		if err != nil {
			return *enum
		}
		enum = e
	}
	return *enum
}
