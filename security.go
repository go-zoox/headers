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

// ContentSecurityPolicyReportOnly (CSP) controls resources the user agent is allowed to load for a given page.
const ContentSecurityPolicyReportOnly = "Content-Security-Policy-Report-Only"

// CrossOriginResourcePolicy (CORP) prevents other domains from reading the response of the resources to which this header is applied.
const CrossOriginResourcePolicy = "Cross-Origin-Resource-Policy"

// UpgradeInsecureRequests sends a signal to the server expressing the client's preference for an encrypted and authenticated response, and that it can successfully handle the upgrade-insecure-requests directive.
const UpgradeInsecureRequests = "Upgrade-Insecure-Requests"

// XContentTypeOptions stops a browser from trying to MIME-sniff the content type and forces it to stick with the declared content-type.
const XContentTypeOptions = "X-Content-Type-Options"

// ReferrerPolicy controls how much referrer information (sent via the Referer header) should be included with requests.
const ReferrerPolicy = "Referrer-Policy"

// PermissionsPolicy allows a site to control which features and APIs can be used in the browser.
const PermissionsPolicy = "Permissions-Policy"

// OriginAgentCluster (OAC) allows a site to control which features and APIs can be used in the browser.
const OriginAgentCluster = "Origin-Agent-Cluster"

// XDNSPrefetchControl controls DNS prefetching.
const XDNSPrefetchControl = "X-DNS-Prefetch-Control"

// XDownloadOptions specifies whether the browser should be allowed to open a download dialog.
const XDownloadOptions = "X-Download-Options"

// XPermittedCrossDomainPolicies defines a policy for fetching content from a cross-domain policy file (crossdomain.xml).
const XPermittedCrossDomainPolicies = "X-Permitted-Cross-Domain-Policies"
