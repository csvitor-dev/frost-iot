package view

import "fmt"

func ShowMenu() {
	fmt.Println("\n==============================")
	fmt.Println("         📋 MENU 📋          ")
	fmt.Println("==============================")
	fmt.Println("1️⃣  Receive Last Temperature Record")
	fmt.Println("2️⃣  Receive Last Stock Level Record")
	fmt.Println("3️⃣  Receive Last Open Port Record")
	fmt.Println("4️⃣  Configure Temperature Limit")
	fmt.Println("5️⃣  Configure Open Port Time")
	fmt.Println("6️⃣  Exit")
	fmt.Println("==============================")
	fmt.Print("👉 Enter your choice: ")
}
func GetUserChoice() int {
	var choice int
	fmt.Scan(&choice)
	return choice
}

func ShowTemperatureRecord(temp float64) {
	fmt.Printf("\n🌡️  Last Temperature Record: %.2f°C\n", temp)
}

func ShowStockLevelRecord(stock float32) {
	fmt.Printf("\n📦 Last Stock Level Record: %.2f units\n", stock)
}

func ShowOpenPortRecord(status bool) {
	fmt.Printf("\n🔌 Last Open Port Record: %s\n", map[bool]string{true: "Open", false: "Closed"}[status])
}

func GetTemperatureLimit() float64 {
	var tempLimit float64
	fmt.Print("\n🌡️  Enter Temperature Limit (°C): ")
	fmt.Scan(&tempLimit)
	return tempLimit
}

func GetOpenPortTime() float64 {
	var openPortTime float64
	fmt.Print("\n⏳ Enter Open Port Time (seconds): ")
	fmt.Scan(&openPortTime)
	return openPortTime
}

func ShowExitMessage() {
	fmt.Println("\n👋 Exiting... Have a great day! 👋")
}

func ShowInvalidChoice() {
	fmt.Println("\n❌ Invalid choice, please try again.")
}
