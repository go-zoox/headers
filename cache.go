package headers

// CacheControl directives for caching mechanisms in both requests and responses
const CacheControl = "Cache-Control"

// Expires is the date/time after which the response is considered stale.
const Expires = "Expires"

// Pragma implementation-specific header that may have various effects anywhere along the request-response chain.
// Used for backwards compatibility with HTTP/1.0 caches where the Cache-Control header is not yet present.
const Pragma = "Pragma"

// Age is the time, in seconds, that the object has been in a proxy cache.
const Age = "Age"

// conditional cache

// LastModified is the last modification date of the resource, used to compare several versions of the same resource.
// It is less accurate than ETag, but easier to calculate in some environments.
// Conditional requests using If-Modified-Since and If-Unmodified-Since use this value to change the behavior of the request.
const LastModified = "Last-Modified"

// ETag is a unique string identifying the version of the resource.
// Conditional requests using If-Match and If-None-Match use this value to change the behavior of the request.
const ETag = "ETag"

// IfMatch makes the request conditional, and applies the method only if the stored resource matches one of the given ETags.
const IfMatch = "If-Match"

// IfNoneMatch makes the request conditional, and applies the method only if the stored resource doesn't match any of the given ETags.
// This is used to update caches (for safe requests), or to prevent uploading a new resource when one already exists.
const IfNoneMatch = "If-None-Match"

// IfModifiedSince makes the request conditional,
// and expects the resource to be transmitted only if it has been modified after the given date.
// This is used to transmit data only when the cache is out of date.
const IfModifiedSince = "If-Modified-Since"

// Vary determines how to match request headers to decide whether a cached response can be used rather than requesting a fresh one from the origin server.
// It is used by CDN Provider.
const Vary = "Vary"
