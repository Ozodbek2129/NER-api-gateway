package client

import (
	"gateway/config"
	"gateway/genproto/contract"
	"gateway/genproto/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceManager interface {
	UserService() user.UserServiceClient
	Productionservice() contract.ContractServiceClient
}

type serviceManagerImpl struct {
	userClient user.UserServiceClient
	productionClient contract.ContractServiceClient
}

func (s *serviceManagerImpl) UserService() user.UserServiceClient {
	return s.userClient
}

func (s *serviceManagerImpl) Productionservice() contract.ContractServiceClient {
	return s.productionClient
}

func NewServiceManager() (ServiceManager, error) {
	connUser, err := grpc.Dial(
		config.Load().USER_SERVICE,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	connDocs, err := grpc.Dial(
		config.Load().PRODUCTION_SERVICE,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return &serviceManagerImpl{
		userClient: user.NewUserServiceClient(connUser),
		productionClient: contract.NewContractServiceClient(connDocs),
	}, nil
}
