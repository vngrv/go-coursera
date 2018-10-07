package person

var (
	// экспортируемая переменная должна начинаться с большой буквы
	Public = 1
	// неэкспортируемая с маленькой
	private = 1
)

type Person struct {
	ID     int
	Name   string
	secret string
}

func (p Person) UpdateSecret(secret string) {
	p.secret = secret
}
