package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github-analyzer/internal/database"
	"github-analyzer/internal/models"
	"github-analyzer/internal/services"

	conectionapigithub "github-analyzer/conectionAPIgithub"
)

type GithubUser struct {
	Login       string `json:"login"`
	Followers   int    `json:"followers"`
	Repos       int    `json:"repos"`
	Language    string `json:"language"`
	PopularRepo string `json:"popularRepo"`
}

type CompareResponse struct {
	User1 GithubUser `json:"user1"`
	User2 GithubUser `json:"user2"`
}

func main() {

	database.ConnectMongo()

	reader := bufio.NewReader(os.Stdin)

	for {

		fmt.Println()
		fmt.Println("=================================")
		fmt.Println("       GITHUB ANALYZER")
		fmt.Println("=================================")
		fmt.Println()
		fmt.Println("1. Login")
		fmt.Println("2. Crear Cuenta")
		fmt.Println("3. Salir")
		fmt.Println()

		fmt.Print("Seleccione una opción: ")

		option, _ := reader.ReadString('\n')

		option = strings.TrimSpace(option)

		switch option {

		case "1":
			loginMenu(reader)

		case "2":
			createAccountMenu(reader)

		case "3":
			fmt.Println("Hasta luego")
			return

		default:
			fmt.Println("Opción inválida")
		}
	}
}
func loginMenu(reader *bufio.Reader) {

	fmt.Println()
	fmt.Println("=== LOGIN ===")

	fmt.Print("Usuario: ")
	username, _ := reader.ReadString('\n')

	fmt.Print("Contraseña: ")
	password, _ := reader.ReadString('\n')

	username = strings.TrimSpace(username)
	password = strings.TrimSpace(password)

	ok := services.Login(
		username,
		password,
	)

	if !ok {

		fmt.Println()
		fmt.Println("Credenciales incorrectas")

		return
	}

	fmt.Println()
	fmt.Println("Login exitoso")

	githubAnalyzerMenu(reader)
}
func githubAnalyzerMenu(reader *bufio.Reader) {

	for {

		fmt.Println()
		fmt.Println("=================================")
		fmt.Println("      GITHUB ANALYZER")
		fmt.Println("=================================")
		fmt.Println()

		fmt.Println("1. Buscar Usuario GitHub")
		fmt.Println("2. Comparar Usuarios")
		fmt.Println("3. Cerrar Sesión")

		fmt.Print("Seleccione una opción: ")

		option, _ := reader.ReadString('\n')

		option = strings.TrimSpace(option)

		switch option {

		case "1":

			fmt.Print(
				"Ingrese usuario GitHub: ",
			)

			username, _ :=
				reader.ReadString('\n')

			username =
				strings.TrimSpace(username)

			fmt.Println()

			fmt.Printf(
				"Buscando %s...\n",
				username,
			)

			var user = conectionapigithub.SearchUser(username)
			conectionapigithub.DescriptionUser(user)

			var repos = conectionapigithub.RepositoryListDeclaration(user.ReposUrl)
			conectionapigithub.RepositoryDescriptionList(repos)

			languages := conectionapigithub.CountLanguages(repos)
			fmt.Println("")
			for language, count := range languages {
				fmt.Printf("%s: %d\n", language, count)
			}

		case "2":

			compareUsersCLI(reader)

		case "3":

			return

		default:

			fmt.Println(
				"Opción inválida",
			)
		}
	}
}
func createAccountMenu(reader *bufio.Reader) {

	fmt.Println()
	fmt.Println("=== CREAR CUENTA ===")

	fmt.Print("Usuario: ")
	username, _ := reader.ReadString('\n')

	fmt.Print("Contraseña: ")
	password, _ := reader.ReadString('\n')

	fmt.Print("Nombre: ")
	name, _ := reader.ReadString('\n')

	fmt.Print("Apellido: ")
	lastname, _ := reader.ReadString('\n')

	fmt.Print("Email: ")
	email, _ := reader.ReadString('\n')

	fmt.Print("Fecha Nacimiento: ")
	birthDate, _ := reader.ReadString('\n')

	user := models.User{

		Username: strings.TrimSpace(username),

		Password: strings.TrimSpace(password),

		Name: strings.TrimSpace(name),

		LastName: strings.TrimSpace(lastname),

		GithubEmail: strings.TrimSpace(email),

		BirthDate: strings.TrimSpace(birthDate),
	}

	_, err := services.CreateUser(user)

	if err != nil {

		fmt.Println()
		fmt.Println(
			"Error al crear usuario:",
			err,
		)

		return
	}

	fmt.Println()
	fmt.Println(
		"Usuario creado correctamente",
	)
}

func compareUsersCLI(reader *bufio.Reader) {

	fmt.Println()
	fmt.Println("=== COMPARAR USUARIOS ===")

	fmt.Print("Primer usuario GitHub: ")
	user1Name, _ := reader.ReadString('\n')

	fmt.Print("Segundo usuario GitHub: ")
	user2Name, _ := reader.ReadString('\n')

	user1Name = strings.TrimSpace(user1Name)
	user2Name = strings.TrimSpace(user2Name)

	user1 :=
		conectionapigithub.SearchUser(user1Name)

	user2 :=
		conectionapigithub.SearchUser(user2Name)

	repos1 :=
		conectionapigithub.RepositoryListDeclaration(
			user1.ReposUrl,
		)

	repos2 :=
		conectionapigithub.RepositoryListDeclaration(
			user2.ReposUrl,
		)

	lang1 :=
		conectionapigithub.CountLanguages(
			repos1,
		)

	lang2 :=
		conectionapigithub.CountLanguages(
			repos2,
		)

	var topLang1 string
	var topCount1 int

	for language, count := range lang1 {

		if count > topCount1 {

			topCount1 = count
			topLang1 = language
		}
	}

	var topLang2 string
	var topCount2 int

	for language, count := range lang2 {

		if count > topCount2 {

			topCount2 = count
			topLang2 = language
		}
	}

	topRepo1 :=
		conectionapigithub.TopRepository(
			repos1,
		)

	topRepo2 :=
		conectionapigithub.TopRepository(
			repos2,
		)

	fmt.Println()
	fmt.Println("========================================")
	fmt.Println("           COMPARACIÓN")
	fmt.Println("========================================")
	fmt.Println()

	fmt.Printf(
		"%-18s %-15s %-15s\n",
		"",
		user1.Login,
		user2.Login,
	)

	fmt.Printf(
		"%-18s %-15d %-15d\n",
		"Followers",
		user1.Followers,
		user2.Followers,
	)

	fmt.Printf(
		"%-18s %-15d %-15d\n",
		"Repositorios",
		user1.PublicRepos,
		user2.PublicRepos,
	)

	fmt.Printf(
		"%-18s %-15s %-15s\n",
		"Lenguaje",
		topLang1,
		topLang2,
	)

	fmt.Printf(
		"%-18s %-15s %-15s\n",
		"Repo Popular",
		topRepo1.Name,
		topRepo2.Name,
	)

	fmt.Println()
	fmt.Println("========================================")
}
