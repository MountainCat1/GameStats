# Game Stats 
### Simple API written with GO and CHI


Application created to store information about users and match data for my other project that is "Feed The Beast".
Game servers will send a request to the Game Stats service whenever match finishes to persist information about end result.
Plan in to later expand on it and maybe create some type of ranking and matchmaking system.

### Included features:
- Authorization via _Basic Auth_
- Domain Driven Design
- Docker support
- Docker Compose support
- Usage of Chi framework
- Usage of GORM

### How to run:
- Clone the repository
- Run `docker-compose up` in the root directory
- Service will be available at `http://localhost:8080`
- Swagger documentation will be available at `http://localhost:8080/swagger/index.html`