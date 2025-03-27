# QMachine - Quantum Computer Simulator

A high-performance, realistic simulation of a quantum computer written in Go, featuring a RISC-V instruction set for quantum operations.

## Features

- 2000 qubit quantum computer simulation
- 100% fidelity quantum operations
- Quantum Volume of 4269
- RISC-V based instruction set
- Interactive REPL interface
- Support for common quantum gates:
  - Pauli gates (X, Y, Z)
  - Hadamard gate (H)
  - Phase gates (S, T)
  - CNOT gate
  - Measurement operations
- Full RISC-V RV32I base integer instruction set support:
  - Arithmetic operations (add, sub, and, or, xor)
  - Shift operations (sll, srl, sra)
  - Comparison operations (slt, sltu)
  - Immediate operations (addi, slli, srli, etc.)
  - Load/Store operations (lw, lh, lb, sw, sh, sb)
  - Branch operations (beq, bne, blt, bge, etc.)
  - Jump operations (jal, jalr)
  - Upper immediate operations (lui, auipc)

## Requirements

- Go 1.21 or later

## Installation

```bash
git clone https://github.com/sneed-group/qmachine.git
cd qmachine
```

## Usage

Run the REPL interface:
```bash
go run .
```

### REPL Commands

- `gate <type> <target> [controls...]` - Apply a quantum gate
- `measure <qubit>` - Measure a qubit
- `state` - Show current quantum state
- `reset` - Reset quantum state
- `riscv <instruction>` - Execute RISC-V instruction
- `load <file>` - Load RISC-V program from file
- `run` - Run loaded RISC-V program
- `registers` - Show RISC-V registers
- `help` - Show help message
- `exit` - Exit REPL

### Example RISC-V Program

Create a file `program.riscv`:
```
# Add two numbers
addi x1, x0, 42
addi x2, x0, 58
add x3, x1, x2
```

Load and run it:
```
qmachine> load program.riscv
qmachine> run
qmachine> registers
```

## Project Structure

- `quantum/state.go`: Quantum state representation and manipulation
- `quantum/gates.go`: Implementation of quantum gates
- `quantum/riscv.go`: RISC-V instruction set implementation
- `main.go`: Main program entry point
- `repl.go`: Interactive REPL implementation

## Performance

The simulator is optimized for high-performance quantum state manipulation using Go's efficient memory management and concurrent execution capabilities. The RISC-V implementation provides a standard instruction set for classical computation alongside quantum operations.

## License

SPL-R5 License

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request. 