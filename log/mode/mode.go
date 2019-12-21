package mode

import (
	"fmt"
)

type Mode uint8

const (
	Development = iota + 1
	Production
)

const modeName = "developmentproduction"

var modeMap = map[Mode]string{
	1: modeName[0:11],
	2: modeName[11:21],
}

func (m Mode) String() string {
	if str, ok := modeMap[m]; ok {
		return str
	}
	return fmt.Sprintf("Mode(%d)", m)
}

func FromString(value string) (m Mode, err error) {
	switch value {
	case "development":
		m = Development
	case "production":
		m = Production
	default:
		err = fmt.Errorf("unknow mode %v", value)
	}
	return
}
