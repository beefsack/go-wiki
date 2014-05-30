package wiki

type Persister interface {
	CreatePage(p Page) (id string, err error)
	UpdatePage(id string, p Page) error
	DeletePage(id string) error
	ReadPage(id string) (Page, error)
}
