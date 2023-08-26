package interfacesegregationprinciple

import "github.com/isfanazha/solid-principle-go/entity"

// PROBLEM
// This single interface handles everything related to a user - registration, login, profile updating, account deletion, and fetching orders.
// If a component is only interested in handling the registration, it still has to know about other methods, which is unnecessary.

type UserManagerNotISP interface {
	Register(username, password string) error
	Login(username, password string) (*entity.User, error)
	UpdateProfile(userID int, profile entity.Profile) error
	DeleteAccount(userID int) error
	GetUserOrders(userID int) ([]entity.Order, error)
}
