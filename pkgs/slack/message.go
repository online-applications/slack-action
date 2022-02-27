package slack

import (
	"fmt"
	"log"
)

type Message struct {
	Channel     string       `json:"channel"`
	Username    string       `json:"username"`
	IconEmoji   string       `json:"icon_emoji"`
	LinkNames   int          `json:"link_names"`
	Attachments []Attachment `json:"attachments"`
}

type Attachment struct {
	Color  string  `json:"color"`
	Title  string  `json:"title"`
	Text   string  `json:"text"`
	Fields []Field `json:"fields"`
}

type Field struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}

type MessageFactory struct {
	ProjectName 	  string `json:"project_name"`
	ProjectUrl        string `json:"project_url"`
	BuildUrl          string `json:"build_url"`
	Commiter          string `json:"commiter"`
	Environment       string `json:"environment"`
	Team              string `json:"team"`
	BuildName         string `json:"buildname"`
	BuildPipelineName string `json:"buildpipelinename"`
	BuildJobName      string `json:"buildjobname"`
	CommitMessage     string `json:"commitmessage"`
	Channel           string `json:"channel"`
  
}

func CreateMessageFactory(projectName, projectUrl, buildUrl, commiterID, environment, team, buildName, commitMeessage, channel string) MessageFactory {
	g := MessageFactory{}
	g.ProjectName 	= projectName
	g.ProjectUrl 	= projectUrl
	g.BuildUrl      = buildUrl
	g.Commiter 		= commiterID
	g.Environment   = environment
	g.Team          = team
	g.BuildName     = buildName
	g.CommitMessage = commitMeessage
	g.Channel       = channel
	return g
}

func (m *MessageFactory) StartMessage() Message {
	log.Println("Start message func started")
	payload := Message{
		Channel:   m.Channel,
		Username:  "GitHub Actions",
		IconEmoji: ":githubactions:",
		LinkNames: 1,
		Attachments: []Attachment{
			{
				Color: "#3465eb",
				Title: fmt.Sprintf("GitHub Actions Build Started: %s", m.ProjectName),
				Text:  fmt.Sprintf("Check it out at: <%s | %s pipeline >", m.BuildUrl, m.ProjectName),
				Fields: []Field{
                    {
					Title: "Commiter",
					Value: "<@" + m.Commiter + ">",
					Short: true,
                    },
                    {
					Title: "Project",
					Value: m.ProjectName,
					Short: true,
					},
                    {
					Title: "Environment",
					Value: m.Environment,
					Short: true,
					},
					{
					Title: "Pipeline",
					Value: m.BuildName,
					Short: true,
					},
                    {
					Title: "Team",
					Value: m.Team,
					Short: true,
					},
					{
					Title: "Commit Message",
					Value: m.CommitMessage,
					Short: true,
					},
				},
			},
		},
	}
	return payload
}

func (m *MessageFactory) SuccessMessage() Message {
	log.Println("Success message func started")
	payload := Message{
		Channel:   m.Channel,
		Username:  "GitHub Actions",
		IconEmoji: ":githubactions:",
		LinkNames: 1,
		Attachments: []Attachment{
			{
				Color: "#36a64f",
				Title: fmt.Sprintf("GitHub Actions Build Succeeded: %s", m.ProjectName),
				Text:  fmt.Sprintf("<%s | %s pipeline >", m.BuildUrl, m.ProjectName),
				Fields: []Field{
                    {
					Title: "Commiter",
					Value: "<@" + m.Commiter + ">",
					Short: true,
                    },
                    {
					Title: "Project",
					Value: m.ProjectName,
					Short: true,
					},
                    {
					Title: "Environment",
					Value: m.Environment,
					Short: true,
					},
					{
					Title: "Pipeline",
					Value: m.BuildName,
					Short: true,
					},
                    {
					Title: "Team",
					Value: m.Team,
					Short: true,
					},
					{
					Title: "Commit Message",
					Value: m.CommitMessage,
					Short: true,
					},
				},
			},
		},
	}
	return payload
}

func (m *MessageFactory) FailedMessage() Message {
	log.Println("Fail message func started")
	payload := Message{
		Channel:   m.Channel,
		Username:  "GitHub Actions",
		IconEmoji: ":githubactions:",
		LinkNames: 1,
		Attachments: []Attachment{
			{
				Color: "#d42a1e",
				Title: fmt.Sprintf(":collision: Build Failed: %s :collision:", m.ProjectName,),
				Text:  fmt.Sprintf("<%s | %s pipeline >", m.BuildUrl, m.ProjectName),
				Fields: []Field{
                    {
					Title: "Commiter",
					Value: "<@" + m.Commiter + ">",
					Short: true,
                    },
                    {
					Title: "Project",
					Value: m.ProjectName,
					Short: true,
					},
                    {
					Title: "Environment",
					Value: m.Environment,
					Short: true,
					},
					{
					Title: "Pipeline",
					Value: m.BuildName,
					Short: true,
					},
                    {
					Title: "Team",
					Value: m.Team,
					Short: true,
					},
					{
					Title: "Commit Message",
					Value: m.CommitMessage,
					Short: true,
					},
				},
			},
		},
	}
	return payload
}