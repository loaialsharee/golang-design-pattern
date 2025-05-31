package main

type Beef struct {
    Shawarma
}

func newBeef() IShawarma {
    return &Beef{
        Shawarma: Shawarma{
            name:  "Beef Shawarma",
			price: 16.20,
			stars: 3,
        },
    }
}