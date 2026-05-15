package client

import (
	"gateway/config"
	"gateway/genproto/contract"
	"gateway/genproto/user"
	"gateway/genproto/services"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceManager interface {
	UserService() user.UserServiceClient
	Productionservice() contract.ContractServiceClient
	Employeeservice() services.ServicesServiceClient
}

type serviceManagerImpl struct {
	userClient user.UserServiceClient
	productionClient contract.ContractServiceClient
	employeeClient services.ServicesServiceClient
}

func (s *serviceManagerImpl) UserService() user.UserServiceClient {
	return s.userClient
}

func (s *serviceManagerImpl) Productionservice() contract.ContractServiceClient {
	return s.productionClient
}

func (s *serviceManagerImpl) Employeeservice() services.ServicesServiceClient {
	return s.employeeClient
}

func NewServiceManager() (ServiceManager, error) {
	connUser, err := grpc.Dial(
		config.Load().USER_SERVICE,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	connproduction, err := grpc.Dial(
		config.Load().PRODUCTION_SERVICE,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return &serviceManagerImpl{
		userClient: user.NewUserServiceClient(connUser),
		productionClient: contract.NewContractServiceClient(connproduction),
		employeeClient: services.NewServicesServiceClient(connproduction),
	}, nil
}
