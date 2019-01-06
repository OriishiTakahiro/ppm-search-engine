package store

type HitItem struct {
	Name     string
	Distance int
}

func AddRow(row interface{}) error {
	return nil
}

func SearchTop(row interface{}, limit uint) ([]HitItem, error) {
	return make([]HitItem, 10), nil
}
