package client

import (
	"math/rand"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/goombaio/namegenerator"
)

type Voter interface {
	GetName() string
	DayVote(names []string) string
	MafiaVote(names []string) string
	SheriffVote(names []string) string
}

type RealVoter struct {
}

func NewRealVoter() *RealVoter {
	return &RealVoter{}
}

func (v *RealVoter) GetName() string {
	answer := struct {
		Name string
	}{}
	qs := []*survey.Question{
		{
			Name:      "name",
			Prompt:    &survey.Input{Message: "To connect to the game session please enter your name:"},
			Validate:  survey.Required,
			Transform: survey.Title,
		},
	}
	survey.Ask(qs, &answer)
	return answer.Name
}

func (v *RealVoter) DayVote(names []string) string {
	vote := ""
	prompt := &survey.Select{
		Message: "Vote for a player:",
		Options: names,
	}
	survey.AskOne(prompt, &vote)
	return vote
}

func (v *RealVoter) SheriffVote(names []string) string {
	vote := ""
	prompt := &survey.Select{
		Message: "Vote for a player to check:",
		Options: names,
	}
	survey.AskOne(prompt, &vote)
	return vote
}

func (v *RealVoter) MafiaVote(names []string) string {
	vote := ""
	prompt := &survey.Select{
		Message: "Vote for a player to kill:",
		Options: names,
	}
	survey.AskOne(prompt, &vote)
	return vote
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
