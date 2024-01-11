package rbac

import (
	"net/http"
	"testing"

	"github.com/casbin/casbin/v2"
)

func mustNewEnforcer() *casbin.Enforcer {
	e, err := casbin.NewEnforcer("./auth_model.conf", "./policy.csv")
	if err != nil {
		panic(err)
	}

	return e
}

func TestCasbin(t *testing.T) {
	e := mustNewEnforcer()

	type args struct {
		role   string
		path   string
		method string
	}

	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "admin can post to /create",
			args: args{
				role:   "admin",
				path:   "/create",
				method: http.MethodPost,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "member cannot post to /create",
			args: args{
				role:   "member",
				path:   "/create",
				method: http.MethodPost,
			},
			want:    false,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := e.Enforce(tt.args.role, tt.args.path, tt.args.method)

			if (err != nil) != tt.wantErr {
				t.Errorf("Enforcer.Enforce() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("Enforcer.Enforce() = %v, want %v", got, tt.want)
			}
		})
	}
}
