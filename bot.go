package main

type BotClient interface {
	Ask(question string) (answer string)
	Name() (name string)
}


