package main

import "fmt"

func getShawarma(shawarmaType string) (IShawarma, error) {
    if shawarmaType == "chicken" {
        return newChicken(), nil
    }
    if shawarmaType == "beef" {
        return newBeef(), nil
    }
    return nil, fmt.Errorf("Wrong Shawarma type passed")
}