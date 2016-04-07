package deputil

import (
	"os/exec"
	"runtime"
	"strings"
	"sync"
)

// Dependency checker object.
type Dependency struct {
	m map[string]string
}

// New dependency creates a new Dependency object.
func New() *Dependency {
	return &Dependency{
		m: make(map[string]string),
	}
}

// Add binary b to the dependency list.
func (d *Dependency) Add(b string) *Dependency {
	return d.AddWithName(b, "")
}

// AddWithName adds the binary b contained in package p to the list.
func (d *Dependency) AddWithName(b, p string) *Dependency {
	if len(strings.TrimSpace(b)) == 0 {
		panic("E: Mandatory binary-name is empty!")
	}
	d.m[b] = p
	return d
}

// Check at runtime if dependencies are met. If not, the returned
// string slice lists the missing package/binary names.
func (d *Dependency) Check() []string {
	errc := make(chan string, poolSize(len(d.m)))
	var wg sync.WaitGroup
	wg.Add(len(d.m))
	for b, p := range d.m {
		go func(bin, pkg string) {
			defer wg.Done()
			checkDependency(bin, pkg, errc)
		}(b, p)
	}
	go func() {
		wg.Wait()
		close(errc)
	}()
	var a []string
	for s := range errc {
		a = append(a, s)
	}
	return a
}

// poolSize evaluates the goroutine-pool-size.
func poolSize(dependencyCount int) int {
	var poolSize = runtime.NumCPU()
	if poolSize == 1 {
		poolSize = 2
	}
	if dependencyCount < poolSize {
		poolSize = dependencyCount
	}
	return poolSize
}

// checkDependency on the system. If dependency is not met,
// an error is send through the error channel.
func checkDependency(bin, pkg string, errc chan<- string) {
	if err := exec.Command("which", bin).Run(); err == nil {
		return
	}
	if len(pkg) == 0 {
		errc <- bin
	} else {
		errc <- pkg
	}
}
