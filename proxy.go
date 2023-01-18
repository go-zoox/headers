package headers

// XForwardedFor identifieds the originating IP addresses of a client
// connecting to a web server through an HTTP proxy or a load balancer.
const XForwardedFor = "X-Forwarded-For"

// XForwardedHost identifies the original host requested that a client used to connect to your proxy or load balancer.
const XForwardedHost = "X-Forwarded-Host"

// XForwardedProto identifies the original protocol (HTTP or HTTPS) requested that a client used to connect to your proxy or load balancer.
const XForwardedProto = "X-Forwarded-Proto"

// Via added by a proxied, both forward and reverse proxies, and can appear in the request headers and the response headers.
const Via = "Via"
