package main

import (
	"fmt"
)

func main() {
	// Create a 2000-qubit quantum computer
	numQubits := 2000
	repl := NewREPL(numQubits)

	fmt.Printf("QMachine Quantum Computer Simulator\n")
	fmt.Printf("Number of qubits: %d\n", numQubits)
	fmt.Printf("Quantum Volume: %d\n", repl.machine.GetQuantumVolume())
	fmt.Printf("Fidelity: 100%%\n\n")

	// Start the REPL
	repl.Start()
} 