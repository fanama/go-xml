package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/fanama/go-xml/cmd/variables"
)

func main() {

	var source string

	fmt.Println("entrezle fichier source : ")

	fmt.Scanf("%s", &source)

	fmt.Println("opening ", source)

	// Open our xmlFile
	xmlFile, err := os.Open(source)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened exemple.xml")
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	// lecture du fichier

	byteValue, _ := ioutil.ReadAll(xmlFile)

	// parsage du fichier

	var list variables.Xlif

	xml.Unmarshal(byteValue, &list)

	// création des nouvelles données

	transcipt := list.File.Body.Trans

	var newtranscript variables.Body

	for _, data := range transcipt {
		if data.Target != "" {
			// ajout de la donnée avec target dans une variable
			newtranscript.Trans = append(newtranscript.Trans, data)

		}
	}

	// fmt.Println(newtranscript)

	// ecriture des nouvelles données

	liste := list

	liste.File.Body = newtranscript

	file, _ := xml.MarshalIndent(liste, "	", "	")

	_ = ioutil.WriteFile("./sorties/"+source, file, 0644)

}
