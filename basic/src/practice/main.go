package main
 
import (
    mypackages "practice/mypackage"
    "fmt"
)
 
func main() {

    //type : bool, string, int(8,16,32,64), uint(8,16,32,64), uintptr, byte(=uint8), rune(=uint32), float(32, 64), complex(64, 128)

    //variable 
    var s1, s2 string = "First", "Second"
    fmt.Printf("%s %s\n", s1, s2)

    //type inference
    s := "string"
    i := 10
    f := 11.11
    fmt.Printf("%s, %d, %f\n", s, i, f)

    //pointer
    var iptr *int
    iptr = &i
    *iptr = *iptr+2
    fmt.Printf("%d\n", *iptr)

    //function
    fmt.Printf("Hello, world. Sqrt(2) = %v\n", mypackages.Sqrt(2))
    add, sub :=  mypackages.AddSubtract(2,3)
    fmt.Printf("Add 2+3= %d, Subtract 2-3=%d\n", add, sub)

    //array 
    myarray := [3]int{1,2,3}
    fmt.Println(len(myarray), myarray[1:2])

    //map
    myMap := map[int]string{1: "first", 2:"second", 3: "third"}
    myMap[4] = "fourth"
    fmt.Println(myMap)

    //struct
    type myStruct struct{
        intField int
        stringField string
        sliceField []int
    }
    var t = myStruct{}
    t.intField = 2
    fmt.Println(t.intField)

    //if
    if x:=5; x==5 {
        fmt.Println("if:",x)
    }

    //switch는 자동 break

    //loop while은 없음
    for i:=1;i<=10;i++{
        fmt.Println("loop:", i)
    }
    myslice := []string{"one", "two", "three", "four", "five"}
    for i,item := range myslice{
        fmt.Println(i, item)
    }

}