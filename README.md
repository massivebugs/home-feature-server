# Home Feature Server
This project is a personal server tailored to meet various personal needs. 

The main idea is to create a api server+frontend with a Desktop like UI to manage household budgets, control homemade robots, and more.

## Getting started for development
### 1. Install dependencies
- Go `1.22.4` (used for VSCode plugin and generating/running tests)
- Docker (https://docs.docker.com/engine/install/ubuntu/)

### 2. Clone the repository
```bash
    git clone https://github.com/massivebugs/home-feature-server.git
    cd home-feature-server
```

### 3. Run containers
- `make start`: Build and start up the DB, API server, and Node front end server/watcher.
- 🧑‍💻️ To check if local api server is up and running, go to http://localhost:1323/api/v1/ping
- ✨ To view the front end, go to http://localhost:5173/

## Utility commands
### Running migrations & rolling back
- `make db-migrate`: Runs migrations listed in `db/migration`
- `make db-rollback`: Rolls back migrations by 1 step

### Generating data access layer
- `make sqlc-generate`: Runs `sqlc generate` based on queries defined in `db/query`

## Contributing
As this is a personal project, contributions are welcome but not expected. If you have any ideas or improvements, feel free to open an issue or submit a pull request.

## License
This project is licensed under the MIT License. See the LICENSE file for more details.