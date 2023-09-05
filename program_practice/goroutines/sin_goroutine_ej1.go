// ejecucion de FORMA LINEAL
package main

import (
	"fmt"
	"time"
)

// funcion para hacer determinada tarea
// agregaremos una pausa, para simular funcionalidad
// que accede a una Base de Datos/ envía información a una API

func procesar(i int) {
	fmt.Println(i, "-Inicia")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(i, "-Termina")
}

// para que inicialice el 2do procesamiento, el 1ero debería haber terminado
func main() {
	for i := 0; i < 10; i++ {
		procesar(i)
	}
}
