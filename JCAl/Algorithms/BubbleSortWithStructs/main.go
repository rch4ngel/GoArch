package main

import (
	"bytes"
	"fmt"
)

type User struct {
	FirstName string
	LastName  string
}

func main() {
	xu := []User{
		{
			FirstName: "Zubei",
			LastName:  "Kibagami",
		},
		{
			FirstName: "Kent",
			LastName:  "Clark",
		},
	}

	fmt.Println("Sorting User List:", xu)
	BubbleSort(xu)
	fmt.Println("Sorted By First name:", xu)
}

func BubbleSort(xu []User) {
	for i := 0; i < len(xu); i++ {
		if !Swap(xu, i) {
			return
		}
	}
}

func Swap(xv []User, np int) bool {
	n := len(xv)
	fi := 0
	si := 1
	ds := false

	for si < n-np {
		ft := xv[fi]
		st := xv[si]

		if isGreaterThan(ft.FirstName[0:1], st.FirstName[0:1]) {
			ds = true
			xv[fi] = st
			xv[si] = ft
		}

		fi++
		si++
	}
	return ds
}

// isGreaterThan is a helper function that will be type agnostic, well for strings and integers.
func isGreaterThan(ft interface{}, st interface{}) bool {
	switch ft.(type) {
	case string:
		fs := fmt.Sprintf("%s", ft)
		ss := fmt.Sprintf("%s", st)

		fxb := []byte(fs)
		sxb := []byte(ss)

		if bytes.Compare(fxb, sxb) == 1 {
			fmt.Println(fxb, "is greater than", sxb)
			return true
		}
	case int:
		fs := fmt.Sprintf("%d", ft)
		ss := fmt.Sprintf("%d", st)

		fxb := []byte(fs)
		sxb := []byte(ss)

		if bytes.Compare(fxb, sxb) == 1 {
			fmt.Println(fxb, "is greater than", sxb)
			return true
		}
	default:
		return true
	}

	return false
}

func typeMultiplexer(t interface{}) interface{} {
	switch t.(type) {
	case string:
		return fmt.Sprintf("%T", t)
	case int:
		return fmt.Sprintf("%T", t)
	case User:
		return fmt.Sprintf("%T", t)
	}

	return t
}
