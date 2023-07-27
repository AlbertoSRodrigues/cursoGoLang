package campaign

type Repository interface {
	Save(Campaign *Campaign) error
}
