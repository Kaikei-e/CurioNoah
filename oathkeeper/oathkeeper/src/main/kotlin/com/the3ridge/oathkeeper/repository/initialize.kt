package com.the3ridge.oathkeeper.repository

import io.github.cdimascio.dotenv.dotenv // Ensure this library is imported correctly
import org.springframework.boot.jdbc.DataSourceBuilder
import org.springframework.context.annotation.Bean
import org.springframework.context.annotation.Configuration
import javax.sql.DataSource

@Configuration
class DataSourceConfig {

    @Bean
    fun dataSource(): DataSource {
        val dotenv = dotenv() // Loads the .env file into environment variables

        return DataSourceBuilder.create()
            .url(dotenv["DB_URL"])
            .username(dotenv["DB_USER"])
            .password(dotenv["DB_PASS"])
            .driverClassName("org.postgresql.Driver")
            .build()
    }
}