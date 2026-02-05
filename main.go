package main

import (
	"fmt"
	"sort"
)

type product struct {
	Name     string
	Quantity int
	Price    float64
}

func main() {
	var (
		n int
		e int
	)
	fmt.Scan(&n)

	type operator struct {
		Name2     string
		operation string
		k         int
	}
	Products := make(map[string]product)
	for i := 0; i != n; i++ {
		var p product
		fmt.Scan(&p.Name, &p.Quantity, &p.Price)
		Products[p.Name] = p
	}
	fmt.Scan(&e)
	na := make([]string, e)
	ope := make([]string, e)
	kk := make([]int, e)
	for i := 0; i != e; i++ {
		var c operator
		fmt.Scan(&c.Name2, &c.operation, &c.k)
		na[i] = c.Name2
		ope[i] = c.operation
		kk[i] = c.k
	}
	Calculate(na, ope, kk, Products)
	keys := make([]string, 0, n)
	for key := range Products {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Printf("%s %d %.2f\n", Products[key].Name, Products[key].Quantity, Products[key].Price*float64(Products[key].Quantity))

	}
}
func Calculate(a, b []string, c []int, d map[string]product) map[string]product {
	for i := 0; i != len(c); i++ {
		if p, exists := d[a[i]]; exists {
			if b[i] == "restock" {
				p.Quantity += c[i]
			} else if b[i] == "sell" {
				if p.Quantity >= c[i] {
					p.Quantity -= c[i]
				}
			}
			d[a[i]] = p
		}

	}

	return d
}
