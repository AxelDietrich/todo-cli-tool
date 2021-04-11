package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type Tasks struct {
	Tasks []Task `json:"tasks"`
}

type Task struct {
	Name string `json:"name"`
	Description string `json:"description"`
	DueDate time.Time `json:"due_date"`
	FinishDate time.Time `json:"finish_date" default:"nil"`
	Finish bool `json:"-" default:"false"`
}

func FindOption(slice []string, val string) bool {
	for _, value := range slice {
		if (value == val) {
			return true
		}
	}
	return false;
}

func FindAndDeleteTask(slice []Task, name string) bool {
	for i, value := range slice {
		if value.Name == name {
			slice = append(slice[:i], slice[i+1:]...)
			return true
		}
	}
	return false
}

func ReadInput(reader *bufio.Reader) (string, []string){
	option, _ := reader.ReadString('\n')
	inputs := strings.Fields(option);
	for len(inputs) > 3 {
		fmt.Println("Too many arguments")
		option, _ = reader.ReadString('\n')
		inputs = strings.Fields(option)
	}
	strings.ToLower(inputs[0])
	return option, inputs
}

func FindAndFinish(slice []Task, name string) bool {
	for _, value := range slice {
		if value.Name == name {
			value.FinishDate = time.Now()
			value.Finish = true
			return true
		}
	}
	return false
}

func SaveChanges(tasks *Tasks){
	file, _ := json.MarshalIndent(tasks, "", " ")
	err := ioutil.WriteFile("tasks.json", file, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func Loop(options []string, inputs []string, tasks *Tasks, option string) {

	found := FindOption(options, inputs[0])

	if found {
		switch inputs[0] {
		case "add":
			if len(inputs) == 1 {
				fmt.Println("Task name empty")

			}
			newTask := &Task{}
			newTask.Name = inputs[1]
			newTask.Description = inputs[2]
			tasks.Tasks = append(tasks.Tasks, *newTask)
			SaveChanges(tasks)
		case "showall":
			for _, value := range tasks.Tasks {
				fmt.Println(value.Name)
				fmt.Println(value.Description)
				fmt.Println(value.DueDate)
				if (value.Finish == true) {
					fmt.Println(value.FinishDate)
				}
				fmt.Println("")
			}
		case "delete":
			FindAndDeleteTask(tasks.Tasks, inputs[1])
			fmt.Println("Task deleted")
			SaveChanges(tasks)
		case "finish":
			FindAndFinish(tasks.Tasks, inputs[1])
			SaveChanges(tasks)
		case "showopen":
			for _, value := range tasks.Tasks {
				if value.Finish == true {
					fmt.Println(value.Name)
					fmt.Println(value.Description)
					fmt.Println(value.DueDate)
					fmt.Println(value.FinishDate)
					fmt.Println("")
				}
			}
		case "exit":
			os.Exit(0)
		}
	}
}


func main () {

	fmt.Println("Welcome to your CLI TO-DO List")
	fmt.Println("")
	fmt.Println("Possible commands:")
	fmt.Println("add <task name> <description>")
	fmt.Println("showall -> Lists all current tasks")
	fmt.Println("showopen -> Lists all open tasks")
	fmt.Println("delete <task name>")
	fmt.Println("finish <task name>")

	options := []string{"add", "showall", "delete", "finish", "showopen"}

	reader := bufio.NewReader(os.Stdin)
	option, inputs := ReadInput(reader)

	if _, err := os.Stat("tasks.json"); err == nil {
		ioutil.WriteFile("tasks.json",nil, 0644)
	}
	jsonFile, err := os.Open("tasks.json")
	if err != nil {
		os.Exit(1)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var tasks Tasks

	json.Unmarshal(byteValue, &tasks)

	Loop(options, inputs, &tasks, option)



}
