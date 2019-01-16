package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"strings"
	"time"

	// The Protobuf generated file
	creator "app/codenamecreator"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

type codenameGenerator struct {
	Adverbs    []string
	Animals    []string
	Scientists []string
}

func newCodenameGenerator() codenameGenerator {
	cg := codenameGenerator{}
	cg.Adverbs = []string{"Anxious", "Artistic", "Bold", "Cheerful", "Curious", "Daring", "Fearless", "Gallant", "Heroic", "Languid", "Lucid", "Mighty", "Nefarious", "Quizzical", "Sleepy", "Tireless", "Vigorous", "Wicked"}
	cg.Animals = []string{"Aardvark", "Badger", "Coyote", "Dolphin", "Fox", "Giraffe", "Heron", "Lizard", "Marmot", "Nighthawk", "Quail", "Shark", "Tiger", "Vulture", "Warthog"}
	cg.Scientists = []string{"Curie", "Dalton", "Davy", "Faraday", "Franklin", "Germain", "Hodgkin", "Hopper", "Lovelace", "Meitner", "Newton", "Salk", "Tesla", "Youyou"}
	return cg
}

func (cg *codenameGenerator) generate(category string) string {
	adverbNumber := rand.Intn(len(cg.Adverbs))

	if strings.ToLower(category) == "science" {
		scientistNumber := rand.Intn(len(cg.Scientists))
		return fmt.Sprintf("%s %s", cg.Adverbs[adverbNumber], cg.Scientists[scientistNumber])
	} else {
		animalNumber := rand.Intn(len(cg.Animals))
		return fmt.Sprintf("%s %s", cg.Adverbs[adverbNumber], cg.Animals[animalNumber])
	}
}

type codenameServer struct{}

func (s *codenameServer) GetCodename(ctx context.Context, request *creator.NameRequest) (*creator.NameResult, error) {
	generator := newCodenameGenerator()
	codename := generator.generate(request.Category)
	return &creator.NameResult{Name: codename}, nil
}

func (s *codenameServer) KeepGettingCodenames(stream creator.CodenameCreator_KeepGettingCodenamesServer) error {
	// get some metadata
	ctx := stream.Context()
	md, ok := metadata.FromIncomingContext(ctx)

	if ok {
		log.Printf("Metadata: %v\n", md)
	}

	// generate a new codename
	generator := newCodenameGenerator()
	categoryChan := make(chan string)

	go func(c chan string) {
		for {
			log.Println("Receiving")
			in, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("%v", err)
			}

			log.Printf("---Updating codename category to: %s---\n", in.Category)
			c <- in.Category
		}
	}(categoryChan)

	var codename string
	var category string

	for {
		select {
		case category = <-categoryChan:
			codename = generator.generate(category)
		default:
			codename = generator.generate(category)
		}

		result := &creator.NameResult{Name: codename}
		err := stream.Send(result)
		if err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}

	return nil
}

func main() {
	address := ":3000"
	crt := "server.crt"
	key := "server.key"

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	creds, err := credentials.NewServerTLSFromFile(crt, key)
	if err != nil {
		log.Fatalf("Failed to load TLS keys")
	}

	grpcServer := grpc.NewServer(grpc.Creds(creds))
	creator.RegisterCodenameCreatorServer(grpcServer, &codenameServer{})

	log.Println("Listening on address ", address)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
