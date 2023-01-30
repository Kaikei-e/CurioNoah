import std/asynchttpserver
import std/asyncdispatch

proc main() {.async.} =
  var server = newAsyncHttpServer()
  proc cb(req: Request) {.async} =
    echo "Request: ", req.url
    await req.respond(Http200, "Hello World!")
  await server.serve(Port(9200), cb = requestHandler)
  echo "Server started"