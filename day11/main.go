package main

import (
	"AoC2025"
	"fmt"
	"slices"
	"strings"
)

type Server struct {
	name        string
	connections []string
}

type Key struct {
	name string
	dac  bool
	fft  bool
}

func main() {
	inputAsSlice := AoC2025.ReadFile("./day11/input.txt")
	var servers []Server

	for _, input := range inputAsSlice {
		input := strings.Split(input, " ")
		name, _ := strings.CutSuffix(input[0], ":")
		servers = append(servers, Server{name: name, connections: input[1:]})
	}

	// part1
	//startIndex := slices.IndexFunc(servers, func(a Server) bool {
	//	return a.name == "you"
	//})
	//count := serverSearch(&servers[startIndex], servers, 0)

	//fmt.Println("All outs:", count)

	// part2
	startIndex := slices.IndexFunc(servers, func(a Server) bool {
		return a.name == "svr"
	})

	passDac := false
	passFft := false
	cache := map[Key]int{}
	count := serverSearch2(&servers[startIndex], servers, passDac, passFft, cache)

	fmt.Println("All svr via dac & fft:", count)

}

func serverSearch(server *Server, servers []Server, count int) int {
	if server.connections[0] == "out" {
		count++
		return count
	}

	for _, connection := range server.connections {
		nextIndex := slices.IndexFunc(servers, func(a Server) bool {
			return a.name == connection
		})
		count = serverSearch(&servers[nextIndex], servers, count)
	}

	return count

}

func serverSearch2(server *Server, servers []Server, passDac bool, passFft bool, cache map[Key]int) int {

	key := Key{
		name: server.name,
		dac:  passDac,
		fft:  passFft,
	}

	if v, ok := cache[key]; ok {
		return v
	}

	count := 0

	if server.connections[0] == "out" && passFft == true && passDac == true {
		count = 1
		cache[key] = count
		return count
	} else if server.connections[0] == "out" {
		cache[key] = count
		return count
	}

	for _, connection := range server.connections {
		nextIndex := slices.IndexFunc(servers, func(a Server) bool {
			return a.name == connection
		})
		if server.name == "dac" {
			passDac = true
		}
		if server.name == "fft" {
			passFft = true
		}
		count += serverSearch2(&servers[nextIndex], servers, passDac, passFft, cache)
	}

	cache[key] = count

	return count

}
