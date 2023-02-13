#!/bin/bash
mockgen -source=internal/microservices/auth/delivery.go \
  -destination=internal/microservices/auth/mock/delivery_mock.go \
  -package=mock

mockgen -source=internal/microservices/auth/repository.go \
  -destination=internal/microservices/auth/mock/repository_mock.go \
  -package=mock

mockgen -source=internal/microservices/auth/usecase.go \
  -destination=internal/microservices/auth/mock/usecase_mock.go \
  -package=mock

######
mockgen -source=internal/microservices/event/usecase.go \
  -destination=internal/microservices/event/mock/usecase_mock.go \
  -package=mock

mockgen -source=internal/microservices/event/repository.go \
  -destination=internal/microservices/event/mock/repository_mock.go \
  -package=mock

#####
mockgen -source=internal/microservices/user/delivery.go \
  -destination=internal/microservices/user/mock/delivery_mock.go \
  -package=mock

mockgen -source=internal/microservices/user/usecase.go \
  -destination=internal/microservices/user/mock/usecase_mock.go \
  -package=mock

mockgen -source=internal/microservices/user/repository.go \
  -destination=internal/microservices/user/mock/repository_mock.go \
  -package=mock

#####
mockgen -source=pkg/password_hasher.go \
  -destination=internal/microservices/auth/mock/password_hasher_mock.go \
  -package=mock

mockgen -source=pkg/token_manager.go \
  -destination=internal/microservices/auth/mock/token_manager_mock.go \
  -package=mock

#####
mockgen -source=internal/middleware/middleware.go \
  -destination=internal/middleware/mock/middleware_mock.go \
  -package=mock
