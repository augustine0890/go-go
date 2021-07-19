# Networking Programming

## Architecture
- The communication between layers is defined by protocols.
- An internet is a connection of two or more distinct networks, typically LANs or WANs.
- The communication between layers in either the TCP/IP stacks is done by sending packets of data from one layer to the next, and then eventually across the network.
- In a distributed system there will be many components running that have to communicate with each other. There are two primary models for this, message passing and remote procedure calls.
- __Message Passing__
  - Message passing is a primitive mechanism for distributed systems. Set up a connection and pump some data down it. At the other end, figure out what the message was and respond to it, possibly sending messages back.
- __Remote Procedure Call__
  - Information placed on a call stack and then control flow is transferred to another part of the program.

### Distributed Computing Models
- The most common occurrence is an asymmetric: a client sends requests to a server, and the server responds --> client-server system. If both able to initiate and to respond to messages, the we have a peer-to-peer system.

### Communication Flows
- Synchronous communication: one party will send a message and block, waiting for a reply.
- Asynchronous communication: on party sends a message and instead of waiting for a reply carries on with other work. When a reply eventually comes, it is handled.
- Streaming communication: on party sends a continuous stream messages. The streaming may need to be handled in real time, may or may not tolerate losses, and can be one-way or allow reverse communication as in control messages.
- Publish/Subscribe: parties subscribe to topics and others post to them. This can be on a small or massive scale.

### Component Distribution
- Presentation component: is responsible for interactions with the user, both displaying data and gathering input.
- Application logic is responsible for interpreting the user's response, for applying business rules, for preparing queries, and for managing responses from the third component.
- Data access is responsible for storing and retrieving data.

### Points of Failure
- Client-side errors:
  - The client side of the application could crash, hardware problems, network card could fail.
- Network errors:
  - Network contention could cause timeouts
  - There may be network address conflicts.
  - Network elements such as routers could fail
  - Transmission errors may lose messages.
- Server errors:
  - Network card could fail, hardware problems, software may crash, database may become corrupted.

  ## Socket-Level Programming
  ### The TCP/IP Stack
  - TCP/IP is the principal UNIX networking protocol.
  - It's standard for Transmission Control Protocol/Internet Protocol
  ```
  Application (OSI 5-7) <-> TCP/UDP (OSI 4) <-> IP (OSI 3) <-> h/w inteface (OSI 1-2)
  ```
  ### Services
  - Services run on host machines. They are typically long lived and are designed to wait for requests and respond to them.
  - There are many standard ports.
    - Telnet typically uses port 23 with the TCP protocol.
    - HTTP usually uses port 80, but it often ports 8000, 8080, and 8088, all with TCP.

## HTTP
- The Web is built on top of the HTTP (Hypertext Transport Protocol), which is layered on top of TCP.
- All the earier verions of HTTP are text-based. The most significant departure for HTTP/2 is that it is a binary format.