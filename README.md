# GOWT
Sample crud web application project using Golang(http, templates, os, sql), Bootstrap 4, DataTables, MySQL.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

What things you need to install the software

* Golang, preferably the latest version (1.11).
* MySQL Database

### Installing

1. Clone this repository

```
$ git clone https://github.com/le4ndro/gowt.git
$ cd gowt
```

2. Run below command and install MySQL driver's

```
$ go get -u github.com/go-sql-driver/mysql
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

4. Make a local copy of set_env.sample.sh

```
$ cp set_env.sample.sh set_env.sh (or rename set_env.sample.bat to set_env.bat on Windows)
```

5. Edit your set_env.sh (or set_env.bat on Windows) file with your parameters and set variables for your environment

```
$ chmod 755 set_env.sh
$ source ./set_env.sh
(or run set_env.bat on Windows)
```

6. Run the application

```
$ go run main.go
```

## Deployment

1. Create an executable

```
$ go build main.go
```

2. Run the application

```
$ main
or main.exe (Windows)
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
