package main

import (
	"fmt"
	"os"
	"testing"
)

type demoAddTest struct{
    arg1, arg2, expected int
}
type demoSubtractTest struct{
    arg1, arg2, expected int
}
var demoAddTests = []demoAddTest{
    {3,2,5},
    {4,6,10},
    {0,5,5},
    {-5,-2,-7},
}
var demoSubtractTests = []demoSubtractTest{
    {3,2,1},
    {4,6,-2},
    {0,5,-5},
    {-5,-2,-3},
}
func TestDemoAdd(t *testing.T){
    for _, test := range demoAddTests{
	if result := DemoAdd(test.arg1, test.arg2); result != test.expected{
	    t.Errorf("Result %d not equal to %d", result, test.expected)
	}
    }
}
func TestDemoSubtract(t *testing.T){
    for _, test := range demoSubtractTests{
	if result := DemoSubtract(test.arg1, test.arg2); result != test.expected{
	    t.Errorf("Result %d not equal to %d", result, test.expected)
	}
    }

}
func ExampleDemoAdd(){
    fmt.Println(DemoAdd(4,6))
    //Output: 10
}
func TestMain(m *testing.M){
    main()
    os.Exit(m.Run())
}
func BenchmarkAdd(b *testing.B){
    for i:= 0; i < b.N; i++{
	DemoAdd(5,4)
    }
}
func BenchmarkSubtract(b *testing.B){
    for i:= 0; i < b.N; i++{
	DemoSubtract(12,10)
    }
}
