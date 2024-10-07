# IPFAS - IP Finder as well as Server Names :)

## Description

**IPFAS** (IP Finder as well as Server Names) is a simple command-line tool written in Go that allows you to fetch IP addresses and name servers related to a domain. You can easily look up this information using the provided commands.

## Features

- Fetch IP addresses for a given domain.
- Retrieve name servers (NS) associated with a domain.

## Installation

### Prerequisites

Ensure you have Go installed in your environment. You can download and install Go from the official site: [https://golang.org/dl/](https://golang.org/dl/).

1. Clone the repository:

   ```bash
   git clone https://github.com/urxfa/ipfas.git
   ```

2. Navigate to the project directory:

   ```bash
   cd ipfas
   ```

3. Install the dependencies:

   ```bash
   go get -u github.com/urfave/cli
   ```

4. Build the application:

   ```bash
   go build
   ```

## Usage

Once you've installed the dependencies and built the project, you can start using the tool.

### Command Syntax

```
./ipfas <command> [options]
```

### Available Commands

- `ip`: Find the IP addresses associated with a domain.
- `nameservers`: Retrieve the name servers associated with a domain.

### Example Usages

1. **Find IP addresses for a domain**:

   ```bash
   ./ipfas ip -host www.google.com
   ```

   Output:

   ```
   List of Ips for this Domain:
   172.217.0.0
   142.250.0.0
   ```

2. **Find name servers for a domain**:

   ```bash
   ./ipfas nameservers -host amazon.com.br
   ```

   Output:

   ```
   List of NS for this Domain:
   ns1.amazon.com
   ns2.amazon.com
   ```

### Flags

- `-host`: Specify the domain name for which you want to find the IP addresses or name servers. Default is \`amazon.com.br\` if no domain is provided.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

## Contributions

Feel free to submit issues or pull requests. Any contribution to improve this tool is welcome!
