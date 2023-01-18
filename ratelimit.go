package headers

// references: https://medium.com/@guillaume.viguierjust/rate-limiting-your-restful-api-3148f8e77248

// RetryAfter indicates how long the user agent should wait before making a follow-up request.
const RetryAfter = "Retry-After"

// XRateLimitLimit indicates the maximum number of requests you're permitted to make per hour.
// Example:
//
//	GitHub: X-RateLimit-Limit
//	Twitter: X-Rate-Limit-Limit
const XRateLimitLimit = "X-RateLimit-Limit"

// XRateLimitRemaining indicates the number of requests remaining in the current rate limit window.
// Example:
//
//	GitHub: X-RateLimit-Remaining
//	Twitter: X-Rate-Limit-Remaining
const XRateLimitRemaining = "X-RateLimit-Remaining"

// XRateLimitReset indicates the time at which the current rate limit window resets in UTC epoch seconds.
// Example:
//
//	GitHub: X-RateLimit-Reset
//	Twitter: X-Rate-Limit-Reset
const XRateLimitReset = "X-RateLimit-Reset"
