package main

import (
	"context"
	"flag"
	"log"
	"net"

	"gitlab.com/atreya2011/grpc-practice/basic-crud/basiccrud"
	"google.golang.org/grpc"
)

type service struct {
	// array of pointers to basiccrud.Fullname
	Fullname []*basiccrud.Fullname
}

// when initializing values in a composite struct
// it should start with the key name followed by :
// then it should have the type of the key name
// followed by the values
// each value should have the type in front of it
var serv = service{
	// first specify the key name,
	// then specify the type,
	// i.e. array of pointers to basiccrud.Fullname
	Fullname: []*basiccrud.Fullname{
		// each value should be a memory location
		// there & in front of the type
		&basiccrud.Fullname{
			Id:         1,
			FirstName:  "Chiaki",
			MiddleName: "",
			LastName:   "Kodani",
		},
		&basiccrud.Fullname{
			Id:         2,
			FirstName:  "Atreya",
			MiddleName: "Ellupai",
			LastName:   "Sridhar",
		},
	},
}

var grpcAddrFlag = flag.String("addr", ":6000", "Address host:port")

func main() {
	log.Printf("grpc server start on port %v", *grpcAddrFlag)
	// Step 1. listen for connections on tcp
	lis, err := net.Listen("tcp", *grpcAddrFlag)
	// Always handle errors
	if err != nil {
		log.Fatalf("frak")
	}
	// Step 2. Create a new grpc server instance
	srv := grpc.NewServer()
	// Step 3. Register the Greeter Service by passing the new server instance
	// and the server struct created above
	// The register function is present in pb.go
	basiccrud.RegisterBasicCrudServer(srv, &serv)
	// Step 4. Serve the listener created above
	srv.Serve(lis)
}

// Create The following is the implementation of the Create service
// as defined in the proto file. It can be any implemention.
// The below finds appends a new item to array based on request
// returns the id by incrementing the id + 1 of the length of slice
func (s *service) Create(ctx context.Context, req *basiccrud.CreateRequest) (*basiccrud.CreateResponse, error) {
	new := &basiccrud.Fullname{
		Id:         int32(len(s.Fullname) + 1),
		FirstName:  req.Fullname.GetFirstName(),
		MiddleName: req.Fullname.GetMiddleName(),
		LastName:   req.Fullname.GetLastName(),
	}
	log.Println(new.Id)
	s.Fullname = append(s.Fullname, new)
	res := &basiccrud.CreateResponse{
		Id: new.Id,
	}
	return res, nil
}

// Read The following is the implementation of the Read service
// as defined in the proto file. It can be any implemention.
// The below finds a fullname based on requested id and returns it
func (s *service) Read(ctx context.Context, req *basiccrud.ReadRequest) (*basiccrud.ReadResponse, error) {
	res := new(basiccrud.ReadResponse)
	for _, f := range s.Fullname {
		if f.Id == req.GetId() {
			res = &basiccrud.ReadResponse{
				Fullname: f,
			}
		}
	}
	if res.Fullname == nil {
		log.Println("not found")
	}
	return res, nil
}

// List lists all the fullnames. Request is empty.
// Below is an example implementation
func (s *service) List(ctx context.Context, req *basiccrud.ListRequest) (*basiccrud.ListResponse, error) {
	res := &basiccrud.ListResponse{
		Fullnames: s.Fullname,
	}
	return res, nil
}
