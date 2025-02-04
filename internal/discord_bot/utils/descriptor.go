package utils

import (
	"github.com/bwmarrin/discordgo"
)

type Predicate func(command string, commandType discordgo.InteractionType) bool

type Method func(*discordgo.Session, *discordgo.InteractionCreate)

type MethodDescriptor struct {
	Predicate Predicate
	Method    Method
}

func CommandEqual(expectedCommand string, expectedType discordgo.InteractionType) Predicate {
	return func(command string, commandType discordgo.InteractionType) bool {
		return command == expectedCommand && commandType == expectedType
	}
}
