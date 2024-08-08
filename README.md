<h1 align="center" id="title">Chat Room Server</h1>

<p align="center"><img src="https://socialify.git.ci/rohankarn35/Chat-Room-Server/image?font=Bitter&amp;issues=1&amp;language=1&amp;name=1&amp;owner=1&amp;pattern=Solid&amp;pulls=1&amp;stargazers=1&amp;theme=Light" alt="project-image"></p>

A Go-based WebSocket server that facilitates real-time communication in chat rooms. This server leverages the `Gin` web framework and `Gorilla WebSocket` to manage WebSocket connections, enabling users to create and join rooms, and exchange messages in real-time.

---

## **Table of Contents**
- [Features](#features)
- [Project Structure](#project-structure)
- [Installation](#installation)
- [Usage](#usage)
- [Makefile Targets](#makefile-targets)
- [Endpoints](#endpoints)
- [Contributing](#contributing)
- [License](#license)

---

## **Features**
- **Room Creation**: Easily create rooms with unique IDs and names.
- **Room Joining**: Join existing rooms by specifying the room ID.
- **Real-Time Messaging**: Supports broadcasting messages in real-time to all room participants.
- **WebSocket Management**: Handles WebSocket connections, message broadcasting, and room management efficiently.

---

## **Project Structure**

```plaintext
golang-websocket/
├── cmd/
│   └── main.go             # Main entry point for the server
├── internals/
│   ├── handlers/           # HTTP routes and WebSocket connections
│   ├── hub/                # Chat room and client management
│   └── models/             # Data structures for messages and rooms
├── Makefile                # Build automation file
├── go.mod                  # Go module dependencies
├── go.sum                  # Go module checksums
└── README.md               # Project documentation
```


<h2>Installation Steps:</h2>

<p>1. Clone the repository</p>

```bash
git clone https://github.com/rohankarn35/Chat-Room-Server
```

<p>2. Install Dependency</p>

```
go mod tidy
```

<p>3. Run the Server</p>

```
go run  cmd/main.go
```

 ## Endpoints

- **`/createRoom/:roomId/:roomName`**: Create a new chat room with a unique ID and name.
- **`/join`**: Join an existing chat room by specifying the room ID.
- **`/ws/:roomId`**: Establish a WebSocket connection to the specified room.

---

## Usage

After running the server, you can interact with it using HTTP clients or WebSocket clients.


#### Create a Room

```bash
GET /createRoom/{roomId}/{roomName}
```
#### Join a Room
```
GET /join?roomId={roomId}
```
#### Connect via WebSocket
```
WS /ws/{roomId}
```
  
<h2>Built with</h2>

Technologies used in the project:

*   Golang
*   GIn
*   Gorilla Websocket

<h2>Contributing</h2>
Contributions are welcome! Please fork the repository and submit a pull request.



