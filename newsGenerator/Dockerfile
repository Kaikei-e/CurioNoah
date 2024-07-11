FROM kotlin/kotlinlang.org:2.0.0

RUN apt-get update && apt-get install -y

ADD ./newsGenerator /app

WORKDIR /app

RUN ./gradlew build

CMD ["java", "-jar", "/app/build/libs/newsGenerator.jar"]