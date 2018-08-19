package MAPL_engine

import (
	"strings"
	"regexp"
)

func IsNumberOfFieldsEqual(generalStruct GeneralStruct, yamlString string) (bool, string) {
	// we convert the structure into a string and count the number of non-empty "fields". Then we compare to the number of nonEmpty fields in the original yaml string
	// It will not work with structure fields with default values (for example ints or floats) so we remove them.
	return compareJsonAndYaml(generalStruct.ToJson(), yamlString)
}


func compareJsonAndYaml(jsonString string, yamlString string) (bool, string) {

	jsonString = strings.Replace(jsonString, "\"", "", -1)    // remove "
	jsonString = strings.Replace(jsonString, "null", "", -1)  // remove null
	jsonString = strings.Replace(jsonString, ",\n", "\n", -1) // remove ',' at the end of a line
	jsonString = strings.Replace(jsonString, "{\n", "\n", -1) // remove '{' at the end of a line
	jsonString = strings.Replace(jsonString, "[\n", "\n", -1) // remove '[' at the end of a line
	jsonString = strings.Replace(jsonString, "}\n", "\n", -1) // remove '}' at the end of a line
	jsonString = strings.Replace(jsonString, "]\n", "\n", -1) // remove ']' at the end of a line

	jsonString = strings.Replace(jsonString, "ValueInt:", "ValueInt-", -1) // change slightly so that the regex will not count it [this field is not part of the input so disregard it]
	jsonString = strings.Replace(jsonString, "ValueFloat:", "ValueFloat-", -1) // change slightly so that the regex will not count it [this field is not part of the input so disregard it]
	jsonString = strings.Replace(jsonString, "ValueRegex:", "ValueRegex-", -1) // change slightly so that the regex will not count it [this field is not part of the input so disregard it]
	jsonString = strings.Replace(jsonString, "ValueTime:", "ValueTime-", -1) // change slightly so that the regex will not count it [this field is not part of the input so disregard it]
	jsonString = strings.Replace(jsonString, "UTC:", "UTC-", -1) // change slightly so that the regex will not count it [this field is not part of the input so disregard it]
	jsonString = strings.Replace(jsonString, "ContextType:", "ContextType-", -1) // change slightly so that the regex will not count it [this field is not part of the input so disregard it]

	jsonString = strings.Replace(jsonString,"Size: 0\n","Size- 0\n",-1) // change slightly so that the regex will not count it [this is an integer field]
	jsonString = strings.Replace(jsonString,"Duration: 0\n","Duration- 0\n",-1) // change slightly so that the regex will not count it [this is an integer field]
	jsonString = strings.Replace(jsonString,"ResponseCode: 0\n","ResponseCode- 0\n",-1) // change slightly so that the regex will not count it [this is an integer field]
	jsonString = strings.Replace(jsonString,"0001-01-01T00:00:00Z\n","\n",-1) // change slightly so that the regex will not count it [this is a timestamp field]

	//re := regexp.MustCompile("[[:alnum:]][:][ ][[:alnum:]]") // catches fieldnam[e: v]alue'
	re := regexp.MustCompile(`(?m)^.*[:][ ]\S+`)

	matches1 := re.FindAllString(yamlString, -1)
	matches2 := re.FindAllString(jsonString, -1)

	L1 := len(matches1)
	L2 := len(matches2)
	if L1 != L2 {
		return false, jsonString
	}

	return true, ""
}