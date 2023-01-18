package headers

// Origin indicates where a fetch originates from.
const Origin = "Origin"

// AccessControlAllowOrigin indicates whether the response can be shared
const AccessControlAllowOrigin = "Access-Control-Allow-Origin"

// AccessControlAllowCredentials indicates whether the response to the request can be exposed when the credentials flag is true.
const AccessControlAllowCredentials = "Access-Control-Allow-Credentials"

// AccessControlAllowHeaders used in response to a preflight request
// to indicate which HTTP headers can be used when making the actual request.
const AccessControlAllowHeaders = "Access-Control-Allow-Headers"

// AccessControlAllowMethods specifies the methods allowed when accessing the resource in response to a preflight request.
const AccessControlAllowMethods = "Access-Control-Allow-Methods"

// AccessControlExposeHeaders indicates which headers can be exposed as a part of the resonse by listing their names.
const AccessControlExposeHeaders = "Access-Control-Expose-Headers"

// AccessControlMaxAge indicates how long the results of a preflight request can be cached.
const AccessControlMaxAge = "Access-Control-Max-Age"

// AccessControlRequestMethod used when issuing a preflight request to let the server know
// which HTTP method will bed used when the actual request is made.
const AccessControlRequestMethod = "Access-Control-Request-Method"
