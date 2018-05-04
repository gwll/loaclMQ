package loaclMQ

import (
	"fmt"
	"testing"
	//"time"
)

func Test_pub(t *testing.T) {

	Pub("t1", "string")

}
func Test_sub(t *testing.T) {
	t1 := Sub("t1")

	fmt.Println(t1)

}
