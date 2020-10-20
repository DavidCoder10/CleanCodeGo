package contracts

//IService ...
type IService interface {
	ValidateUserOfLegalAge(id int) (string, error)
}
