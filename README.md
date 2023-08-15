# SAST HTTP/REST Service

SAST Service is an application exposing REST endpoints for simple CRUD manipulations over defined data scheme.
Application is designed according to DDD(Domain Driven Design) principles to achieve great diversity between business and model logic

Domain-driven design is predicated on the following goals:

placing the project's primary focus on the core domain and domain logic;
basing complex designs on a model of the domain;
initiating a creative collaboration between technical and domain experts to iteratively refine a conceptual model that addresses particular domain problems.

So the main objective is to create a code that is as high abstract, declarative and descriptive as possible to make it easy to understand to refine and scale even without a vast technical knowledge.
(Hiding core implementations and exposing domain driven functions)

Crucial concept of DDD is the CQRS
Command Query Responsibility Segregation (CQRS) is an architectural pattern for separating reading data (a 'query') from writing to data (a 'command'). CQRS derives from Command and Query Separation (CQS)
Application uses this technique for HTTP Server's methods implementation(GET, POST, PUT, DELETE)

- [api](api/) OpenAPI definitions
- [docker](docker/) Dockerfiles
- [internal](internal/) Application code
    - [common](common/) Commonly used tools across sast service,        decorators,logging, errors, http server
    - [sast](sast/) SAST Service code
- [scripts](scripts/) development scripts

### Running locally

```go
> docker-compose up

# ...

 ⠿ Container sast-rest-service-firestore-1  Created                                                                                                                          0.0s
 ⠿ Container sast-rest-service-mysql-1      Created                                                                                                                          0.0s
 ⠿ Container sast-rest-service-sast-http-1  Created                                                                                                                          0.0s
Attaching to sast-rest-service-firestore-1, sast-rest-service-mysql-1, sast-rest-service-sast-http-1
sast-rest-service-mysql-1      | 2023-08-15 12:03:34+00:00 [Note] [Entrypoint]: Entrypoint script for MySQL Server 8.1.0-1.el8 started.
sast-rest-service-firestore-1  | + firebase emulators:start
sast-rest-service-mysql-1      | 2023-08-15 12:03:35+00:00 [Note] [Entrypoint]: Switching to dedicated user 'mysql'
sast-rest-service-mysql-1      | 2023-08-15 12:03:35+00:00 [Note] [Entrypoint]: Entrypoint script for MySQL Server 8.1.0-1.el8 started.
sast-rest-service-mysql-1      | '/var/lib/mysql/mysql.sock' -> '/var/run/mysqld/mysqld.sock'
sast-rest-service-mysql-1      | 2023-08-15T12:03:35.390123Z 0 [System] [MY-015015] [Server] MySQL Server - start.
sast-rest-service-mysql-1      | 2023-08-15T12:03:35.613057Z 0 [Warning] [MY-011068] [Server] The syntax '--skip-host-cache' is deprecated and will be removed in a future release. Please use SET GLOBAL host_cache_size=0 instead.
sast-rest-service-mysql-1      | 2023-08-15T12:03:35.615298Z 0 [System] [MY-010116] [Server] /usr/sbin/mysqld (mysqld 8.1.0) starting as process 1
sast-rest-service-mysql-1      | 2023-08-15T12:03:35.622411Z 1 [System] [MY-013576] [InnoDB] InnoDB initialization has started.
sast-rest-service-mysql-1      | 2023-08-15T12:03:35.786553Z 1 [System] [MY-013577] [InnoDB] InnoDB initialization has ended.
sast-rest-service-mysql-1      | 2023-08-15T12:03:35.898579Z 0 [Warning] [MY-010068] [Server] CA certificate ca.pem is self signed.
sast-rest-service-mysql-1      | 2023-08-15T12:03:35.898882Z 0 [System] [MY-013602] [Server] Channel mysql_main configured to support TLS. Encrypted connections are now supported for this channel.
sast-rest-service-mysql-1      | 2023-08-15T12:03:35.900733Z 0 [Warning] [MY-011810] [Server] Insecure configuration for --pid-file: Location '/var/run/mysqld' in the path is accessible to all OS users. Consider choosing a different directory.
sast-rest-service-mysql-1      | 2023-08-15T12:03:35.912011Z 0 [System] [MY-011323] [Server] X Plugin ready for connections. Bind-address: '::' port: 33060, socket: /var/run/mysqld/mysqlx.sock
sast-rest-service-mysql-1      | 2023-08-15T12:03:35.912289Z 0 [System] [MY-010931] [Server] /usr/sbin/mysqld: ready for connections. Version: '8.1.0'  socket: '/var/run/mysqld/mysqld.sock'  port: 3306  MySQL Community Server - GPL.
sast-rest-service-firestore-1  | ⚠  emulators: You are not currently authenticated so some features may not work correctly. Please run firebase login to authenticate the CLI.
sast-rest-service-firestore-1  | ⚠  emulators: Support for Java version <= 10 will be dropped soon in firebase-tools@11. Please upgrade to Java version 11 or above to continue using the emulators.
sast-rest-service-firestore-1  | i  emulators: Starting emulators: firestore
sast-rest-service-firestore-1  | ⚠  emulators: It seems that you are running multiple instances of the emulator suite for project sast-rest-service. This may result in unexpected behavior.
sast-rest-service-firestore-1  | ⚠  firestore: Did not find a Cloud Firestore rules file specified in a firebase.json config file.
sast-rest-service-firestore-1  | ⚠  firestore: The emulator will default to allowing all reads and writes. Learn more about this option: https://firebase.google.com/docs/emulator-suite/install_and_configure#security_rules_configuration.
sast-rest-service-firestore-1  | i  firestore: Firestore Emulator logging to firestore-debug.log
sast-rest-service-sast-http-1  | [00] Starting service
sast-rest-service-firestore-1  | i  ui: Emulator UI logging to ui-debug.log
sast-rest-service-firestore-1  | 
sast-rest-service-firestore-1  | ┌─────────────────────────────────────────────────────────────┐
sast-rest-service-firestore-1  | │ ✔  All emulators ready! It is now safe to connect your app. │
sast-rest-service-firestore-1  | │ i  View Emulator UI at http://localhost:4000                │
sast-rest-service-firestore-1  | └─────────────────────────────────────────────────────────────┘
sast-rest-service-firestore-1  | 
sast-rest-service-firestore-1  | ┌───────────┬──────────────┬─────────────────────────────────┐
sast-rest-service-firestore-1  | │ Emulator  │ Host:Port    │ View in Emulator UI             │
sast-rest-service-firestore-1  | ├───────────┼──────────────┼─────────────────────────────────┤
sast-rest-service-firestore-1  | │ Firestore │ 0.0.0.0:8787 │ http://localhost:4000/firestore │
sast-rest-service-firestore-1  | └───────────┴──────────────┴─────────────────────────────────┘
sast-rest-service-firestore-1  |   Emulator Hub running at localhost:4400
sast-rest-service-firestore-1  |   Other reserved ports: 4500
sast-rest-service-firestore-1  | 
sast-rest-service-firestore-1  | Issues? Report them at https://github.com/firebase/firebase-tools/issues and attach the *-debug.log files.
sast-rest-service-firestore-1  |  
sast-rest-service-firestore-1  | ⚠  emulators: Support for Java version <= 10 will be dropped soon in firebase-tools@11. Please upgrade to Java version 11 or above to continue using the emulators.
sast-rest-service-sast-http-1  | [00] INFO[0000] Starting HTTP server on port:3000    
```

Docker deamon has to be running in the background. It exposes HTTP port 3000 for performing requests. 

## General Architecture


![Architecture](/arch.jpg)