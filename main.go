// skip-mongodb-collection-generator
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"runtime/debug"
	"strings"
	"time"
	"unsafe"
)

// regex for removing spaces
var space = regexp.MustCompile(`\s+`)

const (
	// temporary file for populates the struct and const
	populateCodePath = "populate_code.txt"
	// the target readme document
	readmePath = "README.md"
	// the target path for reading the struct files
	targetPath = "."
	// the sign on the readme file to start writing the doc
	startCollectionDocSign = "<!-- start collection doc -->"
	// the sign on the readme file to end writing the doc
	endCollectionDocSign = "<!-- end collection doc -->"
)

func main() {
	// remove temporary file
	defer os.Remove(populateCodePath)

	// populates the struct and const
	err := populateEntitiesImports()
	if err != nil {
		log.Fatal(err)
	}

	// get the entities, imports and collections data
	entities, imports, collections, err := getEntitiesImportsCollections(populateCodePath)
	if err != nil {
		log.Fatal(err)
	}

	// generate the new markdown
	markdown, err := generateMarkdown(populateCodePath, entities, imports, collections)
	if err != nil {
		log.Fatal(err)
	}

	// generate the new readme
	newReadme, err := generateNewReadme(markdown)
	if err != nil {
		log.Fatal(err)
	}

	// write the new readme
	err = writeNewReadme(newReadme)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("### Generating all collection finished!! ###")
}

// populateEntitiesImports populates the struct and const
func populateEntitiesImports() error {
	entityImportFileContent := ""

	// read all files under the target path
	err := filepath.Walk(targetPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			// only read the go file
			if strings.Contains(path, ".go") {
				b, err := os.ReadFile(path) // just pass the file name
				if err != nil {
					return err
				}

				// only read file that has bson tag and skip some file if there's any skip comment
				if strings.Contains(string(b), "bson:\"") && !strings.Contains(string(b), "// skip-mongodb-collection-generator") {
					entityImportFileContent += string(b)
				}
			}

			return nil
		})
	if err != nil {
		log.Fatal(err)
	}

	// creates the temporary file
	entityImportFile, err := os.Create(populateCodePath)
	if err != nil {
		log.Fatal(err)
	}
	defer entityImportFile.Close()

	// writes the temporary file content
	_, err = entityImportFile.WriteString(entityImportFileContent)
	if err != nil {
		return err
	}

	return nil
}

// getEntitiesImportsCollections gets the list of struct entity, imported package and collections
// then store it in map in order to get the link pkg of field of struct (built in, internal & external dependencies)
func getEntitiesImportsCollections(filePath string) (map[string]struct{}, map[string]string, map[string]string, error) {
	entities := make(map[string]struct{})
	imports := make(map[string]string)
	collections := make(map[string]string)

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var findStruct, findImport, findCollection, structHaveBSON bool

	scanner := bufio.NewScanner(file)

	lastKey := ""

	for scanner.Scan() {
		// checks if a line is a starting entity of struct or not e.g. type exampleStruct struct {
		// and grab the struct name then put it in map entities
		if checkSubstrings(scanner.Text(), "type", "struct") {
			// grab only the entity name
			entity := strings.ReplaceAll(scanner.Text(), "type ", "")
			entity = strings.ReplaceAll(entity, " struct {", "")

			lastKey = entity

			findStruct = true

			structHaveBSON = false

			entities[entity] = struct{}{}
		}

		// checks if a line is an end of a struct = }
		if findStruct && space.ReplaceAllString(scanner.Text(), "") == "}" {
			if !structHaveBSON {
				delete(entities, lastKey)
			}

			findStruct = false
		}

		// check if the field of struct is containing bson tag
		if findStruct {
			if strings.Contains(scanner.Text(), "bson:\"") {
				structHaveBSON = true
			}
		}

		// checks if a line is a starting of import dependencies e.g. import (
		if checkSubstrings(scanner.Text(), "import (") {
			findImport = true
			continue
		}

		// checks if a line is an end of an import part = )
		if findImport && space.ReplaceAllString(scanner.Text(), "") == ")" {
			findImport = false
		}

		// grab the imported pkg then put in map imports
		if findImport && space.ReplaceAllString(scanner.Text(), "") != "" {
			// remove all spaces of imported package, then it become
			// "errors" or alias"github.com/xxx/yyy" or "github.com/xxx/yyy" or "math/rand"
			importedPkgs := space.ReplaceAllString(scanner.Text(), "")
			// split by double quotes
			pkgs := strings.Split(importedPkgs, "\"")

			var key, pkg string

			// the complete package always on the 2nd position
			pkg = pkgs[1]

			// finds the package part used in the field type
			// e.g. go.mongodb.org/mongo-driver/bson/primitive used in primitive.ObjectID
			// time used in time.Time or using alias pmt "go.mongodb.org/mongo-driver/bson/primitive" used in pmt.ObjectID
			if pkgs[0] != "" {
				// if the pkg name written in alias

				key = pkgs[0]
			} else if strings.Contains(pkgs[1], "/") {
				// if the pkg name has / , it can be external or internal
				// e.g. github.com/xxx/yyy or math/rand

				s := strings.Split(pkgs[1], "/")
				key = s[len(s)-1]
			} else {
				// the rest of case

				key = pkgs[1]
			}

			imports[key] = pkg
		}

		// checks if a line is a starting of const line e.g. const (
		if checkSubstrings(scanner.Text(), "const (") {
			findCollection = true
			continue
		}

		// checks if a line is an end of a constant part = )
		if findCollection && space.ReplaceAllString(scanner.Text(), "") == ")" {
			findCollection = false
		}

		// grab the const for collection with group const format
		if findCollection && strings.Contains(scanner.Text(), "Collection") {
			// remove all spaces
			collectionConst := space.ReplaceAllString(scanner.Text(), "")
			// split by =
			splitCollection := strings.Split(collectionConst, "=")

			collections[splitCollection[0]] = strings.ReplaceAll(splitCollection[1], "\"", "")
		}

		// grab the const for collection with single const format
		if strings.Contains(scanner.Text(), "const") && strings.Contains(scanner.Text(), "Collection") {
			// remove const
			collectionConst := strings.ReplaceAll(scanner.Text(), "const ", "")
			// remove all spaces
			collectionConst = space.ReplaceAllString(collectionConst, "")
			// split by =
			splitCollection := strings.Split(collectionConst, "=")

			collections[splitCollection[0]] = strings.ReplaceAll(splitCollection[1], "\"", "")
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, nil, err
	}

	return entities, imports, collections, nil
}

// generateMarkdown generates markdown documentation for the MongoDB
// collection structure based on the struct that have bson tag.
func generateMarkdown(
	filePath string,
	entities map[string]struct{},
	imports map[string]string,
	collections map[string]string,
) (markdown string, err error) {
	// mapEntity stores the entity as key and the data structure as map
	mapEntity := make(map[string]map[string]string)

	// read the build info to get the service name later
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		log.Fatal("failed to read build info")
	}

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	markdown += "\n"

	var findStruct bool

	// lastEntity and lastEntityDataStructure need to store the entity to map
	var lastEntity string
	var lastEntityDataStructure map[string]string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// checks if a line is a starting entity of struct or not e.g. type exampleStruct struct {
		if checkSubstrings(scanner.Text(), "type", "struct") {
			// grab only the entity name
			entity := strings.ReplaceAll(scanner.Text(), "type ", "")
			entity = strings.ReplaceAll(entity, " struct {", "")

			// skip if the entity is not on the entities key (because don't have bson tag)
			if _, ok := entities[entity]; !ok {
				continue
			}

			// initialize the map entity
			lastEntity = entity

			lastEntityDataStructure = make(map[string]string)
			mapEntity[entity] = lastEntityDataStructure

			// create the heading
			markdown += "### " + entity + "\n"

			// if the entity is the parent of collection, add a description
			if val, ok := collections[entity+"Collection"]; ok {
				markdown += "#### *collection: " + val + "*\n"
				markdown += "```json\n"
				markdown += "{{" + entity + "JSONReplace}}\n"
				markdown += "```\n"
			}

			// create the header table
			markdown += "| Field | Type |\n"
			markdown += "| --- | --- |\n"

			findStruct = true

			continue
		}

		// checks if a line is an end of a struct = }
		if findStruct && space.ReplaceAllString(scanner.Text(), "") == "}" {
			markdown += "\n"
			findStruct = false
		}

		// grab the field & data type then put in table as a row
		if findStruct && strings.Contains(scanner.Text(), "bson:") {
			var field, dataType string

			// split the text between bson tag e.g. `ID primitive.ObjectID bson:"_id,omitempty"`
			bsonString := strings.Split(scanner.Text(), "bson:")

			if len(bsonString) > 1 {
				// split the whitespace just in case there are multiple tags e.g. bson:"xx" json:"yy" sql:"zz"
				splitTags := strings.Split(bsonString[len(bsonString)-1], " ")

				// remove unnecessary text like backtick, omitempty and double quotes
				// in order to get the field name
				// e.g. from e.g.-> ID primitive.ObjectID bson:"_id,omitempty"` gets the _id only
				field = strings.ReplaceAll(splitTags[0], "\"", "")
				field = strings.ReplaceAll(field, "`", "")
				field = strings.ReplaceAll(field, ",omitempty", "")

				findChar := false
				grabChar := false

				// find the data type e.g.-> ID primitive.ObjectID `bson:"_id,omitempty"`
				// getting the primitive.ObjectID part
				// first need to split by front backtick since the previous split is not included it
				// for the example of previous split -> ID primitive.ObjectID `
				fieldAndType := strings.Split(bsonString[0], "`")

				for i := range fieldAndType[0] {
					if findChar && string(fieldAndType[0][i]) == " " {
						grabChar = true
					}

					if string(fieldAndType[0][i]) != "" {
						findChar = true
					}

					if grabChar {
						dataType += string(fieldAndType[0][i])
					}
				}
			}

			// remove all spaces on data type
			dataType = space.ReplaceAllString(dataType, "")
			dataTypeText := dataType
			dataTypeDoc := ""

			importedLib := false

			// check if the data type is not built in means it's internal or external dependencies
			// e.g. time.Time or primitive.ObjectID
			if checkSubstrings(dataType, ".") {
				s := strings.Split(dataType, ".")

				// skip if the len is under 2
				if len(s) < 2 {
					continue
				}

				// assign the first char to dataType
				// and the second char to dataTypeDoc
				// e.g. time.Time -> time (1) and Time (2)
				dataType = s[0]
				dataTypeDoc = s[1]
				importedLib = true
			}

			// to add the prefix of data type, is it has array, pointer or both
			prefix := ""

			if strings.Contains(dataType, "[]*") {
				prefix = "[]*"
			} else if strings.Contains(dataType, "*") {
				prefix = "*"
			} else if strings.Contains(dataType, "[]") {
				prefix = "[]"
			}

			// remove * and [] to match the key in imports & entities map
			dataTypeLink := strings.ReplaceAll(dataType, "*", "")
			dataTypeLink = strings.ReplaceAll(dataTypeLink, "[]", "")

			// if the field is from imported lib
			if importedLib {
				// create the dataType text and link based on imports map
				if val, ok := imports[dataTypeLink]; ok {
					if strings.Contains(val, bi.Main.Path) {
						dataType = "<a href=\"#" + dataTypeDoc + "\">" + dataTypeDoc + "</a>"
					} else {
						split := strings.Split(val, "/")

						dataType = "<a href=\"https://pkg.go.dev/" + val + "#" + dataTypeDoc + "\">" + prefix + split[len(split)-1] + "." + dataTypeDoc + "</a>"
					}
				}
			} else {
				// create the dataType text and link based on entities map,
				// if the key not found, pointing to built in library
				if _, ok := entities[dataTypeLink]; ok {
					dataType = "<a href=\"#" + dataTypeLink + "\">" + dataTypeText + "</a>"
				} else {
					dataType = "<a href=\"https://pkg.go.dev/builtin#" + dataTypeLink + "\">" + dataTypeText + "</a>"
				}
			}

			// wrapped the row
			markdown += "|" + field + "|" + dataType + "|\n"

			// store data structure to the map entity
			lastEntityDataStructure[field] = dataTypeText
			mapEntity[lastEntity] = lastEntityDataStructure
		}
	}

	if err = scanner.Err(); err != nil {
		return "", err
	}

	// generates JSON structure from collection entity
	for keyEntity, valueEntity := range mapEntity {
		if _, ok := collections[keyEntity+"Collection"]; ok {
			newDataStructure := make(map[string]interface{})

			for keyDataStructure, valueDataStructure := range valueEntity {
				newDataStructure[keyDataStructure] = mapFieldToValue(mapEntity, valueDataStructure)
			}

			jsonStr, err := json.MarshalIndent(newDataStructure, "", "  ")
			if err != nil {
				return "", err
			}

			markdown = strings.ReplaceAll(markdown, "{{"+keyEntity+"JSONReplace}}", string(jsonStr))
		}
	}

	return markdown, nil
}

// generateNewReadme generates the content of the new readme (including the old readme + collection readme)
func generateNewReadme(readmeCollection string) (string, error) {
	readmeFile, err := os.Open(readmePath)
	if err != nil {
		return "", err
	}
	defer readmeFile.Close()

	scanner := bufio.NewScanner(readmeFile)

	skipLine := false
	newReadme := ""

	for scanner.Scan() {
		if skipLine {
			// continue reads the content if end collection label found
			if scanner.Text() == endCollectionDocSign {
				skipLine = false

				newReadme += scanner.Text() + "\n"
			}

			continue
		}

		newReadme += scanner.Text()

		// start skips the content if start collection label found
		// and put the readme collection on the content
		if scanner.Text() == startCollectionDocSign {
			skipLine = true

			newReadme += readmeCollection
		}

		newReadme += "\n"
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return newReadme, nil
}

// writeNewReadme writes the new readme file.
func writeNewReadme(newReadme string) error {
	newReadmeFile, err := os.Create(readmePath)
	if err != nil {
		log.Fatal(err)
	}
	defer newReadmeFile.Close()

	_, err = newReadmeFile.WriteString(newReadme)
	if err != nil {
		return err
	}

	return nil
}

// checks some strings contains in a string or not.
func checkSubstrings(str string, subs ...string) bool {
	isCompleteMatch := true

	for _, sub := range subs {
		if !strings.Contains(str, sub) {
			isCompleteMatch = false
		}
	}

	return isCompleteMatch
}

// mapFieldToValue convert the data type text into real value
func mapFieldToValue(mapEntity map[string]map[string]string, field string) interface{} {
	var val interface{}

	type sample struct {
		a int
		b string
	}

	s := &sample{
		a: 1,
		b: "test",
	}

	field = strings.ReplaceAll(field, "*", "")

	switch field {
	case "primitive.ObjectID":
		val = "633fe183150872f57b930ac8"
	case "[]primitive.ObjectID":
		val = []string{"633fe183150872f57b930ac8"}
	case "string":
		val = "lorem ipsum"
	case "[]string":
		val = []string{"lorem ipsum"}
	case "byte":
		val = 97
	case "[]byte":
		val = []byte{98}
	case "rune":
		val = '♄'
	case "[]rune":
		val = []rune{'♄'}
	case "int", "int8", "int16", "int32", "int64":
		val = 100
	case "[]int", "[]int8", "[]int16", "[]int32", "[]int64":
		val = []int{200}
	case "uintptr":
		val = uintptr(unsafe.Pointer(s))
	case "[]uintptr":
		val = []uintptr{uintptr(unsafe.Pointer(s))}
	case "uint", "uint8", "uint16", "uint32", "uint64":
		val = 300
	case "[]uint", "[]uint8", "[]uint16", "[]uint32", "[]uint64":
		val = []uint{400}
	case "float32", "float64":
		val = 99.99
	case "[]float32", "[]float64":
		val = []float32{199.99}
	case "bool":
		val = true
	case "[]bool":
		val = []bool{true}
	case "time.Time":
		val = time.Now()
	case "[]time.Time":
		val = []time.Time{time.Now()}
	default:
		// split if the field is from internal imported struct
		if strings.Contains(field, ".") {
			splitDot := strings.Split(field, ".")
			field = splitDot[1]
		}

		// checks if the field entity is on the map, either it's an array or not
		if dataStructure, ok := mapEntity[strings.ReplaceAll(field, "[]", "")]; ok {
			// checks if the fields has the same type, for avoid error
			// due to infinite loop, then assign null as the value
			newDataStructure := make(map[string]interface{})
			for keyDataStructure, valueDataStructure := range dataStructure {
				basicValueDataStructure := strings.ReplaceAll(valueDataStructure, "[]", "")
				basicValueDataStructure = strings.ReplaceAll(basicValueDataStructure, "*", "")
				if field == basicValueDataStructure {
					newDataStructure[keyDataStructure] = nil
				} else {
					newDataStructure[keyDataStructure] = mapFieldToValue(mapEntity, valueDataStructure)
				}
			}

			if strings.Contains(field, "[]") {
				val = []interface{}{newDataStructure}
			} else {
				val = newDataStructure
			}
		} else {
			val = nil
		}
	}

	return val
}
