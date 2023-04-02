# Todo.md to Jira CSV Importer

A simple CLI tool to convert your Todo.md file into a CSV format that can be imported into Jira, written in Golang.

## Features

- Convert your Todo.md tasks into Jira-compatible CSV format
- Maintain task hierarchy and descriptions
- Include task priorities and tags as custom fields

## Requirements

- Golang 1.15 or later

## Installation

1. Clone this repository or download the source code:

```
git clone https://github.com/yourusername/todo-md-to-jira-csv.git
```

2. Change to the project directory:

```
cd todo-md-to-jira-csv
```

3. Build the executable:

```
go build -o todo_md_to_jira_csv
```

## Usage

1. Create a Todo.md file or use an existing one.

2. Run the following command to convert the Todo.md file into a Jira-compatible CSV file:

```
./todo_md_to_jira_csv --input todo.md --output jira_import.csv
```

Replace `todo.md` with the path to your Todo.md file and `jira_import.csv` with the desired output file name.

3. Import the generated CSV file into Jira using the External System Import feature. Make sure to map the fields correctly during the import process.

## Support and Contributions

If you encounter any issues or have suggestions for improvements, please open an issue on the GitHub repository. Contributions are welcome! Feel free to submit a pull request with your changes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Inspired by the simplicity of Todo.md format and the need to integrate it with Jira for better project management.