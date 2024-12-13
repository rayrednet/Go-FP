# 🌟 Final Project Go - Scarlett's Cafe  

## 📝 Group Information  
- *Group Number:* 9  
- *Class:* Pemrograman Berbasis Kerangka Kerja (D)  

## 👥 Group Members  

| Name                       | NRP         |  
|----------------------------|-------------|  
| Helsa Sriprameswari Putri  | 5025221154  |  
| Rayssa Ravelia             | 5025211219  |  

## 🌍 Project Overview  
For our final project in the Go framework course, we developed a website called **Scarlett's Cafe**, a cafe catalog. This website serves as an **ADMIN** role platform, allowing full **create, read, update, and delete (CRUD)** operations for every model. We chose a pink, girly theme to represent the cafe's branding.  

## 📽️ Demo Video  
Below is the demo video showcasing the key functionalities of this project:  
[Watch the demo video here](https://youtu.be/yO-bhszipl0)  

## 💻 Tech Stack  
1. **Language:** GoLang  
2. **Framework:** Gin  
3. **Database:** MySQL  
4. **ORM:** GORM  

## ⚙️ Installation  
1. Clone this project from GitHub:  
   ```bash  
   https://github.com/rayrednet/Go-FP  
   ```  

2. Navigate to the directory:  
   ```bash  
   cd Go-FP  
   ```  

3. Install project dependencies:  
   ```bash  
   go mod tidy  
   ```  

4. Setup MySQL and configure your username and password in `.env`:  
   ```env  
   DB_USER=yourusername  
   DB_PASSWORD=yourpassword  
   DB_HOST=127.0.0.1  
   DB_PORT=3306  
   DB_NAME=fp_pbkk  
   ```  

5. Create the `fp_pbkk` database in MySQL:  
   ```sql  
   CREATE DATABASE fp_pbkk;  
   ```  

6. Run the project:  
   ```bash  
   go run main.go  
   ```  
   After this, your project should be running at [http://localhost:8080](http://localhost:8080).  

## 🛠️ Web App Features  

1. Create, update, and delete menus.  
2. Create, update, and delete menu categories.  
3. Add, update, and delete baristas.  
4. Create, update, and delete reviews.  

## 📊 Conceptual Data Model  
![Conceptual Data Model](https://github.com/user-attachments/assets/33186796-38f8-4a03-8dcc-deb3354ac259)  

## 🖼️ Web Pages Screenshot  

### 1️⃣ Homepage  
![homepage](img/home.png)


### 2️⃣ Category Section  
- **Index Page:** Shows all available categories. 
![category-index](img/category-index.png)

- **Create Page:** Allows admins to add a new category. 
![create-page](img/category-create.png)
![create-page](img/category-create-after.png)

- **Delete:** Removes an unwanted category.  
![delete](img/category-delete.png)
![delete](img/category-delete-after.png)

- **Update Page:** Enables editing of a category.  
![update-page](img/category-edit-before.png)
![update-page](img/category-edit.png)
![update-page](img/category-edit-after.png)

### 3️⃣ Menu Section  
- **Index Page:** Lists all menu items.  
![menu-index](img/menu/index.png)
![menu-index](img/menu/detail.png)

- **Create Page:** Adds a new menu item.  
![create-page](img/menu/create-form.png)
![create-page](img/menu/create-form-filled.png)
![create-page](img/menu/create-after.png)

- **Delete:** Deletes a specific menu item.  
![delete](img/menu/delete.png)
![delete](img/menu/delete-after.png)

- **Update Page:** Updates details of a menu item.  
![update-page](img/menu/edit-before.png)
![update-page](img/menu/edit.png)
![update-page](img/menu/edit-after.png)
![update-page](img/menu/edit-after-v2.png)

### 4️⃣ Barista Section  
- **Index Page:** Lists all baristas. 
![index](img/barista/index.png)

- **Create Page:** Adds a new barista.  
![create-page](img/barista/create-form.png)
![create-page](img/barista/create-form-filled.png)
![create-page](img/barista/create-after.png)

- **Delete:** Removes a barista. 
 ![delete](img/barista/delete.png)
![delete](img/barista/delete-after.png)

- **Update Page:** Updates barista details.  
![update-page](img/barista/edit-before.png)
![update-page](img/barista/edit.png)
![update-page](img/barista/edit-after.png)

### 5️⃣ Review Section  
- **Index Page:** Lists all reviews.  
![index](img/review/index.png)

- **Create Page:** Adds a new review.  
![create-page](img/review/create-form.png)
![create-page](img/review/create-form-filled.png)
![create-page](img/review/create-after.png)

- **Delete:** Removes a review.  
![delete](img/review/delete.png)
![delete](img/review/delete-after.png)

- **Update Page:** Updates review details.  
![update-page](img/review/edit-before.png)
![update-page](img/review/edit.png)
![update-page](img/review/edit-after.png)
