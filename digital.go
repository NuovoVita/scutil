package scutil

import (
	"fmt"
	"strconv"
	"strings"
)

func Grading(money float64, seq byte, grading int) string {
	strMoney := strconv.FormatFloat(money, 'f', -1, 64)
	arrayMoney := strings.Split(strMoney, ".")

	length := len(arrayMoney[0])
	if length <= grading {
		return strMoney
	}

	sInt := make([]byte, 0, 100)
	for i, v := range arrayMoney[0] {
		sInt = append(sInt, byte(v))
		if (length-i-1) > 0 && (length-i-1)%grading == 0 {
			sInt = append(sInt, seq)
		}
	}
	if len(arrayMoney) == 1 {
		return string(sInt)
	}
	return fmt.Sprintf("%s.%s", string(sInt), arrayMoney[1])
}

func CancelGrading(money string, seq byte) float64 {
	newMoney := strings.Replace(money, string(seq), "", -1)
	float, err := strconv.ParseFloat(newMoney, 64)
	if err != nil {
		return 0
	}
	return float
}