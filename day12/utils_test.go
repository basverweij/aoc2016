package main

import "testing"
import "reflect"

func assertNil(t *testing.T, actual interface{}, msg string) {
	if !nilCheck(actual) {
		t.Fatalf("%s: expected nil, but got '%v' (%T)", msg, actual, actual)
	}
}

func assertNotNil(t *testing.T, actual interface{}, msg string) {
	if nilCheck(actual) {
		t.Fatalf("%s: expected not nil, but got '%v' (%T)", msg, actual, actual)
	}
}

func nilCheck(actual interface{}) bool {
	return actual == nil || reflect.ValueOf(actual).IsNil()
}

func assertEquals(t *testing.T, expected, actual interface{}, msg string) {
	if !equalCheck(expected, actual) {
		t.Fatalf("%s: expected '%v' (%T), but got '%v' (%T)", msg, expected, expected, actual, actual)
	}
}

func assertNotEquals(t *testing.T, notExpected, actual interface{}, msg string) {
	if equalCheck(notExpected, actual) {
		t.Fatalf("%s: did not expect '%v'", msg, notExpected)
	}
}

func equalCheck(expected, actual interface{}) bool {
	return expected == actual
}
