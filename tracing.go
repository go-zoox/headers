package headers

// references:
//   - https://uptrace.dev/opentelemetry/opentelemetry-traceparent.html
//   - https://github.com/opentracing/specification/blob/master/rfc/trace_identifiers.md

// XRequestID ...
const XRequestID = "X-Request-ID"

// Traceparent contains information about the incoming request in a tracing system.
// Example:
//
// # {version}-{trace_id}-{span_id}-{trace_flags}
// traceparent: 00-80e1afed08e019fc1110464cfa66635c-7a085853722dc6d2-01
const Traceparent = "Traceparent"
