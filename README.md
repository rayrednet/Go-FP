| Name           | NRP        | Kelas     |
| ---            | ---        | ----------|
| Helsa Sriprameswari Putri | 5025221154 | PBKK (D) |
| |  | |

# Final Project Go - Scralett's Cafe

## Tech Stack
1. Language: GoLang
2. Framework: Gin
3. Database: MySQL
4. ORM: GORM

## Installation
1. Clone this project for GitHub
```
https://github.com/rayrednet/Go-FP
```

2. Navigate to the directory
```
cd Go-Fp
```
3. Install project dependencies
```
go mod tidy
```
4. Setup my SQL and use your own username and password on config/db.go
```
dsn := "root:password@tcp(127.0.0.1:3306)/fp_pbkk?charset=utf8mb4&parseTime=True&loc=Local"
```

5. Create database fp_pbkk on mySQL
```
CREATE DATABASE fp_pbkk;
```

6. Run this project
```
go run .
```
After this, you should see your project running on http://localhost:8080/

# Web App Features

1. User can create new menu for Scarlett's Cafe.
2. User can update an existing menu.
3. User can delete a desired menu.
4. User can add categories for the menu.
5. User can update an existing category.
6. User can delete a category.


## Overview

