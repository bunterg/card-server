# Fruno backend

### Hub 

The code for the `Hub` type is in
[hub.go](https://github.com/gorilla/websocket/blob/master/examples/chat/hub.go). 
The application's `main` function starts the hub's `run` method as a goroutine.
Clients send requests to the hub using the `register`, `unregister` and
`broadcast` channels.

The hub registers clients by adding the client pointer as a key in the
`clients` map. The map value is always true.

The unregister code is a little more complicated. In addition to deleting the
client pointer from the `clients` map, the hub closes the clients's `send`
channel to signal the client that no more messages will be sent to the client.

The hub handles messages by looping over the registered clients and sending the
message to the client's `send` channel. If the client's `send` buffer is full,
then the hub assumes that the client is dead or stuck. In this case, the hub
unregisters the client and closes the websocket.

### Client

The code for the `Client` type is in [client.go](https://github.com/gorilla/websocket/blob/master/examples/chat/client.go).

The `serveWs` function is registered by the application's `main` function as
an HTTP handler. The handler upgrades the HTTP connection to the WebSocket
protocol, creates a client, registers the client with the hub and schedules the
client to be unregistered using a defer statement.

Next, the HTTP handler starts the client's `writePump` method as a goroutine.
This method transfers messages from the client's send channel to the websocket
connection. The writer method exits when the channel is closed by the hub or
there's an error writing to the websocket connection.

Finally, the HTTP handler calls the client's `readPump` method. This method
transfers inbound messages from the websocket to the hub.

WebSocket connections [support one concurrent reader and one concurrent
writer](https://godoc.org/github.com/gorilla/websocket#hdr-Concurrency). The
application ensures that these concurrency requirements are met by executing
all reads from the `readPump` goroutine and all writes from the `writePump`
goroutine.

To improve efficiency under high load, the `writePump` function coalesces
pending chat messages in the `send` channel to a single WebSocket message. This
reduces the number of system calls and the amount of data sent over the
network.