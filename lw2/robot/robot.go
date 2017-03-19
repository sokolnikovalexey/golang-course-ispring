package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Command interface {
	Execute()
}

type ShowHelpCommand struct {
	menu *Menu
}

func (c *ShowHelpCommand) Execute() {
	c.menu.ShowInstructions()
}

type ExitMenuCommand struct {
	menu *Menu
}

func (c *ExitMenuCommand) Execute() {
	c.menu.Exit()
}

type OnCommand struct {
	robot *Robot
}

func (c *OnCommand) Execute() {
	c.robot.TurnOn()
}

type OffCommand struct {
	robot *Robot
}

func (c *OffCommand) Execute() {
	c.robot.TurnOff()
}

type WalkCommand struct {
	robot     *Robot
	direction int64
}

func (c *WalkCommand) Execute() {
	c.robot.Walk(c.direction)
}

type CompositeCommand struct {
	commands []Command
}

func (c *CompositeCommand) Execute() {
	for i := 0; i < len(c.commands); i++ {
		c.commands[i].Execute()
	}
}

func (c *CompositeCommand) AddCommand(com Command) {
	c.commands = append(c.commands, com)
}

type HorseMovingCommand struct {
	CompositeCommand
}

type StatusCommand struct {
	robot *Robot
}

func (c *StatusCommand) Execute() {
	c.robot.Status()
}

type StopCommand struct {
	robot *Robot
}

func (c *StopCommand) Execute() {
	c.robot.Stop()
}

const (
	UP           int64 = iota
	DOWN
	LEFT
	RIGHT
	NO_DIRECTION
)

type Robot struct {
	direction int64
	turnedOn  bool
}

func (r *Robot) TurnOn() {
	if !r.turnedOn {
		r.turnedOn = true
		fmt.Println("It am waiting for your commands")
	}
}

func (r *Robot) TurnOff() {
	if r.turnedOn {
		r.turnedOn = false
		r.direction = NO_DIRECTION
		fmt.Println("It is a pleasure to serve you")
	}
}

func (r *Robot) Walk(direction int64) {
	if r.turnedOn {
		r.direction = direction
		directions := make(map[int64]string)
		directions[UP] = "up"
		directions[DOWN] = "down"
		directions[LEFT] = "left"
		directions[RIGHT] = "right"
		fmt.Printf("Walking %v\n", directions[direction])
	} else {
		fmt.Println("The robot should be turned on first")
	}
}

func (r *Robot) Status() {
	if r.turnedOn {
		directionStr := "to Hell"
		if r.direction == UP {
			directionStr = "Up"
		} else if r.direction == DOWN {
			directionStr = "Down"
		} else if r.direction == LEFT {
			directionStr = "Left"
		} else if r.direction == RIGHT {
			directionStr = "Right"
		}
		fmt.Printf("I moving to %v\n", directionStr)
	} else {
		fmt.Println("The robot should be turned on first")
	}
}

func (r *Robot) Stop() {
	if r.turnedOn {
		if r.direction != NO_DIRECTION {
			r.direction = NO_DIRECTION
			fmt.Printf("Stopped\n")
		} else {
			fmt.Printf("I am staying still\n")
		}
	} else {
		fmt.Println("The robot should be turned on first")
	}
}

func NewRobot() *Robot {
	return &Robot{NO_DIRECTION, false}
}

type Menu struct {
	exit  bool
	items map[string]Item
}

type Item struct {
	shortcut    string
	description string
	command     Command
}

func (m *Menu) AddItem(shortcut, description string, command Command) {
	m.items[shortcut] = Item{shortcut, description, command}
}

func (m *Menu) Run(input *bufio.Reader) {
	for {
		s, isPrefix, err := input.ReadLine()
		if err == io.EOF {
			break
		}
		if isPrefix {
			fmt.Println("Command is too long, try again")
			continue
		}
		if !m.executeCommand(string(s)) {
			break
		}
	}
}

func (m *Menu) executeCommand(word string) bool {
	m.exit = false
	item, ok := m.items[word]
	if !ok {
		fmt.Println("Unknown command")
	} else {
		item.command.Execute()
	}
	return !m.exit
}

func (m *Menu) ShowInstructions() {
	fmt.Println("Commands list:")
	for _, item := range m.items {
		fmt.Printf("\t%v: %v\n", item.shortcut, item.description)
	}
}

func NewMenu() *Menu {
	m := &Menu{}
	m.items = make(map[string]Item)
	return m
}

func (m *Menu) Exit() {
	m.exit = true
}

func main() {
	robot := NewRobot()
	menu := NewMenu()

	menu.AddItem("on", "Turns the Robot on", &OnCommand{robot})
	menu.AddItem("off", "Turns the Robot off", &OffCommand{robot})

	menu.AddItem("up", "Makes the Robot walk up", &WalkCommand{robot, UP})
	menu.AddItem("down", "Makes the Robot walk down", &WalkCommand{robot, DOWN})
	menu.AddItem("left", "Makes the Robot walk left", &WalkCommand{robot, LEFT})
	menu.AddItem("right", "Makes the Robot walk right", &WalkCommand{robot, RIGHT})

	compositeCommand := &HorseMovingCommand{}
	compositeCommand.AddCommand(&WalkCommand{robot, UP})
	compositeCommand.AddCommand(&WalkCommand{robot, UP})
	compositeCommand.AddCommand(&WalkCommand{robot, LEFT})
	menu.AddItem("horse_moving", "Makes the Robot walk like horse", compositeCommand)

	menu.AddItem("status", "Prints Robot status (turned on/off, walk direction)", &StatusCommand{robot})
	menu.AddItem("stop", "Stops the Robot", &StopCommand{robot})

	menu.AddItem("help", "Show instructions", &ShowHelpCommand{menu})
	menu.AddItem("exit", "Exit from this menu", &ExitMenuCommand{menu})

	menu.executeCommand("help")
	menu.Run(bufio.NewReader(os.Stdin))
}
