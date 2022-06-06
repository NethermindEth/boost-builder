package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

var GOBIN, _ = filepath.Abs(filepath.Join("build", "bin"))

func executablePath(name string) string {
	if runtime.GOOS == "windows" {
		name += ".exe"
	}
	return filepath.Join(GOBIN, name)
}

// MustRun executes the given command and exits the host process for
// any error.
func MustRun(cmd *exec.Cmd) {
	fmt.Println(">>>", strings.Join(cmd.Args, " "))
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

// buildFlags returns the go tool flags for building.
func buildFlags() (flags []string) {
	var ld []string
	// Strip DWARF on darwin. This used to be required for certain things,
	// and there is no downside to this, so we just keep doing it.
	if runtime.GOOS == "darwin" {
		ld = append(ld, "-s")
	}
	// Enforce the stacksize to 8M, which is the case on most platforms apart from
	// alpine Linux.
	if runtime.GOOS == "linux" {
		ld = append(ld, "-extldflags", "-Wl,-z,stack-size=0x800000")
	}
	if len(ld) > 0 {
		flags = append(flags, "-ldflags", strings.Join(ld, " "))
	}
	return flags
}

func buildEnv() (env []string) {
	skip := map[string]struct{}{"GOROOT": {}, "GOBIN": {}}
	for _, e := range os.Environ() {
		if i := strings.IndexByte(e, '='); i >= 0 {
			if _, ok := skip[e[:i]]; ok {
				continue
			}
		}
		env = append(env, e)
	}
	return env
}

func main() {
	log.SetFlags(log.Lshortfile)

	if len(os.Args) < 2 {
		log.Fatal("need subcommand as first argument")
	}
	switch os.Args[1] {
	case "install":
		doInstall()
	case "test":
		doTest()
	default:
		log.Fatal("unknown command ", os.Args[1])
	}
}

func doInstall() {
	root := runtime.GOROOT()
	cmd := exec.Command(filepath.Join(root, "bin", "go"), "build")
	cmd.Args = append(cmd.Args, buildFlags()...)
	cmd.Env = append(cmd.Env, "GOROOT="+root)
	cmd.Env = append(cmd.Env, buildEnv()...)

	// We use -trimpath to avoid leaking local paths into the built executables.
	cmd.Args = append(cmd.Args, "-trimpath")

	// Show packages during build.
	cmd.Args = append(cmd.Args, "-v")

	cmd.Args = append(cmd.Args, "-o", executablePath(path.Base("builder")))
	cmd.Args = append(cmd.Args, "./cmd")
	MustRun(cmd)
}

// Running The Tests
func doTest() {
	root := runtime.GOROOT()
	cmd := exec.Command(filepath.Join(root, "bin", "go"), "test")

	// Test a single package at a time.
	cmd.Args = append(cmd.Args, "-p", "1")

	packages := []string{"./..."}
	cmd.Args = append(cmd.Args, packages...)
	MustRun(cmd)
}
