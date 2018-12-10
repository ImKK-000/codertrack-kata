package couple

import (
	"fmt"
	"sort"
)

type Person struct {
	Name, Destination string
}

func SortName(personStrings []string) []string {
	sort.Strings(personStrings)
	return personStrings
}

func GetPerson(personString string) (string, string) {
	var name, destination string
	fmt.Sscanf(personString, "%s %s", &name, &destination)
	return name, destination
}

func GetPersonSingle(personStrings []string) []string {
	destinations := map[string][]string{}
	for _, personString := range personStrings {
		personName, personDestination := GetPerson(personString)
		destinations[personDestination] = append(destinations[personDestination], personName)
	}
	var personSingle []string
	for _, value := range destinations {
		if len(value) == 1 {
			personSingle = append(personSingle, value[0])
		}
	}

	return SortName(personSingle)
}
