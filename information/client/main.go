package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/ss810n/go-grpc/information/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main(){

	grpcOptions :=[]grpc.DialOption{}
	creds, err := credentials.NewClientTLSFromFile("../../cert/server.crt", "")
  if err != nil {
        log.Fatalf("could not load tls cert: %s", err)
  }
	grpcOptions = append(grpcOptions, grpc.WithTransportCredentials(creds))
  conn, err := grpc.NewClient("localhost:9091", grpcOptions...)
  if err != nil {
        log.Fatalf("did not connect: %v", err)
	}

	//----------- client with out security ----------
	// conn, err :=grpc.NewClient(":9091",grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if err!=nil{
	// 	log.Fatal("failed to connect server from client. error : ",err)
	// }

	defer conn.Close()

	client := pb.NewPublicInformationServiceClient(conn)
	//------------unidirectional streaming--------------
	// LetsCommunicate(client)
	//---------------server streaming-------------------
	// LetsCommunicateForServerStreaming(client)
	//---------------client streaming------------------- 
	// LetsCommunicateForClientStreaming(client)
	//------------bidirectional streaming---------------
	LetsCommunicateForBiDirectionalStreaming(client)
}

func LetsCommunicate(c pb.PublicInformationServiceClient){
	log.Printf("we are inside letsCommunicate .... ")
	res, err :=c.GetDetails(context.Background(), &pb.BasicInfo{
		Firstname: "Soumyadip",
		Lastname: "Sarkar",
		Location: "Kolkata",
	})
	if err !=nil{
		log.Fatal("unable to communicate ... ",err)
	}
	log.Printf("Name: %s , Country: %s , Message: %s ",res.Name,res.Country,res.Message)
}

func LetsCommunicateForServerStreaming(c pb.PublicInformationServiceClient){
	log.Print("we are waiting for server streaming ..... ")
	req :=&pb.BasicInfo{
		Firstname: "Soumyadip",
		Lastname: "Sarkar",
		Location: "Kolkata",
	}
	stream, err :=c.ServerStreamGetDetails(context.Background(),req)
	if err!=nil{
		log.Fatal("unable to connect for server streaming .... error : ",err)
	}
	for {
		message, err :=stream.Recv()
		if err== io.EOF{
			break
		}
		if err!=nil{
			log.Fatal(" unable to read data for server streaming")
		}
		log.Printf("Server Streaming --> Name: %s , Country: %s , Message: %s ",message.Name,message.Country,message.Message)
	}
}

func LetsCommunicateForClientStreaming(c pb.PublicInformationServiceClient){
	log.Print("we are waiting for client streaming ..... ")
	infos :=[]*pb.BasicInfo{
		{
			Firstname: "Soumyadip",Lastname: "Sarkar",Location: "Kolkata",
		},
		{
			Firstname: "Arun",Lastname: "Sharma",Location: "Jammu",
		},
		{
			Firstname: "Shubham",Lastname: "Patil",Location: "Pune",
		},
	}

	stream, err :=c.ClientStreamGetDetails(context.Background())
	if err!=nil{
		log.Fatal("unable start client streaming ....")
	}
	for _, info :=range infos{
		stream.Send(info)
	}
	res, err :=stream.CloseAndRecv()
	if err!=nil{
		log.Fatal("unable to get the response from server for clientstreaming ....")
	}
	log.Print("final result for client streaming ... ",res)

}

func LetsCommunicateForBiDirectionalStreaming(c pb.PublicInformationServiceClient){
	log.Print(" we are ready for bidirectional streaming from client side ")
	infos :=[]*pb.BasicInfo{
		{	Firstname: "Soumyadip",Lastname: "Sarkar",Location: "Kolkata",},
		{Firstname: "Arun",Lastname: "Sharma",Location: "Jammu",},
		{Firstname: "Shubham",Lastname: "Patil",Location: "Pune",},
	}
	stream, err :=c.BiDirectionalStreamGetDetails(context.Background())
	if err!=nil{
		log.Fatal("unable to start the bidirectional stream from client side ")
	}
	waitChannel :=make(chan struct{})

	go func ()  {
		for _, info := range infos {
			log.Print(" send request ",info)
			stream.Send(info)
			time.Sleep(50*time.Second)
		}
		stream.CloseSend()
	}()
	go func ()  {
		for {
			res , err :=stream.Recv()
			if err==io.EOF{
				break
			}
			if err!=nil{
				log.Print("unable to receive response from client side for bidirectional stream")
				break
			}
			log.Printf("BiDirectional Streaming --> Name: %s , Country: %s , Message: %s ",res.Name,res.Country,res.Message)
		}
		// close(waitChannel)
		waitChannel <- struct{}{}
	}()
	<- waitChannel
	close(waitChannel)
}