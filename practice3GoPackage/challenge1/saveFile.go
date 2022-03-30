// Una empresa que se encarga de vender productos de limpieza necesita:
// 1. Implementar una funcionalidad para guardar un archivo de texto, con la información de
// productos comprados, separados por punto y coma (csv).
// 2. Debe tener el id del producto, precio y la cantidad.
// 3. Estos valores pueden ser hardcodeados o escritos en duro en una variable.
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// type sliceProductos struct {
// 	Productos []producto `json:"productos"`
// }

type producto struct {
	Id       int    `json:"id"`
	Precio   int    `json:"precio"`
	Cantidad int    `json:"cantidad"`
	Nombre   string `json:"nombre"`
}

func main() {
	// newSliceProducts := &sliceProductos{}
	// newSliceProducts.Productos = []producto{}
	newProduct := []producto{{1, 5000, 2, "galleta"}, {2, 200, 2, "bombom"}, {3, 8000, 2, "gaseosa"}}
	newProductJson, err1 := json.Marshal(newProduct)
	if err1 != nil {
		log.Fatal(err1)
	}
	d1 := []byte(string(newProductJson))
	err2 := os.WriteFile("./file.txt", d1, 0644)
	if err2 != nil {
		log.Fatal(err2)
	}
	data, err := os.Open("./file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()

	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		fmt.Println(scanner.Text())

	}

	if err := scanner.Err(); err != nil {
		log.Fatal()
	}

}
