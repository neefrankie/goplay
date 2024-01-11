package rbac

import (
	"github.com/casbin/casbin/v2"
)

type Role string

const (
	Admin     Role = "admin"
	Staff     Role = "staff"
	Anonymous Role = "anonymous"
)

type Accessor struct {
	Role   string // sub
	Path   string // obj
	Method string // act
}

type RBAC struct {
	enforcer *casbin.Enforcer
}

func New() (RBAC, error) {
	e, err := casbin.NewEnforcer("./auth_model.conf", "./policy.csv")
	if err != nil {
		return RBAC{}, err
	}

	return RBAC{
		enforcer: e,
	}, nil
}

func MustNew() RBAC {
	ac, err := New()
	if err != nil {
		panic(err)
	}

	return ac
}

func (ac RBAC) Grant(a Accessor) (bool, error) {
	return ac.enforcer.Enforce(a.Role, a.Path, a.Method)
}
