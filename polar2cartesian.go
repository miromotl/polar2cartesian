package main

import (
    "bufio"
    "fmt"
    "math"
    "os"
    "runtime"
)

// Strings for user interaction
const result = "Polar radius=%.02f θ=%.02f → Cartesian x=%.02f y=%.02f\n"

var prompt = "Enter a radius and an angle (in degrees), e.g., 12.5 90, " +
             "or %s to quit."

// Structs for polar and cartesian coordinates
type polar struct {
    radius float64
    θ      float64
}

type cartesian struct {
    x float64
    y float64
}

// Determine OS at startup time and adjust user prompt accordingly
func init() {
    if runtime.GOOS == "windows" {
        prompt = fmt.Sprintf(prompt, "Ctrl+Z, Enter")
    } else { // Unix-like
        prompt = fmt.Sprintf(prompt, "Ctrl+D")
    }
}

// The main function
func main() {
    // Create the input questions channel and defer its closing to the end
    questions := make(chan polar)
    defer close(questions)
    
    // Create the output answers channel and defer its closing to the end
    answers := createSolver(questions)
    defer close(answers)
    
    // Interact with the user
    interact(questions, answers)
}

// Create the solver that returns the channel with the answers
func createSolver(questions chan polar) chan cartesian {
    // Create the channel
    answers := make(chan cartesian)
    
    // Start a go routine that reads the questions and writes the answers
    // in an eternal loop
    go func() {
        for {
            // get next question from the questions channel
            polarCoord:= <- questions
            
            // solve the question
            θ := polarCoord.θ * math.Pi / 180.0 // degrees to radians
            x := polarCoord.radius * math.Cos(θ)
            y := polarCoord.radius * math.Sin(θ)
            
            // put the answer into the answers channel
            answers <- cartesian{x, y}
        }
    } ()
    
    // return the answers channel to the caller
    return answers
}

// The interaction with the user
func interact(questions chan polar, answers chan cartesian) {
    // Read the questions from stdio
    reader := bufio.NewReader(os.Stdin)
    
    // Present the prompt to the user
    fmt.Println(prompt)
    
    // As the user in an endless loop for the questions
    for {
        fmt.Printf("Radius and angle: ")
        line, err := reader.ReadString('\n')
        if err != nil {
            break
        }
        
        var radius, θ float64
        
        if _, err := fmt.Sscanf(line, "%f %f", &radius, &θ); err != nil {
            fmt.Fprintln(os.Stderr, "invalid input")
            continue
        }
        
        questions <- polar{radius, θ}
        coord := <- answers
        
        fmt.Printf(result, radius, θ, coord.x, coord.y)
    }
    
    fmt.Println()
}