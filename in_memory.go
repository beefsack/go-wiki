package wiki

import "errors"

type InMemory struct {
	Pages map[string]Page
}

func NewInMemory() InMemory {
	return InMemory{
		Pages: map[string]Page{},
	}
}

func (i InMemory) CreatePage(p Page) (id string, err error) {
	i.Pages[p.Title] = p
	return p.Title, nil
}

func (i InMemory) UpdatePage(id string, p Page) error {
	i.Pages[id] = p
	return nil
}

func (i InMemory) DeletePage(id string) error {
	delete(i.Pages, id)
	return nil
}

func (i InMemory) ReadPage(id string) (Page, error) {
	p, ok := i.Pages[id]
	if !ok {
		return p, errors.New("could not find page")
	}
	return p, nil
}
