package endpoints

func AssignRoutesTo(endpoint *Endpoint) {
	endpoint.Group.Get("/", endpoint.Index)
	endpoint.Group.Post("/auth/register", endpoint.AuthRegister)
	endpoint.Group.Post("/auth/login", endpoint.AuthLogin)
}
