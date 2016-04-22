package mail

type Contacts []*Contact

func NewContacts() *Contacts {
	contacts := Contacts(make([]*Contact, 0))
	return &contacts
}

// Formatted to RFC 2822 specs
func (c *Contacts) String() string {
	ret := ""
	for i, contact := range *c {
		ret += contact.String()
		if (i < len(*c)-2) {
			ret += ","
		}
	}
	return ret
}

func (c *Contacts) Add(contact *Contact) {
	*c = append(*c, contact)
}

func (c *Contacts) Len() int {
	return len(*c)
}

type Contact struct {
	Name   string
	Email  string
}

// Formatted to RFC 2822 specs
func (c Contact) String() string {
	if c.Name != "" {
		return c.Name + " <" + c.Email + ">"
	} else {
		return c.Email
	}
}

