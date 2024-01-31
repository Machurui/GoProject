package main

import (
	"fmt"
	"log"
	"os"
	"projet/client"
	"projet/databases"
	"projet/dossiers"
	"projet/fichiers"
)

const path = "C:\\GoEstiamProjet\\src\\data\\"

const config = "online"

type Manager interface {
	CreateFolder(name string) error
	ReadFolder(name string) error
	RenameFolder(oldName, newName string) error
	DeleteFolder(name string) error
	CreateFile(name, text string) error
	ReadFile(name string) error
	RenameFile(oldName, newName string) error
	DeleteFile(name string) error
	UpdateText(name, text string) error
	History() error
}

// Implémentation Offline
type OfflineManager struct{}

func (fm OfflineManager) CreateFolder(name string) error {
	// Créer un dossier
	folderPath, err := dossiers.CreateFolder(name, path)
	if err != nil {
		fmt.Println("Erreur lors de la création du dossier :", err)
	} else {
		fmt.Println("Voici le path du nouveau dossier :", folderPath)
	}

	return nil
}

func (fm OfflineManager) ReadFolder(name string) error {
	// Lire le dossier
	datas, err := dossiers.ReadFolder(name, path)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du dossier :", err)
	} else {
		fmt.Println("Contenu du dossier :", datas)
	}

	return nil
}

func (fm OfflineManager) RenameFolder(oldName, newName string) error {
	// Renommer un dossier
	folderPath, err := dossiers.RenameFolder(oldName, newName, path)
	if err != nil {
		fmt.Println("Erreur lors du renommage du dossier :", err)
	} else {
		fmt.Println("Voici le nouveau path du dossier :", folderPath)
	}

	return nil
}

func (fm OfflineManager) DeleteFolder(name string) error {
	// Delete un dossier
	folderPath, err := dossiers.DeleteFolder(name, path)
	if err != nil {
		fmt.Println("Erreur lors de la suppression du dossier :", err)
	} else {
		fmt.Println("Ancien path du dossier :", folderPath)
	}

	return nil
}

func (fm OfflineManager) CreateFile(name, text string) error {
	// Créer un fichier
	filePath, err := fichiers.CreateFile(name, text, path)
	if err != nil {
		fmt.Println("Erreur lors de la création du fichier :", err)
	} else {
		fmt.Println("Voici le path du nouveau fichier :", filePath)
	}

	return nil
}

func (fm OfflineManager) ReadFile(name string) error {
	// Lire un fichier
	content, err := fichiers.ReadFile(name, path)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier :", err)
	} else {
		fmt.Println("Voici le contenu du fichier :", content)
	}

	return nil
}

func (fm OfflineManager) RenameFile(oldName, newName string) error {
	// Renommer un fichier
	filePath, err := fichiers.UpdateNameFile(oldName, newName, path)
	if err != nil {
		fmt.Println("Erreur lors du renommage du fichier :", err)
	} else {
		fmt.Println("Voici le path du nouveau fichier :", filePath)
	}

	return nil
}

func (fm OfflineManager) UpdateText(name, text string) error {
	// Ajouter du texte dans un fichier
	content, err := fichiers.UpdateTextFile(name, text, path)
	if err != nil {
		fmt.Println("Erreur lors de l'écriture du fichier :", err)
	} else {
		fmt.Println("Voici le contenu du fichier :", content)
	}

	return nil
}

func (fm OfflineManager) DeleteFile(name string) error {
	// Ajouter du texte dans un fichier
	filePath, err := fichiers.DeleteFile(name, path)
	if err != nil {
		fmt.Println("Erreur lors de la suppression du fichier :", err)
	} else {
		fmt.Println("Voici le path du fichier supprimé :", filePath)
	}

	return nil
}

func (fm OfflineManager) History() error {
	databases.ConnectDataBase()

	journaux, err := databases.LastJournal()
	if err != nil {
		log.Fatal(err)
	}

	if len(journaux) > 0 {
		fmt.Printf("Voici l'historique des 50 dernières commandes :\n\n")
		for _, entry := range journaux {
			fmt.Println(entry.ID, " | ", entry.DH, " | ", entry.MF, " | ", entry.Argument, " | ", entry.Statut)
		}
	}

	return nil
}

// // Implémentation Online
type OnlineManager struct{}

func (fm OnlineManager) CreateFolder(name string) error {
	folderPath, err := client.CreateFolder(name)
	if err != nil {
		fmt.Println("Erreur lors de la création du dossier :", err)
	} else {
		fmt.Println("Voici le path du nouveau dossier :", folderPath)
	}

	return nil
}

func (fm OnlineManager) ReadFolder(name string) error {
	datas, err := client.ReadFolder(name)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du dossier :", err)
	} else {
		fmt.Println("Contenu du dossier :", datas)
	}

	return nil
}

func (fm OnlineManager) RenameFolder(oldName, newName string) error {
	folderPath, err := client.RenameFolder(oldName, newName)
	if err != nil {
		fmt.Println("Erreur lors du renommage du dossier :", err)
	} else {
		fmt.Println("Voici le nouveau path du dossier :", folderPath)
	}

	return nil
}

func (fm OnlineManager) DeleteFolder(name string) error {
	folderPath, err := client.DeleteFolder(name)
	if err != nil {
		fmt.Println("Erreur lors de la suppression du dossier :", err)
	} else {
		fmt.Println("Ancien path du dossier :", folderPath)
	}

	return nil
}

func (fm OnlineManager) CreateFile(name, text string) error {
	filePath, err := client.CreateFile(name, text)
	if err != nil {
		fmt.Println("Erreur lors de la création du fichier :", err)
	} else {
		fmt.Println("Voici le path du nouveau fichier :", filePath)
	}

	return nil
}

func (fm OnlineManager) ReadFile(name string) error {
	datas, err := client.ReadFile(name)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier :", err)
	} else {
		fmt.Println("Nouveau contenu du fichier :", datas)
	}

	return nil
}

func (fm OnlineManager) RenameFile(oldName, newName string) error {
	filePath, err := client.UpdateNameFile(oldName, newName)
	if err != nil {
		fmt.Println("Erreur lors du renommage du fichier :", err)
	} else {
		fmt.Println("Nouveau path du fichier :", filePath)
	}

	return nil
}

func (fm OnlineManager) UpdateText(name, text string) error {
	datas, err := client.UpdateTextFile(name, text)
	if err != nil {
		fmt.Println("Erreur lors de l'écriture du fichier :", err)
	} else {
		fmt.Println("Contenu du fichier :", datas)
	}

	return nil
}

func (fm OnlineManager) DeleteFile(name string) error {
	filePath, err := client.DeleteFile(name)
	if err != nil {
		fmt.Println("Erreur lors de la suppression du fichier :", err)
	} else {
		fmt.Println("Ancien path du fichier :", filePath)
	}

	return nil
}

func (fm OnlineManager) History() error {
	journaux, err := client.Hist()
	if err != nil {
		fmt.Println("Erreur lors de l'affichage de l'historique :", err)
	}

	if len(journaux) > 0 {
		fmt.Printf("Voici l'historique des 50 dernières commandes :\n\n")
		for _, entry := range journaux {
			fmt.Println(entry.ID, " | ", entry.DH, " | ", entry.MF, " | ", entry.Argument, " | ", entry.Statut)
		}
	}

	return nil
}

func main() {
	if len(os.Args) > 1 {
		var manager Manager

		// Choix de l'implémentation basé sur la configuration
		if config == "offline" {
			manager = OfflineManager{}
		} else {
			manager = OnlineManager{}
		}

		commandName := os.Args[1]

		switch commandName {
		case "dir":
			if len(os.Args) > 2 {
				sousCommandName := os.Args[2]
				switch sousCommandName {
				case "create":
					if len(os.Args) > 3 {
						manager.CreateFolder(os.Args[3])
					} else {
						fmt.Println("Nom du dossier manquant")
					}

				case "read":
					if len(os.Args) > 3 {
						manager.ReadFolder(os.Args[3])
					} else {
						fmt.Println("Nom du dossier manquant")
					}

				case "rename":
					if len(os.Args) > 4 {
						manager.RenameFolder(os.Args[3], os.Args[4])
					} else {
						fmt.Println("Nom des dossiers manquant")
					}

				case "delete":
					if len(os.Args) > 3 {
						manager.DeleteFolder(os.Args[3])
					} else {
						fmt.Println("Nom du dossier manquant")
					}
				default:
					fmt.Println("Commande inconnue")
				}
			} else {
				fmt.Println("Veuillez saisir une sous-commande.")
			}

		case "file":
			if len(os.Args) > 2 {
				sousCommandName := os.Args[2]
				switch sousCommandName {
				case "create":
					if len(os.Args) > 4 {
						manager.CreateFile(os.Args[3], os.Args[4])
					} else {
						fmt.Println("Le nom du fichier ou le texte est manquant")
					}

				case "read":
					if len(os.Args) > 3 {
						manager.ReadFile(os.Args[3])
					} else {
						fmt.Println("Nom du fichier manquant")
					}

				case "rename":
					if len(os.Args) > 4 {
						manager.RenameFile(os.Args[3], os.Args[4])
					} else {
						fmt.Println("Nom des fichiers manquants")
					}

				case "update":
					if len(os.Args) > 4 {
						manager.UpdateText(os.Args[3], os.Args[4])
					} else {
						fmt.Println("Le nom du fichier ou le texte est manquant")
					}

				case "delete":
					if len(os.Args) > 3 {
						manager.DeleteFile(os.Args[3])
					} else {
						fmt.Println("Nom du fichier manquant")
					}
				default:
					fmt.Println("Commande inconnue")
				}
			} else {
				fmt.Println("Veuillez saisir une sous-commande.")
			}

		case "help":

			// Liste de toutes les commandes disponibles
			if len(os.Args) > 2 {
				switch os.Args[2] {
				case "dir":
					// Commande dir
					command := []string{"create", "read", "rename", "delete"}
					if len(command) > 0 {
						fmt.Println("Voici les sous-commandes disponibles pour la commande dir:")
						for _, entry := range command {
							fmt.Println("-", entry)
						}
					}

				case "file":
					// Commande file
					command := [...]string{"create", "read", "rename", "update", "delete"}
					if len(command) > 0 {
						fmt.Println("Voici les sous-commandes disponibles pour la commande file:")
						for _, entry := range command {
							fmt.Println("-", entry)
						}
					}

				default:
					fmt.Println("Aucune commande ne correspond à votre saisie.")

				}
			} else if len(os.Args) == 2 {
				// Commande de base
				command := [...]string{"dir", "file", "hist"}
				if len(command) > 0 {
					fmt.Println("Voici les commandes disponibles :")
					for _, entry := range command {
						fmt.Println("-", entry)
					}
				}
			}

		case "hist":
			manager.History()

		default:
			fmt.Println("Commande inconnue")
		}

	} else {
		fmt.Println("Exécutez la commande \"help\" pour obtenir de l'aide.")
	}
}
