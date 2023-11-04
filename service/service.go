package service

type (
	ServiceDependencies struct {
	}

	templateService struct {
	}
)

func NewService() Service {
	return &templateService{}
}
