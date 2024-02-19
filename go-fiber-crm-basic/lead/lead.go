package lead

import (
	"go-fiber-crm-basic/database" // Importing database package for database connection

	"github.com/gofiber/fiber"                 // Importing fiber package for building web applications
	"github.com/jinzhu/gorm"                   // Importing gorm package for ORM
	_ "github.com/jinzhu/gorm/dialects/sqlite" // Importing SQLite dialect for GORM
)

// Lead represents the structure of a lead in the database
type Lead struct {
	gorm.Model
	Name    string `json:"name"`    // Name field represents the name of the lead
	Company string `json:"company"` // Company field represents the company of the lead
	Email   string `json:"email"`   // Email field represents the email of the lead
	Phone   int    `json:"phone"`   // Phone field represents the phone number of the lead
}

// GetLeads retrieves all leads from the database and returns them as JSON
func GetLeads(c *fiber.Ctx) {
	db := database.DBCon
	var leads []Lead
	db.Find(&leads)
	c.JSON(leads)
}

// GetLead retrieves a specific lead by its ID from the database and returns it as JSON
func GetLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBCon
	var lead Lead
	db.Find(&lead, id)
	c.JSON(lead)
}

// NewLead creates a new lead record in the database based on the JSON data received in the request body
func NewLead(c *fiber.Ctx) {
	db := database.DBCon
	lead := new(Lead)
	if err := c.BodyParser(lead); err != nil {
		c.Status(503).Send(err)
		return
	}
	db.Create(&lead)
	c.JSON(lead)
}

// DeleteLead deletes a specific lead from the database based on its ID
func DeleteLead(c *fiber.Ctx) {
	id := c.Params(("id"))
	db := database.DBCon
	var lead Lead
	db.First(&lead, id)
	if lead.Name == "" {
		c.Status(500).Send("No lead found with id")
		return
	}
	db.Delete(&lead)
	c.Status(204).Send()
}
