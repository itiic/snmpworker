package async

import (
	"bufio"
	"log"
	"os"
	"sync"

	"github.com/itiic/snmpworker/pkg/conf"
)

// Load data to the channel
func Load(path string, ch chan string) {
	defer close(ch)

	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		ch <- scanner.Text()
	}
}

// Fun Out load across workers and put result out channel for future processing
func FanOutFanIn(inChan chan string, outChan chan string, config conf.Config) {
	defer close(outChan)
	var wg sync.WaitGroup
	wg.Add(config.Worker)
	for i := 0; i < config.Worker; i++ {
		go func() {
			for ip := range inChan {
				func(ip string, config conf.Config) {
					for _, v := range Run(ip, config) {
						outChan <- v
					}
				}(ip, config)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

// Run - return empty slice for now
func Run(ip string, config conf.Config) []string {
	return []string{}
}
