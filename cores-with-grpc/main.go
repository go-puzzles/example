package main

import (
	"github.com/go-puzzles/cores"
	"github.com/go-puzzles/example/cores-with-grpc/examplepb"
	srv "github.com/go-puzzles/example/cores-with-grpc/service"
	"github.com/go-puzzles/example/cores-with-grpc/testpb"
	grpcpuzzle "github.com/go-puzzles/grpc-puzzle"
	grpcuipuzzle "github.com/go-puzzles/grpcui-puzzle"
	"github.com/go-puzzles/plog"
	"github.com/go-puzzles/plog/level"
	"google.golang.org/grpc"
)

func main() {
	plog.Enable(level.LevelDebug)

	example := srv.NewExampleService()
	test := srv.NewTestService()

	srv := cores.NewPuzzleCore(
		grpcuipuzzle.WithCoreGrpcUI(),
		grpcpuzzle.WithCoreGrpcPuzzle(func(srv *grpc.Server) {
			examplepb.RegisterExampleHelloServiceServer(srv, example)
			testpb.RegisterExampleHelloServiceServer(srv, test)
		}),
	)

	if err := cores.Start(srv, 0); err != nil {
		panic(err)
	}
}
