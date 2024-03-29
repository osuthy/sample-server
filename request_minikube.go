package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"sync"
)

func main() {

	namespaces := []string{"hoge", "fuga"}
	minikube_ip_out, _ := exec.Command("minikube", "ip").Output()
	minikube_ip := strings.TrimRight(string(minikube_ip_out), "\n")
	wg := &sync.WaitGroup{}

	for _, ns := range namespaces {
		wg.Add(1)
		go func(ns string) {
			nodePort, _ := exec.Command(
				"kubectl", "get", "svc", "-n", ns,
				"--output=jsonpath='{.items[0].spec.ports[0].nodePort}'").Output()
			port, _ := strconv.Atoi(strings.Trim(string(nodePort), "'"))
			addr := fmt.Sprintf("%s:%d", minikube_ip, port)
			out, _ := exec.Command("curl", addr).Output()
			fmt.Print(string(out))
			wg.Done()
		}(ns)
		wg.Wait()
	}
}
