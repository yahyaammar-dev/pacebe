package db

import (
	"fmt"
	"strings"

	"math/rand"

	"github.com/yahyaammar-dev/pacebe/services/auth"
	"github.com/yahyaammar-dev/pacebe/types"
	"gorm.io/gorm"
)

type Seeder struct {
	db *gorm.DB
}

func NewSeeder(db *gorm.DB) *Seeder {
	return &Seeder{
		db: db,
	}
}

func (s *Seeder) CreateRoles() error {
	roles := []string{"admin", "operator", "professional", "customer"}

	for i := 0; i < len(roles); i++ {
		role := &types.Role{}
		role.Name = roles[i]
		s.db.Create(&role)
	}
	return nil
}

func (s *Seeder) CreateUsers() error {
	// Sample data arrays for variety
	firstNames := []string{"John", "Jane", "Michael", "Sarah", "David", "Emma", "James", "Lisa", "Robert", "Mary"}
	lastNames := []string{"Smith", "Johnson", "Williams", "Brown", "Jones", "Garcia", "Miller", "Davis", "Rodriguez", "Martinez"}

	// Create 20 users
	for i := 0; i < 20; i++ {
		// Generate random indices for names
		firstNameIndex := i % len(firstNames)
		lastNameIndex := i % len(lastNames)

		// Create password hash
		hashedPassword, err := auth.HashPassword("password123")
		if err != nil {
			return err
		}

		user := &types.User{
			FirstName: firstNames[firstNameIndex],
			LastName:  lastNames[lastNameIndex],
			Email: fmt.Sprintf("%s.%s%d@example.com",
				strings.ToLower(firstNames[firstNameIndex]),
				strings.ToLower(lastNames[lastNameIndex]),
				i+1),
			Password: hashedPassword,
		}

		if err := s.db.Create(user).Error; err != nil {
			return err
		}
	}

	return nil
}

func (s *Seeder) AssignRoles() {
	var users []types.User
	usersResult := s.db.Find(&users)
	if usersResult.Error != nil {
		panic("Failed to fetch Users")
	}

	var roles []types.Role
	rolesResult := s.db.Find(&roles)
	if rolesResult.Error != nil {
		panic("Failed to fetch Roles")
	}

	for _, user := range users {
		user.Roles = append(user.Roles, roles[rand.Intn(len(roles))])
		if err := s.db.Save(&user).Error; err != nil {
			panic(err)
		}
	}
}

func (s *Seeder) CreateProducts() error {
	products := []types.Product{
		{
			Name:        "Sony WH-1000XM5",
			Description: "Noise-cancelling over-ear headphones",
			Price:       349.99,
			Stock:       75,
			Category:    "Audio",
			ImageURL:    "https://example.com/sony-headphones.jpg",
		},
		{
			Name:        "Dell XPS 13",
			Description: "Ultra-thin and lightweight laptop",
			Price:       1299.99,
			Stock:       40,
			Category:    "Computers",
			ImageURL:    "https://example.com/dell-xps.jpg",
		},
		{
			Name:        "Google Pixel 8",
			Description: "Flagship Android phone with great camera",
			Price:       799.99,
			Stock:       55,
			Category:    "Electronics",
			ImageURL:    "https://example.com/pixel8.jpg",
		},
		{
			Name:        "Apple Watch Series 9",
			Description: "Smartwatch with health and fitness tracking",
			Price:       399.99,
			Stock:       80,
			Category:    "Wearables",
			ImageURL:    "https://example.com/apple-watch.jpg",
		},
		{
			Name:        "Samsung QLED 4K TV",
			Description: "65-inch 4K QLED Smart TV",
			Price:       1499.99,
			Stock:       25,
			Category:    "Home Appliances",
			ImageURL:    "https://example.com/samsung-tv.jpg",
		},
		{
			Name:        "Amazon Echo Dot",
			Description: "Smart speaker with Alexa",
			Price:       49.99,
			Stock:       200,
			Category:    "Smart Home",
			ImageURL:    "https://example.com/echo-dot.jpg",
		},
		{
			Name:        "PlayStation 5",
			Description: "Next-gen gaming console",
			Price:       499.99,
			Stock:       20,
			Category:    "Gaming",
			ImageURL:    "https://example.com/ps5.jpg",
		},
		{
			Name:        "Xbox Series X",
			Description: "Powerful gaming console",
			Price:       499.99,
			Stock:       15,
			Category:    "Gaming",
			ImageURL:    "https://example.com/xbox-series-x.jpg",
		},
		{
			Name:        "Nintendo Switch",
			Description: "Hybrid gaming console",
			Price:       299.99,
			Stock:       50,
			Category:    "Gaming",
			ImageURL:    "https://example.com/nintendo-switch.jpg",
		},
		{
			Name:        "Canon EOS R5",
			Description: "Mirrorless camera with 8K video",
			Price:       3899.99,
			Stock:       10,
			Category:    "Cameras",
			ImageURL:    "https://example.com/canon-r5.jpg",
		},
		{
			Name:        "Bose QuietComfort 45",
			Description: "Noise-cancelling headphones",
			Price:       329.99,
			Stock:       60,
			Category:    "Audio",
			ImageURL:    "https://example.com/bose-qc45.jpg",
		},
		{
			Name:        "LG UltraFine 4K Monitor",
			Description: "27-inch 4K monitor",
			Price:       699.99,
			Stock:       35,
			Category:    "Computers",
			ImageURL:    "https://example.com/lg-monitor.jpg",
		},
		{
			Name:        "Fitbit Charge 5",
			Description: "Advanced fitness tracker",
			Price:       179.99,
			Stock:       90,
			Category:    "Wearables",
			ImageURL:    "https://example.com/fitbit-charge5.jpg",
		},
		{
			Name:        "Instant Pot Duo",
			Description: "7-in-1 multi-functional pressure cooker",
			Price:       99.99,
			Stock:       120,
			Category:    "Home Appliances",
			ImageURL:    "https://example.com/instant-pot.jpg",
		},
		{
			Name:        "Dyson V11 Vacuum",
			Description: "Cordless vacuum cleaner",
			Price:       599.99,
			Stock:       40,
			Category:    "Home Appliances",
			ImageURL:    "https://example.com/dyson-v11.jpg",
		},
	}

	for _, product := range products {
		if err := s.db.Create(&product).Error; err != nil {
			return err
		}
	}
	for _, product := range products {
		if err := s.db.Create(&product).Error; err != nil {
			return err
		}
	}
	for _, product := range products {
		if err := s.db.Create(&product).Error; err != nil {
			return err
		}
	}
	for _, product := range products {
		if err := s.db.Create(&product).Error; err != nil {
			return err
		}
	}
	for _, product := range products {
		if err := s.db.Create(&product).Error; err != nil {
			return err
		}
	}
	for _, product := range products {
		if err := s.db.Create(&product).Error; err != nil {
			return err
		}
	}
	for _, product := range products {
		if err := s.db.Create(&product).Error; err != nil {
			return err
		}
	}
	for _, product := range products {
		if err := s.db.Create(&product).Error; err != nil {
			return err
		}
	}
	for _, product := range products {
		if err := s.db.Create(&product).Error; err != nil {
			return err
		}
	}
	for _, product := range products {
		if err := s.db.Create(&product).Error; err != nil {
			return err
		}
	}
	for _, product := range products {
		if err := s.db.Create(&product).Error; err != nil {
			return err
		}
	}
	for _, product := range products {
		if err := s.db.Create(&product).Error; err != nil {
			return err
		}
	}
	for _, product := range products {
		if err := s.db.Create(&product).Error; err != nil {
			return err
		}
	}
	for _, product := range products {
		if err := s.db.Create(&product).Error; err != nil {
			return err
		}
	}
	for _, product := range products {
		if err := s.db.Create(&product).Error; err != nil {
			return err
		}
	}
	for _, product := range products {
		if err := s.db.Create(&product).Error; err != nil {
			return err
		}
	}
	for _, product := range products {
		if err := s.db.Create(&product).Error; err != nil {
			return err
		}
	}

	return nil
}
