# Kisumu Wrapping Paper Calculator

This Go program calculates the total amount of wrapping paper needed for a list of presents based on their dimensions. The dimensions are provided in a file named `dimensions.txt`. Each line in the file should contain the dimensions of a present in the format `LxWxH`, where `L`, `W`, and `H` are the length, width, and height of the present, respectively.

The program reads the file, parses the dimensions, and calculates the total surface area required for wrapping paper, including extra paper for the smallest side of each present.

## Features

- **Reads Input from a File**: Reads a list of dimensions from a file named `dimensions.txt`.
- **Calculates Wrapping Paper Surface Area**: Calculates the total surface area required for each present, including an extra area for the smallest side.
- **Error Handling**: Handles invalid dimension formats and conversion errors gracefully.

## How to Use

### Prerequisites

- Go 1.16 or later installed on your machine.
- A file named `dimensions.txt` with present dimensions in the format `LxWxH`, one per line.

### Installation

1. **Clone the Repository**: If you haven't already, clone this repository to your local machine:

   ```bash
   git clone https://github.com/yourusername/kisumu-wrapping-paper-calculator.git
   cd kisumu-wrapping-paper-calculator
