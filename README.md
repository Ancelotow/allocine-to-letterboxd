# AlloCiné → Letterboxd Sync Reviews

[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Platform](https://img.shields.io/badge/Platform-Console-lightgrey.svg)](#)
[![Language](https://img.shields.io/badge/Go-1.25-blue.svg?logo=go)](https://go.dev/doc/install)
[![IDE](https://img.shields.io/badge/IDE-GoLand-7161e9.svg?logo=jetbrains)](https://www.jetbrains.com/fr-fr/go/)


A Go script to export movie reviews from AlloCiné and import them into Letterboxd.

## ✨ Features
- Export movie reviews from AlloCiné in CSV format rideable by Letterboxd
```CSV
Title,Year,WatchedDate,Rating,Review
"Thor: The Dark World",2013,2025-09-03,2.5,"My review" 
```

## 🚀 Getting Started

### Requirements
- [Go 1.25](https://go.dev/doc/install)
- [GoLand IDE](https://www.jetbrains.com/fr-fr/go/) (optional)

### Installing
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

### Creating the CSV file from your AlloCiné reviews

1. Run `main.go`
```shell
go run cmd/app/main.go
```

2. The script will generate a `reviews-to-imported.csv` file in the `output` directory, located in the root directory of the project.

### Importing the CSV file into Letterboxd

1. Go to the [Letterboxd import page](https://letterboxd.com/import/)
2. Click on the **Select a file** button
3. Select the `reviews-to-imported.csv` file
4. Click on the **Import** button
5. Wait for the import to complete
6. Go to the [Letterboxd diary page](https://letterboxd.com/ancelotow/films/diary/)
7. Now you can enjoy your reviews!

## ⚖️ License
This project is licensed under the MIT License.
See the [LICENSE](https://github.com/Ancelotow/allocine-to-letterboxd/blob/main/LICENSE) file for details.


## 📬 Contact
- Author: Owen ANCELOT
- Email: ancelotow@icloud.com
- GitHub: [Ancelotow](https://github.com/Ancelotow)