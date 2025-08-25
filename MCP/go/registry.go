package main

import (
	"github.com/amplifybackend/mcp-server/config"
	"github.com/amplifybackend/mcp-server/models"
	tools_backend "github.com/amplifybackend/mcp-server/tools/backend"
	tools_s3buckets "github.com/amplifybackend/mcp-server/tools/s3buckets"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_backend.CreateCreatebackendTool(cfg),
		tools_backend.CreateListbackendjobsTool(cfg),
		tools_backend.CreateImportbackendstorageTool(cfg),
		tools_backend.CreateGetbackendauthTool(cfg),
		tools_backend.CreateCreatebackendapiTool(cfg),
		tools_backend.CreateDeletebackendTool(cfg),
		tools_backend.CreateDeletebackendstorageTool(cfg),
		tools_backend.CreateGetbackendapimodelsTool(cfg),
		tools_backend.CreateCreatebackendconfigTool(cfg),
		tools_backend.CreateDeletetokenTool(cfg),
		tools_backend.CreateCreatebackendstorageTool(cfg),
		tools_backend.CreateUpdatebackendapiTool(cfg),
		tools_backend.CreateUpdatebackendauthTool(cfg),
		tools_backend.CreateGetbackendjobTool(cfg),
		tools_backend.CreateUpdatebackendjobTool(cfg),
		tools_backend.CreateGetbackendapiTool(cfg),
		tools_backend.CreateGeneratebackendapimodelsTool(cfg),
		tools_backend.CreateDeletebackendapiTool(cfg),
		tools_backend.CreateImportbackendauthTool(cfg),
		tools_backend.CreateGettokenTool(cfg),
		tools_backend.CreateGetbackendTool(cfg),
		tools_backend.CreateUpdatebackendconfigTool(cfg),
		tools_backend.CreateCreatebackendauthTool(cfg),
		tools_backend.CreateClonebackendTool(cfg),
		tools_backend.CreateDeletebackendauthTool(cfg),
		tools_backend.CreateUpdatebackendstorageTool(cfg),
		tools_backend.CreateGetbackendstorageTool(cfg),
		tools_backend.CreateRemovebackendconfigTool(cfg),
		tools_backend.CreateRemoveallbackendsTool(cfg),
		tools_backend.CreateCreatetokenTool(cfg),
		tools_s3buckets.CreateLists3bucketsTool(cfg),
	}
}
