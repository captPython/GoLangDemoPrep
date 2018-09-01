package main

import (
	"fmt"
)

const usixteenbitmax float64 = 10
const data_mul float64 = 2

type cert struct {
	id    int16
	name  string
	data  int16
	value float32
}

func (c cert) getData() int16 {
	return int16((float64(c.data) * usixteenbitmax) / data_mul)
}

func (c *cert) updateData(newData int16) {
	c.data = newData
}

func updateDataNew(c cert, newData int16) {
	c.data = newData
}

func main() {
	a_cert := cert{id: 1, name: "Dev", data: -10, value: 10.10}
	fmt.Println(a_cert.data, a_cert.id, a_cert.name)
	fmt.Println(a_cert.getData())

	a_cert.updateData(10) // Updating the value of struct using pointer
	fmt.Println(a_cert.data, a_cert.id, a_cert.name)
	fmt.Println(a_cert.getData())

	updateDataNew(a_cert, 10)
	fmt.Println(a_cert.data, a_cert.id, a_cert.name)
	fmt.Println(a_cert.getData())
}
