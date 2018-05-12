package main

import (
	"context"
	"log"
	"os"

	pb "github.com/koffsson/shipper/user-service/proto/user"
	microClient "github.com/micro/go-micro/client"
)

func main() {
	// Create new greeter client
	client := pb.NewUserServiceClient("go.micro.srv.user", microClient.DefaultClient)

	name := "koffsson"
	email := "koffsson@gmail.com"
	password := "test123"
	company := "?"

	log.Println(name, email, password)

	r, err := client.Create(context.TODO(), &pb.User{
		Name:     name,
		Email:    email,
		Password: password,
		Company:  company,
	})
	if err != nil {
		log.Fatalf("Could not create: %v", err)
	}
	log.Printf("Created: %s", r.User.Id)

	getAll, err := client.GetAll(context.Background(), &pb.Request{})
	if err != nil {
		log.Fatalf("Could not list users: %v", err)
	}
	for _, v := range getAll.Users {
		log.Println(v)
	}

	authResponse, err := client.Auth(context.TODO(), &pb.User{
		Email:    email,
		Password: password,
	})

	log.Printf("Your access token is: %s\n", authResponse.Token)

	os.Exit(0)
}
