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

## Developer mode (live reload with Air)

```sh
docker-compose up --force-recreate
```

```
change-me_1  | 
change-me_1  |   __    _   ___  
change-me_1  |  / /\  | | | |_) 
change-me_1  | /_/--\ |_| |_| \_ v1.12.1 // live reload for Go apps, with Go1.14.0
change-me_1  | 
change-me_1  | watching .
change-me_1  | watching configs
change-me_1  | watching controllers
change-me_1  | watching controllers/healthcheck
change-me_1  | watching controllers/version
change-me_1  | watching models
change-me_1  | watching models/foo
change-me_1  | watching pkg
change-me_1  | watching pkg/configuration
change-me_1  | watching pkg/system
change-me_1  | !exclude tmp
change-me_1  | building...
```
### Testing 

```sh
curl 0.0.0.0:8080/version -i
```

```
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Sun, 02 Aug 2020 22:06:54 GMT
Content-Length: 16

{"version":"v1"}
```

## Unit Tests 

```bash
ENVIRONMENT=test go test -v -cover
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

## Features 

### Chaos Monkey 

This boilerplate works with a side-project called [gin-chaos-monkey](https://github.com/msfidelis/gin-chaos-monkey).

You can enable chaos assaults in your app setting some environment variables in runtime, like:

```yml 
// ... 
    environment:
      - ENVIRONMENT=dev
      - CHAOS_MONKEY_ENABLED=true
      - CHAOS_MONKEY_MODE=critical
      - CHAOS_MONKEY_LATENCY=true
      - CHAOS_MONKEY_EXCEPTION=false
      - CHAOS_MONKEY_APP_KILLER=true
      - CHAOS_MONKEY_MEMORY=false
```

You can enable assaults in production setting this variables in your platform cluster runtime. For more information, [read the docs](https://github.com/msfidelis/gin-chaos-monkey)! 

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