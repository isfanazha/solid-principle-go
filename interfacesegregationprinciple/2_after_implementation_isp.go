package interfacesegregationprinciple

import "github.com/isfanazha/solid-principle-go/entity"

// Solution
// With this breakdown, every component or module in the system can depend solely on the aspects it truly requires,
// which makes the application cleaner, more maintainable, and potentially less error-prone.

type UserRegistrar interface {
	Register(username, password string) error
}

type UserAuthenticator interface {
	Login(username, password string) (*entity.User, error)
}

type UserProfileManager interface {
	UpdateProfile(userID int, profile entity.Profile) error
}

type UserAccountManager interface {
	DeleteAccount(userID int) error
}

type UserOrderViewer interface {
	GetUserOrders(userID int) ([]entity.Order, error)
}
