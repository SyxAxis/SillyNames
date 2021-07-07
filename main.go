package main

import (
	"bufio"
	"embed"
	"flag"
	"fmt"
	"math/rand"
	"time"
)

// to make this more efficient the names are loaded from files just once
// then the data structs are read over and over
type sillyNames struct {
	randomName []string
}

// attach all the external files into the EXE build
// datacontent is the file handle reference
//
//  NOTE this feature requires v1.16 or above
//
//go:embed data_sources/*
var datacontent embed.FS

// this is the folder which is still relevant even with the embedded data files
const folderpath string = "data_sources/"

func main() {

	clfNameType := flag.String("t", "character", "<string> - { acme | band | business | character | drug | eatery | fantasy | morpheme | team }")
	clfNumNames := flag.Int("n", 5, "<int> - number of entries")
	clfPrefixHonorific := flag.Bool("h", false, "Prefix honorific")
	flag.Parse()

	runNameGenerator(*clfNameType, *clfNumNames, *clfPrefixHonorific)

}

func runNameGenerator(clfNameType string, clfNumNames int, clfPrefixHonorific bool) {

	filePrefix := ""
	fileNumParts := 1
	fileSpaces := false

	switch clfNameType {
	case "acme":
		filePrefix = "acme_names_part"
		fileNumParts = 2
		fileSpaces = true
	case "band":
		filePrefix = "band_names_part"
		fileNumParts = 1
		fileSpaces = false
	case "business":
		filePrefix = "business_names_part"
		fileNumParts = 3
		fileSpaces = true
	case "character":
		filePrefix = "character_names_part"
		fileNumParts = 2
		fileSpaces = true
	case "drug":
		filePrefix = "drug_names_part"
		fileNumParts = 3
		fileSpaces = false
	case "eatery":
		filePrefix = "eatery_names_part"
		fileNumParts = 3
		fileSpaces = true
	case "fantasy":
		filePrefix = "fantasy_names_part"
		fileNumParts = 3
		fileSpaces = false
	case "morpheme":
		filePrefix = "morpheme_names_part"
		fileNumParts = 2
		fileSpaces = false
	case "team":
		filePrefix = "team_names_part"
		fileNumParts = 3
		fileSpaces = true
	default:
		filePrefix = "character_names_part"
		fileNumParts = 2
		fileSpaces = true
	}

	generateRandomNames(folderpath+filePrefix, fileNumParts, fileSpaces, clfNumNames, clfPrefixHonorific)
}

func generateRandomNames(filenamePrefix string, numFiles int, insSpace bool, numNames int, addHonorific bool) {

	rand.Seed(time.Now().UnixNano())

	var honorificsArray sillyNames

	// add honorific if requested
	if addHonorific {
		honorificsArray = getRandomNamesFromFile(folderpath + "honorifics.txt")
	}

	// loop the source files
	sillyNamesSlice := []sillyNames{}
	for x := 1; x < (numFiles + 1); x++ {
		sillyNamesSlice = append(sillyNamesSlice, getRandomNamesFromFile(filenamePrefix+fmt.Sprintf("%d", x)+".txt"))
	}

	for y := 0; y < numNames; y++ {

		var retName string

		// add honorific if requested
		if addHonorific {
			rndInt := rand.Intn(len(honorificsArray.randomName) - 1)
			retName = honorificsArray.randomName[rndInt] + " "
		}

		// loop arrays of filedata
		for fileNum := 0; fileNum < numFiles; fileNum++ {
			rndInt := rand.Intn(len(sillyNamesSlice[fileNum].randomName) - 1)
			retName = retName + sillyNamesSlice[fileNum].randomName[rndInt]
			if insSpace {
				retName = retName + " "
			}
		}
		fmt.Println(retName)
		retName = ""
	}

}

func getRandomNamesFromFile(filename string) sillyNames {

	cntEntries := 0
	f, err := datacontent.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	namelist := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		namelist = append(namelist, scanner.Text())
		cntEntries++
	}
	return sillyNames{namelist}

}
