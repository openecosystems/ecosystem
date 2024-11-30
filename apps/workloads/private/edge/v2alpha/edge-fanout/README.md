# Fanout Starter Kit for Rust

[![Deploy to Fastly](https://deploy.edgecompute.app/button)](https://deploy.edgecompute.app/deploy)

Install this starter kit to use Fanout. It routes incoming requests through the Fanout GRIP proxy and on to an origin. It also provides some endpoints for testing subscriptions without an origin.

There is no need to modify this kit. It is production-ready for typical Fanout usage that coordinates with an origin. However, if you would like to implement Fanout logic at the edge, this kit is also a good starting point for that and you can modify it to fit your needs. See the test endpoints for inspiration.

**For more details about this and other starter kits for Compute, see the [Fastly Documentation Hub](https://www.fastly.com/documentation/solutions/starters/)**.

## Setup

The app expects a configured backend named "origin" where Fanout-capable requests should be forwarded to.

Additionally, for the test endpoints to work, the app expects a configured backend named "self" that points back to app itself. For example, if the service has a domain `foo.edgecompute.app`, then you'll need to create a backend on the service named "self" with the destination host set to `foo.edgecompute.app` and port 443. Also set "Override Host" to the same host value.

## Test Endpoints

For requests made to domains ending in `.edgecompute.app`, the app will handle requests to the following endpoints without forwarding to the origin:

* `/test/websocket`: bi-directional WebSocket
* `/test/stream`: HTTP streaming of `text/plain`
* `/test/sse`: SSE (streaming of `text/event-stream`)
* `/test/long-poll`: Long-polling

Connecting to any of these endpoints will subscribe the connection to channel "test". The WebSocket endpoint echos back any messages it receives from the client.

Data can be sent to the connections via the GRIP publish endpoint at `https://api.fastly.com/service/{service-id}/publish/`. For example, here's a curl command to send a WebSocket message:

```sh
curl \
  -H "Authorization: Bearer {fastly-api-token}" \
  -d '{"items":[{"channel":"test","formats":{"ws-message":{"content":"hello"}}}]}' \
  https://api.fastly.com/service/{service-id}/publish/
```

## How it works

Non-test requests are simply forwarded through the Fanout proxy and on to the origin.

For test requests, the app is actually invoked twice.

1. Initially, a client request arrives at the app without having been routed through the Fanout proxy yet. The app checks for this via the presence of a `Grip-Sig` header. If that header is not present, the app calls `req.handoff_fanout("self")` and exits. This tells the subsystem that the connection should be routed through Fanout, and is used for HTTP requests controlled by GRIP.

2. Since `self` refers to the same app, a second request is made to the same app, this time coming through Fanout. The app checks for this, and then handles the request accordingly (in `handle_test()`).

## Note

This app is not currently supported in Fastly's [local development server](https://www.fastly.com/documentation/guides/compute/testing/#running-a-local-testing-server), as the development server does not support Fanout features. To experiment with Fanout, you will need to publish this project to your Fastly Compute service. using the `fastly compute publish` command.

## Security issues

Please see [SECURITY.md](SECURITY.md) for guidance on reporting security-related issues.
