package interfaces

type Finder interface {
	Find(id string) (Viewer, error)
}
