FROM maven:3.6.3-openjdk-15-slim as build
WORKDIR /app
COPY pom.xml .
RUN mvn dependency:go-offline
COPY src src
RUN mvn clean package -Dmaven.test.skip

FROM openjdk:15-jdk-alpine
COPY --from=build /app/target/key_generation_service-0.0.1-SNAPSHOT.jar /app/key_generation_service-0.0.1-SNAPSHOT.jar
EXPOSE 8080
ENTRYPOINT ["java","-jar","/app/key_generation_service-0.0.1-SNAPSHOT.jar"]
