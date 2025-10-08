// cmd/bridgeprotocol/main.go
package main

import (
    "flag"
    "log"
    "os"

    "bridgeprotocol/internal/bridgeprotocol"
)

func main() {
    // Define command line flags with descriptive names and usage
    verbose := flag.Bool("verbose", false, "Enable verbose logging")
    input := flag.String("input", "", "Path to the input file")
    output := flag.String("output", "", "Path to the output file")
    flag.Parse()

    // Create a new app instance with the verbosity level
    app := bridgeprotocol.NewApp(*verbose)

    // Run the app with input and output paths, handle any errors
    if err := app.Run(*input, *output); err != nil {
        log.Printf("Error: %v", err)
        os.Exit(1)
    }
}