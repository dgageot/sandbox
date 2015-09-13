package main

import "testing"
import "github.com/assertgo/assert"

func TestFooBarQix(t *testing.T) {
	assert := assert.New(t)

	assert.That(foobarqix(1)).IsEqualTo("1")
	assert.That(foobarqix(2)).IsEqualTo("2")
	assert.That(foobarqix(3)).IsEqualTo("FooFoo")
	assert.That(foobarqix(5)).IsEqualTo("BarBar")
	assert.That(foobarqix(7)).IsEqualTo("QixQix")
	assert.That(foobarqix(9)).IsEqualTo("Foo")
	assert.That(foobarqix(10)).IsEqualTo("Bar")
	assert.That(foobarqix(14)).IsEqualTo("Qix")
	assert.That(foobarqix(15)).IsEqualTo("FooBarBar")
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
