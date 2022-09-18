package service

import "github.com/AliasYermukanov/proxy-server/src/util"

func (s *service) checkMethod(method string) error {
	switch method {
	case "GET":
	case "POST":
	case "PUT":
	case "DELETE":
	case "OPTIONS":
	default:
		return util.CommonError.SetDevMessage("unacceptable request method")
	}
	return nil
}
