# Notify Key Value Storage (nkv) client for golang

### What is it for? 
nkv lets you share state between your services using client server topology. 
it provides you with following API:

- get the state for the given key
- put the state for the given key
- delete the state for the given key
- subscribe to the updates for the given key
- unsubscribe from the updated for the given key

nkv is a client/server solution. This repository is client for golang. For other clients, refer to [this](https://github.com/uncleDecart/nkv) repository.
Also note that API is intented to be straightforward, so it shouldn't take you time to implement client using any programming language you like.

### What protocol does it use to communicate with server?

Check out [this](https://github.com/uncleDecart/nkv/blob/main/docs/CLIENT_SERVER_PROTOCOL.md) document

### When should I use it?
When you have some shared state between services/processes and you also want to be notified when the value is changed

### How do I use it?

You can run it as a separate binary (checkout cmd folder) or you can use it in your golang service like this:


```golang

import (
	"github.com/uncleDecart/go-nkv/pkg/client"
)

// ...

client := client.NewClient("127.0.0.1:4222")
key := "key1"
value := "bazinga"

printUpdate := func(msg protocol.Notification) {
  fmt.Printf("Received update:\n%s\n", protocol.MarshalNotification(&msg))
}

resp, err := client.Subscribe(key, printUpdate)
// print response, handle error
resp, err := client.Put(key, []byte(value))
// print response, handle error
resp, err := client.Get(key)
resp, err := client.Unsubscribe(key)
// print response, handle error
resp, err := client.Delete(key)
// print response, handle error
```

