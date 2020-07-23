# Gotrav
Gotrav is a tool written in Go that can help determine whether a web application is vulnerable to a directory traversal attack.

### Getting Started
#### Clone the repository
``` Bash
git clone https://github.com/umkhan23/Gotrav.git
cd Gotrav/src/gotrav
```
### Build the binary
`go build gotrav.go`

### Usage

```
                          .-') _   _  .-')     ('-.          (`-.   
                         (  OO) ) ( \\( -O )   ( OO ).-.    _(OO  )_ 
  ,----.     .-'),-----. /     '._ ,------.   / . --. /,--(_/   ,. \\
 '  .-./-') ( OO'  .-.  '|'--...__)|   /'. '  | \\-.  \\ \\   \\   /(__/
 |  |_( O- )/   |  | |  |'--.  .--'|  /  | |.-'-'  |  | \\   \\ /   / 
 |  | .--, \\\\_) |  |\\|  |   |  |   |  |_.' | \\| |_.'  |  \\   '   /, 
(|  | '. (_/  \\ | | |  |  |   |   |  .  '.'  |  .-.  |   \\     /__)
 |  '--'  |    ''  '-'  '  |   |   |  |\  \   |  | |  |    \   /    
  '------'       '-----'    '--'   '--' '--'  '--' '--'     '-'     
==========================================================================
                             Author: Usman Q Khan

Gotrav is a tool written in Go that can help determine whether a web application is vulnerable to a directory traversal attack

Usage: ./gotrav -u <url> -p <path> -o <outfile>

Example: ./gotrav -u http://10.10.10.10/ -p /etc/passwd -o users.txt

Flags:
-u
  Enter a URL
-p
  Enter the full path of the file you would like to find
-o
  Enter name of outfile

```
### Future plans
Currently, gotrav only features one directory traversal payload to test out. I plan to include multiple common payloads in the future to make gotrav a lot more versatile and effective than it is right now.
