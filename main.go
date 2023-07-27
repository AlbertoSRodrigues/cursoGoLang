package main

import (
	"fmt"
	"projeto/internal/domain/campaign"

	"github.com/go-playground/validator/v10"
)

func main() {
	contacts := []campaign.Contact{{Email: "asd@g.com"}}
	campaign := campaign.Campaign{Name: "Projeto Inicial", Content: "Body HTML", Contacts: contacts}
	validate := validator.New()
	err := validate.Struct(campaign)
	if err == nil {
		println("Nenhum erro")
	} else {
		validationErrors := err.(validator.ValidationErrors)
		for _, v := range validationErrors {
			switch v.Tag() {
			case "required":
				fmt.Println(v.StructField() + " is required")
			case "min":
				fmt.Println(v.StructField() + " requires " + v.Param() + " minimum characters")
			case "max":
				fmt.Println(v.StructField() + " requires " + v.Param() + " maximum characters")
			case "email":
				fmt.Println(v.StructField() + " is an invalid email")
			}
		}
	}
}
