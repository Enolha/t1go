package annuaire

import (
	"fmt"
)

var carnet = make(map[string]string)

func Ajouter(nom, tel string) {
	if _, existe := carnet[nom]; existe {
		fmt.Printf("Le contact '%s' existe déjà.\n", nom)
		return
	}
	carnet[nom] = tel
	fmt.Printf("%s a bien été ajouté avec le numéro %s.\n", nom, tel)
}

func Supprimer(nom string) {
	if _, existe := carnet[nom]; !existe {
		fmt.Printf("Impossible de supprimer : aucun contact trouvé avec le nom '%s'.\n", nom)
		return
	}
	delete(carnet, nom)
	fmt.Printf("%s a été supprimé de l'annuaire.\n", nom)
}

func Modifier(nom, tel string) {
	if _, existe := carnet[nom]; !existe {
		fmt.Printf("Impossible de modifier : '%s' n'existe pas dans l'annuaire.\n", nom)
		return
	}
	carnet[nom] = tel
	fmt.Printf("Le numéro de %s a été mis à jour avec succès.\n", nom)
}

func Rechercher(nom string) {
	tel, existe := carnet[nom]
	if !existe {
		fmt.Printf("Aucun contact trouvé pour '%s'.\n", nom)
		return
	}
	fmt.Printf("%s : %s\n", nom, tel)
}

func Lister() {
	if len(carnet) == 0 {
		fmt.Println("L'annuaire est vide")
		return
	}
	fmt.Println("Voici la liste des contacts :")
	for nom, tel := range carnet {
		fmt.Printf("- %s : %s\n", nom, tel)
	}
}
