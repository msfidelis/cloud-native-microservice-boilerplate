# Welcome to Easy Golang API Bootstrap 👋
![Version](https://img.shields.io/badge/version-v1-blue.svg?cacheSeconds=2592000)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![Twitter: fidelissauro](https://img.shields.io/twitter/follow/fidelissauro.svg?style=social)](https://twitter.com/fidelissauro)

> Project boilerplate to start new cloud native golang projects quickly

## How to start

1. Clone this boilerplate

```bash
git clone git@github.com:msfidelis/cloud-native-microservice-boilerplate.git
```

2. Change project name

> Use your CTRL + F or Search/Replace

* Search for `change-me` in project folder
* Change for your project name

## Developer mode

```sh
docker-compose up --force-recreate
```

## Production Image

```sh
docker build -t newproject .
```

## Project Structure 

```
.
├── Dockerfile (Production golang Dockerfile - Multistage Build)
├── Dockerfile.dev (Development golang Dockerfile - With Air)
├── LICENSE
├── README.md
├── configs (unsensitive configuration files)
│   ├── dev.json
│   ├── prod.json
│   └── test.json
├── controllers (controller folder)
│   ├── healthcheck (controller name / context)
│   │   └── main.go (main module for controller)
│   └── version
│       └── main.go
├── docker-compose.yml (docker-compose for development environment)
├── go.mod
├── go.sum
├── pkg (Shared libs folder)
│   ├── configuration (Lib Context / Name)
│   │   └── main.go (Main module for package)
│   └── system
│       └── main.go
└── main.go (Entrypoint / Routes)
```

## Author

👤 **Matheus Fidelis**

* Website: https://raj.ninja
* Twitter: [@fidelissauro](https://twitter.com/fidelissauro)
* Github: [@msfidelis](https://github.com/msfidelis)
* LinkedIn: [@msfidelis](https://linkedin.com/in/msfidelis)

## 🤝 Contributing

Contributions, issues and feature requests are welcome!

Feel free to check [issues page](/issues). 

## Show your support

Give a ⭐️ if this project helped you!


## 📝 License

Copyright © 2020 [Matheus Fidelis](https://github.com/msfidelis).

This project is [MIT](LICENSE) licensed.

***
_This README was generated with ❤️ by [readme-md-generator](https://github.com/kefranabg/readme-md-generator)_