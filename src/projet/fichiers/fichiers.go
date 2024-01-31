package fichiers

import (
	"errors"
	"fmt"
	"os"
	"projet/databases"
	"strings"
	"time"
)

const module = "file"

func logCommand(command, argument, statut string) {
	databases.ConnectDataBase()
	mf := fmt.Sprintf("%s %s", module, command)
	log := databases.LogData{
		DH:       time.Now(),
		MF:       mf,
		Argument: argument,
		Statut:   statut,
	}

	_, err := databases.AddLog(log)
	if err != nil {
		fmt.Println(err)
	}
}

func containsNoSpecificChars(s string) bool {
	// Retourne `false` si `s` contient au moins un des caractères dans `chars`
	chars := "/:*?\"<>|"

	return strings.ContainsAny(s, chars)
}

func CreateFile(name string, text string, path string) (string, error) {
	command := "create"
	filePath := path + name
	// Si le nom du dossier contient des caractères bloquant
	if containsNoSpecificChars(name) {
		logCommand(command, name+" "+filePath+" "+text, "La chaîne contient au moins un caractère bloquant.")
		return "", errors.New("la chaîne contient au moins un caractère bloquant")
	}

	// Vérifie si le fichier existe déjà
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if err != nil {
			logCommand(command, name+" "+filePath+" "+text, "impossible de créer le fichier")
			return "", errors.New("impossible de créer le fichier")
		}
		defer file.Close()

		// Si du texte est fourni, écrivez-le dans le fichier
		if text != "" {
			errorWrite := os.WriteFile(filePath, []byte(text), 0644)
			if errorWrite != nil {
				logCommand(command, name+" "+filePath+" "+text, "le fichier ne s'est pas remplie")
				return "", errors.New("le fichier ne s'est pas remplie")
			}
		}
		logCommand(command, name+" "+filePath+" "+text, "Fichier créé")
		fmt.Println("Fichier créé")
		
		return filePath, nil
	} else {
		logCommand(command, name+" "+filePath+" "+text, "le fichier existe deja")
		return "", errors.New("le fichier existe deja")
	}
}

func ReadFile(name string, path string) (string, error) {
	command := "read"
	filePath := path + name
	data, err := os.ReadFile(filePath)

	if err != nil {
		logCommand(command, name+" "+filePath, "ce fichier n'existe pas")
		return "", errors.New("ce fichier n'existe pas")

	}
	logCommand(command, name+" "+filePath, "Affichage du contenue du fichier")
	fmt.Println("Contenu du fichier:", string(data))
	return string(data), nil
}

func UpdateTextFile(name string, data string, path string) (string, error) {
	filePath := path + name
	command := "updatetext"
	info, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			logCommand(command, name+" "+filePath+" "+data, "aucun fichier n'existe avec ce nom")
			return "", errors.New("aucun fichier n'existe avec ce nom")
		}
		// Autres erreurs de système de fichiers
		logCommand(command, name+" "+filePath+" "+data, "impossible de vérifier l'existence du fichier")
		return "", errors.New("impossible de vérifier l'existence du fichier")
	}

	// Vérifier si le chemin est un dossier
	if info.IsDir() {
		logCommand(command, name+" "+filePath+" "+data, "le chemin correspond à un dossier, pas à un fichier")
		return "", errors.New("le chemin correspond à un dossier, pas à un fichier")
	}

	err = os.WriteFile(filePath, []byte(data), 0644)
	if err != nil {
		logCommand(command, name+" "+filePath+" "+data, "impossible de mettre à jour le fichier")
		return "", errors.New("impossible de mettre à jour le fichier")
	}
	logCommand(command, name+" "+filePath+" "+data, "Fichier mis à jour.")
	fmt.Println("Fichier mis à jour.")

	return string(data), nil
}

func UpdateNameFile(oldName string, newName string, path string) (string, error) {
	command := "rename"
	oldPath := path + oldName
	newPath := path + newName
	// Si le nom du dossier contient des caractères bloquant
	if containsNoSpecificChars(newName) {
		logCommand(command, oldName+" "+newName+" "+newPath, "la chaîne contient au moins un caractère bloquant")
		return "", errors.New("la chaîne contient au moins un caractère bloquant")
	}

	err := os.Rename(oldPath, newPath)
	if err != nil {
		logCommand(command, oldName+" "+newName+" "+newPath, "impossible de mettre à jour le nom fichier")
		return "", errors.New("impossible de mettre à jour le nom fichier")
	}
	logCommand(command, oldName+" "+newName+" "+newPath, "Nom fichier mis à jour.")
	fmt.Println("Nom fichier mis à jour.")

	return newPath, nil
}

func DeleteFile(name string, path string) (string, error) {
	filePath := path + name
	command := "delete"
	err := os.Remove(filePath)
	if err != nil {
		logCommand(command, name+" "+filePath, "impossible de supprimer le fichier")
		return "", errors.New("impossible de supprimer le fichier")

	}
	logCommand(command, name+" "+filePath, "Fichier supprimé.")
	fmt.Println("Fichier supprimé.")

	return filePath, nil
}
