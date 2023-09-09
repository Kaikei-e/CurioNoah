import cats.effect.unsafe.implicits.global
import server.APIServer

object Main {
	def main(args: Array[String]): Unit = {
		APIServer.run(args.toList).unsafeRunSync()
	}
}