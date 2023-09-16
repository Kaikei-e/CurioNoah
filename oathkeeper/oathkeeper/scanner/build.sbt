ThisBuild / version := "0.1.0-SNAPSHOT"

ThisBuild / scalaVersion := "2.13.12"

lazy val root = (project in file("."))
  .settings(
    name := "scanner"
  )

val http4sVersion = "0.23.23"

libraryDependencies ++= Seq(
  "org.http4s" %% "http4s-ember-client" % http4sVersion,
  "org.http4s" %% "http4s-ember-server" % http4sVersion,
  "org.http4s" %% "http4s-dsl"          % http4sVersion,
)

libraryDependencies += "org.scalatra" %% "scalatra" % "2.8.4"
libraryDependencies += "org.scalatra" %% "scalatra-scalatest" % "2.8.4" % "test"
libraryDependencies += "org.scalatra" %% "scalatra-json" % "2.8.4"
libraryDependencies += "org.json4s" %% "json4s-jackson" % "4.0.6"
