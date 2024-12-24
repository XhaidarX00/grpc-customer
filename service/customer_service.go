package service

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "microservice/proto/customer"
)

type CustomerService struct {
	pb.UnimplementedUserServiceServer
}

type Customer struct {
	Name     string
	Email    string
	Password string
}

func (c *CustomerService) CheckDetailByEmail(ctx context.Context, req *pb.CheckEmailRequest) (*pb.CheckEmailResponse, error) {
	log.Println("CheckDetailByEmail Called")

	if req.GetEmail() == "" {
		return nil, fmt.Errorf("email is required")
	}

	customer := Customer{
		Name:     "Haidar",
		Email:    "haidar@example.com",
		Password: "pass",
	}

	response := &pb.CheckEmailResponse{
		Message: "success get customer",
		Data: &pb.UserData{
			Id:        1,
			Name:      customer.Name,
			Email:     customer.Email,
			Role:      "admin",
			CreatedAt: time.Now().Format(time.RFC3339),
			UpdatedAt: time.Now().Format(time.RFC3339),
		},
	}

	return response, nil
}

// UpdateUserByEmail updates the user's details by email
func (c *CustomerService) UpdateUserByEmail(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	log.Println("UpdateUserByEmail Called")

	// Validate the request
	if req.GetEmail() == "" {
		return nil, fmt.Errorf("email is required")
	}
	if req.GetName() == "" {
		return nil, fmt.Errorf("name is required")
	}
	if req.GetPassword() == "" {
		return nil, fmt.Errorf("password is required")
	}

	// Simulating database update (you should replace this with actual DB logic)
	updatedCustomer := Customer{
		Name:     req.GetName(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}

	// Log the update operation (simulated)
	log.Printf("Customer updated: %+v\n", updatedCustomer)

	response := &pb.UpdateUserResponse{
		Message: "success update customer",
	}

	return response, nil
}
