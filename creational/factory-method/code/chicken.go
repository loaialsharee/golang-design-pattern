package main

type Chicken struct {
    Shawarma
}

func newChicken() IShawarma {
    return &Chicken{
        Shawarma: Shawarma{
            name:  "Chicken Shawarma",
			price: 20.50,
			stars: 5,
        },
    }
}