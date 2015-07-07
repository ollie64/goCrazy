
// Go has strange behavior -- in one case it doesn't support generics
// In over case it can't convert []AnyType to []interface[] even pointers
// So it don't support sub-typing that makes generic coding really hard


package main

import (
	"fmt"
	"io/ioutil"
	"strings"
//	"container/list"
	"./Vector"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func StringsToArray(arr0 []string) []interface{} {
	res := make([]interface{}, len(arr0))
	for i,x := range arr0 {
		res[i] = x
	}
	return res
}

func ArrayJoin(arr []interface{}, sep interface{}) []interface{} {
	if len(arr) == 0 {
		return []interface{}{}
	} else {
		var res []interface{}
		res = append(res, arr[0])
		for _, z := range arr[1:] {
			res = append(res, sep)
			res = append(res, z)
		}
		return res
	}
}





func main() {


//	dbg0 := StringsToArray(strings.Split("((dsad))", "("))
//
//	dbg1 := ArrayJoin(dbg0, "[")
//
//	fmt.Println("dbg =")
//	for _,x := range dbg1 {
//		fmt.Println(x)
//	}
//	fmt.Println("----------")

	dat, err := ioutil.ReadFile("/Users/obaskakov/IdeaProjects/goCrazy/code.scm")
	check(err)
	code := string(dat)
	lines := strings.Split(code, "\n")


	lines1 := StringsToArray(lines)


//	fmt.Println(len(lines1))

	tokens := Vector.FromArray(lines1).FlatMap(func (x interface{}) []interface{} {
		lll := x.(string)
		tmp := strings.Split(lll, "\t")
		return StringsToArray(tmp)
	}).FlatMap(func (x interface{}) []interface{} {
		return StringsToArray(strings.Split(x.(string), " "))
	}).FlatMap(func (x interface{}) []interface{} {
		tmp := StringsToArray(strings.Split(x.(string), "("))
		return ArrayJoin(tmp, "(")
	}).FlatMap(func (x interface{}) []interface{} {
		tmp := StringsToArray(strings.Split(x.(string), ")"))
		return ArrayJoin(tmp, ")")
	}).Filter(func (x interface{}) bool {
		return len(x.(string)) != 0
	}).ToArray()


	fmt.Println("tokens:")
	for _,x := range tokens {
		fmt.Println(x)
	}
	fmt.Println("----")

//	fmt.Println(strings.Join(append(str0, tokens), "\n"))



//	tmp2 := flatMap(tmp1, func(x Object) []Object {return []Object{1,2,3, x}})

//	xxx := []string{"1", "2222"}

//	tmp2 := Map(func(x string) int {return len(x)}, xxx)

//	fmt.Println(tmp2)

	// Create a new list and put some numbers in it.
	ll := Vector.New()
	ll.PushBack("dasdsa")
	ll.PushBack("das")


//	fmt.Println("----")
//	fmt.Println(e4.Value)

	tmp2 := ll.
		Map(func (x interface{}) interface{} {
			return len(x.(string))
		}).Map(func (x interface{}) interface{} {
			return x.(int) + 1
		})

	fmt.Println("----")
	// Iterate through list and print its contents.
	for e := tmp2.Front(); e != nil; e = e.Next() {
		fmt.Println("val =", e.Value)
	}

	fmt.Println("----")
	fmt.Println(tmp2.ToArray())

}
