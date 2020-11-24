package main

import (
	"fmt"
	"net/rpc"
)

type Registro struct {
	Alumno       string
	Materia      string
	Calificacion float64
}

func client() {
	c, err := rpc.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	var op int64
	for {
		fmt.Println("1) Agregar Calificacion")
		fmt.Println("2) Mostrar Promedio Alumno")
		fmt.Println("3) Mostrar Promedio General")
		fmt.Println("4) Mostrar Promedio Materia")
		fmt.Println("0) Salir")
		fmt.Scanln(&op)

		switch op {
		case 1:
			var datos Registro
			fmt.Print("Nombre Alumno: ")
			fmt.Scanln(&datos.Alumno)
			fmt.Print("Materia: ")
			fmt.Scanln(&datos.Materia)
			fmt.Print("Calificacion: ")
			fmt.Scanln(&datos.Calificacion)

			var result string
			err = c.Call("Server.Agregar", datos, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Server.Agregar =", result)
			}
		case 2:
			var Alumno string
			fmt.Print("Nombre Alumno: ")
			fmt.Scanln(&Alumno)
			var result float64
			err = c.Call("Server.PromAlum", Alumno, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Server.PromAlum", result)
			}
		case 3:
			var result map[string][]float64
			err = c.Call("Server.PromTodos", true, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				for key, element := range result {
					fmt.Println(key, ": ", element[2])
				}
			}
		case 4:
			var Materia string
			fmt.Print("Nombre Materia: ")
			fmt.Scanln(&Materia)
			var result float64
			err = c.Call("Server.PromMateria", Materia, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Server.PromMateria", result)
			}
		case 0:
			return
		}
	}
}

func main() {
	client()
}
