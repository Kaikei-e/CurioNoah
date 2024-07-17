FROM gradle:jdk21

RUN apt-get update && apt-get install -y

ADD ./news-generator /app

WORKDIR /app

RUN ./gradlew build

CMD ["java", "-jar", "/app/build/libs/newsGenerator.jar"]