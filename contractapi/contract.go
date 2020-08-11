package contractapi

type ContractInterface interface {
	GetTransactionContextHandler() SettableTransactionContextInterface
}

type Contract struct {
	TransactionContextHandler SettableTransactionContextInterface
}

// GetTransactionContextHandler returns the current transaction context set for
// the contract. If none has been set then TransactionContext will be returned
func (c *Contract) GetTransactionContextHandler() SettableTransactionContextInterface {
	if c.TransactionContextHandler == nil {
		return new(TransactionContext)
	}
	return c.TransactionContextHandler
}
