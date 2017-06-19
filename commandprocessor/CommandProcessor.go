package commandprocessor

import (
	"strings"

	"github.com/vmednis/tgbot/tgtype"
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

// RegisterCommand registers command for recieving callbacks
func (cp *CommandPorcessor) RegisterCommand(command string, callback func(arguments string, m tgtype.Message)) {
	if cp.callbacks == nil {
		cp.callbacks = make(map[string]func(string, tgtype.Message))
	}

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
	botIdentifierIndex := strings.Index(command, "@")
	if botIdentifierIndex != -1 {
		botIdentifier := command[botIdentifierIndex:]
		command = command[:botIdentifierIndex-1]

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
		if (botIdentifierIndex != -1 || m.Chat.TypeString == "private") && cp.OnUnregisteredTragetedCommand != nil {
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
