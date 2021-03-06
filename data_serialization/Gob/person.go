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

// Implement fmt.Stringer interface
func (p Person) String() string {
	s := p.Name.Personal + " " + p.Name.Family
	for _, v := range p.Email {
		s += "\n" + v.Kind + ": " + v.Address
	}
	return s
}

var person = Person{
	Name: Name{
		Family:   "de Oliveira",
		Personal: "Alexandro"},
	Email: []Email{
		Email{
			Kind:    "home",
			Address: "alexandro.deoliveira@icloud.com"},
		Email{
			Kind:    "school",
			Address: "alexandro.oliveira@holbertonschool.com"}}}
