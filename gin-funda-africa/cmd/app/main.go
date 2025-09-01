package main

import (
	"fmt"
	"gin-funda-africa/internal/modules/users"
	"gin-funda-africa/internal/modules/users/models"
	"gin-funda-africa/internal/modules/users/routes"
	"io"
	"os"

	// "path/filepath"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	Log "github.com/sirupsen/logrus"
	"gorm.io/gorm"

	docs "gin-funda-africa/internal/docs"
)

var server *gin.Engine

var appConfig *users.Config
var adapter *gormadapter.Adapter

func init() {
	logFile, err := os.OpenFile("funda-africa.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Could not open log file, defaulting to stdout only:", err)
		logFile = nil
	}

	// Set formatter and level
	Log.SetFormatter(&Log.JSONFormatter{})
	Log.SetLevel(Log.DebugLevel)

	if logFile != nil {
		// Send logs to both stdout and file
		multiWriter := io.MultiWriter(os.Stdout, logFile)
		Log.SetOutput(multiWriter)
	} else {
		// Fallback to stdout
		Log.SetOutput(os.Stdout)
	}

	// rootPath, err := filepath.Abs(filepath.Join("..", "..")) // go 2 dirs up to reach project root

	// if err != nil {
	// 	Log.Fatalf("Error resolving root path: %v", err)
	// }

	cfg, err := users.LoadConfig(".")
	if err != nil {
		Log.Fatalf("Error loading config: %v", err)
	}

	appConfig = &cfg
	Log.Infof("Environment configuration loaded successfully.")

	databaseMigrate(users.ConnectDB(appConfig))

	setupGin()

	rbacConfig()
}

func databaseMigrate(database *gorm.DB) {
	err := database.AutoMigrate(
		&models.User{},
		&models.Organization{},
		&models.SubscriptionPlan{},
		&models.UserOrganization{},
		&models.UserSubscription{})

	if err != nil {
		Log.Fatalf("❌ Failed to run migrations: %v", err)
	}
	fmt.Println("✅ Database migrations completed")

	adapter, err = gormadapter.NewAdapterByDB(database)
	if err != nil {
		Log.Fatalf("❌ Failed to initialize Casbin adapter: %v", err)
	}
	fmt.Println("✅ Casbin adapter initialized")
}

func setupGin() {
	server = gin.New()
	server.Use(gin.LoggerWithWriter(Log.New().Writer()), gin.Recovery())

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8000", appConfig.ClientOrigin}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))
}

// Helper to convert []string to []interface{}
func stringSliceToInterface(slice []string) []interface{} {
	result := make([]interface{}, len(slice))
	for i, v := range slice {
		result[i] = v
	}
	return result
}

func addPoliciesToDB(enforcer, csvEnforcer *casbin.Enforcer) {
	policies, err := csvEnforcer.GetPolicy()
	if err != nil {
		Log.Fatalf("❌ Failed to get policies from CSV: %v", err)
	}
	for _, p := range policies {
		if _, err := enforcer.AddPolicy(stringSliceToInterface(p)...); err != nil {
			Log.Warnf("⚠️ Failed to add policy %v: %v", p, err)
		}
	}

	groupings, err := csvEnforcer.GetGroupingPolicy()
	if err != nil {
		Log.Fatalf("❌ Failed to get grouping policies from CSV: %v", err)
	}
	for _, g := range groupings {
		if _, err := enforcer.AddGroupingPolicy(stringSliceToInterface(g)...); err != nil {
			Log.Warnf("⚠️ Failed to add grouping policy %v: %v", g, err)
		}
	}

	if err := enforcer.SavePolicy(); err != nil {
		Log.Fatalf("❌ Failed to save policies to DB: %v", err)
	}
}

func loadCsvPoliciesToDb(enforcer *casbin.Enforcer, configFile, csvPolicyFile string, policies, groupPolicies [][]string) {
	if len(policies) == 0 && len(groupPolicies) == 0 {
		Log.Info("📦 No policies found in DB. Seeding from CSV...")

		csvEnforcer, err := casbin.NewEnforcer(configFile, csvPolicyFile)
		if err != nil {
			Log.Fatalf("❌ Failed to load CSV policies: %v", err)
		}
		if err := csvEnforcer.LoadPolicy(); err != nil {
			Log.Fatalf("❌ Failed to read CSV policies: %v", err)
		}

		addPoliciesToDB(enforcer, csvEnforcer)
		Log.Info("✅ Policies seeded successfully.")
	} else {
		Log.Info("✅ Policies already exist in DB. Skipping seeding.")
	}
}

func rbacConfig() {
	configFile := "internal/pkg/rbac/rbac_with_domains_model.conf"
	csvPolicyFile := "internal/pkg/rbac/policy.csv"
	enforcer, err := casbin.NewEnforcer(configFile, adapter)

	if err != nil {
		Log.Fatalf("❌ Failed to initialize Casbin enforcer: %v", err)
	}

	if err := enforcer.LoadPolicy(); err != nil {
		Log.Fatalf("❌ Failed to load Casbin policies: %v", err)
	}

	policies, err := enforcer.GetPolicy()
	if err != nil {
		Log.Fatalf("❌ Failed to load policies from db: %v", err)
	}

	groupPolicies, err := enforcer.GetGroupingPolicy()
	if err != nil {
		Log.Fatalf("❌ Failed to load group policies from db: %v", err)
	}

	loadCsvPoliciesToDb(enforcer, configFile, csvPolicyFile, policies, groupPolicies)

	docs.SwaggerInfo.BasePath = "/api/v1"

	routes.RegisterAllUserServiceRoutes(server, enforcer, []byte(appConfig.JwtSecretKey))
}

// @title Funda Africa APIs
// @version 1.0
// @description Testing Swagger APIs.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apiKey CookieAuth
// @in cookie
// @name session_token

// @schemes http
func main() {
	server.Static("/uploads", "./uploads")

	Log.Fatal(server.Run(":" + appConfig.ServerPort))
}
