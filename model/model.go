package model

type Order struct {
	ID             int
	VIP            bool
	ProcessingTime int
}

type Bot struct {
	ID   int
	Idle chan bool
	Stop chan bool
}
