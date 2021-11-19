# Project: Banking Application
- Open new account for customer
- Making a deposit or withdraw transaction
- Role based access control (RBAC)

## Objective
- Hexagonal architecture
- Dependency inversion in Go and working with stubs
- JWT Tokens
- Banking-Auth microservice based on OAuth standard
- Unit test for routes and other components using mocks and state based tests

## Structure
### Router Basic
- Incoming Request (`/customer`) --> Request Multiplexer/Router (matches registered patterns) --> Customer Route

### Hexagonal Architecture
- [Hexagonal Architecture](https://www.qwan.eu/2020/08/20/hexagonal-architecture.html)
- User Side --> REST Handler - Test Agent --> Service <interface> <-- Business Logic --> Repository<interface> <-- Database Adapter (DB) - Mock Adapter (Mock Implementation) --> Server Side (DB, external WS, filesystem)
