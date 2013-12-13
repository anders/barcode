package codabar

import (
	"fmt"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/utils"
	"regexp"
)

var encodingTable map[rune][]bool = map[rune][]bool{
	'0': []bool{true, false, true, false, true, false, false, true, true},
	'1': []bool{true, false, true, false, true, true, false, false, true},
	'2': []bool{true, false, true, false, false, true, false, true, true},
	'3': []bool{true, true, false, false, true, false, true, false, true},
	'4': []bool{true, false, true, true, false, true, false, false, true},
	'5': []bool{true, true, false, true, false, true, false, false, true},
	'6': []bool{true, false, false, true, false, true, false, true, true},
	'7': []bool{true, false, false, true, false, true, true, false, true},
	'8': []bool{true, false, false, true, true, false, true, false, true},
	'9': []bool{true, true, false, true, false, false, true, false, true},
	'-': []bool{true, false, true, false, false, true, true, false, true},
	'$': []bool{true, false, true, true, false, false, true, false, true},
	':': []bool{true, true, false, true, false, true, true, false, true, true},
	'/': []bool{true, true, false, true, true, false, true, false, true, true},
	'.': []bool{true, true, false, true, true, false, true, true, false, true},
	'+': []bool{true, false, true, true, false, false, true, true, false, false, true, true},
	'A': []bool{true, false, true, true, false, false, true, false, false, true},
	'B': []bool{true, false, true, false, false, true, false, false, true, true},
	'C': []bool{true, false, false, true, false, false, true, false, true, true},
	'D': []bool{true, false, true, false, false, true, true, false, false, true},
}

func Encode(content string) (barcode.Barcode, error) {
	checkValid, _ := regexp.Compile(`[ABCD][0123456789\-\$\:/\.\+]*[ABCD]$`)
	if content == "!" || checkValid.ReplaceAllString(content, "!") != "!" {
		return nil, fmt.Errorf("can not encode \"%s\"", content)
	}
	resBits := new(utils.BitList)
	for i, r := range content {
		if i > 0 {
			resBits.AddBit(false)
		}
		bits, ok := encodingTable[r]
		if !ok {
			return nil, fmt.Errorf("can not encode \"%s\"", content)
		}
		for _, bit := range bits {
			resBits.AddBit(bit)
		}
	}
	return &codabarcode{resBits, content}, nil
}
