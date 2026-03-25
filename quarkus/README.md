quarkus create && cd code-with-quarkus



config.json -->

{
  "BaseVolumeSz": "500m",
  "Args": ["-jar", "target/quarkus-app/quarkus-run.jar"],
  "Dirs": ["src", "target"]
}

mvn clean package -Dmaven.compiler.release=21


ops pkg load eyberg/java:21.0.4 -c config.json -p 8080

