package main

import(
	"os"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"sync"
)

func main() {

	namespaces := []string{"hoge", "fuga"}
	minikube_ip := os.Getenv("minikube_ip")
	connectionMaps := make(map[string](chan string))
	wg := &sync.WaitGroup{}

	for _, ns := range namespaces {
		connectionMaps[ns] = make(chan string)
	}

	for ns, channel := range connectionMaps {
		wg.Add(2)
		go func(ch chan string, ns string) {
			nodePort, _ := exec.Command(
				"kubectl", "get", "svc", "-n", ns,
				"--output=jsonpath='{.items[0].spec.ports[0].nodePort}'").Output()
			port, _ := strconv.Atoi(strings.Trim(string(nodePort), "'"))
			addr := fmt.Sprintf("%s:%d", minikube_ip, port)
			out, _ := exec.Command("curl", addr).Output()
			ch <- string(out)
			wg.Done()
		}(channel, ns)

		go func(ch chan string, ns string) {
			fmt.Println(<-ch)
			wg.Done()
		}(channel, ns)
	}
	wg.Wait()
}