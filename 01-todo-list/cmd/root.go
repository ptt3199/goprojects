package cmd

import (
	"fmt"
	"todo-list/internal"
	"github.com/spf13/cobra"
	"strconv"
)

var taskCSVRepo = internal.NewTaskRepositoryCSV("tasks.csv")

var addCmd = &cobra.Command{
	Use:   "add [task description]",
	Short: "Add a new task to the todo list",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskDescription := args[0]
		fmt.Printf("Adding new task: %s\n", taskDescription)
		err := taskCSVRepo.AddTask(taskDescription)
		if err != nil {
			fmt.Printf("Error adding task: %s\n", err)
		}
		fmt.Printf("Task added successfully\n")
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks in the todo list",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Listing all tasks")
		tasks, err := taskCSVRepo.ListTasks(false)
		if err != nil {
			fmt.Printf("Error listing tasks: %s\n", err)
			return
		}
		for _, task := range tasks {
			fmt.Printf("Task ID: %d, Description: %s, Created At: %s, Is Complete: %t\n", task.ID, task.Description, task.CreatedAt, task.IsComplete)
		}
	},
}

var completeCmd = &cobra.Command{
	Use:   "complete [task id]",
	Short: "Complete a task in the todo list",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskID := args[0]
		id, err := strconv.Atoi(taskID)
		if err != nil {
			fmt.Printf("Invalid task ID: %s\n", err)
			return
		}
		fmt.Printf("Completing task with ID: %d\n", id)
		err = taskCSVRepo.CompleteTask(id)
		if err != nil {
			fmt.Printf("Error completing task: %s\n", err)
			return
		}
		fmt.Printf("Task completed successfully\n")
	},
}

var rootCmd = &cobra.Command{
	Use:   "tasks",
	Short: "A simple CLI todo list manager",
	Long: `Tasks is a CLI application that helps you manage your todo list
directly from the terminal. You can add, list, complete, and delete tasks.`,
}

func Execute() error {
	rootCmd.AddCommand(addCmd, listCmd, completeCmd)
	return rootCmd.Execute()
}


