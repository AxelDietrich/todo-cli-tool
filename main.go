//go:generate goversioninfo
package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type LastID struct {
	ID uint
}

type Tasks struct {
	Tasks []Task `json:"tasks"`
}

type Task struct {
	ID uint `json:"ID"`
	Name string `json:"name"`
	Description string `json:"description"`
	FinishDate time.Time `json:"finish_date" default:"nil"`
	Finish bool `json:"finished" default:"false"`
}

func FindOption(slice []string, val string) bool {
	for _, value := range slice {
		if (value == val) {
			return true
		}
	}
	return false;
}

func FindAndDeleteTask(slice *[]Task, ID string) (error) {
	findID, err := strconv.Atoi(ID)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	for i, value := range *slice {
		if value.ID == uint(findID) {
			*slice = append((*slice)[:i], (*slice)[i+1:]...)
			return nil
		}
	}
	return errors.New("There's not a task with that ID")
}

func ReadInput(reader *bufio.Reader) []string {
	option, _ := reader.ReadString('\n')
	r := regexp.MustCompile(`[^\s"']+|"([^"]*)"|'([^']*)`)
	inputs := r.FindAllString(option, -1)
	for len(inputs) > 3 {
		fmt.Println("Too many arguments")
		option, _ = reader.ReadString('\n')
		inputs = strings.Fields(option)
	}
	for i, value := range inputs {
		inputs[i] = strings.Trim(value, "\"")
	}
	strings.ToLower(inputs[0])
	return inputs
}

func FindAndFinish(slice *[]Task, ID string) error {
	findID, err := strconv.Atoi(ID)
	if err != nil {
		return err
	}
	for i, value := range *slice {
		if value.ID == uint(findID) {
			value.FinishDate = time.Now()
			value.Finish = true
			(*slice)[i] = value
			return nil
		}
	}
	return errors.New("There's no taks with that ID")
}
func SaveLastID(lastID *LastID) error {
	file, _ := json.MarshalIndent(lastID, "", " ")
	err := ioutil.WriteFile("ID.json", file, 0644)
	if err != nil {
		return err
	}
	return nil
}

func SaveChanges(tasks *Tasks) error {
	file, _ := json.MarshalIndent(tasks, "", " ")
	err := ioutil.WriteFile("tasks.json", file, 0644)
	if err != nil {
		return err
	}
	return nil
}

func Loop(options []string, tasks *Tasks, lastID *LastID) {

	reader := bufio.NewReader(os.Stdin)
	inputs := ReadInput(reader)

	found := FindOption(options, inputs[0])

	if found {
		switch inputs[0] {
		case "add":
			if len(inputs) == 1 {
				fmt.Println("Task name empty")
			}
			newTask := &Task{}
			newTask.ID = lastID.ID + 1
			lastID.ID = newTask.ID
			newTask.Name = inputs[1]
			if len(inputs) > 2 {
				newTask.Description = inputs[2]
			}
			tasks.Tasks = append(tasks.Tasks, *newTask)
			errID := SaveLastID(lastID)
			if errID != nil {
				fmt.Println(errID.Error())
			} else {
				err := SaveChanges(tasks)
				if err != nil {
					fmt.Println(err.Error())
					fmt.Println("")
				} else {
					fmt.Println("Task added")
					fmt.Println("")
				}
			}
			fmt.Println("-------------------------------------------------------------------------------------")
			fmt.Println("")

		case "showall":
			if len(inputs) > 1 {
				fmt.Println("Too many arguments")
				break
			}
			if len(tasks.Tasks) == 0{
				fmt.Println("There isn't tasks yet")
			} else {
				fmt.Println("")
				for _, value := range tasks.Tasks {
					fmt.Printf("ID: %s\n", strconv.FormatUint(uint64(value.ID), 10))
					fmt.Printf("Name: %s\n", value.Name)
					if (value.Description != "") {
						fmt.Printf("Description: %s\n", value.Description)
					}
					if value.Finish {
						fmt.Printf("Finish date: %s\n", value.FinishDate)
					}
					fmt.Println("")
				}
			}
			fmt.Println("-------------------------------------------------------------------------------------")
			fmt.Println("")
		case "delete":
			if len(inputs) > 2 {
				fmt.Println("Too many arguments")
				break
			}
			err := FindAndDeleteTask(&tasks.Tasks, inputs[1])
			if err != nil {
				fmt.Println(err.Error())
				break
			}
			err = SaveChanges(tasks)
			if err != nil {
				fmt.Println(err.Error())
			} else{
				fmt.Println("Task deleted")
				fmt.Println("")
			}
			fmt.Println("-------------------------------------------------------------------------------------")
			fmt.Println("")
		case "finish":
			if len(inputs) > 2 {
				fmt.Println("Too many arguments")
				fmt.Println("")
				break
			}
			err := FindAndFinish(&tasks.Tasks, inputs[1])
			if err != nil {
				fmt.Print(err.Error())
			}
			err = SaveChanges(tasks)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("Tasks finished")
				fmt.Println("")
			}
			fmt.Println("-------------------------------------------------------------------------------------")
			fmt.Println("")
		case "showopen":
			if len(inputs) > 1 {
				fmt.Println("")
				fmt.Println("Too many arguments")
				fmt.Println("")
				break
			}
			var count int
			fmt.Println("")
			for _, value := range tasks.Tasks {
				if value.Finish == false {
					count++
					fmt.Println("")
					fmt.Printf("ID: %s\n", strconv.FormatUint(uint64(value.ID), 10))
					fmt.Printf("Name: %s\n", value.Name)
					if (value.Description != "") {
						fmt.Printf("Description: %s\n", value.Description)
					}
					fmt.Println("")
					fmt.Println("-------------------------------------------------------------------------------------")
					fmt.Println("")
				}
			}
			if count == 0 {
				fmt.Println("There isn't open tasks")
			}
		case "showfinished":
			if len(inputs) > 1 {
				fmt.Println("Too many arguments")
				fmt.Println("")
				break
			}
			var count int
			fmt.Println("")
			for _, value := range tasks.Tasks {
				if value.Finish {
					count++
					fmt.Printf("ID: %s\n", strconv.FormatUint(uint64(value.ID), 10))
					fmt.Printf("Name: %s\n", value.Name)
					if (value.Description != "") {
						fmt.Printf("Description: %s\n", value.Description)
					}
					fmt.Printf("Finish date: %s\n", value.FinishDate)
					fmt.Println("")
				}
			}
			if count == 0 {
				fmt.Println("There isn't open tasks")
			}
			fmt.Println("-------------------------------------------------------------------------------------")
			fmt.Println("")
		case "exit":
			os.Exit(0)
		}
	} else {
		fmt.Println("Not a valid command")
		fmt.Println("-------------------------------------------------------------------------------------")
		fmt.Println("")
	}
}


func main () {

	fmt.Println("Welcome to your CLI TO-DO List")
	fmt.Println("")
	fmt.Println("Possible commands:")
	fmt.Println("")
	fmt.Println("add \"task name\" \"description\" -> Adds new task")
	fmt.Println("showall -> Lists all current tasks")
	fmt.Println("showopen -> Lists all open tasks")
	fmt.Println("showfinished -> Lists all finished tasks")
	fmt.Println("delete <task ID> -> Deletes task")
	fmt.Println("finish <task name> -> Marks a task as finished")
	fmt.Println("exit -> Closes the app")
	fmt.Println("")

	options := []string{"add", "showall", "delete", "finish", "showopen", "showfinished", "exit"}


	homeDir, _ := os.UserHomeDir()
	todoDir := homeDir + "\\todo\\"
	if _, err := os.Stat(todoDir); os.IsNotExist(err) {
		err = os.Mkdir(todoDir, 0755)
		if err != nil {
			os.Exit(1)
		}
	}

	if _, err := os.Stat(todoDir + "tasks.json"); err != nil {
		ioutil.WriteFile(todoDir + "tasks.json",nil, 0644)
	}
	jsonFile, err := os.Open(todoDir + "tasks.json")
	if err != nil {
		os.Exit(1)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var tasks Tasks

	json.Unmarshal(byteValue, &tasks)

	//Parse IDs
	if _, err := os.Stat(todoDir + "ID.json"); err != nil {
		ioutil.WriteFile(todoDir + "ID.json",nil, 0644)
	}
	jsonFileID, err := os.Open(todoDir + "ID.json")
	if err != nil {
		os.Exit(1)
	}

	byteValueID, _ := ioutil.ReadAll(jsonFileID)

	var lastID LastID

	json.Unmarshal(byteValueID, &lastID)

	for {
		Loop(options, &tasks, &lastID)
	}


}
