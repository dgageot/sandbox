package main

import "testing"
import "github.com/assertgo/assert"

func TestCountCorrectColorAndPosition(t *testing.T) {
	assert := assert.New(t)

	noPegIsCorrectPositionOrColor := CountCorrectColorAndPosition([...]string{"B", "B", "B", "B"}, [...]string{"N", "N", "N", "N"})
	assert.That(noPegIsCorrectPositionOrColor).IsEqualTo(0)

	onePegIsCorrectPositionAndColor := CountCorrectColorAndPosition([...]string{"B", "B", "B", "B"}, [...]string{"B", "N", "N", "N"})
	assert.That(onePegIsCorrectPositionAndColor).IsEqualTo(1)

	allPegsAreCorrectPositionAndColor := CountCorrectColorAndPosition([...]string{"B", "B", "B", "B"}, [...]string{"B", "B", "B", "B"})
	assert.That(allPegsAreCorrectPositionAndColor).IsEqualTo(4)
}

func TestCountCorrectColorWrongPosition(t *testing.T) {
	assert := assert.New(t)

	noPegIsCorrectColor := CountCorrectColorWrongPosition([...]string{"B", "B", "B", "B"}, [...]string{"N", "N", "N", "N"})
	assert.That(noPegIsCorrectColor).IsEqualTo(0)

	onePegIsCorrectColorAndWrongPosition := CountCorrectColorWrongPosition([...]string{"V", "B", "B", "B"}, [...]string{"N", "V", "N", "N"})
	assert.That(onePegIsCorrectColorAndWrongPosition).IsEqualTo(1)

	ignoreAditionalPegsOfSameColor := CountCorrectColorWrongPosition([...]string{"V", "B", "B", "B"}, [...]string{"N", "V", "V", "N"})
	assert.That(ignoreAditionalPegsOfSameColor).IsEqualTo(1)

	ignorePegsWithCorrectColorAndPosition := CountCorrectColorWrongPosition([...]string{"V", "B", "B", "V"}, [...]string{"V", "V", "N", "N"})
	assert.That(ignorePegsWithCorrectColorAndPosition).IsEqualTo(1)

	recognizePegsAreCorrectColorsForMultipleColors := CountCorrectColorWrongPosition([...]string{"V", "R", "B", "B"}, [...]string{"R", "V", "N", "N"})
	assert.That(recognizePegsAreCorrectColorsForMultipleColors).IsEqualTo(2)
}
