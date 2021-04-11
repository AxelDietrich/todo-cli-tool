package TO_DO

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Task struct {
	Name string `json:"name"`
	Description string `json:"description"`
	DueDate time.Time `json:due_date`
	FinishDate time.Time `json:finish_date`
}
func Find(slice []string, val string) bool {
	for _, value := range slice {
		if (value == val) {
			return true
		}
	}
	return false;
}

func ReadInput(reader *bufio.Reader) {
	option, _ := reader.ReadString('\n')
	inputs := strings.Fields(option);
	for len(inputs) > 2 {
		fmt.Println("Too many arguments")
		option, _ = reader.ReadString('\n')
	}
}

func main () {
	fmt.Println("Welcome to your CLI TO-DO List")
	fmt.Println("")
	fmt.Println("Possible commands:")
	fmt.Println("add <task name>")
	fmt.Println("   description <task description>")
	fmt.Println("	due <due date>")
	fmt.Println("showall -> Lists all current tasks")
	fmt.Println("delete <task name>")
	fmt.Println("finish <task name>")

	options := []string{"add", "showall", "delete", "finish", "description", "due"}

	reader := bufio.NewReader(os.Stdin)
	option, _ := reader.ReadString('\n')
	inputs := strings.Fields(option);
	for len(inputs) > 2 {
		fmt.Println("Too many inputs provided")
		option, _ = reader.ReadString('\n')
	}
	strings.ToLower(option)


	_, found := Find(options)

	if found {
		switch option {
		case "add":
			if len(inputs) == 1 {
				fmt.Println("Task name empty")

			}
			newTask := &Task{}
			newTask.Name =
		}
	}

}
