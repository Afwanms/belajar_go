package identities

type Identity struct {
	Name    string
	Address string
	Age     int
}

func (params Identity) Identities(name, address string, Age int) {
	params.Name = name
	params.Address = address
	params.Age = Age
}
