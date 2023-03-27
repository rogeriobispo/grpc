package main

import (
	"database/sql"
	"net"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rogeriobispo/grpc/internal/database"
	"github.com/rogeriobispo/grpc/internal/pb"
	"github.com/rogeriobispo/grpc/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	categoryDb := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDb)
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}

}
