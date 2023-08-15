package typo

import "strings"

// specialStringsInStructureTag service words in a structure tag that are not a name
var specialStringsInStructureTag = []string{`-`, `omitempty`}

// writeNullInString variable containing text, which means nil when output as text
var writeNullInString = `null`

// structureTagContainingNameTheElementJson a variable containing the name of the structure tag in which the name of the element's json is specified
var structureTagContainingNameTheElementJson = `json`

// SetStringToWriteNull specify text that stands for nil when outputting it as text
func SetStringToWriteNull(str string) {
	writeNullInString = str
}

// SetStructureTagContainingNameTheElementJson specifying a struct tag that specifies the json name of the element
func SetStructureTagContainingNameTheElementJson(name string) {
	structureTagContainingNameTheElementJson = name
}

// SetSpecialStringsInStructureTag indication of service words of a structural tag that are not a name
func SetSpecialStringsInStructureTag(words []string) {
	specialStringsInStructureTag = words
}

// inSpecialStringsInStructureTag search for a string among the service words of a structure tag
func inSpecialStringsInStructureTag(str string) bool {
	for i := 0; i < len(specialStringsInStructureTag); i++ {
		if str == specialStringsInStructureTag[i] {
			return true
		}
	}

	return false
}

// searchNameInStructureTag search for a possible field name in a structure tag
func searchNameInStructureTag(tag string) string {
	if `` == tag {
		return tag
	}
	sliceString := strings.Split(tag, ",")
	var temp string
	for i := 0; i < len(sliceString); i++ {
		temp = strings.TrimSpace(sliceString[i])
		if !inSpecialStringsInStructureTag(temp) {
			return temp
		}
	}
	return ``
}
