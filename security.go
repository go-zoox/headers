package headers

// ContentSecurityPolicy (CSP) controls resources the user agent is allowed to load for a given page.
const ContentSecurityPolicy = "Content-Security-Policy"

// StrictTransportSecurity (HSTS) forces communication using HTTPS insted of HTTP.
const StrictTransportSecurity = "Strict-Transport-Security"

// XFrameOptions (XFO) indicates whether a browser should be allowed to render a page
// in a <frame>, <iframe>, <embed>, or <object>.
// Options: DENY, SAMEORIGIN, ALLOW-FROM=URL
// See more: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Frame-Options
const XFrameOptions = "X-Frame-Options"

// XXSSProtection enables cross-site scripting filtering.
const XXSSProtection = "X-XSS-Protection"

// CrossOriginEmbederPolicy (COEP) allows a server to declare an embedder policy for a given document.
const CrossOriginEmbederPolicy = "Cross-Origin-Embedder-Policy"

// CrossOriginOpenerPolicy (COOP) prevents other domain from opening/controlling a window.
const CrossOriginOpenerPolicy = "Cross-Origin-Opener-Policy"

// CrossOriginResourcePolicy (CORP) prevents other domains from reading the response of the resources to which this header is applied.
const CrossOriginResourcePolicy = "Cross-Origin-Resource-Policy"

// UpgradeInsecureRequests sends a signal to the server expressing the client's preference for an encrypted and authenticated response, and that it can successfully handle the upgrade-insecure-requests directive.
const UpgradeInsecureRequests = "Upgrade-Insecure-Requests"
