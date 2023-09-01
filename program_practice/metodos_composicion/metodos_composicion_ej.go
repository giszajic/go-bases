/* Enunciado:
1- Tener una estructura llamada Product con los campos ID,
Name, Price, Description y Category.
2- Tener un slice global de Product llamado Products
instanciadocon valores.
3- 2 métodos asociados a la estructura Product: Save(),
GetAll(). El método Save() deberá tomar el slice de Products
y añadir el producto desde el cual se llama al método.
El método GetAll() deberá imprimir todos los productos guardados
en el slice Products.
4- Una función getById() al cual se le deberá pasar un INT como
parámetro y retorna el producto correspondiente al parámetro pasado.
5- Ejecutar al menos una vez cada método y función definido desde main().
*/

package main

import "fmt"

type Product struct {
	Id          int
	Name        string
	Price       float64
	Description string
	Category    string
}

var Products []Product

func (p *Product) Save() {
	Products = append(Products, *p)
}

func (p *Product) GetAll() {
	for _, product := range Products {
		fmt.Printf("ID: %d\nName: %s\nPrice: %.2f\nDescription: %s\nCategory: %s\n\n", product.Id, product.Name, product.Price, product.Description, product.Category)
	}
}

func getById(id int) *Product {
	for _, product := range Products {
		if product.Id == id {
			return &product
		}
	}
	return nil
}

func main() {
	product1 := Product{
		Id:          1,
		Name:        "Chocolate",
		Price:       200.00,
		Description: "Delicioso chocolate negro",
		Category:    "Alimentos",
	}

	product2 := Product{
		Id:          2,
		Name:        "Libro",
		Price:       15.99,
		Description: "Novela emocionante",
		Category:    "Libros",
	}

	// Guardar productos en el slice
	product1.Save()
	product2.Save()

	// Obtener todos los productos
	fmt.Println("Todos los productos:")
	product1.GetAll()

	// Obtener un producto por ID
	fmt.Println("\nObtener producto por ID:")
	id := 1
	result := getById(id)
	if result != nil {
		fmt.Printf("ID: %d\nName: %s\nPrice: %.2f\nDescription: %s\nCategory: %s\n", result.Id, result.Name, result.Price, result.Description, result.Category)
	} else {
		fmt.Printf("Producto con ID %d no encontrado\n", id)
	}
}
