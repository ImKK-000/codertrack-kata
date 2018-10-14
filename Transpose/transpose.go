package transpose

import (
	"strings"
)

func Transpose(chords []string) []string {
	for index, chord := range chords {
		chords[index] = TransposeChord(chord)
	}
	return chords
}

func TransposeChord(chords string) string {
	transposeMapping := map[string]string{
		"C":  "D",
		"Dm": "Em",
		"Em": "F#m",
		"F":  "G",
		"G":  "A",
		"Am": "Bm",
	}
	chordArray := strings.Split(chords, " ")
	for index, chord := range chordArray {
		chordArray[index] = transposeMapping[chord]
	}
	return strings.Join(chordArray, " ")
}
