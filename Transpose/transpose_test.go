package Transpose_test

import (
	"testing"
	. "transpose"
)

func TestTransposeInputChordArrayShoudBeTransposeChord(t *testing.T) {
	expectedTransposeChord := []string{
		"F#m F#m Bm F#m", "A D Bm D", "F#m Em Em Em", "G G D Em",
		"F#m G G D", "Bm Bm Em Bm", "D D G A", "D A Em F",
		"Bm F#m D F#m", "Em D D Bm", "A Bm Em D", "Em Em Bm F",
		"D F#m D F", "F#m Em A F#m", "F#m A A F#m", "D F#m A Em",
		"D F#m F#m F#m", "D A Em D", "A D G Em", "F#m D A A",
		"A A G A", "Em A F#m F#m", "F#m D F#m Bm", "Bm G D A",
		"G Bm F#m Bm", "G D D A", "A G F#m F#m", "D G D Bm",
		"F#m Bm F#m F", "A G A F#m", "D A Bm F#m", "A D G F#m",
		"A Em Em A", "Em Bm G Bm", "Bm F#m Em A", "A G G F",
		"F#m Em G F", "A F#m Em F", "F#m Em A Bm", "Em D F#m F",
		"G Bm A F", "Em Bm F#m F#m", "Em G D F#m", "A D A F#m",
		"G Em Em Bm", "A D G F#m", "Em Em D Em", "A G F#m Bm",
		"A G Bm A", "Bm A F#m Em",
	}
	chords := []string{
		"Em Em Am Em", "G C Am C", "Em Dm Dm Dm", "F F C Dm",
		"Em F F C", "Am Am Dm Am", "C C F G", "C G Dm F",
		"Am Em C Em", "Dm C C Am", "G Am Dm C", "Dm Dm Am F",
		"C Em C F", "Em Dm G Em", "Em G G Em", "C Em G Dm",
		"C Em Em Em", "C G Dm C", "G C F Dm", "Em C G G",
		"G G F G", "Dm G Em Em", "Em C Em Am", "Am F C G",
		"F Am Em Am", "F C C G", "G F Em Em", "C F C Am",
		"Em Am Em F", "G F G Em", "C G Am Em", "G C F Em",
		"G Dm Dm G", "Dm Am F Am", "Am Em Dm G", "G F F F",
		"Em Dm F F", "G Em Dm F", "Em Dm G Am", "Dm C Em F",
		"F Am G F", "Dm Am Em Em", "Dm F C Em", "G C G Em",
		"F Dm Dm Am", "G C F Em", "Dm Dm C Dm", "G F Em Am",
		"G F Am G", "Am G Em Dm",
	}

	actualTransposeChord := Transpose(chords)

	for index, _ := range actualTransposeChord {
		if string(expectedTransposeChord[index]) != string(actualTransposeChord[index]) {
			t.Errorf("expect %v but it got %v", expectedTransposeChord[index], actualTransposeChord[index])
		}
	}
}
