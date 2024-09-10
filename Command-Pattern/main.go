package main

import "fmt"

type IDevice interface {
	on()
	off()
}
type TV struct {
}

func (t *TV) on() {
	fmt.Println("On TV")
}
func (t *TV) off() {
	fmt.Println("Offing the TV")
}

type Stereo struct {
}

func (s *Stereo) on() {
	fmt.Println("On Stereo")
}
func (s *Stereo) off() {
	fmt.Println("Off Stereo")
}

type ICommand interface {
	execute()
}
type OnCommand struct {
	device IDevice
}

func (o *OnCommand) execute() {
	o.device.on()
}

type OfCommand struct {
	device IDevice
}

func (o *OfCommand) execute() {
	o.device.off()
}

func main() {
	for {
		fmt.Println("1. TV off\n2. TV of\n3. Stereo On\n4. Stereo off")
		var input int
		fmt.Scanf("%d", &input)

		switch input {
		case 1:
			{
				device := &TV{}
				command := &OnCommand{device: device}
				command.execute()
				break
			}
		case 2:
			{
				device := &TV{}
				command := &OfCommand{device: device}
				command.execute()
				break
			}
		case 3:
			{
				device := &Stereo{}
				command := &OnCommand{device: device}
				command.execute()
				break
			}
		case 4:
			{
				device := &Stereo{}
				command := &OfCommand{device: device}
				command.execute()
				break
			}
		}
	}
}
