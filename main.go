package main

import (
	"flag"
	"fmt"
	"os"

	"qmachine/quantum"
)

func main() {
	// Define command-line flags
	numQubits := flag.Int("qubits", 2000, "Number of qubits for the quantum computer")
	quantumFile := flag.String("quantum", "", "Path to quantum RISC-V file to execute")
	hostQuantumFile := flag.String("host-quantum", "", "Path to quantum RISC-V file to execute on host")
	flag.Parse()

	// Create the quantum computer
	repl := NewREPL(*numQubits)

	if *hostQuantumFile != "" {
		// Execute quantum RISC-V file on host
		fmt.Printf("Executing quantum RISC-V file on host: %s\n", *hostQuantumFile)

		// Create host-optimized quantum machine
		hostMachine := quantum.NewHostQuantumMachine(*numQubits)

		// Load and parse the program
		if err := repl.machine.LoadRISCProgram(*hostQuantumFile); err != nil {
			fmt.Printf("Error loading quantum RISC-V program: %v\n", err)
			os.Exit(1)
		}

		// Execute quantum instructions on host
		for _, inst := range repl.machine.GetRISCProgram() {
			if isQuantumInstruction(inst.Opcode) {
				if err := hostMachine.ExecuteQuantumRISCV(inst); err != nil {
					fmt.Printf("Error executing quantum instruction: %v\n", err)
					os.Exit(1)
				}
			} else {
				// Execute classical RISC-V instructions using the standard machine
				instruction := fmt.Sprintf("%s x%d, x%d, x%d", inst.Opcode, inst.Rd, inst.Rs1, inst.Rs2)
				if err := repl.machine.ExecuteRISCInstruction(instruction); err != nil {
					fmt.Printf("Error executing RISC-V instruction: %v\n", err)
					os.Exit(1)
				}
			}
		}

		// Display final register state
		registers := hostMachine.GetRegisters()
		fmt.Println("\nFinal register state:")
		for i, reg := range registers {
			fmt.Printf("  x%d: %d\n", i, reg)
		}
		os.Exit(0)
	}

	if *quantumFile != "" {
		// Execute quantum RISC-V file in VM mode
		fmt.Printf("Executing quantum RISC-V file in VM mode: %s\n", *quantumFile)
		if err := repl.machine.LoadRISCProgram(*quantumFile); err != nil {
			fmt.Printf("Error loading quantum RISC-V program: %v\n", err)
			os.Exit(1)
		}
		if err := repl.machine.ExecuteRISCProgram(); err != nil {
			fmt.Printf("Error executing quantum RISC-V program: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Quantum RISC-V program executed successfully")
		os.Exit(0)
	}

	// Interactive REPL mode
	fmt.Printf("QMachine Quantum Computer Simulator\n")
	fmt.Printf("Number of qubits: %d\n", *numQubits)
	fmt.Printf("Quantum Volume: %d\n", repl.machine.GetQuantumVolume())
	fmt.Printf("Fidelity: 100%%\n\n")

	// Start the REPL
	repl.Start()
}

// isQuantumInstruction checks if an instruction is a quantum instruction
func isQuantumInstruction(opcode string) bool {
	switch opcode {
	case "qinit", "qapply", "qmeasure", "qentangle":
		return true
	default:
		return false
	}
}
