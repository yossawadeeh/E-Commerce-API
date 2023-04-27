package database

import (
	"e-commerce-api/models"
	"e-commerce-api/utils"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() error {

	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		utils.ViperGetString("postgres.host"),
		utils.ViperGetString("postgres.user"),
		utils.ViperGetString("postgres.password"),
		utils.ViperGetString("postgres.dbname"),
		utils.ViperGetString("postgres.port"))

	var err error
	DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	//DB.Debug()

	// Migrate Data
	//migrateDB()
	// migrateData()

	return err
}

func migrateDB() {
	fmt.Println("Migrating DB...")

	DB.AutoMigrate(&models.Test{})

	DB.AutoMigrate(&models.Permission{})
	DB.AutoMigrate(&models.Role{})
	DB.AutoMigrate(&models.RolePermission{})
	DB.AutoMigrate(&models.ShopOwner{})
	DB.AutoMigrate(&models.Employee{})

	DB.AutoMigrate(&models.ProductCategory{})
	DB.AutoMigrate(&models.Product{})

	DB.AutoMigrate(&models.Customer{})
	DB.AutoMigrate(&models.Address{})
	DB.AutoMigrate(&models.Cart{})

	DB.AutoMigrate(&models.Shipper{})

	DB.AutoMigrate(&models.PaymentType{})
	DB.AutoMigrate(&models.Payment{})

	DB.AutoMigrate(&models.OrderStatus{})
	DB.AutoMigrate(&models.Order{})
	DB.AutoMigrate(&models.OrderDetail{})

	fmt.Println("Migrate DB Done !")
}

func migrateData() {
	fmt.Println("Migrating Data...")

	// Test
	DB.Create(&models.Test{
		Description: "Test1",
	})

	// Permission
	permissionDetail1 := "View shop report"
	permission1 := models.Permission{
		PermissionName: "View Report",
		Detail:         &permissionDetail1}
	DB.Create(&permission1)

	permissionDetail2 := "Setup shop"
	permission2 := models.Permission{
		PermissionName: "View Report",
		Detail:         &permissionDetail2}
	DB.Create(&permission2)

	permissionDetail3 := "Manage Shop"
	permission3 := models.Permission{
		PermissionName: "View Report",
		Detail:         &permissionDetail3}
	DB.Create(&permission3)

	// Role
	roleAdmin := models.Role{
		RoleName: "Administator",
	}
	DB.Create(&roleAdmin)

	// RolePermission
	DB.Create(&[]models.RolePermission{
		{
			PermissionId: permission1.ID,
			RoleId:       roleAdmin.ID,
		},
		{
			PermissionId: permission2.ID,
			RoleId:       roleAdmin.ID,
		},
		{
			PermissionId: permission3.ID,
			RoleId:       roleAdmin.ID,
		},
	})

	// ShopOwner
	shopDes := "Sell everything you want !"
	shopOwner := models.ShopOwner{
		Name:        "Yeo Shop",
		Description: &shopDes,
	}
	DB.Create(&shopOwner)

	// ProductCategory
	productCat1 := models.ProductCategory{
		Name: "Food",
	}
	DB.Create(&productCat1)
	DB.Create(&[]models.ProductCategory{
		{Name: "Drink"},
		{Name: "Cat"},
		{Name: "Furniture"},
		{Name: "Clothes"},
		{Name: "Electronic"},
	})

	// Product
	proDes := "Cat meat, fresh, new from farm :)"
	product1 := models.Product{
		Name:              "Cat meat",
		Description:       &proDes,
		Amount:            100,
		Price:             1500,
		IsActive:          true,
		ShopOwnerId:       shopOwner.ID,
		ProductCategoryId: productCat1.ID,
	}
	DB.Create(&product1)

	// Shipper
	DB.Create(&models.Shipper{
		CompanyName: "Kerry",
		Phone:       nil,
	})

	// PaymentType
	DB.Create(&[]models.PaymentType{
		{TypeName: "Cash"},
		{TypeName: "Credit Card"},
	})

	// OrderStatus
	waitPay := "Please pay within 24 hours."
	paid := "Payment done, waiting for delivery."
	delivery := "The product is being delivered. It takes about 2-3 days."
	successDelivery := "The product has been sent."
	DB.Create(&[]models.OrderStatus{
		{
			StatusName:  "Waiting for payment",
			Description: &waitPay,
		},
		{
			StatusName:  "Paid",
			Description: &paid,
		},
		{
			StatusName:  "Delivery",
			Description: &delivery,
		},
		{
			StatusName:  "Successful delivery",
			Description: &successDelivery,
		},
	})

	fmt.Println("Migrate Data Done !")
}
