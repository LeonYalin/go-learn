package main

import (
	languageBasics "github.com/LeonYalinAgentVI/go-learn/src/1.languageBasics"
	buildingabasicwebapp "github.com/LeonYalinAgentVI/go-learn/src/2.buildingABasicWebApp"
	routingandmiddleware "github.com/LeonYalinAgentVI/go-learn/src/3.routingAndMiddleware"
	managingsessions "github.com/LeonYalinAgentVI/go-learn/src/4.managingSessions"
	buildingprojecttemplate "github.com/LeonYalinAgentVI/go-learn/src/5.buildingProjectTemplate"
	convertingtogotemplates "github.com/LeonYalinAgentVI/go-learn/src/6.convertingToGoTemplates"
)

func main() {
	languageBasics.LanguageBasics()
	buildingabasicwebapp.BasicWebApp()
	routingandmiddleware.RoutingMiddleware()
	managingsessions.Sessions()
	buildingprojecttemplate.ProjectTemplate()
	convertingtogotemplates.ConvertToGoTemplates()
}
