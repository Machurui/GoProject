package main

import (
	"net/http"
	"projet/databases"
	"projet/dossiers"
	"projet/fichiers"

	"github.com/gin-gonic/gin"
)

const path = "C:\\GoEstiamProjet\\src\\data\\"

func main() {
	r := gin.Default()

	// Groupe pour les opérations sur les dossiers
	dossierGroup := r.Group("/dossiers")
	{
		dossierGroup.POST("/", createFolder)
		dossierGroup.GET("/:name", getFolder)
		dossierGroup.PUT("/:name", renameFolder)
		dossierGroup.DELETE("/:name", deleteFolder)
	}

	// Groupe pour les opérations sur les fichiers
	fichierGroup := r.Group("/fichiers")
	{
		fichierGroup.POST("/", createFile)
		fichierGroup.GET("/:name", getFile)
		fichierGroup.PUT("/rename/:name", renameFile)
		fichierGroup.PUT("/update/:name", updateTextFile)
		fichierGroup.DELETE("/:name", deleteFile)
	}

	// Groupe pour les opérations sur les diverses requêtes
	diversGroup := r.Group("/divers")
	{
		diversGroup.GET("/hist", historiqueCommand)
	}

	// Démarrer le serveur Gin sur le port 8080
	r.Run(":8080")
}

// Handlers pour les opérations sur les dossiers
func createFolder(c *gin.Context) {
	var requestData struct {
		Name string `json:"name"`
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if requestData.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Le nom du dossier ne peut pas être vide"})
		return
	}

	folderPath, err := dossiers.CreateFolder(requestData.Name, path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Dossier créé avec succès", "folderPath": folderPath})
}

func getFolder(c *gin.Context) {
	name := c.Param("name")

	// Vérifier si le nom n'est pas vide
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Le nom du dossier ne peut pas être vide"})
		return
	}

	// Le chemin doit être celui où les dossiers sont stockés
	content, err := dossiers.ReadFolder(name, path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, content)
}

func renameFolder(c *gin.Context) {
	oldName := c.Param("name")

	var requestData struct {
		NewName string `json:"name"` // Nouveau nom pour le dossier
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Le corps de la requête n'est pas valide"})
		return
	}

	// Vérifier si le nouveau nom n'est pas vide
	if requestData.NewName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Le nouveau nom ne peut pas être vide"})
		return
	}

	folderPath, err := dossiers.RenameFolder(oldName, requestData.NewName, path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Dossier renommé avec succès", "folderPath": folderPath})
}

func deleteFolder(c *gin.Context) {
	name := c.Param("name")

	// Vérifier si le nom n'est pas vide
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Le nom du dossier ne peut pas être vide"})
		return
	}

	folderPath, err := dossiers.DeleteFolder(name, path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Dossier supprimé avec succès", "folderPath": folderPath})
}

// Handlers pour les opérations sur les fichiers
func createFile(c *gin.Context) {
	var requestData struct {
		Name    string `json:"name"`    // Nom du fichier à créer
		Content string `json:"content"` // Contenu du fichier
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Vérifier si le nom du fichier n'est pas vide
	if requestData.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Le nom du fichier ne peut pas être vide"})
		return
	}

	// Appeler la fonction CreateFile du package fichiers
	filePath, err := fichiers.CreateFile(requestData.Name, requestData.Content, path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Fichier créé avec succès", "filePath": filePath})
}

func getFile(c *gin.Context) {
	name := c.Param("name")

	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Le nom du fichier ne peut pas être vide"})
		return
	}

	content, err := fichiers.ReadFile(name, path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"content": content})
}

func renameFile(c *gin.Context) {
	var requestData struct {
		NewName string `json:"name"` // Nouveau nom pour le fichier
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Le corps de la requête n'est pas valide"})
		return
	}

	// Récupérer l'ancien nom du fichier à partir du paramètre d'URL
	oldName := c.Param("name")
	if oldName == "" || requestData.NewName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Les anciens et nouveaux noms ne peuvent pas être vides"})
		return
	}

	filePath, err := fichiers.UpdateNameFile(oldName, requestData.NewName, path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Fichier renommé avec succès", "filePath": filePath})
}

func updateTextFile(c *gin.Context) {
	var requestData struct {
		Content string `json:"content"` // Nouveau contenu du fichier
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Le corps de la requête n'est pas valide"})
		return
	}

	// Récupérer le nom du fichier à partir du paramètre d'URL
	name := c.Param("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Le nom du fichier ne peut pas être vide"})
		return
	}

	content, err := fichiers.UpdateTextFile(name, requestData.Content, path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Contenu du fichier mis à jour avec succès", "content": content})
}

func deleteFile(c *gin.Context) {
	name := c.Param("name")

	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Le nom du fichier ne peut pas être vide"})
		return
	}

	filePath, err := fichiers.DeleteFile(name, path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Fichier supprimé avec succès", "filePath": filePath})
}

func historiqueCommand(c *gin.Context) {

	databases.ConnectDataBase()

	journaux, err := databases.LastJournal()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, journaux)
}
