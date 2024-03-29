package hotspot

type trackedFile struct {
	path  string
	err   error
	dates []string
}

type dataSlice []*trackedFile

// Len is part of sort.Interface.
func (d dataSlice) Len() int {
	return len(d)
}

// Swap is part of sort.Interface.
func (d dataSlice) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

// Less is part of sort.Interface. We use count as the value to sort by
func (d dataSlice) Less(i, j int) bool {
	return len(d[i].dates) > len(d[j].dates)
}

type GitCommands interface {
	Config() ([]byte, error)
	Files(string) ([]byte, error)
}
