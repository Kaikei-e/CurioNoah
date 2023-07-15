package com.the3ridge.oathkeeper.repository

import io.github.cdimascio.dotenv.dotenv
import org.springframework.boot.jdbc.DataSourceBuilder
import org.springframework.context.annotation.Bean
import org.springframework.context.annotation.Configuration
import javax.sql.DataSource

@Configuration
class DataSourceConfig {

    @Bean
    fun dataSource(): DataSource {
        val dotenv = dotenv()

        return DataSourceBuilder.create()
            .url(dotenv["DB_URL"])
            .username(dotenv["DB_USERNAME"])
            .password(dotenv["DB_PASSWORD"])
            .driverClassName("org.postgresql.Driver")
            .build()
    }
}

//public fun initializeConnection() {
//    println("Initializing connection...")
//
//    val dotenv = dotenv()
//    val host = dotenv["DB_HOST"]
//    val port = dotenv["DB_PORT"]
//    val database = dotenv["DB_NAME"]
//    val user = dotenv["DB_USER"]
//    val password = dotenv["DB_PASSWORD"]
//
//    val url = "jdbc:postgresql://$user:$password@$host:$port/$database"
//
//    println("Required information was loaded.")
//    println("Connecting to database...")
//
//    val connection = java.sql.DriverManager.getConnection(url)
//
//    println("Connected to database.")
//
//    connection.close()
//
//}