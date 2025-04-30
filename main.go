package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	// _ "go-swagger-example/docs" // import generated docs
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Allow all origins (you can specify specific origins if needed)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle OPTIONS preflight request
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
	})

	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/accounts/:id", getAccount)
		v1.GET("/insights/:id", getInsight)
	}

	// Swagger docs route
	r.StaticFile("/swagger.yaml", "./docs/swagger.yaml")

	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":3001")
}

type Account struct {
	ID   int    `json:"id" example:"1"`
	Name string `json:"name" example:"Current account"`
	Balance float32 `json:"balance" example:"100.0"`
	SortCode string `json:"sortCode" example:"12-34-56"`
	AccountNumber string `json:"accountNumber" example:"12345687"`
}

type Insight struct {
	ID   int    `json:"id" example:"1"`
	Insight string `json:"insight" example:"An insight"`
}

func getAccount(c *gin.Context) {
	id := c.Param("id")

	switch id {
		case "1":
			c.JSON(http.StatusOK, Account{ID: 1, Name: "Personal Current Account", Balance: 2023.00, SortCode: "12-45-78", AccountNumber: "12345678"})
		case "2":
			c.JSON(http.StatusOK, Account{ID: 2, Name: "Family Savings", Balance: 10000.00, SortCode: "12-45-78", AccountNumber: "87654321"})
		case "3":
			c.JSON(http.StatusOK, Account{ID: 3, Name: "Mortgage Account", Balance: 155685.00, SortCode: "", AccountNumber: "987650234"})
		case "4":
			c.JSON(http.StatusOK, Account{ID: 4, Name: "Sam's Children's Account", Balance: 100.00, SortCode: "12-45-78", AccountNumber: "98766443"})
		case "5":
			c.JSON(http.StatusOK, Account{ID: 5, Name: "Jess' Children's Account", Balance: 130.00, SortCode: "12-45-78", AccountNumber: "12345668"})
		case "6":
			c.JSON(http.StatusOK, Account{ID: 6, Name: "Max's Children's Account", Balance: 600.00, SortCode: "12-45-78", AccountNumber: "12344678"})
		case "7":
			c.JSON(http.StatusOK, Account{ID: 7, Name: "Children's Current Account (A Jones)", Balance: 500.00, SortCode: "12-45-78", AccountNumber: "85274169"})
		case "8":
			c.JSON(http.StatusOK, Account{ID: 8, Name: "Junior ISA (A Jones)", Balance: 2023.00, SortCode: "12-45-78", AccountNumber: "45679888"})
		default:
			c.JSON(http.StatusBadRequest, nil)

	}
}

func getInsight(c *gin.Context) {
	id := c.Param("id")

	switch id {
		case "1":
			c.JSON(http.StatusOK, Insight{ID: 1, Insight: `The gap between what you earn and what you spend is where wealth is built.
Prioritize saving a portion of your income before spending—ideally 20%—to create a buffer for emergencies, long-term goals, and financial freedom.
Tracking expenses often reveals small habits that, if adjusted, can significantly boost your savings over time.`})
		default:
			c.JSON(http.StatusBadRequest, nil)

	}
}