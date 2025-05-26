package main

import (
	"flag"
	"fmt"
	"os"
	"tp1-annuaire/annuaire"
)

var action, nom, tel string

func main() {
	flag.StringVar(&action, "action", "", "Action à exécuter")
	flag.StringVar(&nom, "nom", "", "Nom")
	flag.StringVar(&tel, "tel", "", "Téléphone")
	flag.Parse()

	switch action {
	case "ajouter":
		annuaire.Ajouter(nom, tel)
	case "supprimer":
		annuaire.Supprimer(nom)
	case "modifier":
		annuaire.Modifier(nom, tel)
	case "rechercher":
		annuaire.Rechercher(nom)
	case "lister":
		annuaire.Lister()
	default:
		fmt.Println("action invalide")
		os.Exit(1)
	}
}
