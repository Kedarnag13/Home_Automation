package models

type Light struct {
	Pin_number int
	Status bool
}

type LightMessage struct {
	Success string
	Message string
}
