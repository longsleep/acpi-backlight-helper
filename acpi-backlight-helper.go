package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
)

const backlightPath = "/sys/class/backlight/intel_backlight"

func join(n string) string {
	return path.Join(backlightPath, n)
}

func backlightHelper(args []string) error {
	switch len(args) {
	default:
		return fmt.Errorf("no arguments provided")
	case 1:
	case 2:
	}

	switch args[0] {
	case "-h":
		fallthrough
	case "--help":
		usage()
	case "--get-max-brightness":
		maxBrightness, err := getMaxBrightness()
		if err != nil {
			return err
		}
		fmt.Fprintf(os.Stdout, "%d\n", maxBrightness)
	case "--get-brightness":
		maxBrightness, err := getBrightness()
		if err != nil {
			return err
		}
		fmt.Fprintf(os.Stdout, "%d\n", maxBrightness)
	case "--set-brightness":
		if os.Getuid() != 0 || os.Geteuid() != 0 {
			fmt.Fprintf(os.Stderr, "This program can only be used by the root user\n")
			os.Exit(4)
		}

		if len(args) != 2 {
			return fmt.Errorf("missing brightness value")
		}
		value, err := strconv.Atoi(args[1])
		if err != nil {
			return err
		}
		err = setBrightness(value)
		return err
	}

	return nil
}

func usage() {
	fmt.Fprintf(os.Stdout, `Usage:
  %s [OPTION...]

ACPI Backlight Helper (%s)

Help Options:
  -h, --help                Show help options

Application Options:
  --set-brightness          Set the current brightness
  --get-brightness          Get the current brightness
  --get-max-brightness      Get the number of brighness levels supported

`,
		os.Args[0], backlightPath)
}

func readInt(n string) (int, error) {
	raw, err := ioutil.ReadFile(join(n))
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(strings.TrimSpace(string(raw)))
}

func writeInt(n string, value int) error {
	raw := []byte(fmt.Sprintf("%d\n", value))
	return ioutil.WriteFile(join(n), raw, 0644)
}

func getMaxBrightness() (int, error) {
	return readInt("max_brightness")
}

func getBrightness() (int, error) {
	return readInt("brightness")
}

func setBrightness(value int) error {
	return writeInt("brightness", value)
}

func main() {
	args := os.Args[1:]

	err := backlightHelper(args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(3)
	}

	os.Exit(0)
}
