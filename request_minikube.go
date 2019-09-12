package main

import(
	"os"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	minikube_ip := os.Getenv("minikube_ip")
	nodePort, _ := exec.Command(
		"kubectl", "get", "svc", "-n", "micro-test-ns",
		"--output=jsonpath='{.items[0].spec.ports[0].nodePort}'",
		).Output()

	port, _ := strconv.Atoi(strings.Trim(string(nodePort), "'"))
	addr := fmt.Sprintf("%s:%d", minikube_ip, port)
	
	out, _ := exec.Command("curl", addr).Output()
	fmt.Print(string(out))
}