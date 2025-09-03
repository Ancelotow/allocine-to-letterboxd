# AlloCiné → Letterboxd Sync Reviews

[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Platform](https://img.shields.io/badge/Platform-Console-lightgrey.svg)](#)
[![Language](https://img.shields.io/badge/Go-1.25-blue.svg?logo=go)](https://go.dev/doc/install)
[![IDE](https://img.shields.io/badge/IDE-GoLand-7161e9.svg?logo=jetbrains)](https://www.jetbrains.com/fr-fr/go/)


A Go script to export movie reviews from AlloCiné and import them into Letterboxd.


## ⚠️ ️**DISCLAIMER**
This project is not affiliated with or endorsed by AlloCiné or Letterboxd. 
It is a personal project created to facilitate the transfer of reviews data between the two platforms.
Please use it responsibly and respect the terms of service of both platforms.


## ✨ Features
- Export movie reviews from AlloCiné in CSV format.
- Import reviews into Letterboxd using their import feature. _**(in progress...)**_


## 🚀 Getting Started

### Requirements
- [Go 1.25](https://go.dev/doc/install)
- [GoLand IDE](https://www.jetbrains.com/fr-fr/go/) (optional)

### Run with Console
1. Clone the repository: 
```shell
git clone https://github.com/Ancelotow/allocine-to-letterboxd.git
```

2. Go to the project directory
```shell
cd allocine-to-letterboxd/src
```

3. Install dependencies
```shell
go install
```

5. Add **.env** file in the root directory of the project with your JWT token from AlloCiné
```ENV
JWT_TOKEN=<YOUR_TOKEN>
```

6. Run `main.go`
```shell
go run cmd/app/main.go
```

7. The script will generate a `allocine-reviews.csv` file in the `output` directory, located in the root directory of the project.

8. Enjoy !

## ⚖️ License
This project is licensed under the MIT License.
See the [LICENSE](https://github.com/Ancelotow/allocine-to-letterboxd/blob/main/LICENSE) file for details.


## 📬 Contact
- Author: Owen ANCELOT
- Email: ancelotow@icloud.com
- GitHub: [Ancelotow](https://github.com/Ancelotow)