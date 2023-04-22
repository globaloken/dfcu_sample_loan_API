## Local Setup
1. Install Go on your machine
2. In the root folder create a .env file from .env-example
3. Run `go mod tidy` for base dependency installation
4. We'll use `make` for the remaining process.
    
    ```bash
    build                          Build server
    db_docs                        Generate database docs
    db_schema                      Generate database schema
    key                            Generate assymeric key
    migratedown                    Run all migrations down
    migrateup                      Run all migrations up
    mock                           Generate test mocks
    server                         Run server
    simulation                     Run simulation script
    sqlc                           Generate sqlc files
    test                           Run tests
    ```

5. Start a postgres instance by using `docker-compose up`
6. Run migrations using `make migrateup` to setup the base database. This will also setup a base admin account for the API with username: *admin* and password: *admin@admin.com*
7. Start development server with `make server`

## Simulations
To be able to run simulations, after the base setup above. Simply run `make simulation` that will run simulations and generate a report in `simulations.json`

## Admin
The admin portal is built in vue
Simply add the particular configurations in .env for tthe API end-point and access ttoken storage key then run `yarn or npm install` to install dependencies then build the frontend with `yarn build` or `npm run build`

The admin is automatically bundled with go and can be accessed at `http://<server>:<port>/admin`

## API Documentation
The API documentation is setup with swagger at  `http://<server>:<port>/swagger`

## Deployment
I've setup a docker-compose file for quick deployment with docker. First setup the environment variables then you can uncomment the api docker-compose configuration and get it running alongside the databse using `docker-compose up` 

> Note: I've also built the server executable that can be launched with `./server`, only make sure to configure the .env well.

> Note 2: All documentation is made with an assumption of a unix operating system for development.
