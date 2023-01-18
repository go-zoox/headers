package headers

// ContentType indicates the data type, both in request and response.
const ContentType = "Content-Type"

// ContentLength indicates the size of the resource, in decimal number of bytes, both in request and response.
const ContentLength = "Content-Length"

// ContentEncoding used to specifiy the compression algorithm.
const ContentEncoding = "Content-Encoding"

// ContentLanguage describes the human language(s) intended for the audience, so that it allows a user to differentiate according to the users' own preferred language.
const ContentLanguage = "Content-Language"

// ContentLocation indicates an alternate location for the returned data.
const ContentLocation = "Content-Location"

// Accept informs the server about the types of data that can be sent back.
const Accept = "Accept"

// AcceptEncoding is the encoding algorithm, usually a compression algorithm, that can be used on the resource sent back.
const AcceptEncoding = "Accept-Encoding"

// AcceptLanguage informs the server about the human language the server is expected to send back.
// This is a hint and is not necessarily under the full control of the user: the server should always pay attention not to override an explicit user choice (like selecting a language from a dropdown).
const AcceptLanguage = "Accept-Language"
