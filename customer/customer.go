package customer

type Customer struct {
	noHp       string
	idEmployee int
	name       string
}

func (c *Customer) SetNohp(newNohp string) {
	c.noHp = newNohp
}
func (c *Customer) SetIdEmployee(newIdEmployee int) {
	c.idEmployee = newIdEmployee
}
func (c *Customer) SetName(newName string) {
	c.name = newName
}

// method get

func (c *Customer) GetNohp() string {
	return c.noHp
}

func (c *Customer) GetIdEmployee() int {
	return c.idEmployee
}

func (c *Customer) GetName() string {
	return c.name
}
