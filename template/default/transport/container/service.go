package container

// ServiceContainer handle all service used in project
type ServiceContainer struct {
}

// CreateServiceContainer construct all services available in the app
func CreateServiceContainer(repositories *RepositoryContainer, clients *ClientContainer) *ServiceContainer {
	return &ServiceContainer{}
}
