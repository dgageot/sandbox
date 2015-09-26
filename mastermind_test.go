package main

import "testing"
import "github.com/assertgo/assert"

func TestMasterMind(t *testing.T) {
	assert := assert.New(t)

	// should recognize when no peg is correct position or color
	assert.That(countCorrectColorAndPosition([...]string{"B", "B", "B", "B"}, [...]string{"N", "N", "N", "N"})).IsEqualTo(0)

	// should recognize when one peg is correct position and color
	assert.That(countCorrectColorAndPosition([...]string{"B", "B", "B", "B"}, [...]string{"B", "N", "N", "N"})).IsEqualTo(1)

	// should recognize when all pegs are correct position and color
	assert.That(countCorrectColorAndPosition([...]string{"B", "B", "B", "B"}, [...]string{"B", "B", "B", "B"})).IsEqualTo(4)

	// should recognize when no peg is correct color
	assert.That(countCorrectColorWrongPosition([...]string{"B", "B", "B", "B"}, [...]string{"N", "N", "N", "N"})).IsEqualTo(0)

	// should recognize when a peg is correct color and wrong position
	assert.That(countCorrectColorWrongPosition([...]string{"V", "B", "B", "B"}, [...]string{"N", "V", "N", "N"})).IsEqualTo(1)

	// should recognize when there are more pegs of a given color in the guess than in the secret, and ignore these pegs
	assert.That(countCorrectColorWrongPosition([...]string{"V", "B", "B", "B"}, [...]string{"N", "V", "V", "N"})).IsEqualTo(1)

	// should ignore pegs with correct position and color
	assert.That(countCorrectColorWrongPosition([...]string{"V", "B", "B", "V"}, [...]string{"V", "V", "N", "N"})).IsEqualTo(1)

	// should recognize when pegs are correct colors for multiple colors
	assert.That(countCorrectColorWrongPosition([...]string{"V", "R", "B", "B"}, [...]string{"R", "V", "N", "N"})).IsEqualTo(2)
}
