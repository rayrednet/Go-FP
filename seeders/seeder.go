package seeders

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"

	"go-final-project/models" 
	"gorm.io/gorm"
)

func SeedDatabase(db *gorm.DB) {
	seedCategories(db)
	seedProducts(db)
	seedReviews(db)
	SeedBaristas(db)

	log.Println("Database seeding completed!")
}

func seedCategories(db *gorm.DB) {
	var categories = []models.Category{
		{Name: "Breakfast", Description: "Morning delights to start your day right."},
		{Name: "Ala Carte", Description: "Individually crafted dishes for your appetite."},
		{Name: "Dessert", Description: "Sweet treats to satisfy your cravings."},
		{Name: "Appetizer", Description: "Small bites to get you started."},
		{Name: "Beverages", Description: "Refreshing drinks to pair with your meal."},
	}

	for _, category := range categories {
		if err := db.FirstOrCreate(&category, models.Category{Name: category.Name}).Error; err != nil {
			log.Printf("Failed to seed category %s: %v", category.Name, err)
		}
	}
	log.Println("Categories seeded successfully!")
}

func seedProducts(db *gorm.DB) {
	var categories []models.Category
	if err := db.Find(&categories).Error; err != nil {
		log.Fatalf("Failed to fetch categories: %v", err)
	}

	rand.Seed(time.Now().UnixNano())
	productNames := []string{
		"Matcha Pancakes", "Avocado Toast", "Chocolate Lava Cake",
		"Caesar Salad", "Mushroom Soup", "Egg Benedict",
		"Berry Smoothie Bowl", "Tiramisu", "French Onion Soup",
		"Green Tea Latte", "Iced Mocha", "Caramel Frappuccino",
		"Truffle Fries", "Shrimp Cocktail", "Matcha Soufflé",
	}
	productDescriptions := []string{
		"A delightful dish made with premium matcha powder.",
		"Perfectly toasted bread topped with fresh avocado and seasoning.",
		"A rich and decadent chocolate dessert with a gooey center.",
		"Crisp romaine lettuce with a creamy Caesar dressing.",
		"A warm and hearty soup made with fresh mushrooms.",
		"Classic eggs benedict served with hollandaise sauce.",
		"A healthy and refreshing smoothie bowl with assorted berries.",
		"Classic Italian dessert with layers of coffee-soaked biscuits and mascarpone.",
		"A savory soup with caramelized onions and a cheese crust.",
		"A creamy and aromatic latte with a hint of green tea.",
		"A refreshing iced coffee with chocolate and whipped cream.",
		"A sweet and creamy blended coffee topped with caramel.",
		"Crispy golden fries with a hint of truffle oil.",
		"Succulent shrimp served with a tangy cocktail sauce.",
		"A fluffy soufflé infused with premium matcha flavor.",
	}

	for i := 0; i < 50; i++ {
		
		startHour := rand.Intn(12) + 6 
		endHour := startHour + rand.Intn(5) + 1 
		
		availableHour := fmt.Sprintf("%d AM to %d PM", startHour, endHour)

		product := models.Product{
			Name:          productNames[rand.Intn(len(productNames))],
			Description:   productDescriptions[rand.Intn(len(productDescriptions))],
			
			Price:         math.Round((rand.Float64()*50+10)*100) / 100, 
			CategoryID:    categories[rand.Intn(len(categories))].ID,
			AvailableHour: availableHour,
			Rating:        math.Round((rand.Float64()*4+1)*100) / 100,  
			Stock:         rand.Intn(100) + 1,   
		}

		if err := db.Create(&product).Error; err != nil {
			log.Printf("Failed to seed product %s: %v", product.Name, err)
		}
	}
	log.Println("Products seeded successfully!")
}

func seedReviews(db *gorm.DB) {
	var products []models.Product
	if err := db.Find(&products).Error; err != nil {
		log.Fatalf("Failed to fetch products: %v", err)
	}

	rand.Seed(time.Now().UnixNano())
	customerNames := []string{
		"John Doe", "Jane Smith", "Alex Johnson", "Chris Lee", "Emily Davis",
	}
	customerEmails := []string{
		"johndoe@example.com", "janesmith@example.com", "alexjohnson@example.com", 
		"chrislee@example.com", "emilydavis@example.com",
	}
	reviews := []string{
		"Delicious! Highly recommend this dish.",
		"The taste was amazing, will definitely order again.",
		"Not what I expected, but still good.",
		"Great flavor, but a bit too salty for my taste.",
		"Absolutely loved it, one of the best meals I've had.",
	}

	// Loop to create 15 reviews
	for i := 0; i < 15; i++ {
		review := models.Review{
			CustName:  customerNames[rand.Intn(len(customerNames))],
			CustEmail: customerEmails[rand.Intn(len(customerEmails))],
			Review:    reviews[rand.Intn(len(reviews))],
			ProductID: products[rand.Intn(len(products))].ID,
		}

		if err := db.Create(&review).Error; err != nil {
			log.Printf("Failed to seed review: %v", err)
		}
	}
	log.Println("Reviews seeded successfully!")
}

func SeedBaristas(db *gorm.DB) {
	rand.Seed(time.Now().UnixNano())

	baristaNames := []string{
		"John Doe", "Jane Smith", "Emily Davis", "Michael Brown", "Sophia Wilson",
		"James Johnson", "Olivia Taylor", "William Martinez", "Charlotte Anderson", "Benjamin Thomas",
		"Amelia White", "Alexander Harris", "Mia Clark", "Ethan Lewis", "Harper Robinson",
	}

	for i := 0; i < 20; i++ {
		barista := models.Barista{
			Name:       baristaNames[rand.Intn(len(baristaNames))],
			Experience: rand.Intn(11), // Random years of experience (0–10 years)
			ProfilePic: fmt.Sprintf("https://picsum.photos/seed/%d/200/200", rand.Intn(1000)), // Random image from Picsum
			Rating:     math.Round((rand.Float64()*4+1)*100) / 100, // Rating between 1.0 and 5.0
		}

		if err := db.Create(&barista).Error; err != nil {
			log.Printf("Failed to seed barista %s: %v", barista.Name, err)
		}
	}
	log.Println("Baristas seeded successfully!")
}
