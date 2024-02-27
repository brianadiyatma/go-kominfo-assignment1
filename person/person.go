package person

type Person struct {
	Name      string
	Alamat    string
	Pekerjaan string
	Reason    string
}

func NewPerson(name string, alamat string, pekerjaan string, reason string) *Person {
	return &Person{
		Name:      name,
		Alamat:    alamat,
		Pekerjaan: pekerjaan,
		Reason:    reason,
	}
}
