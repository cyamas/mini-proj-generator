package main

import (
	"reflect"
	"testing"
)

var testFile = "test.json"
var testUpdateFile = "testupdate.json"

func TestReadAndUnmarshal(t *testing.T) {
	expected := FeaturesFile{
		Tool:      []string{"1", "2"},
		StdLib:    []string{"3", "4", "5"},
		ExtLib:    []string{"6"},
		DataType:  []string{"7", "8"},
		Algorithm: []string{},
	}
	var testFeatFile FeaturesFile
	ReadAndUnmarshal(testFile, &testFeatFile)
	if !reflect.DeepEqual(testFeatFile, expected) {
		t.Fail()
	}
}

func TestAdd(t *testing.T) {
	var testFeatFile FeaturesFile
	ReadAndUnmarshal(testFile, &testFeatFile)

	testFeatFile.Add("tool", "9")
	testFeatFile.Add("std", "10")
	testFeatFile.Add("ext", "11")
	testFeatFile.Add("algo", "12")
	testFeatFile.Add("dtype", "13")

	expected := FeaturesFile{
		Tool:      []string{"1", "2", "9"},
		StdLib:    []string{"3", "4", "5", "10"},
		ExtLib:    []string{"6", "11"},
		DataType:  []string{"7", "8", "13"},
		Algorithm: []string{"12"},
	}

	if !reflect.DeepEqual(testFeatFile, expected) {
		t.Fail()
	}
}

func TestRemove(t *testing.T) {
	var testFeatFile FeaturesFile
	ReadAndUnmarshal(testFile, &testFeatFile)

	testFeatFile.Remove("tool", "2")
	testFeatFile.Remove("std", "4")
	testFeatFile.Remove("ext", "6")
	testFeatFile.Remove("dtype", "8")

	expected := FeaturesFile{
		Tool:      []string{"1"},
		StdLib:    []string{"3", "5"},
		ExtLib:    []string{},
		DataType:  []string{"7"},
		Algorithm: []string{},
	}
	if !reflect.DeepEqual(expected, testFeatFile) {
		t.Fail()
	}
}

func TestGetEntryIndex(t *testing.T) {
	testSlice := []string{"a", "b", "c"}

	test1, err := getEntryIndex("a", testSlice)
	test2, err := getEntryIndex("b", testSlice)
	test3, err := getEntryIndex("c", testSlice)
	test4, err := getEntryIndex("d", testSlice)

	if test1 != 0 ||
		test2 != 1 ||
		test3 != 2 ||
		(test4 != -1 || err == nil) {
		t.Fail()
	}
}

func TestRemoveEntryByIndex(t *testing.T) {
	testSlice1 := []string{"a", "b", "c", "d", "e"}
	result1 := removeEntryByIndex(0, testSlice1)
	if len(result1) != 4 ||
		result1[0] != "e" ||
		result1[3] != "d" {
		t.Fail()
	}
	testSlice2 := []string{"a"}
	result2 := removeEntryByIndex(0, testSlice2)
	if len(result2) != 0 {
		t.Fail()
	}
}

func TestUpdateFile(t *testing.T) {
	testFeatFile := FeaturesFile{
		Tool:      []string{"z", "y"},
		StdLib:    []string{},
		ExtLib:    []string{"x"},
		Algorithm: []string{},
		DataType:  []string{"w", "v", "u"},
	}
	UpdateFile(&testFeatFile, testUpdateFile)

	var resultFeatFile FeaturesFile
	ReadAndUnmarshal(testUpdateFile, &resultFeatFile)
	if !reflect.DeepEqual(testFeatFile, resultFeatFile) {
		t.Fail()
	}
	var reverted FeaturesFile
	ReadAndUnmarshal(testFile, &reverted)
	UpdateFile(&reverted, testUpdateFile)

}
