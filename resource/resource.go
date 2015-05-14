package resource

import "github.com/usabilla/gobilla/auth"

// Resource provides auth for every API resource
type Resource struct {
	Auth auth.Auth
}
