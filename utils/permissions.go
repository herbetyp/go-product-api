package utils

import "fmt"

func CheckAdminPermissions(requestHttpMethod string, requestPath string, sub string) bool {
	denyHttpMethods := []string{"DELETE", "PATCH"}
	basePath := "/v1/admin"
	denyPaths := []string{fmt.Sprintf("%s/users/%s/status", basePath, sub), fmt.Sprintf("%s/users/%s", basePath, sub)}

	resourcePath := fmt.Sprintf("%s:%s", requestHttpMethod, requestPath)

	for _, path := range denyPaths {
		for _, httpMethod := range denyHttpMethods {
			denyResource := fmt.Sprintf("%s:%s", httpMethod, path)
			if resourcePath == denyResource {
				return false
			}
		}
	}
	return true
}
