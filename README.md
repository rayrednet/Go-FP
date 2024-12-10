# 🌟 Final Project Go - Scralett's Cafe

## 📝 Group Information
- *Group Number:* 9
- *Class:* Pemrograman Berbasis Kerangka Kerja (D)

## 👥 Group Members

| Name                       | NRP         |
|----------------------------|-------------|
| Helsa Sriprameswari Putri  | 5025221154  |
| Rayssa Ravelia             | 5025211219  |

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
7. User can view and add a new review

# Table Structures

## Overview

### Home Page
![image](https://github.com/user-attachments/assets/b26a0c26-3492-4309-b079-a6a312f650e8)

## Menu Page
![image](https://github.com/user-attachments/assets/864630c4-3cf7-456a-937c-4387aecf040a)

## View Menu
![image](https://github.com/user-attachments/assets/f4bfb9a0-d53d-4b83-ab96-32130499d0e3)

## Detail Menu 
![image](https://github.com/user-attachments/assets/e3104226-246f-4804-8b85-9e076405e70e)

## Delete Menu
![image](https://github.com/user-attachments/assets/cfa5d60d-6d15-4782-9379-3cada8d27988)

## Edit Menu
![image](https://github.com/user-attachments/assets/bc3e263f-83fd-4f2c-a987-e9d34e6acbf6)

## Category Page
![image](https://github.com/user-attachments/assets/64aed20f-660a-4e67-b69c-da041ff31fdb)

## Add New Category
![image](https://github.com/user-attachments/assets/bcf4539f-ce58-4d51-a258-035eb528bdf2)

## Delete Category
![image](https://github.com/user-attachments/assets/4367dad3-98a6-446d-881f-cd923c7f782d)

## Review Page


![image](https://github.com/user-attachments/assets/ba128f6b-bf12-464e-9c4e-cb7a3d91529a)

## Add Review

![image](https://github.com/user-attachments/assets/f01bc9d6-4938-4798-81f0-1321a1a350e0)