package withdrawmoney

import "transactions/model"

type Env interface {
	GetUserByID(id int) (*model.User, error)
}

func Resolve(e *Env, input model.WithdrawMoneyInput) {

}
