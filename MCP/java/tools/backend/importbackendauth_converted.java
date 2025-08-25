/**
 * MCP Server function for Imports an existing backend authentication resource.
 */

import java.io.IOException;
import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.util.*;
import java.util.function.Function;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.JsonNode;

class Post_Backend_App_Id_Auth_Backend_Environment_Name_ImportMCPTool {
    
    public static Function<MCPServer.MCPRequest, MCPServer.MCPToolResult> getPost_Backend_App_Id_Auth_Backend_Environment_Name_ImportHandler(MCPServer.APIConfig config) {
        return (request) -> {
            try {
                Map<String, Object> args = request.getArguments();
                if (args == null) {
                    return new MCPServer.MCPToolResult("Invalid arguments object", true);
                }
                
                List<String> queryParams = new ArrayList<>();
        if (args.containsKey("X-Amz-Content-Sha256")) {
            queryParams.add("X-Amz-Content-Sha256=" + args.get("X-Amz-Content-Sha256"));
        }
        if (args.containsKey("X-Amz-Date")) {
            queryParams.add("X-Amz-Date=" + args.get("X-Amz-Date"));
        }
        if (args.containsKey("X-Amz-Algorithm")) {
            queryParams.add("X-Amz-Algorithm=" + args.get("X-Amz-Algorithm"));
        }
        if (args.containsKey("X-Amz-Credential")) {
            queryParams.add("X-Amz-Credential=" + args.get("X-Amz-Credential"));
        }
        if (args.containsKey("X-Amz-Security-Token")) {
            queryParams.add("X-Amz-Security-Token=" + args.get("X-Amz-Security-Token"));
        }
        if (args.containsKey("X-Amz-Signature")) {
            queryParams.add("X-Amz-Signature=" + args.get("X-Amz-Signature"));
        }
        if (args.containsKey("X-Amz-SignedHeaders")) {
            queryParams.add("X-Amz-SignedHeaders=" + args.get("X-Amz-SignedHeaders"));
        }
        if (args.containsKey("appId")) {
            queryParams.add("appId=" + args.get("appId"));
        }
        if (args.containsKey("backendEnvironmentName")) {
            queryParams.add("backendEnvironmentName=" + args.get("backendEnvironmentName"));
        }
        if (args.containsKey("identityPoolId")) {
            queryParams.add("identityPoolId=" + args.get("identityPoolId"));
        }
        if (args.containsKey("nativeClientId")) {
            queryParams.add("nativeClientId=" + args.get("nativeClientId"));
        }
        if (args.containsKey("userPoolId")) {
            queryParams.add("userPoolId=" + args.get("userPoolId"));
        }
        if (args.containsKey("webClientId")) {
            queryParams.add("webClientId=" + args.get("webClientId"));
        }
                
                String queryString = queryParams.isEmpty() ? "" : "?" + String.join("&", queryParams);
                String url = config.getBaseUrl() + "/api/v2/post_backend_app_id_auth_backend_environment_name_import" + queryString;
                
                HttpClient client = HttpClient.newHttpClient();
                HttpRequest httpRequest = HttpRequest.newBuilder()
                    .uri(URI.create(url))
                    .header("Authorization", "Bearer " + config.getApiKey())
                    .header("Accept", "application/json")
                    .GET()
                    .build();
                
                HttpResponse<String> response = client.send(httpRequest, HttpResponse.BodyHandlers.ofString());
                
                if (response.statusCode() >= 400) {
                    return new MCPServer.MCPToolResult("API error: " + response.body(), true);
                }
                
                // Pretty print JSON
                ObjectMapper mapper = new ObjectMapper();
                JsonNode jsonNode = mapper.readTree(response.body());
                String prettyJson = mapper.writerWithDefaultPrettyPrinter().writeValueAsString(jsonNode);
                
                return new MCPServer.MCPToolResult(prettyJson);
                
            } catch (IOException | InterruptedException e) {
                return new MCPServer.MCPToolResult("Request failed: " + e.getMessage(), true);
            } catch (Exception e) {
                return new MCPServer.MCPToolResult("Unexpected error: " + e.getMessage(), true);
            }
        };
    }
    
    public static MCPServer.Tool createPost_Backend_App_Id_Auth_Backend_Environment_Name_ImportTool(MCPServer.APIConfig config) {
        Map<String, Object> parameters = new HashMap<>();
        parameters.put("type", "object");
        Map<String, Object> properties = new HashMap<>();
        Map<String, Object> X-Amz-Content-Sha256Property = new HashMap<>();
        X-Amz-Content-Sha256Property.put("type", "string");
        X-Amz-Content-Sha256Property.put("required", false);
        X-Amz-Content-Sha256Property.put("description", "");
        properties.put("X-Amz-Content-Sha256", X-Amz-Content-Sha256Property);
        Map<String, Object> X-Amz-DateProperty = new HashMap<>();
        X-Amz-DateProperty.put("type", "string");
        X-Amz-DateProperty.put("required", false);
        X-Amz-DateProperty.put("description", "");
        properties.put("X-Amz-Date", X-Amz-DateProperty);
        Map<String, Object> X-Amz-AlgorithmProperty = new HashMap<>();
        X-Amz-AlgorithmProperty.put("type", "string");
        X-Amz-AlgorithmProperty.put("required", false);
        X-Amz-AlgorithmProperty.put("description", "");
        properties.put("X-Amz-Algorithm", X-Amz-AlgorithmProperty);
        Map<String, Object> X-Amz-CredentialProperty = new HashMap<>();
        X-Amz-CredentialProperty.put("type", "string");
        X-Amz-CredentialProperty.put("required", false);
        X-Amz-CredentialProperty.put("description", "");
        properties.put("X-Amz-Credential", X-Amz-CredentialProperty);
        Map<String, Object> X-Amz-Security-TokenProperty = new HashMap<>();
        X-Amz-Security-TokenProperty.put("type", "string");
        X-Amz-Security-TokenProperty.put("required", false);
        X-Amz-Security-TokenProperty.put("description", "");
        properties.put("X-Amz-Security-Token", X-Amz-Security-TokenProperty);
        Map<String, Object> X-Amz-SignatureProperty = new HashMap<>();
        X-Amz-SignatureProperty.put("type", "string");
        X-Amz-SignatureProperty.put("required", false);
        X-Amz-SignatureProperty.put("description", "");
        properties.put("X-Amz-Signature", X-Amz-SignatureProperty);
        Map<String, Object> X-Amz-SignedHeadersProperty = new HashMap<>();
        X-Amz-SignedHeadersProperty.put("type", "string");
        X-Amz-SignedHeadersProperty.put("required", false);
        X-Amz-SignedHeadersProperty.put("description", "");
        properties.put("X-Amz-SignedHeaders", X-Amz-SignedHeadersProperty);
        Map<String, Object> appIdProperty = new HashMap<>();
        appIdProperty.put("type", "string");
        appIdProperty.put("required", true);
        appIdProperty.put("description", "The app ID.");
        properties.put("appId", appIdProperty);
        Map<String, Object> backendEnvironmentNameProperty = new HashMap<>();
        backendEnvironmentNameProperty.put("type", "string");
        backendEnvironmentNameProperty.put("required", true);
        backendEnvironmentNameProperty.put("description", "The name of the backend environment.");
        properties.put("backendEnvironmentName", backendEnvironmentNameProperty);
        Map<String, Object> identityPoolIdProperty = new HashMap<>();
        identityPoolIdProperty.put("type", "string");
        identityPoolIdProperty.put("required", false);
        identityPoolIdProperty.put("description", "Input parameter: The ID of the Amazon Cognito identity pool.");
        properties.put("identityPoolId", identityPoolIdProperty);
        Map<String, Object> nativeClientIdProperty = new HashMap<>();
        nativeClientIdProperty.put("type", "string");
        nativeClientIdProperty.put("required", true);
        nativeClientIdProperty.put("description", "Input parameter: The ID of the Amazon Cognito native client.");
        properties.put("nativeClientId", nativeClientIdProperty);
        Map<String, Object> userPoolIdProperty = new HashMap<>();
        userPoolIdProperty.put("type", "string");
        userPoolIdProperty.put("required", true);
        userPoolIdProperty.put("description", "Input parameter: The ID of the Amazon Cognito user pool.");
        properties.put("userPoolId", userPoolIdProperty);
        Map<String, Object> webClientIdProperty = new HashMap<>();
        webClientIdProperty.put("type", "string");
        webClientIdProperty.put("required", true);
        webClientIdProperty.put("description", "Input parameter: The ID of the Amazon Cognito web client.");
        properties.put("webClientId", webClientIdProperty);
        parameters.put("properties", properties);
        
        MCPServer.ToolDefinition definition = new MCPServer.ToolDefinition(
            "post_backend_app_id_auth_backend_environment_name_import",
            "Imports an existing backend authentication resource.",
            parameters
        );
        
        return new MCPServer.Tool(definition, getPost_Backend_App_Id_Auth_Backend_Environment_Name_ImportHandler(config));
    }
    
}