package main

func main() {
	server := APIServerInstance(":3000")
	server.Run()
}