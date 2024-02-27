package main

import (
	"first/person"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var personsArray = [5]*person.Person{person.NewPerson("Budi", "JL Pahlawan", "Software Engineer", "Karena Suka"), person.NewPerson("Joko", "JL Kebangkitan", "Software Analyst", "Karena Pingin"), person.NewPerson("Jeremy", "JL Nusa", "Manager", "Karena Pingin Tau"), person.NewPerson("Jason", "JL Nusa", "Politikus", "Karena Pingin Belajar"), person.NewPerson("Loni", "JL Nusa", "Pekerja kantor", "Karena sudah tau bahasa lain")}

	args := os.Args[1]
	if i, err := strconv.Atoi(args); err == nil && i < cap(personsArray) {
		fmt.Printf("Nama : %s\n", personsArray[i].Name)
		fmt.Printf("Alamat : %s\n", personsArray[i].Alamat)
		fmt.Printf("Pekerjaan : %s\n", personsArray[i].Pekerjaan)
		fmt.Printf("Alasan : %s\n", personsArray[i].Reason)
	} else {
		fmt.Println("Invalid input or index out of range")
	}
}
