## Characteristics of a Distributed System
- Service discovery
- Load balancing
- Distributed tracing and logging
- Service monitoring

## Types of Distributed Systems
- Hub and Spoke
- Peer to Peer
- Message Queues
- Hybrid

## Architectural Elements
- Languages: single, multiple, JS
- Frameworks
  - Are architectures compatible?
  - Do they support required transports and protocols?
  - How stable are they? Are they easy to upgrade?
- Transports
  - HTTP, gRPC, other RPC, Mixed
- Protocol
  - Language specific: encoding/gob
  - JSON
  - Protocol Buffers
  - XML, SOAP, ...

## Project
- Type: hybrid - central registry and monitoring
- Language: Golang
- Frameworks: Go standard library, Fiber (Go-Kit, Go-Micro)
- Transports: HTTP
- Protocol: JSON

### Components
- Service Registry:
  - Service registration
  - Health monitoring
- Teacher Portal
  - Web app
  - API Gateway
- Log Service
  - Centralized logging
- Grading Service
  - Business logic
  - Data persistence