package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"

	"net/http"
	_ "net/http/pprof"

	pb "github.com/ss810n/go-grpc/information/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Server struct{
	pb.PublicInformationServiceServer
}

func main(){
	//  strat profiling for the server 
	go func(){
		http.ListenAndServe(":10001", nil)
	}()

	// start grpc server
	lis , err :=net.Listen("tcp","localhost:9091")
	if err!=nil{
		log.Fatal("Unable to listen with port 9091. error : ",err)
	}

	log.Print("Server is listning on port 9091")

	//generate server side credentials 
	creds, err := credentials.NewServerTLSFromFile("../../cert/server.crt", "../../cert/server.pem")
    if err != nil {
        log.Fatalf("Failed to generate credentials %v", err)
    }
  grpcOpts :=[]grpc.ServerOption{}
	grpcOpts = append(grpcOpts,grpc.Creds(creds) )
	server := grpc.NewServer(grpcOpts...)

	//----- server without credentials -----
	// server :=grpc.NewServer()

	// register the new grpc service
	pb.RegisterPublicInformationServiceServer(server,&Server{})


	if err =server.Serve(lis);err!=nil{
		log.Fatal("server is unable to server. error : ",err)
	}

}

func (server *Server)GetDetails(c context.Context,info *pb.BasicInfo) (*pb.DetailedInfo, error){
	log.Print("we are inside GetDetails ....... ")
	return &pb.DetailedInfo{
		Name: fmt.Sprintf("%s %s",info.Firstname,info.Lastname),
		Country: "India",
		Message: "you are from "+info.Location,
	},nil
}

func (server *Server)ServerStreamGetDetails(info *pb.BasicInfo, s pb.PublicInformationService_ServerStreamGetDetailsServer) error{
	log.Print("we are inside ServerStreamGetDetails ....... ")
	num :=rand.Intn(20)
	for i:=0;i<=num;i++{
		s.Send(&pb.DetailedInfo{
			Name: fmt.Sprintf("%s %s -%d",info.Firstname,info.Lastname,i),
			Country: "India",
			Message: fmt.Sprintf("Hello, %s %s-%d . You are from %s",info.Firstname,info.Lastname,i,info.Location),
		})
	}
	return nil
}

func (server *Server)ClientStreamGetDetails(stream pb.PublicInformationService_ClientStreamGetDetailsServer) error{
	log.Print("we are inside Client streaming ....... ")
	final_response :=""
	for {
		req , err :=stream.Recv()
		if err==io.EOF{
			return stream.SendAndClose(&pb.DetailedInfo{
				Message: final_response,
			})
		}else if err!=nil{
			log.Fatal("unable to receave from server end for client streaming ....")
		}
		final_response +=fmt.Sprintf("Name : %s %s , Location : %s \n",req.Firstname,req.Lastname,req.Location)
	}
}

func (server *Server)BiDirectionalStreamGetDetails(stream pb.PublicInformationService_BiDirectionalStreamGetDetailsServer) error{
	log.Print("we are inside BiDirectional Server streaming ..... ")
	for {
		req , err :=stream.Recv()
		if err==io.EOF{
			return nil
		}
		if err!=nil{
			log.Fatal("unable to receive stream data from server end ")
		}
		log.Print(" server received ",req)
		err =stream.Send(&pb.DetailedInfo{
			Name: fmt.Sprintf("%s %s",req.Firstname,req.Lastname),
			Country: "India",
			Message: "From "+req.Location,
		})
		if err!=nil{
			log.Fatal(" unable to send the response from server side for bidirectional stream")
		}
	}
}