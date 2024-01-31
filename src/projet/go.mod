module projet

go 1.21.6

replace projet/fichiers => ./fichiers

replace projet/dossiers => ./dossiers

replace projet/databases => ./databases

replace projet/server => ./server

replace projet/client => ./client

require (
	projet/client v0.0.0-00010101000000-000000000000
	projet/databases v0.0.0-00010101000000-000000000000
	projet/dossiers v0.0.0-00010101000000-000000000000
	projet/fichiers v0.0.0-00010101000000-000000000000
)

require github.com/go-sql-driver/mysql v1.7.1 // indirect
