package main

import (
	"log"
	"slack-action/pkgs/slack"
	"slack-action/pkgs/utils"
	"slack-action/pkgs/s3"

)

func main() {
	log.Println("Sending Slack Message")
	// Getting env variables
	projectName 	 			:= utils.GetEnv("PROJECT_NAME")
	repositoryUrl 		 		:= utils.GetEnv("REPOSITORY_URL")
	environment 	 			:= utils.GetEnv("ENVIRONMENT")
	url 						:= utils.GetEnv("SLACK_URL")
	channelID 					:= utils.GetEnv("CHANNEL_ID")
	buildName   	 			:= utils.GetEnv("GITHUB_WORKFLOW")
	team 			 			:= utils.GetEnv("TEAM")
	prBuildUrlRaw    			:= utils.GetEnv("PR_BUILD_URL")
	pushBuildUrl     			:= utils.GetEnv("PUSH_BUILD_URL")
	runId   		 			:= utils.GetEnv("RUN_ID")
	usersFile		 			:= utils.GetEnv("USERS_FILE")
	commiter 	     			:= utils.GetEnv("COMMITER")
	commitMessageRaw 			:= utils.GetEnv("COMMIT_MESSAGE")
	customPayload    			:= utils.GetEnv("CUSTOM_PAYLOAD_PATH")
	commitSha        			:= utils.GetEnv("COMMIT_SHA")
	s3FilePath 		 			:= utils.GetEnv("USERS_S3_FILE_PATH")
	// Get CLI arguments
	jobStatus 					:= utils.GetCliArg(1)
	// Getting slack variables
	buildUrl		            := slack.GetBuildUrl(prBuildUrlRaw, pushBuildUrl, runId)
	commitMessage, err_commit   := slack.GetCommit(commitSha, commitMessageRaw)
	// Get slack ID
	err_s3 						:= s3.DownloadS3(s3FilePath, usersFile)
	slackID, err_json 			:= utils.GetJsonValue(commiter, usersFile)
	// Error handing
	if err_json != nil || err_s3 != nil || err_commit != nil{
        log.Fatal("Found error!", err_json, err_s3, err_commit)
    }

	// Create slack message payload
	factory := slack.CreateMessageFactory(projectName, repositoryUrl, buildUrl, slackID, environment, team, buildName, commitMessage, channelID)

	// Send message
	switch jobStatus {
	case "started":
		payloadRaw := factory.StartMessage()
		payload, err_json := utils.JsonMarshal(payloadRaw)
		if err_json != nil {
			log.Fatal(err_json)
		}
		err := slack.SendMessage(payload, url)
		if err != nil {
			log.Fatal(err)
		}
	case "success":
		payloadRaw := factory.SuccessMessage()
		payload, err_json := utils.JsonMarshal(payloadRaw)
		if err_json != nil {
			log.Fatal(err_json)
		}
		err := slack.SendMessage(payload, url)
		if err != nil {
			log.Fatal(err)
		}
	case "failed":
		payloadRaw := factory.FailedMessage()
		payload, err_json := utils.JsonMarshal(payloadRaw)
		if err_json != nil {
			log.Fatal(err_json)
		}
		err := slack.SendMessage(payload, url)
		if err != nil {
			log.Fatal(err)
		}
	case "custom":
		payload, err_json := utils.ReadFile(customPayload)
		if err_json != nil {
			log.Fatal(err_json)
		}
		err := slack.SendMessage(payload, url)
		if err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatalln("No / Incorrect parameter provided!! \n Valid parameters are: [started, success, failed, custom]")
	}
}