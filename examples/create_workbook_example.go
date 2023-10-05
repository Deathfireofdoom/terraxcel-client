package main

import (
	"fmt"
	"log"

	"github.com/Deathfireofdoom/excel-client-go/pkg/models"
	"github.com/Deathfireofdoom/terraxcel-client/client"
)

func main() {
	// Initialize client
	c, err := client.NewClient(&client.ClientConfig{
		BaseURL:   "http://localhost:8080",
		AuthToken: "your-auth-token",
	})
	if err != nil {
		log.Fatal("Error creating client:", err)
	}

	// Create workbook
	workbook := &models.Workbook{
		FileName:   "TestWorkbook",
		FolderPath: "/output/test",
		Extension:  "xlsx",
	}
	newWorkbook, err := c.CreateWorkbook(workbook)
	if err != nil {
		log.Fatal("Error creating workbook:", err)
	}
	fmt.Println("Workbook created with ID:", newWorkbook.ID)

	// Create sheet
	sheet := &models.Sheet{
		WorkbookID: newWorkbook.ID,
		Name:       "TestSheet",
	}

	newSheet, err := c.CreateSheet(sheet)
	if err != nil {
		log.Fatal("Error creating sheet:", err)
	}
	fmt.Println("Sheet created with ID:", newSheet.ID)

	// Create cell
	cell := &models.Cell{
		WorkbookID: newWorkbook.ID,
		SheetID:    newSheet.ID,
		Row:        1,
		Column:     "A",
		Value:      "Hello, world!",
	}

	newCell, err := c.CreateCell(cell)
	if err != nil {
		log.Fatal("Error creating cell:", err)
	}
	fmt.Println("Cell created with ID:", newCell.ID)

	// Update cell
	newCell.Value = "Hello, world! Updated"
	updatedCell, err := c.UpdateCell(newCell)
	if err != nil {
		log.Fatal("Error updating cell:", err)
	}
	fmt.Println("Cell updated with ID:", updatedCell.ID)
}
