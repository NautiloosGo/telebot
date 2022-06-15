package product

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const catalogFilename = "etc/catalog.json"

var catalog Catalog

func GetJsonCatalog() {

	// открыть(создать) и закрыть, чтобы проставить правила доступа к файлу
	f, err := os.OpenFile(catalogFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

	//считываем из файла json
	rawDataIn, err := ioutil.ReadFile(catalogFilename)
	if err != nil {
		log.Println("Cannot load catalog:", err)
	}

	err = json.Unmarshal(rawDataIn, &catalog)
	if err != nil {
		log.Println("Invalid catalogs format:", err)
	}
}

func getSku() {
	// добавляем продукт
	newProduct := Sku{
		Title:       "Cherry",
		Description: "berry",
		Price:       11,
	}

	//проверка на повтор
	i, ok := catalog.findSku(newProduct.Title)
	if ok == true {
		catalog.Products[i] = newProduct
		fmt.Printf("New Sku %s is already exist with id = #%d. Replaced.", newProduct.Title, i)
	} else {
		catalog.Products = append(catalog.Products, newProduct)
		fmt.Printf("New Sku %s added with id = #%d.", newProduct.Title, i)
	}
}
func RewriteStorage() error {
	//заливаем catalog обратно в json
	rawDataOut, err := json.MarshalIndent(&catalog, "", "  ")
	if err != nil {
		log.Fatal("JSON marshaling failed:", err)
	}

	err = ioutil.WriteFile(catalogFilename, rawDataOut, 0)
	if err != nil {
		log.Fatal("Cannot write updated catalog file:", err)
	}
	return nil
}

func (cat *Catalog) editSku(s *Sku) (int, bool) {
	for i, unit := range cat.Products {
		if s.Title == unit.Title {
			cat.Products[i] = *s
			return i, true
		}
	}
	return 0, false
}
func (cat *Catalog) findSku(title string) (int, bool) {
	for i, unit := range cat.Products {
		if unit.Title == title {
			return i, true
		}
	}
	return len(cat.Products), false
}
