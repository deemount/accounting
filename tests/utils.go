package tests

import (
	"fmt"
	"log"
	"math"
	"reflect"
	"strconv"
	"unsafe"
)

// GetSliceHeader inspecting the header values of each slice
func GetSliceHeader(slice *interface{}) string {
	sh := (*reflect.SliceHeader)(unsafe.Pointer(slice))
	return fmt.Sprintf("%+v", sh)
}

// GetSize obtains the size of v using reflect and then, for the supported types
// (slices, maps, strings, and structs), it computes the memory required by the content
// stored in them. You would need to add here other types that you need to support.
//
// There are a few details to work out:
// - Private fields are not counted.
// - For structs we are double-counting the basic types.
//
// For number two, filter them out before doing the recursive call when handling structs,
// also check the kinds in the documentation for the reflect package.
func GetSize(v interface{}) int {

	size := int(reflect.TypeOf(v).Size())

	switch reflect.TypeOf(v).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(v)
		for i := 0; i < s.Len(); i++ {
			size += GetSize(s.Index(i).Interface())
		}

	case reflect.Map:

		s := reflect.ValueOf(v)
		keys := s.MapKeys()
		size += int(float64(len(keys)) * 10.79) // approximation from https://golang.org/src/runtime/hashmap.go
		for i := range keys {
			size += GetSize(keys[i].Interface()) + GetSize(s.MapIndex(keys[i]).Interface())
		}

	case reflect.String:
		size += reflect.ValueOf(v).Len()

	case reflect.Struct:
		s := reflect.ValueOf(v)
		for i := 0; i < s.NumField(); i++ {
			if s.Field(i).CanInterface() {
				size += GetSize(s.Field(i).Interface())
			}
		}

	}
	return size
}

// ShowSize ...
func ShowSize(result interface{}) {

	suffixes[0] = "Bytes"
	suffixes[1] = "KB"
	suffixes[2] = "MB"
	suffixes[3] = "GB"
	suffixes[4] = "TB"

	size, _ := strconv.ParseFloat(strconv.Itoa(GetSize(result)), 64)
	base := math.Log(size) / math.Log(1024)
	getSize := Round(math.Pow(1024, base-math.Floor(base)), .5, 2)
	getSuffix := suffixes[int(math.Floor(base))]
	log.Printf(strconv.FormatFloat(getSize, 'f', -1, 64) + " " + string(getSuffix))

}

// Round ...
func Round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}
