package resolver

import (
	"github.com/bishal-dd/go-server/graph/resolvers/receipt"
	"github.com/bishal-dd/go-server/graph/resolvers/user"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	*receipt.ReceiptResolver
	*user.UserResolver
}
