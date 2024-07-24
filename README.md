# 🛡️ File Integrity Checker

A simple tool to monitor and verify the integrity of files in a directory using SHA-256 hashes.

## 📜 Description

This project helps you ensure that the files in a specified directory have not been tampered with. It calculates the SHA-256 hash of each file and stores it in a JSON database. On subsequent runs, it compares the current hashes with the stored ones and logs any discrepancies.

## ✨ Features

- 🔍 **File Integrity Check**: Detects changes in file content.
- 🆕 **New File Detection**: Logs new files added to the directory.
- 🗑️ **File Deletion Detection**: Logs files that have been deleted.
- 📄 **JSON Database**: Stores file hashes in a JSON file.
- 📋 **Verbose Logging**: Option to enable verbose logging for detailed output.

## 🛠️ Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/yourusername/file-integrity-checker.git
   cd file-integrity-checker
   ```

2. Build the project:
   ```sh
   go build -o file-integrity-checker
   ```

## 🚀 Usage

1. Run the script:
   ```sh
   ./file-integrity-checker -v
   ```

2. Check the log file for results:
   ```sh
   cat integrity_check.log
   ```

### Log Levels

- `INFO`: General information about the integrity check.
- `WARN`: New files detected.
- `ALERT`: File integrity check failed or file deleted.
- `ERROR`: Errors encountered during the process.

## 🤝 Contributing

Contributions are welcome! Please fork the repository and submit a pull request.

## 📄 License

This project is licensed under the MIT License.
