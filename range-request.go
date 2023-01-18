package headers

// AcceptRanges indicates if the server supports range requests, and if so in which unit the range can be expressed.
const AcceptRanges = "Accept-Ranges"

// Range indicates the part of a document that the server should return.
const Range = "Range"

// IfRange creates a conditional range request that is only fulfilled if the given etag or date matches the remote resource.
// Used to prevent downloading two ranges from incompatible version of the resource.
const IfRange = "If-Range"

// ContentRange indicates where in a full body message a partial message belongs.
const ContentRange = "Content-Range"
