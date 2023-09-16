package server

import org.scalatra._
import org.json4s.{DefaultFormats, Formats}
import org.scalatra.json._

class APIServer extends ScalatraServlet with JacksonJsonSupport {

	// Sets up automatic case class to JSON output serialization
	protected implicit lazy val jsonFormats: Formats = DefaultFormats

	// Before every action runs, set the content type to be in JSON format.
	before() {
		contentType = formats("json")
	}

	def apiRoute(route: String)(block: => Any): Any = {
			val fullRoutes = s"/api/v1$route"
			context(fullRoutes)(block)
	}

	def feedsRoute(route: String)(block: => Any): Any = {
		val fullRoute = s"/feeds$route"
	}

	// API v1 group
	apiRoute("") {
		feedsRoute("/fetch") {
			get("/all") {
				// Fetch all feeds
				List("feed1", "feed2")
			}
			get("/:page") {
				// Fetch feeds by page
				val page = params.getOrElse("page", "1")
			}
		}
		feedsRoute("/register") {
			post("/one") {
				// Register a single feed
				"New feed registered"
			}
		}
	}
}



//import cats.effect.{ExitCode, IO, IOApp, Resource}
//import com.comcast.ip4s.{Host, Port}
//import org.http4s.ember.server.EmberServerBuilder
//import server.system.System
//import org.http4s.server.Server
//
//object APIServer extends IOApp{
//	private val host: Host = Host.fromString("localhost").get
//	private val port: Port = Port.fromInt(8800).get
//
//	val server: Resource[IO, Server] = EmberServerBuilder.default[IO]
//			.withHttpApp(System.healthCheck.orNotFound)
//			.withHost(host)
//			.withPort(port)
//			.build
//
//	def run(args: List[String]): IO[ExitCode] = server.use(_ => IO.never).as(ExitCode.Success)
//}
