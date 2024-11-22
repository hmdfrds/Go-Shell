# Go Shell
![Go Version](https://img.shields.io/github/go-mod/go-version/hmdfrds/go-shell)
![License](https://img.shields.io/github/license/hmdfrds/go-shell)  

Just a super basic **command-line shell** written in Go. Cmd/shell wannabe.  

## Description  
I Just started learning GO. So I decide to GO to build something with GO.

## Features 

- [x] **Command Prompt**  
  Type stuff, hit Enter, and get responses.

- [x] **`cd` (Change Directory)**  
  Type `cd <directory name>` to change the directory.

- [x] **`pwd` (Print Working Directory)**  
  Type `pwd` will show the working directory.

- [x] **`exit`**  
   Type `exit` to leave.  

- [x]  **`ls`**  
  Type `ls [directory name]` will how directory list in a directory.

- [x] **Tab Auto Completion**  
  Click tab will auto complete the directory name that exist in **working directory**.

- [x] **Run External Commands**  
  Run commands that available in the system. Got it from environment variables or in the working directory.

- [ ] **Pipe Stuff**  
  Send output from another command to another. This thing `|`.

- [ ] **Redirection**  
  Send output to file. Something like `echo "hi" > text.txt`.

- [ ] **Environment Variables**  
  Set and get environment variables.


## How to Run This Thing

1. Clone this repo.
2. Run it with `go run .`.
3. Have fun typing commands. If you break something, no worries, just restart it.

### Prerequisites

- You need **Go**. Just Go here https://go.dev/doc/install.
- I write this in **Go 1.23.2**. I think **Go 1.23+** should be okay.

### Installation

```bash
git clone https://github.com/hmdfrds/go-shell.git
cd go-shell
go run .
```  

## License  
This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for the detail.
