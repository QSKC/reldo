package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("TODO List for Your Pleasure")
	fmt.Println("Type 'help' for commands, 'quit' to Exit")
	// listTasks() -> Does not function

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("ToDo> ")
		if !scanner.Scan() {
			break
		}
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		args := strings.Fields(line)
		command := args[0]

		switch command {
		case "add":
			if len(args) < 2 {
				fmt.Printf("❌ Missing task description.")
				continue
			}
			task := strings.Join(args[1:], " ")
			if err := addTask(task); err != nil {
				fmt.Printf("❌", err)
				continue
			}
			fmt.Printf("✅ Added task: %s \n", task)
		case "edit":

			if len(args) < 3 {
				fmt.Printf("❌ Please to make sure to follow the formula")
				fmt.Println("Edit: edit <task desc> <id>")
				return
			}
			id, err := strconv.Atoi(args[2])
			if err != nil {
				fmt.Println("Invalid command, ID must be number")
				continue
			}
			newDesc := strings.Join(args[1:], " ")
			if err := editTask(id, newDesc); err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Printf("Updated task #%d\n", id)

		case "list", "ls":
			todos, err := listTasks()
			if err != nil {
				fmt.Printf("❌", err)
				continue
			}
			if len(todos) == 0 {
				fmt.Println("📭 No tasks found! Use 'add <task>' to add one")
				continue
			}
			fmt.Printf("📝Todo List:")
			for _, t := range todos {
				status := " "
				if t.Done {
					status = "x"
				}
				fmt.Printf(" %d. [%s] %s\n", t.ID, status, t.Task)
			}

		case "done":
			if len(args) < 2 {
				fmt.Printf("❌ Missing task ID.")
				continue
			}
			id, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Printf("❌ Invalid ID. Must be a number")
				continue
			}
			if err := markDone(id); err != nil {
				fmt.Printf("❌", err)
				continue
			}
			fmt.Printf("✅ Marked task #%d as done.\n", id)

		case "remove":
			if len(args) < 2 {
				fmt.Printf("❌ Missing task ID.")
				continue
			}
			id, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Printf("❌ Invalid ID. Must be a number.")
				continue
			}
			if err := removeTask(id); err != nil {
				fmt.Printf("❌", err)
				continue
			}
			fmt.Printf("✅ Removed task #%d.\n", id)

		case "c", "clear", "clr":
			clear()
			continue
		case "help":
			printUsage()
		case "exit", "quit", "q":
			fmt.Println("Make this thing better, Keep Working At IT !!!")
			fmt.Println("✌ Adios Amigos✌")
			return
		default:
			fmt.Printf("❌ Unknown command:", command)
			printUsage()
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}

func printUsage() {
	fmt.Println("=========== Commands ===========")
	fmt.Println(" add <task>			add a new todo task")
	fmt.Println(" list 					List all tasks")
	fmt.Println(" edit <id> <newDesc> 	Edit a task description")
	fmt.Println(" done <id>				Mark a task as done")
	fmt.Println(" remove <id>			Remove a task")
	fmt.Println(" help					Shows this help")
	fmt.Println(" exit, quit			Quit the program")
}

func clear() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
