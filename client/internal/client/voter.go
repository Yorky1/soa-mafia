package client

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/goombaio/namegenerator"
	"github.com/pterm/pterm"
)

type MainMenuChoise int32

const (
	FindGameSession MainMenuChoise = iota
	Exit
)

type VoteMenuChoise int32

const (
	GoToChat VoteMenuChoise = iota
	Vote
)

type Voter interface {
	GetName() string
	DayVote(names []string) string
	MafiaVote(names []string) string
	SheriffVote(names []string) string
	GetMainMenuChoise() MainMenuChoise
	GetVoteMenuChoise() VoteMenuChoise
	GetMessage() string
}

type RealVoter struct {
}

func NewRealVoter() *RealVoter {
	return &RealVoter{}
}

func (v *RealVoter) GetName() string {
	res, err := pterm.DefaultInteractiveTextInput.WithMultiLine(false).WithDefaultText("To connect to the game session please enter your name").Show()
	if err != nil {
		log.Panic(err)
	}
	pterm.Println()
	return res
}

func (v *RealVoter) DayVote(names []string) string {
	printer := pterm.DefaultInteractiveSelect.WithOptions(names).WithDefaultText("Vote for a player")
	vote, _ := printer.Show()
	return vote
}

func (v *RealVoter) SheriffVote(names []string) string {
	printer := pterm.DefaultInteractiveSelect.WithOptions(names).WithDefaultText("Vote for a player to check:")
	vote, _ := printer.Show()
	return vote
}

func (v *RealVoter) MafiaVote(names []string) string {
	printer := pterm.DefaultInteractiveSelect.WithOptions(names).WithDefaultText("Vote for a player to kill:")
	vote, _ := printer.Show()
	return vote
}

func (v *RealVoter) GetMainMenuChoise() MainMenuChoise {
	printer := pterm.DefaultInteractiveSelect.WithOptions([]string{"find game session", "exit"})
	selectedOption, _ := printer.Show()

	if selectedOption == "exit" {
		return Exit
	} else {
		return FindGameSession
	}
}

func (v *RealVoter) GetVoteMenuChoise() VoteMenuChoise {
	printer := pterm.DefaultInteractiveSelect.WithOptions([]string{"vote for a player", "go to chat"})
	selectedOption, _ := printer.Show()

	if selectedOption == "vote for a player" {
		return Vote
	} else {
		return GoToChat
	}
}

func (v *RealVoter) GetMessage() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return text[:len(text)-1]
}

type BotVoter struct {
	nameGenerator namegenerator.Generator
}

func NewBotVoter() *BotVoter {
	seed := time.Now().UTC().UnixNano()
	rand.Seed(seed)
	nameGenerator := namegenerator.NewNameGenerator(seed)
	return &BotVoter{nameGenerator: nameGenerator}
}

func (v *BotVoter) GetName() string {
	return v.nameGenerator.Generate()
}

func (v *BotVoter) DayVote(names []string) string {
	return names[rand.Int31n(int32(len(names)))]
}

func (v *BotVoter) SheriffVote(names []string) string {
	return names[rand.Int31n(int32(len(names)))]
}

func (v *BotVoter) MafiaVote(names []string) string {
	return names[rand.Int31n(int32(len(names)))]
}

func (v *BotVoter) GetMainMenuChoise() MainMenuChoise {
	return FindGameSession
}

func (v *BotVoter) GetMessage() string {
	return "/exit"
}

func (v *BotVoter) GetVoteMenuChoise() VoteMenuChoise {
	return Vote
}
