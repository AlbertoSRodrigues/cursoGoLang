package campaign

type Repository interface {
	Save(Campaign *Campaign) error
	Get() ([]Campaign, error)
	GetBy(id string) (*Campaign, error)
}
