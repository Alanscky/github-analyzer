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

			fmt.Println(
				"Comparación pendiente",
			)

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
