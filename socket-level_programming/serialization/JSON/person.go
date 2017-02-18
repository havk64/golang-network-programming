package main

// Person struct
type Person struct {
	Name  Name    `json:"name"`
	Email []Email `json:"emails"`
}

// Name struct
type Name struct {
	Family   string `json:"family"`
	Personal string `json:"personal"`
}

// Email struct
type Email struct {
	Kind    string `json:"kind"`
	Address string `json:"address"`
}
