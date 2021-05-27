# GOWT
Sample crud web application project using Golang(http, templates, os, sql), Bootstrap 4, DataTables, MySQL, Docker.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

What things you need to install the software

* Golang, preferably the latest version (1.16).
* MySQL Database
* Docker (optional)

### Installing

1. Clone this repository

```
git clone https://github.com/le4ndro/gowt.git
cd gowt
```

2. Run below command and install dependencies

```
go mod download
```

3. Create database on MySQL

```
CREATE DATABASE gowtdb CHARACTER SET utf8 COLLATE utf8_unicode_ci;

USE gowtdb;

CREATE TABLE tools (
  id int(11) NOT NULL AUTO_INCREMENT,
  name varchar(80) COLLATE utf8_unicode_ci DEFAULT NULL,
  category varchar(80) COLLATE utf8_unicode_ci DEFAULT NULL,
  url varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  rating int(11) DEFAULT NULL,
  notes text COLLATE utf8_unicode_ci,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
```

4. Create a .env file with the variables listed bellow and change values as needed

```
DATABASE_NAME="gowtdb"
DATABASE_USERNAME="user"
DATABASE_PASSWORD="pass"
DATABASE_SERVER="localhost"
DATABASE_PORT="3306"
```

6. Run the application

```
make run
```

## Deployment

1. Create an executable

```
make build
```

2. Run the application

```
./out/bin/gowt
\out\bin\main.exe (Windows)
```
## Create Docker image

1. To build and tag your image locally

```
make docker-build
```

2. To push your image to registry

```
make docker-release
```

## Run Docker image locally

```
make docker-run
```


## Built With

* [Golang](https://golang.org/doc/) - The programming language 
* [Bootstrap 4](https://getbootstrap.com/docs/4.2/getting-started/introduction/) - Html framework
* [jQuery](https://api.jquery.com/) - Javascript library
* [DataTables](https://datatables.net/manual/) - Advanced tables - Plug-in for jQuery

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/your/project/tags).

## Authors

* **Leandro Souza** - *Initial work*

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details

## Acknowledgments

* This project is in development
