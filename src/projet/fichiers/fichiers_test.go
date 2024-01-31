package fichiers

import (
	"os"
	"testing"
)

const path = "C:\\GoEstiamProjet\\src\\test\\"

func TestCreateFile_Success(t *testing.T) {
	// Nom du fichier pour le test
	testFileName := "TestCreateFile_Success.txt"
	textTest := "test"

	filePath := path + testFileName

	// Appel de la fonction CreateFile
	_, err := CreateFile(testFileName, textTest, path)

	if err != nil {
		t.Error("Erreur :", err)
	}

	// Vérification de l'existence du fichier
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Errorf("Fonction de création de fichier ne marche pas. %s", testFileName)
	}

	// Nettoyage: Supprimer le fichier après le test
	os.Remove(filePath)
}

func TestCreateFile_FileExist(t *testing.T) {
	// Nom du fichier pour le test
	testFileName := "TestCreateFile_FileExist.txt"
	textTest := "test"
	filePath := path + testFileName
	file, _ := os.Create(filePath)
	file.Close()

	// Appel de la fonction CreateFile
	_, err := CreateFile(testFileName, textTest, path)

	if err != nil {
		t.Error("Erreur :", err)
	}
	// Vérification de l'existence du fichier
	if _, err2 := os.Stat(filePath); err2 == nil {
		t.Error("Le Fichier existe déjà")
	}

	// Nettoyage: Supprimer le fichier après le test
	os.Remove(filePath)
}

func TestReadFile_Success(t *testing.T) {
	// Nom du fichier pour le test
	testFileName := "TestReadFile_Success.txt"
	textTest := "test"

	filePath := path + testFileName
	file, _ := os.Create(filePath)
	file.Close()

	//Ecriture dans le fichier
	os.WriteFile(filePath, []byte(textTest), 0644)

	// Appel de la fonction ReadFile
	data, err := ReadFile(testFileName, path)

	if err != nil {
		t.Error("Erreur :", err)
	}

	// Vérification de l'existence du fichier
	if data != textTest {
		t.Error("La fonction read ne marche pas")
	}

	// Nettoyage: Supprimer le fichier après le test
	os.Remove(filePath)
}

func TestReadFile_FileNotExist(t *testing.T) {
	// Nom du fichier pour le test
	testFileName := "TestReadFile_FileNotExist.txt"

	// Appel de la fonction ReadFile
	_, err := ReadFile(testFileName, path)

	// Vérification de l'existence du fichier
	if err != nil {
		t.Error("erreur :", err)
	}
}

func TestUpdateTextFile_Success(t *testing.T) {
	// Nom du fichier pour le test
	testFileName := "TestUpdateTextFile_Success.txt"
	textTest := "test"

	filePath := path + testFileName
	file, _ := os.Create(filePath)
	file.Close()

	//Ecriture dans le fichier
	dataWrite, err := UpdateTextFile(testFileName, textTest, path)

	if err != nil {
		t.Error("Erreur :", err)
	}

	data, _ := os.ReadFile(filePath)

	// Vérification de l'existence du fichier
	if string(data) != dataWrite {
		t.Error("Le fichier n'existe pas donc l'updateText n'écrit pas")
	}

	// Nettoyage: Supprimer le fichier après le test
	os.Remove(filePath)
}

func TestUpdateTextFile_FileNotExist(t *testing.T) {
	// Nom du fichier pour le test
	testFileName := "TestUpdateTextFile_FileNotExist.txt"
	textTest := "test"
	filePath := path + testFileName

	//ecriture data
	dataWrite, err := UpdateTextFile(testFileName, textTest, path)

	if err != nil {
		t.Error("Erreur :", err)
	}
	// Appel de la fonction ReadFile
	data, _ := os.ReadFile(filePath)

	// Vérification de l'existence du fichier
	if string(data) != dataWrite {
		t.Error("Le fichier n'existe pas")
	}
}

func TestUpdateNameFile_Success(t *testing.T) {
	// Nom du fichier pour le test
	oldTestFileName := "TestUpdateNameFile_Success.txt"
	newNameTest := "newTestFile.txt"

	filePath := path + oldTestFileName
	file, _ := os.Create(filePath)
	file.Close()

	//Ecriture dans le fichier
	_, err := UpdateNameFile(oldTestFileName, newNameTest, path)
	newfilePath := path + newNameTest

	if err != nil {
		t.Error("Erreur :", err)
	}

	// Vérification de l'existence du fichier
	if _, err := os.Stat(newfilePath); os.IsNotExist(err) {
		t.Error("La fonction updateNameFile ne fonctionne pas")
	}

	// Nettoyage: Supprimer le fichier après le test
	os.Remove(filePath)
	os.Remove(newfilePath)
}

func TestUpdateNameFile_NotExist(t *testing.T) {
	// Nom du fichier pour le test
	oldTestFileName := "TestUpdateNameFile_NotExist.txt"
	newNameTest := "newTestFile.txt"

	filePath := path + oldTestFileName

	//Rename dans le fichier
	_, err := UpdateNameFile(oldTestFileName, newNameTest, path)
	newfilePath := path + newNameTest

	if err != nil {
		t.Error("Erreur :", err)
	}

	// Vérification de l'existence du fichier
	if _, err := os.Stat(oldTestFileName); os.IsNotExist(err) {
		t.Error("Le fichier n'existe pas")
	}

	// Nettoyage: Supprimer les fichiers après le test
	os.Remove(filePath)
	os.Remove(newfilePath)
}

func TestDeleteFile_Success(t *testing.T) {
	// Nom du fichier pour le test
	testFileName := "TestDeleteFile_Success.txt"

	filePath := path + testFileName
	file, _ := os.Create(filePath)
	file.Close()

	// Appel de la fonction DeleteFile
	_, err := DeleteFile(testFileName, path)
	if err != nil {
		t.Error("Erreur :", err)
	}

	// Vérification de l'existence du fichier
	if _, err2 := os.Stat(filePath); err2 == nil {
		t.Errorf("Fonction de supression de fichier ne marche pas. %s", testFileName)
	}

	// Nettoyage: Supprimer le fichier après le test
	os.Remove(filePath)
}

func TestDeleteFile_NotExist(t *testing.T) {
	// Nom du fichier pour le test
	testFileName := "TestDeleteFile_NotExist.txt"

	filePath := path + testFileName

	// Appel de la fonction DeleteFile
	_, err := DeleteFile(testFileName, path)
	if err != nil {
		t.Error("Erreur :", err)
	}
	// Vérification de l'existence du fichier
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Errorf("Le fichier n'existe pas donc pas de supression %s", testFileName)
	}

	// Nettoyage: Supprimer le fichier après le test
	os.Remove(filePath)
}
