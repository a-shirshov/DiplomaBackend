#!/bin/bash
mockgen -source=internal/microservices/auth/repository.go \
  -destination=internal/microservices/auth/mock/repository_mock.go \
  -package=mock

mockgen -source=pkg/password_hasher.go \
  -destination=internal/microservices/auth/mock/password_hasher_mock.go \
  -package=mock

mockgen -source=pkg/token_manager.go \
  -destination=internal/microservices/auth/mock/token_manager_mock.go \
  -package=mock

mockgen -source=internal/microservices/auth/usecase.go \
  -destination=internal/microservices/auth/mock/usecase_mock.go \
  -package=mock