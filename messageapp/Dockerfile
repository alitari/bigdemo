FROM maven:3.5.4-jdk-10-slim AS mavenrepo
WORKDIR /app
COPY pom.xml .
RUN mvn clean package --fail-never

FROM mavenrepo AS mavenbuild 
COPY src/ ./src
RUN mvn clean package

FROM openjdk:10-jre-slim
COPY --from=mavenbuild /app/target/*.jar app.jar
COPY entrypoint.sh /entrypoint.sh
RUN chmod u+x /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]
