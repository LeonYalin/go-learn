package main

import (
	languageBasics "github.com/LeonYalinAgentVI/go-learn/src/1.languageBasics"
	buildingabasicwebapp "github.com/LeonYalinAgentVI/go-learn/src/2.buildingABasicWebApp"
	routingandmiddleware "github.com/LeonYalinAgentVI/go-learn/src/3.routingAndMiddleware"
	managingsessions "github.com/LeonYalinAgentVI/go-learn/src/4.managingSessions"
	buildingprojecttemplate "github.com/LeonYalinAgentVI/go-learn/src/5.buildingProjectTemplate"
	convertingtogotemplates "github.com/LeonYalinAgentVI/go-learn/src/6.convertingToGoTemplates"
	writingtests "github.com/LeonYalinAgentVI/go-learn/src/7.writingTests"
	designingTheDbStructure "github.com/LeonYalinAgentVI/go-learn/src/8.designingTheDbStructure"
	connectingTheAppToDB "github.com/LeonYalinAgentVI/go-learn/src/9.connectingTheAppToDB"
)

func main() {
	languageBasics.LanguageBasics()
	buildingabasicwebapp.BasicWebApp()
	routingandmiddleware.RoutingMiddleware()
	managingsessions.Sessions()
	buildingprojecttemplate.ProjectTemplate()
	convertingtogotemplates.ConvertToGoTemplates()
	writingtests.Tests()
	designingTheDbStructure.DbStructure()
	connectingTheAppToDB.AppToDb()
}
