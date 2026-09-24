package library

import "fmt"

type Basket struct {
	name    string
	color   []Clothing
	content string
}

type Clothing struct {
	name  string
	color string
	state bool
}

func InitPanierBlanc() {
	basket := Basket{
		name:    "Panier Blanc",
		color:   []Clothing{{name: "Blanc", color: "Blanc", state: true}},
		content: "Contient des vêtements blancs",
	}
	fmt.Println(basket)
}

func InitPanierNoir() {
	basket := Basket{
		name:    "Panier Noir",
		color:   []Clothing{{name: "Noir", color: "Noir", state: true}},
		content: "Contient des vêtements noirs",
	}
	fmt.Println(basket)
}

func InitPanierCouleur() {
	basket := Basket{
		name:    "Panier Couleur",
		color:   []Clothing{{name: "Couleur", color: "Couleur", state: true}},
		content: "Contient des vêtements de toutes les couleurs",
	}
	fmt.Println(basket)
}
