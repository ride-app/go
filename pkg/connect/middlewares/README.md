# Middlewares

`github.com/dragonfish-tech/pkg/middlewares` provides the following middlewares

## Firebase auth middleware

Authenticate requests with firebase auth jwt token.

### Exmaple:

```go
package main

import (
  "net/http"

  "github.com/dragonfish-tech/pkg/connect/middlewares"

  "connectrpc.com/authn"
  "connectrpc.com/authn/internal/gen/authn/ping/v1/pingv1connect"
)

func main() {
  mux := http.NewServeMux()
  service := &pingv1connect.UnimplementedPingServiceHandler{}
  mux.Handle(pingv1connect.NewPingServiceHandler(service))

  middleware := authn.NewMiddleware(middlewares.FirebaseAuth)
  handler := middleware.Wrap(mux)
  http.ListenAndServe("localhost:8080", handler)
}
```
