package crud

import (
	"net/http"
)

// type Handler interface {
//   ServeHTTP(ResponseWriter, *Request)
// }
//
// handler func(ResponseWriter, *Request)

type Crud struct {
	// Requests a representation of the specified resource. Requests using GET should only retrieve data and should have no other effect. (This is also true of some other HTTP methods.)[1] The W3C has published guidance principles on this distinction, saying, "Web application design should be informed by the above principles, but also by the relevant limitations."[12] See safe methods below.
	Get func(http.ResponseWriter, *http.Request)

	// Asks for the response identical to the one that would correspond to a GET request, but without the response body. This is useful for retrieving meta-information written in response headers, without having to transport the entire content.
	Head func(http.ResponseWriter, *http.Request)

	// Requests that the server accept the entity enclosed in the request as a new subordinate of the web resource identified by the URI. The data POSTed might be, as examples, an annotation for existing resources; a message for a bulletin board, newsgroup, mailing list, or comment thread; a block of data that is the result of submitting a web form to a data-handling process; or an item to add to a database.[13]
	Post func(http.ResponseWriter, *http.Request)

	// Requests that the enclosed entity be stored under the supplied URI. If the URI refers to an already existing resource, it is modified; if the URI does not point to an existing resource, then the server can create the resource with that URI.[14]
	Put func(http.ResponseWriter, *http.Request)

	// Deletes the specified resource.
	Delete func(http.ResponseWriter, *http.Request)

	// Echoes back the received request so that a client can see what (if any) changes or additions have been made by intermediate servers.
	Trace func(http.ResponseWriter, *http.Request)

	// Returns the HTTP methods that the server supports for the specified URL. This can be used to check the functionality of a web server by requesting '*' instead of a specific resource.
	Options func(http.ResponseWriter, *http.Request)

	// Converts the request connection to a transparent TCP/IP tunnel, usually to facilitate SSL-encrypted communication (HTTPS) through an unencrypted HTTP proxy.[15][16] See HTTP CONNECT Tunneling.
	Connect func(http.ResponseWriter, *http.Request)

	// Is used to apply partial modifications to a resource.[17]
	Patch func(http.ResponseWriter, *http.Request)
}

func (c *Crud) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		if c.Get != nil {
			c.Get(rw, req)
		}
	case "HEAD":
		if c.Get != nil {
			c.Head(rw, req)
		}
	case "POST":
		if c.Get != nil {
			c.Post(rw, req)
		}
	case "PUT":
		if c.Get != nil {
			c.Put(rw, req)
		}
	case "DELETE":
		if c.Get != nil {
			c.Delete(rw, req)
		}
	case "TRACE":
		if c.Get != nil {
			c.Trace(rw, req)
		}
	case "OPTIONS":
		if c.Get != nil {
			c.Options(rw, req)
		}
	case "CONNECT":
		if c.Get != nil {
			c.Connect(rw, req)
		}
	case "PATCH":
		if c.Get != nil {
			c.Patch(rw, req)
		}
	default:
		break
	}

}
