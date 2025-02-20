## FireGo  

FireGo is a Go-based tool for scanning and exploiting misconfigured Firebase databases. It helps identify publicly accessible Firebase databases and enables takeover if vulnerabilities are found.  

### Features  
- Automated scanning of Firebase databases from a domain list  
- Detection of vulnerabilities in Firebase endpoints  
- Displays data snippets from vulnerable databases  
- Automated exploitation for database takeover  

### Installation  
Ensure Go is installed, then run:  

```sh
go install github.com/rhyru9/firego@latest
```

### Usage  
Run FireGo with a list of Firebase domains:  

```sh
firego -l firebase_list.txt
```
### Demo
[![asciicast](https://asciinema.org/a/b67veNLmI301TuHs8UcPGZElU.svg)](https://asciinema.org/a/b67veNLmI301TuHs8UcPGZElU)

### License  
[https://creativecommons.org/licenses/by/4.0/](https://creativecommons.org/licenses/by/4.0/)  

### Disclaimer  
This tool is created for educational and security testing purposes with proper authorization. Unauthorized use is the sole responsibility of the user.
