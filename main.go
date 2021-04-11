package TO_DO

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type Task struct {
	Name string `json:"name"`
	Description string `json:"description"`
	DueDate time.Time `json:due_date`
	FinishDate time.Time `json:finish_date`
}
func Find(slice []string, val string) (int, bool) {
	for slice {
		if
	}
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
		return -1, false
	}

func main () {
	fmt.Println("Welcome to your CLI TO-DO List")
	fmt.Println("")
	fmt.Println("Possible commands:")
	fmt.Println("add <task name>")
	fmt.Println("   description <task description>")
	fmt.Println("	due <due date>")
	fmt.Println("show all -> Lists all current tasks")
	fmt.Println("delete <task name>")
	fmt.Println("finish <task name>")

	options := []string{"add", "show all", "delete", "finish", "description", "due"}

	reader := bufio.NewReader(os.Stdin)
	option, _ := reader.ReadString('\n')

	_, found := Find(options)

}
