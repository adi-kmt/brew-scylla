package utils

import (
	"reflect"
	"testing"
)

func TestSliceContains(t *testing.T) {
	// Test cases for integers
	intSlice := []int{1, 2, 3, 4, 5}
	if !SliceContains(intSlice, 3) {
		t.Errorf("Expected 3 to be found in slice, but it was not")
	}
	if SliceContains(intSlice, 6) {
		t.Errorf("Expected 6 not to be found in slice, but it was found")
	}

	// Test cases for strings
	strSlice := []string{"apple", "banana", "cherry", "date"}
	if !SliceContains(strSlice, "banana") {
		t.Errorf("Expected 'banana' to be found in slice, but it was not")
	}
	if SliceContains(strSlice, "orange") {
		t.Errorf("Expected 'orange' not to be found in slice, but it was found")
	}

	// Test cases for custom struct
	type person struct {
		Name string
		Age  int
	}
	personSlice := []person{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
		{Name: "Charlie", Age: 35},
	}
	if !SliceContains(personSlice, person{Name: "Bob", Age: 25}) {
		t.Errorf("Expected person{Name: 'Bob', Age: 25} to be found in slice, but it was not")
	}
	if SliceContains(personSlice, person{Name: "Eve", Age: 28}) {
		t.Errorf("Expected person{Name: 'Eve', Age: 28} not to be found in slice, but it was found")
	}
}

func TestGetEntityThatMatchesInSlice(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	type Product struct {
		ID    int
		Name  string
		Price float64
	}

	// Test case for struct Person
	people := []Person{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
		{Name: "Charlie", Age: 35},
	}
	personField := "Name"
	personValue := "Bob"

	person, err := GetEntityThatMatchesInSlice(people, personField, personValue)
	if err != nil {
		t.Errorf("Expected person with Name '%s' to be found in slice, but got error: %v", personValue, err)
	} else if person.Name != personValue {
		t.Errorf("Expected person with Name '%s', but got Name '%s'", personValue, person.Name)
	}

	// Test case for struct Product
	products := []Product{
		{ID: 1, Name: "Apple", Price: 1.25},
		{ID: 2, Name: "Banana", Price: 0.75},
		{ID: 3, Name: "Orange", Price: 1.50},
	}
	productField := "Name"
	productValue := "Banana"

	product, err := GetEntityThatMatchesInSlice(products, productField, productValue)
	if err != nil {
		t.Errorf("Expected product with Name '%s' to be found in slice, but got error: %v", productValue, err)
	} else if product.Name != productValue {
		t.Errorf("Expected product with Name '%s', but got Name '%s'", productValue, product.Name)
	}

	// Test case for entity not found
	nonexistentValue := "Nonexistent"
	_, err = GetEntityThatMatchesInSlice(products, productField, nonexistentValue)
	if err == nil {
		t.Errorf("Expected error for entity with Name '%s' not found, but got nil", nonexistentValue)
	} else if err.Error() != "entity not found" {
		t.Errorf("Expected error message 'entity not found', but got '%v'", err)
	}
}

func TestGetFieldSliceFromEntitySlice(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	type Product struct {
		ID    int
		Name  string
		Price float64
	}

	// Test case for struct Person
	people := []Person{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
	}
	personField := "Name"
	expectedPeopleNames := []string{"Alice", "Bob"}

	names := GetFieldSliceFromEntitySlice(people, personField)
	if !reflect.DeepEqual(names, expectedPeopleNames) {
		t.Errorf("Expected names %v, but got %v", expectedPeopleNames, names)
	}

	// Test case for struct Product
	products := []Product{
		{ID: 1, Name: "Apple", Price: 1.25},
		{ID: 2, Name: "Banana", Price: 0.75},
	}
	productField := "Price"
	expectedProductPrices := []string{"1.25", "0.75"}

	prices := GetFieldSliceFromEntitySlice(products, productField)
	if !reflect.DeepEqual(prices, expectedProductPrices) {
		t.Errorf("Expected prices %v, but got %v", expectedProductPrices, prices)
	}

	// Test case for non-existent field
	nonexistentField := "NonexistentField"
	emptyResult := GetFieldSliceFromEntitySlice(products, nonexistentField)
	if len(emptyResult) != 0 {
		t.Errorf("Expected empty result for non-existent field, but got %v", emptyResult)
	}
}

func TestSearchEntityFieldFromSlice(t *testing.T) {
	type Entity struct {
		Name string
	}
	// Test case 1: Matching "apple"
	entities := []Entity{
		{Name: "apple"},
		{Name: "orange"},
		{Name: "banana"},
		{Name: "pineapple"},
	}

	pattern := "aple"
	expected := []Entity{
		{Name: "apple"},
	}

	result, err := SearchEntityFieldFromSlice(entities, "Name", pattern)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(result) != len(expected) {
		t.Fatalf("Expected %d entities, got %d", len(expected), len(result))
	}

	for i, expectedEntity := range expected {
		if result[i].Name != expectedEntity.Name {
			t.Errorf("Expected entity %v at index %d, got %v", expectedEntity, i, result[i])
		}
	}

	// Test case 2: No match
	pattern = "xyz"
	expected = nil

	result, err = SearchEntityFieldFromSlice(entities, "Name", pattern)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if result != nil {
		t.Fatalf("Expected nil result, got %v", result)
	}
}
