package transpose

import "strings"

func ConvertChord(chord string) string {
	mappingChords := map[string]string{
		"C":  "D",
		"Dm": "Em",
		"Em": "F#m",
		"F":  "G",
		"G":  "A",
		"Am": "Bm",
	}
	return mappingChords[chord]
}

func Transpose(chords string) string {
	var newChords []string
	chordsSlice := strings.Split(chords, " ")
	for _, chord := range chordsSlice {
		newChords = append(newChords, ConvertChord(chord))
	}
	return strings.Join(newChords, " ")
}

func TransposeArray(chordsArray []string) string {
	var newChordsArray []string
	for _, chords := range chordsArray {
		newChordsArray = append(newChordsArray, Transpose(chords))
	}
	return strings.Join(newChordsArray, "\n")
}
