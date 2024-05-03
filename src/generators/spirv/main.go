package main

import (
	"bufio"
	"flag"
	"kaiju/klib"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func compile(args ...string) {
	cmd := exec.Command("glslc", args...)
	outPipe := klib.MustReturn(cmd.StdoutPipe())
	scanner := bufio.NewScanner(outPipe)
	err := cmd.Start()
	if err != nil {
		vp := os.Getenv("VK_SDK_PATH")
		if vp != "" {
			cmd = exec.Command(filepath.Join(vp, "Bin", "glslc"), args...)
			outPipe = klib.MustReturn(cmd.StdoutPipe())
			scanner = bufio.NewScanner(outPipe)
			err = cmd.Start()
		}
		if err != nil {
			panic("Failed to run glslc, make sure you have the Vulkan 'Bin' folder in your environment path")
		}
	}
	for scanner.Scan() {
		println(scanner.Text())
	}
	klib.Must(cmd.Wait())
}

func hasOIT(path string) bool {
	data, err := os.ReadFile(path)
	if err != nil {
		return false
	}
	return strings.Contains(string(data), "#ifdef OIT")
}

func main() {
	fs := flag.NewFlagSet("Kaiju Spir-V compile", flag.ContinueOnError)
	dbg := fs.Bool("d", false, "Compile the shader for debugging")
	out := fs.String("o", "", "The output path for the compiled shader")
	in := fs.String("i", "", "The path of the shader to be compiled")
	fs.Parse(os.Args[1:])
	if *in == "" {
		panic("Expected -i=... input, run with arg -h for help")
	}
	outName := *out
	if outName == "" {
		outName = filepath.Dir(*in)
	}
	if !strings.HasSuffix(*out, ".spv") {
		outName = filepath.Join(*out, filepath.Base(*in)+".spv")
	}
	args := []string{*in,
		"-o", outName,
	}
	if *dbg {
		args = append(args, "-g")
	}
	compile(args...)
	if hasOIT(*in) {
		args[2] = strings.TrimSuffix(args[2], ".spv") + ".oit.spv"
		args = append(args, "-DOIT")
		compile(args...)
	}
}
