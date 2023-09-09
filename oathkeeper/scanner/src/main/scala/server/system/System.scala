package server.system

import cats.effect.IO
import org.http4s.HttpRoutes
import org.http4s.dsl.io._

object System {
	val healthCheck: HttpRoutes[IO] = HttpRoutes.of[IO] {
		case GET -> Root / "system" / "health" => Ok("System is healthy")
	}
}


