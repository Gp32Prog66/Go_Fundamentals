/*
package filedemo

import (
	"bufio"
	"fmt"
	"os"
)

func filedemo() {
	userFile, err := os.Open("BasicText.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	//Create Reader
	r := bufio.NewReader(userFile)

	//Go Through Each Line
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}

		fmt.Print(line)
	}

	//Close the File
	defer userFile.Close()
}
*/