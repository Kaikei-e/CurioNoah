package com.the3ridge.oathkeeper

import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.boot.runApplication

@SpringBootApplication
class OathkeeperApplication

fun main(args: Array<String>) {
    runApplication<OathkeeperApplication>(*args)
}
