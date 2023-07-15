package com.the3ridge.oathkeeper.repository

import io.github.cdimascio.dotenv.dotenv

public fun initializeConnection() {
    println("Initializing connection...")

//    val url = System.getenv("DATABASE_URL")
//    val username = System.getenv("DATABASE_USERNAME")
//    val password = System.getenv("DATABASE_PASSWORD")
    val dotenv = dotenv()
    val url = dotenv["DATABASE_URL"]
    val username = dotenv["DATABASE_USERNAME"]
    val password = dotenv["DATABASE_PASSWORD"]





}