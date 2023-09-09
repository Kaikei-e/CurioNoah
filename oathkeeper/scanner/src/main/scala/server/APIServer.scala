package server

import cats.effect.{ExitCode, IO, IOApp, Resource}
import com.comcast.ip4s.{Host, Port}
import org.http4s.ember.server.EmberServerBuilder
import server.system.System
import org.http4s.server.Server

object APIServer extends IOApp{
	private val host: Host = Host.fromString("localhost").get
	private val port: Port = Port.fromInt(8800).get

	val server: Resource[IO, Server] = EmberServerBuilder.default[IO]
			.withHttpApp(System.healthCheck.orNotFound)
			.withHost(host)
			.withPort(port)
			.build

	def run(args: List[String]): IO[ExitCode] = server.use(_ => IO.never).as(ExitCode.Success)
}
