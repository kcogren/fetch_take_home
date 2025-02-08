package handler

import (
	"fetch_api/api/model"
	"testing"
)

func TestMorning(t *testing.T) {

	tempStorage := make(map[string]uint64)
	rh := &ReceiptHandler{Storage: tempStorage}

	var input model.ReceiptBody
	input.Retailer = "Walgreens"
	input.PurchaseDate = "2022-01-02"
	input.PurchaseTime = "08:13"
	input.Total = "2.65"
	input.Items = []model.ReceiptItem{{ShortDescription: "Pepsi - 12-oz", Price: "1.25"}, {ShortDescription: "Dasani", Price: "1.40"}}

	result, _ := calcPoints(rh, input, "TestMorning")

	expected := uint64(15)

	if result != expected {
		t.Errorf("TestMorning expected %d got %d", expected, result)
	}

}

func TestSimple(t *testing.T) {

	tempStorage := make(map[string]uint64)
	rh := &ReceiptHandler{Storage: tempStorage}

	var input model.ReceiptBody
	input.Retailer = "Target"
	input.PurchaseDate = "2022-01-02"
	input.PurchaseTime = "13:13"
	input.Total = "1.25"
	input.Items = []model.ReceiptItem{{ShortDescription: "Pepsi - 12-oz", Price: "1.25"}}

	result, _ := calcPoints(rh, input, "TestSimple")

	expected := uint64(31)

	if result != expected {
		t.Errorf("TestSimple expected %d got %d", expected, result)
	}

}

func TestReadme_1(t *testing.T) {

	tempStorage := make(map[string]uint64)
	rh := &ReceiptHandler{Storage: tempStorage}

	var input model.ReceiptBody
	input.Retailer = "Target"
	input.PurchaseDate = "2022-01-01"
	input.PurchaseTime = "13:01"
	input.Total = "35.35"
	input.Items = []model.ReceiptItem{
		{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
		{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
		{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
		{ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
		{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: "12.00"},
	}

	result, _ := calcPoints(rh, input, "TestReadme_1")

	expected := uint64(28)

	if result != expected {
		t.Errorf("test_example_1 expected %d got %d", expected, result)
	}

}

func TestReadme_2(t *testing.T) {

	tempStorage := make(map[string]uint64)
	rh := &ReceiptHandler{Storage: tempStorage}

	var input model.ReceiptBody
	input.Retailer = "M&M Corner Market"
	input.PurchaseDate = "2022-03-20"
	input.PurchaseTime = "14:33"
	input.Total = "9.00"
	input.Items = []model.ReceiptItem{
		{ShortDescription: "Gatorade", Price: "2.25"},
		{ShortDescription: "Gatorade", Price: "2.25"},
		{ShortDescription: "Gatorade", Price: "2.25"},
		{ShortDescription: "Gatorade", Price: "2.25"},
	}

	result, _ := calcPoints(rh, input, "TestReadme_2")

	expected := uint64(109)

	if result != expected {
		t.Errorf("TestReadme_2 expected %d got %d", expected, result)
	}

}
