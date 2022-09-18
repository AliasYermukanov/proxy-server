package middleware

import "github.com/AliasYermukanov/proxy-server/src/api/service"

// Middleware describes a service middleware
type Middleware func(service service.ProxyService) service.ProxyService
