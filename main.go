package main

import (
	"fmt"
)

type Clothing struct {
	name  string
	color string
	state bool
}

type Basket struct {
	name    string
	color   string
	content []Clothing
}

type Washroom struct {
	white Basket
	black Basket
	color Basket
}

func (basket Basket) displayBasket() {
	fmt.Println("=====", basket.name, "=====")

	if len(basket.content) == 0 {
		fmt.Println("Le panier est vide")
		return
	}

	for _, clothing := range basket.content {
		fmt.Println("- ", clothing.name)
		fmt.Println("couleur : ", clothing.color)

		if clothing.state {
			fmt.Println("état : propre")
		} else {
			fmt.Println("état : sale")
		}
	}
}

func (basket *Basket) addToBasket(clothing Clothing) {

	if clothing.color != basket.color {
		fmt.Println("Oupss ! Ajout impossible la couleur du vêtement ne correspond pas à la corbeille")
		return
	}

	if clothing.state {
		fmt.Println("Oupss ! Ajout impossible le vêtement est déjà propre")
		return
	}

	basket.content = append(basket.content, clothing)

	fmt.Println("Ajout d'un nouveau vêtement :", clothing.name, "au", basket.name)
}

func (basket *Basket) cleanBasket() {
	count := 0

	fmt.Println("=== Nettoyage du", basket.name, "===")

	if len(basket.content) == 0 {
		fmt.Println("Impossible de nettoyer : le panier est vide")
		return
	}

	for i := range basket.content {
		if !basket.content[i].state {
			basket.content[i].state = true
			count++
		}
	}

	fmt.Println("Le lavage du", basket.name, "est terminé,", count, "vêtements ont été lavés")
}

func (basket *Basket) emptyCleanLaundry() {
	var dirtyClothes []Clothing
	var cleanClothes []Clothing

	fmt.Println("===== Vidage du", basket.name, "=====")

	if len(basket.content) == 0 {
		fmt.Println("Impossible de vider le panier car le panier est déjà vide")
		return
	}

	for _, clothing := range basket.content {
		if clothing.state {
			cleanClothes = append(cleanClothes, clothing)
		} else {
			dirtyClothes = append(dirtyClothes, clothing)
		}
	}

	if len(cleanClothes) == 0 {
		fmt.Println("Aucun vêtement propre à retirer")
		return
	}

	fmt.Println("Les vêtements suivants ont été retirés :")

	for _, clothing := range cleanClothes {
		fmt.Println("-", clothing.name)
	}

	basket.content = dirtyClothes
}

func (washroom Washroom) displayWashRoom() {
	fmt.Println("============================================================")
	fmt.Println("========== Toutes les corbeilles de la BUANDERIE ===========")
	fmt.Println("============================================================")

	washroom.white.displayBasket()
	washroom.black.displayBasket()
	washroom.color.displayBasket()
}

func main() {
	// Création des paniers
	basketWhite := Basket{"Panier blanc", "white", []Clothing{}}
	basketBlack := Basket{"Panier noir", "black", []Clothing{}}
	basketColor := Basket{"Panier couleur", "color", []Clothing{}}

	// Création des vêtements
	whiteShirt := Clothing{"T-shirt blanc", "white", false}
	whitePants := Clothing{"Pantalon blanc", "white", false}

	blackShirt := Clothing{"T-shirt noir", "black", false}
	blackPants := Clothing{"Pantalon noir", "black", false}

	colorShirt := Clothing{"T-shirt couleur", "color", false}
	colorShort := Clothing{"Short couleur", "color", false}

	cleanClothing := Clothing{"Pull propre", "white", true}

	fmt.Println("========== TEST 1 : PANIERS VIDES ==========")

	basketWhite.displayBasket()
	basketBlack.displayBasket()
	basketColor.displayBasket()

	fmt.Println("========== TEST 2 : AJOUTS VALIDES ==========")

	basketWhite.addToBasket(whiteShirt)
	basketWhite.addToBasket(whitePants)

	basketBlack.addToBasket(blackShirt)
	basketBlack.addToBasket(blackPants)

	basketColor.addToBasket(colorShirt)
	basketColor.addToBasket(colorShort)

	fmt.Println("========== TEST 3 : MAUVAISE COULEUR ==========")
	basketWhite.addToBasket(blackShirt)
	basketBlack.addToBasket(colorShirt)
	basketColor.addToBasket(whiteShirt)

	fmt.Println("========== TEST 4 : VÊTEMENT DÉJÀ PROPRE ==========")
	basketWhite.addToBasket(cleanClothing)

	fmt.Println("========== TEST 5 : AFFICHAGE DES PANIERS ==========")
	basketWhite.displayBasket()
	basketBlack.displayBasket()
	basketColor.displayBasket()

	fmt.Println("========== TEST 6 : NETTOYAGE ==========")
	basketColor.cleanBasket()
	basketColor.displayBasket()

	fmt.Println("========== TEST 7 : VIDAGE DU LINGE PROPRE ==========")
	basketColor.emptyCleanLaundry()
	basketColor.displayBasket()

	fmt.Println("========== TEST 8 : NETTOYAGE PANIER VIDE ==========")
	basketColor.cleanBasket()

	fmt.Println("========== TEST 9 : VIDAGE PANIER VIDE ==========")
	basketColor.emptyCleanLaundry()

	fmt.Println("========== TEST 10 : AUCUN VÊTEMENT PROPRE ==========")
	basketBlack.emptyCleanLaundry()

	fmt.Println("========== TEST 11 : BUANDERIE COMPLÈTE ==========")
	washroom := Washroom{
		white: basketWhite,
		black: basketBlack,
		color: basketColor,
	}
	washroom.displayWashRoom()

	fmt.Println("========== TEST 12 : NETTOYAGE COMPLET ==========")
	washroom.white.cleanBasket()
	washroom.black.cleanBasket()
	washroom.color.cleanBasket()
	washroom.displayWashRoom()

	fmt.Println("========== TEST 13 : VIDAGE COMPLET ==========")
	washroom.white.emptyCleanLaundry()
	washroom.black.emptyCleanLaundry()
	washroom.color.emptyCleanLaundry()
	washroom.displayWashRoom()
}
