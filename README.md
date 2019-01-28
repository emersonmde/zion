## Synopsis

A Go API for Matrix Chat Client-Server API.

## Details

There are a few really good full featured Go Matrix API's (like those found in GO-NEB). However, I
wanted to create a very slim and simple API implementation mostly for the creation of simple chat bots. So 
far it's not much, but I hope to add features for connecting/disconnecting, joining rooms, receiving messages,
and sending messages.

Currently each API call will be using an ACCESS_TOKEN (which you can easily get from Riot). Perhaps
in the future I will be able to add a feature for logging into a homeserver.

## Useful Information

- [Matrix.org's How to use the Client-Server API](https://matrix.org/docs/guides/client-server.html)
- [GO-NEB - An amazon Matrix bot written in Go](https://github.com/matrix-org/go-neb)
- [Matrix Spec Overview](https://matrix.org/docs/spec/#matrix-apis)
- [Matrix Client-Server API Spec](https://matrix.org/docs/spec/client_server/r0.4.0.html)

## Author

Matthew Emerson

## License

Released under MIT License.
