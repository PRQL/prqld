package main

func main() {
  PopulateDatabasePool()
  PopulateTokenPool()

  defer CloseDatabaseConnections()

  StartServer(&Config{Port: 1999})
}
