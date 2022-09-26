package usecase

import (
	"fmt"
	"testing"
	"time"
)

func TestDate(tt *testing.T) {
	layout := "02/01/2006 15:04"
	str := "21/09/2022 22:07"
	t, err := time.Parse(layout, str)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t.String())
	fmt.Println(int(t.Month()))
}
