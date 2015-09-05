package main

import "testing"

func TestFooBarQix(t *testing.T) {
	if foobarqix(1) != "1" {
		t.Fail()
	}

	if foobarqix(2) != "2" {
		t.Fail()
	}

	if foobarqix(3) != "FooFoo" {
		t.Fail()
	}

	if foobarqix(5) != "BarBar" {
		t.Fail()
	}

	if foobarqix(7) != "QixQix" {
		t.Fail()
	}

	if foobarqix(9) != "Foo" {
		t.Fail()
	}

	if foobarqix(10) != "Bar" {
		t.Fail()
	}

	if foobarqix(14) != "Qix" {
		t.Fail()
	}

	if foobarqix(15) != "FooBarBar" {
		t.Fail()
	}
}

func ExampleFooBarQix() {
	main()
	// Output: 1
	// 2
	// FooFoo
	// 4
	// BarBar
	// Foo
	// QixQix
	// 8
	// Foo
	// Bar
	// 11
	// Foo
	// Foo
	// Qix
	// FooBarBar
	// 16
	// Qix
	// Foo
	// 19
	// Bar
	// FooQix
	// 22
	// Foo
	// Foo
	// BarBar
	// 26
	// FooQix
	// Qix
	// 29
	// FooBarFoo
	// Foo
	// Foo
	// FooFooFoo
	// Foo
	// BarQixFooBar
	// FooFoo
	// FooQix
	// Foo
	// FooFoo
	// Bar
	// 41
	// FooQix
	// Foo
	// 44
	// FooBarBar
	// 46
	// Qix
	// Foo
	// Qix
	// BarBar
	// FooBar
	// Bar
	// BarFoo
	// FooBar
	// BarBarBar
	// QixBar
	// FooBarQix
	// Bar
	// Bar
	// FooBar
	// 61
	// 62
	// FooQixFoo
	// 64
	// BarBar
	// Foo
	// Qix
	// 68
	// Foo
	// BarQixQix
	// Qix
	// FooQix
	// QixFoo
	// Qix
	// FooBarQixBar
	// Qix
	// QixQixQix
	// FooQix
	// Qix
	// Bar
	// Foo
	// 82
	// Foo
	// FooQix
	// BarBar
	// 86
	// FooQix
	// 88
	// 89
	// FooBar
	// Qix
	// 92
	// FooFoo
	// 94
	// BarBar
	// Foo
	// Qix
	// Qix
	// Foo
	// Bar
}
