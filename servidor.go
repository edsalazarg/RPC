package main

import (
	"fmt"
	"net"
	"net/rpc"
)

var GlobalMaterias map[string]map[string]float64
var GlobalAlumno map[string]map[string]float64
var PromTodos map[string][]float64

type Server struct{}

type Registro struct {
	Alumno       string
	Materia      string
	Calificacion float64
}

func (this *Server) Agregar(datos Registro, reply *string) error {
	materiasDelAlumno, existeElAlumno := GlobalAlumno[datos.Alumno]
	materia, existeLaMateria := GlobalMaterias[datos.Materia]

	if existeElAlumno {
		materiasDelAlumno[datos.Materia] = datos.Calificacion
	} else {
		materiasDelAlumno = make(map[string]float64)
		materiasDelAlumno[datos.Materia] = datos.Calificacion
		GlobalAlumno[datos.Alumno] = materiasDelAlumno
	}

	if existeLaMateria {
		materia[datos.Alumno] = datos.Calificacion
	} else {
		materia = make(map[string]float64)
		materia[datos.Alumno] = datos.Calificacion
		GlobalMaterias[datos.Materia] = materia
	}

	for key, element := range GlobalMaterias {
		fmt.Println("Materia:", key)
		for keyAl, Calificacion := range element {
			fmt.Println("	Alumno:", keyAl, "=>", "Calificacion:", Calificacion)
		}
	}
	fmt.Println("*****************")
	*reply = "Califacion Agregada"
	return nil
}

func (this *Server) PromAlum(alumno string, reply *float64) error {
	total := 0.0
	cantMaterias := 0.0
	for key, element := range GlobalMaterias {
		fmt.Println("Materia:", key)
		for keyAl, Calificacion := range element {
			if keyAl == alumno {
				fmt.Println("	Alumno:", keyAl, "=>", "Calificacion:", Calificacion)
				cantMaterias++
				total += Calificacion
			}
		}
	}
	promedio := total / cantMaterias
	fmt.Println("*****************")
	*reply = promedio
	return nil
}

func (this *Server) PromTodos(signal bool, reply *map[string][]float64) error {
	if signal {
		PromTodos = make(map[string][]float64)
		for _, element := range GlobalMaterias {
			for keyAl, Calificacion := range element {
				_, existeElPromedio := PromTodos[keyAl]
				if existeElPromedio {
					PromTodos[keyAl][0] += Calificacion
					PromTodos[keyAl][1]++
				} else {
					PromTodos[keyAl] = append(PromTodos[keyAl], Calificacion)
					PromTodos[keyAl] = append(PromTodos[keyAl], 1)
				}

			}
		}

		for key, element := range PromTodos {
			fmt.Println("Alumno", key)
			fmt.Println("	Calificacion Total", element[0])
			fmt.Println("	Cantidad de Materias", element[1])
			PromTodos[key] = append(PromTodos[key], element[0]/element[1])
		}
		fmt.Println("*****************")
		*reply = PromTodos
	}
	return nil
}

func (this *Server) PromMateria(materia string, reply *float64) error {
	total := 0.0
	cantAlumnos := 0.0
	for key, element := range GlobalMaterias {
		if key == materia {
			for keyAl, Calificacion := range element {
				fmt.Println("	Alumno:", keyAl, "=>", "Calificacion:", Calificacion)
				cantAlumnos++
				total += Calificacion

			}
		}
	}
	promedio := total / cantAlumnos
	fmt.Println("*****************")
	*reply = promedio
	return nil
}

func server() {
	rpc.Register(new(Server))
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go rpc.ServeConn(c)
	}
}

func main() {
	go server()

	GlobalAlumno = make(map[string]map[string]float64)
	GlobalMaterias = make(map[string]map[string]float64)

	var input string
	fmt.Scanln(&input)
}
