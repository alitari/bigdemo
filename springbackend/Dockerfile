FROM maven:3.5.4-jdk-10-slim AS mavenrepo
WORKDIR /app
COPY pom.xml .
RUN mvn clean package --fail-never

FROM mavenrepo AS mavenbuild 
COPY . .
RUN mvn clean package

FROM tomcat:9.0.10-jre10-slim
COPY --from=mavenbuild /app/target/*.war $CATALINA_HOME/webapps/
