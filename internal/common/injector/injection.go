package injector

import (
	"github.com/adi-kmt/brew-scylla/internal/authflow"
	"github.com/adi-kmt/brew-scylla/internal/db"
	"github.com/adi-kmt/brew-scylla/internal/user"
	"github.com/scylladb/gocqlx/v3"
)

type InjectedVars struct {
	Session gocqlx.Session
	Product *user.ProductService
	Order   *user.OrderService
	User    *authflow.AuthService
}

func Inject() (InjectedVars, error) {
	db_session, err := db.Init()
	if err != nil {
		return InjectedVars{}, err
	}
	productRepository := user.NewProductRepository(db_session)
	orderRepository := user.NewOrderRepository(db_session)
	userRepository := authflow.NewUserRepository(db_session)

	userService := authflow.NewAuthService(userRepository)

	productService := user.NewProductService(productRepository)
	orderService := user.NewOrderService(orderRepository)

	return InjectedVars{db_session, productService, orderService, userService}, nil
}
