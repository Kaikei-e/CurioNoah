FROM gradle:jdk21 AS build

RUN apt-get update && apt-get install -y

ADD ./news-generator /app

WORKDIR /app

RUN ./gradlew build


FROM openjdk:21-jdk-slim

RUN apt-get update && apt-get install -y curl

WORKDIR /app

COPY --from=build /app/.env /app/.env
COPY --from=build /app/build/libs/*.jar /app/newsGenerator.jar

CMD ["java", "-jar", "/app/newsGenerator.jar"]