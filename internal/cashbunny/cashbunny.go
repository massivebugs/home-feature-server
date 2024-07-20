package cashbunny

import "context"

type Cashbunny struct{}

func NewCashbunny() *Cashbunny {
	return &Cashbunny{}
}

func (s *Cashbunny) ListAccounts(ctx context.Context, userID uint32) error {
	return nil
}
