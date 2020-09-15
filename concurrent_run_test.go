package async

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func function1() {
	println("function1")
	fmt.Println("function1")
}

func functionWithError() {
	println("functionWithError")
	fmt.Println("functionWithError")

	// return errors.New("fake error")
	panic("fake error")
}

func TestConcurrentRun1Function(t *testing.T) {
	err := ConcurrentRun(function1)
	Convey("1 task", t, func() {
		So(err, ShouldBeNil)
	})
}

func TestConcurrentRun2Functions(t *testing.T) {
	err := ConcurrentRun(function1, function1)
	Convey("2 tasks", t, func() {
		So(err, ShouldBeNil)
	})
}

func TestConcurrentRun3Functions(t *testing.T) {
	err := ConcurrentRun(function1, function1, functionWithError)
	fmt.Println(err)
	// Convey("2 tasks", t, func() {
	// 	So(err, ShouldBeNil)
	// })
}
