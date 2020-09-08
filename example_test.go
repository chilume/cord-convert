package convert_test

import (
	"context"
	"fmt"

	convert "github.com/chilume/cord-convert"
)

func ExampleLatLnglToBNG() {

	c := convert.NewClient(nil)
	resp, _, _ := c.LatLnglToBNG(context.Background(), -5.55, -1.54)
	fmt.Println(*resp)

	// Output: {{-1 24 32} 451030.444044407 -1.54 false {-5 0 33} -6.14106483570885e+06 -5.55}
}

func ExampleBNGToLatLng() {
	c := convert.NewClient(nil)

	resp, _, _ := c.BNGToLatLng(context.Background(), 429157, 623009)
	fmt.Println(*resp)

	// Output: {{-1 24.028476096768 32} 429157 -1.54000791002688 false {55 59.99859710664 29} 623009 55.4999996103074}
}
