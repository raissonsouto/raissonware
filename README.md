# Raissonware

Raissonware is a proof of concept ransomware application built in Go. It demonstrates how a simple ransomware might encrypt files and transfer the report file between a victim and its attacker.

> Note: This project is purely academic and intended for educational purposes only.

## Features

- **File Encryption**: Encrypts files using AES encryption with randomly generated keys and nonce's to each file.
- **CSV Logging**: Logs encrypted file details (path, encrypted key, nonce, file mode) in a CSV file.

## Components

- **Server**: Listens for incoming connections and stores the received reports.
- **Client**: Encrypt files and sends a report to the server.
- **Unlocker**: Provides a mechanism to revert the encryption attack using the generated report and unlocker tool.

### How to run it

1. Clone the repository:
   ```bash
   git clone https://github.com/raissonsouto/raissonware.git
   cd raissonware
   ```

2. Build the components:
   ```bash
   make build
   ```

3. Run the server:
   ```bash
   ./server
   ```
   
4. Make the victim run the client:
   ```bash
   ./client
   ```

### Usage

1. **Client Usage**:

2. **Server Usage**:

3. **Ransom Negotiation**:

### Contributing

Contributions are welcome! If you find a bug or have suggestions for improvements, please open an issue or create a pull request.