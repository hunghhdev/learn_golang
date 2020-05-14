package services

var (
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	GetItem()
	SaveItem()
}

type itemsService struct {
}

func (s *itemsService) GetItem() {

}

func (s *itemsService) SaveItem() {

}
