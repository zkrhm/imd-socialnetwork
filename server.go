package main

func main() {
	app := NewApp()
	app.Initialize()

	app.Run("localhost:8000")
}
