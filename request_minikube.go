package main

import(
	"os"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	// namespcaeの数分、channelを作成する

	namespaces := []string{"hoge", "fuga"}
	minikube_ip := os.Getenv("minikube_ip")
	connectionMaps := make(map[string](chan string))

	for _, ns := range namespaces {
		connectionMaps[ns] = make(chan string)
	}

	for ns, channel := range connectionMaps {
			// nodePort, _ := exec.Command(
			// 	"kubectl", "get", "svc", "-n", ns,
			// 	"--output=jsonpath='{.items[0].spec.ports[0].nodePort}'").Output()
			// port, _ := strconv.Atoi(strings.Trim(string(nodePort), "'"))
			// addr := fmt.Sprintf("%s:%d", minikube_ip, port)
			// out, _ := exec.Command("curl", addr).Output()
			// fmt.Print(string(out))


		go func(ch chan string, ns string) {
			nodePort, _ := exec.Command(
				"kubectl", "get", "svc", "-n", ns,
				"--output=jsonpath='{.items[0].spec.ports[0].nodePort}'").Output()
			port, _ := strconv.Atoi(strings.Trim(string(nodePort), "'"))
			addr := fmt.Sprintf("%s:%d", minikube_ip, port)
			out, _ := exec.Command("curl", addr).Output()
			ch <- string(out)
		}(channel, ns)

		go func(ch chan string) {
			fmt.Println(<-ch)
		}(channel)
	}
	// minikube_ip := os.Getenv("minikube_ip")
	// nodePort, _ := exec.Command(
	// 	"kubectl", "get", "svc", "-n", "micro-test-ns",
	// 	"--output=jsonpath='{.items[0].spec.ports[0].nodePort}'",
	// 	).Output()

	// port, _ := strconv.Atoi(strings.Trim(string(nodePort), "'"))
	// addr := fmt.Sprintf("%s:%d", minikube_ip, port)
	
	// out, _ := exec.Command("curl", addr).Output()
	// fmt.Print(string(out))
}