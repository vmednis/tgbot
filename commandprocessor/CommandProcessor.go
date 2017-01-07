package commandprocessor

import (
	"strings"
	"tgbot/tgtype"
)

// CommandPorcessor - Processes incomming commands, runs callbacks for them
type CommandPorcessor struct {
	BotName                       string
	SwallowRegisteredBotCommands  bool
	SwallowOtherCommands          bool
	IgnoreCommandCase             bool
	OnUnregisteredTragetedCommand func(command string, arguments string, m tgtype.Message)
	OnBeforeCommand               func(m tgtype.Message)

	callbacks map[string]func(arguments string, m tgtype.Message)
}

// RegisterCommad registers command for recieving callbacks
func (cp *CommandPorcessor) RegisterCommad(command string, callback func(arguments string, m tgtype.Message)) {
	if cp.IgnoreCommandCase {
		cp.callbacks[strings.ToLower(command)] = callback
	} else {
		cp.callbacks[command] = callback
	}
}

// ExecuteCommand executes command returns true if the command should be swallowed
func (cp *CommandPorcessor) ExecuteCommand(m tgtype.Message) bool {
	if len(m.Text) < 2 || m.Text[0] != '/' {
		//Not a command
		return false
	}

	//Split the command from arguments
	commandSplitPoint := strings.Index(m.Text, " ")
	if commandSplitPoint == -1 {
		commandSplitPoint = len(m.Text)
	}
	command := m.Text[1:commandSplitPoint]
	arguments := m.Text[commandSplitPoint:]

	//Check if the command is meant for this bot
	botIdentifierIndex := strings.Index(m.Text, "@")
	if botIdentifierIndex != -1 {
		botIdentifier := command[botIdentifierIndex+1:]
		command = command[:botIdentifierIndex]

		if strings.ToLower(botIdentifier) != strings.ToLower(cp.BotName) {
			//This command is not meant for this bot
			if cp.SwallowOtherCommands {
				return true
			}
			return false
		}
	}

	//If we can ignore command case convert command to lower case
	if cp.IgnoreCommandCase {
		command = strings.ToLower(command)
	}

	//Check if the command exists
	callback := cp.callbacks[command]

	if callback == nil {
		if cp.OnUnregisteredTragetedCommand != nil {
			cp.OnUnregisteredTragetedCommand(command, arguments, m)
		}

		if cp.SwallowOtherCommands {
			return true
		}
		return false
	}

	if cp.OnBeforeCommand != nil {
		cp.OnBeforeCommand(m)
	}
	callback(arguments, m)

	if cp.SwallowRegisteredBotCommands {
		return true
	}
	return false
}