package dossiers

import (
	"os"
	"testing"
)

const path = "C:\\GoEstiamProjet\\src\\test\\"

func TestCreateFolder(t *testing.T) {
	name := "TestCreateFolder"
	finalPath := path + name

	_, err := CreateFolder(name, path)
	if err != nil {
		t.Error("Erreur :", err)
	}

	os.RemoveAll(finalPath)
}

func TestCreateFolderChar(t *testing.T) {
	name := "TestCreateFolder|*"

	_, err := CreateFolder(name, path)
	if err != nil {
		t.Error("Erreur :", err)
	}
}

func TestReadFolder(t *testing.T) {
	name1 := "TestReadFolder"
	name2 := "TestReadFolder\\Alire"
	finalPath1 := path + name1
	finalPath2 := path + name2

	os.Mkdir(finalPath1, 0755)
	os.Mkdir(finalPath2, 0755)

	_, err := ReadFolder(name1, path)
	if err != nil {
		t.Error("Erreur :", err)
	}

	os.RemoveAll(finalPath1)
}

func TestReadFolderNotExist(t *testing.T) {
	name := "TestReadFolderNotExist"

	_, err := ReadFolder(name, path)
	if err != nil {
		t.Error("Erreur :", err)
	}
}

func TestUpdateFolder(t *testing.T) {
	name1 := "TestUpdateFolder1"
	name2 := "TestUpdateFolder2"
	finalPath1 := path + name1
	finalPath2 := path + name2

	os.Mkdir(finalPath1, 0755)

	_, err := RenameFolder(name1, name2, path)
	if err != nil {
		t.Error("Erreur :", err)
	}

	os.RemoveAll(finalPath2)
}

func TestUpdateFolderChar(t *testing.T) {
	name1 := "TestUpdateFolderAccent"
	name2 := "TestUpdateFolderAccent\\\""
	finalPath := path + name1

	os.Mkdir(finalPath, 0755)

	_, err := RenameFolder(name1, name2, path)
	if err != nil {
		t.Error("Erreur :", err)
	}

	os.RemoveAll(finalPath)
}

func TestDelete(t *testing.T) {
	name := "TestDelete"

	finalPath := path + name

	os.Mkdir(finalPath, 0755)

	_, err := DeleteFolder(name, path)
	if err != nil {
		t.Error("Erreur :", err)
	}

	os.RemoveAll(finalPath)
}

func TestDeleteNotExist(t *testing.T) {
	name := "TestDeleteNotExist"

	_, err := DeleteFolder(name, path)
	if err != nil {
		t.Error("Erreur :", err)
	}
}
