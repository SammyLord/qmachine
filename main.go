package main

import (
	"flag"
	"fmt"
	"os"

	"qmachine/quantum"
	"qmachine/repl"
)

func main() {
	// Define command-line flags
	numQubits := flag.Int("qubits", 2000, "Number of qubits for the quantum computer")
	quantumFile := flag.String("quantum", "", "Path to quantum RISC-V file to execute")
	hostQuantumFile := flag.String("host-quantum", "", "Path to quantum RISC-V file to execute on host")
	flag.Parse()

	// Create the quantum computer REPL
	replInstance := repl.New(*numQubits)

	// Handle file execution modes
	if *hostQuantumFile != "" {
		fmt.Printf("Executing quantum RISC-V file on host: %s\n", *hostQuantumFile)
		if err := executeHostQuantumFile(*hostQuantumFile, *numQubits); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Quantum RISC-V program executed successfully using host-native execution")
		os.Exit(0)
	}

	if *quantumFile != "" {
		fmt.Printf("Executing quantum RISC-V file in VM mode: %s\n", *quantumFile)
		machine := quantum.NewQuantumRISCVMachine(*numQubits)

		// Load and execute the program
		if err := machine.LoadRISCProgram(*quantumFile); err != nil {
			fmt.Printf("Error loading quantum RISC-V program: %v\n", err)
			os.Exit(1)
		}

		// Print initial state
		fmt.Printf("\nInitial register state:\n")
		printRegisters(machine.GetRegisters())

		// Execute the program
		if err := machine.ExecuteRISCProgram(); err != nil {
			fmt.Printf("Error executing quantum RISC-V program: %v\n", err)
			os.Exit(1)
		}

		// Print final state
		fmt.Printf("\nFinal register state:\n")
		printRegisters(machine.GetRegisters())

		fmt.Println("\nQuantum RISC-V program executed successfully")
		os.Exit(0)
	}

	// Start interactive REPL mode
	replInstance.Start()
}

// executeHostQuantumFile executes a quantum RISC-V file using host-native execution
func executeHostQuantumFile(filename string, numQubits int) error {
	// Create a VM just to parse the program
	machine := quantum.NewQuantumRISCVMachine(numQubits)
	if err := machine.LoadRISCProgram(filename); err != nil {
		return fmt.Errorf("error loading quantum RISC-V program: %v", err)
	}

	// Create host machine for native execution
	hostMachine := quantum.NewHostQuantumMachine(numQubits)

	// Execute each instruction directly on the host
	for _, inst := range machine.GetRISCProgram() {
		if isQuantumInstruction(inst.Opcode) {
			// Execute quantum instructions using host-native execution
			if err := hostMachine.ExecuteQuantumRISCV(inst); err != nil {
				return fmt.Errorf("error executing quantum instruction on host: %v", err)
			}
		} else {
			// Execute classical RISC-V instructions
			var instruction string
			switch inst.Opcode {
			case "add", "sub", "and", "or", "xor", "sll", "srl", "sra", "slt", "sltu":
				instruction = fmt.Sprintf("%s x%d, x%d, x%d", inst.Opcode, inst.Rd, inst.Rs1, inst.Rs2)
			case "addi", "slli", "srli", "srai", "andi", "ori", "xori", "slti", "sltiu":
				instruction = fmt.Sprintf("%s x%d, x%d, %d", inst.Opcode, inst.Rd, inst.Rs1, inst.Imm)
			case "lui", "auipc":
				instruction = fmt.Sprintf("%s x%d, %d", inst.Opcode, inst.Rd, inst.Imm)
			case "jal":
				instruction = fmt.Sprintf("%s x%d, %d", inst.Opcode, inst.Rd, inst.Offset)
			case "jalr":
				instruction = fmt.Sprintf("%s x%d, x%d, %d", inst.Opcode, inst.Rd, inst.Rs1, inst.Offset)
			case "beq", "bne", "blt", "bge", "bltu", "bgeu":
				instruction = fmt.Sprintf("%s x%d, x%d, %d", inst.Opcode, inst.Rs1, inst.Rs2, inst.Offset)
			case "lw", "lh", "lb", "lwu", "lhu", "lbu":
				instruction = fmt.Sprintf("%s x%d, %d(x%d)", inst.Opcode, inst.Rd, inst.Offset, inst.Rs1)
			case "sw", "sh", "sb":
				instruction = fmt.Sprintf("%s x%d, %d(x%d)", inst.Opcode, inst.Rs2, inst.Offset, inst.Rs1)
			default:
				return fmt.Errorf("unknown instruction type: %s", inst.Opcode)
			}
			if err := machine.ExecuteRISCInstruction(instruction); err != nil {
				return fmt.Errorf("error executing RISC-V instruction on host: %v", err)
			}
		}
	}

	return nil
}

// printRegisters prints the current state of the registers
func printRegisters(registers [128]uint64) {
	for i, reg := range registers {
		if reg != 0 { // Only print non-zero registers to reduce noise
			fmt.Printf("  x%d: %d\n", i, reg)
		}
	}
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
