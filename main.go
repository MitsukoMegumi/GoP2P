package main

import (
	"flag"
	"net/http"

	"github.com/mitsukomegumi/GoP2P/cli"
	"github.com/mitsukomegumi/GoP2P/common"
	node "github.com/mitsukomegumi/GoP2P/rpc/node"
	proto "github.com/mitsukomegumi/GoP2P/rpc/proto"
	"github.com/mitsukomegumi/GoP2P/upnp"
)

var (
	terminalFlag = flag.Bool("terminal", false, "launch GoP2P in terminal mode")                   // Init term flag
	upnpFlag     = flag.Bool("upnp", false, "launch GoP2P without automatic UPnP port forwarding") // Init upnp flag
)

func main() {
	flag.Parse() // Parse flags

	if !*upnpFlag {
		go upnp.ForwardPortSilent(8080) // Forward port 8080
	}

	startRPCServer() // Start RPC server

	if *terminalFlag {
		cli.NewTerminal() // Initialize terminal
	}

	startNode() // Attempt to start GoP2P in node mode

	go common.Forever() // Prevent main from closing
	select {}           // Prevent main from closing
}

// startRPCServer - start RPC server
func startRPCServer() {
	nodeHandler := proto.NewNodeServer(&node.Server{}, nil) // Init handler

	mux := http.NewServeMux() // Init mux

	mux.Handle(proto.NodePathPrefix, nodeHandler) // Start mux node handler

	go http.ListenAndServe(":8080", mux) // Start server
}

// startNode - attempt to execute attachnode, starthandler commands
func startNode() {
	terminal := cli.Terminal{Variables: []cli.Variable{}} // Init terminal

	terminal.HandleCommand("node.Attach()") // Attach node

	terminal.HandleCommand("node.StartHandler()") // Start handler
}

/* TODO:
- Add custom RPC addresses
- On attachment, make sure to write changes to a discovered node's db.
*/
