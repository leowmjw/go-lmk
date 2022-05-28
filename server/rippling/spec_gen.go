// This file is generated by github.com/mikekonan/go-oas3. DO NOT EDIT.

package rippling

import "net/http"

var spec = []byte("{\"components\":{\"schemas\":{\"Error\":{\"properties\":{\"code\":{\"format\":\"int32\",\"type\":\"integer\"},\"message\":{\"type\":\"string\"}},\"required\":[\"code\",\"message\"],\"type\":\"object\"},\"Pet\":{\"properties\":{\"id\":{\"format\":\"int64\",\"type\":\"integer\"},\"name\":{\"type\":\"string\"},\"tag\":{\"type\":\"string\"}},\"required\":[\"id\",\"name\"],\"type\":\"object\"},\"Pets\":{\"items\":{\"$ref\":\"#/components/schemas/Pet\"},\"type\":\"array\"}}},\"info\":{\"license\":{\"name\":\"MIT\"},\"title\":\"Swagger Petstore\",\"version\":\"1.0.0\"},\"openapi\":\"3.0.0\",\"paths\":{\"/pets\":{\"get\":{\"operationId\":\"listPets\",\"parameters\":[{\"description\":\"How many items to return at one time (max 100)\",\"in\":\"query\",\"name\":\"limit\",\"schema\":{\"format\":\"int32\",\"type\":\"integer\"}}],\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"$ref\":\"#/components/schemas/Pets\"}}},\"description\":\"A paged array of pets\",\"headers\":{\"x-next\":{\"description\":\"A link to the next page of responses\",\"schema\":{\"type\":\"string\"}}}},\"default\":{\"content\":{\"application/json\":{\"schema\":{\"$ref\":\"#/components/schemas/Error\"}}},\"description\":\"unexpected error\"}},\"summary\":\"List all pets\",\"tags\":[\"pets\"]},\"post\":{\"operationId\":\"createPets\",\"responses\":{\"201\":{\"description\":\"Null response\"},\"default\":{\"content\":{\"application/json\":{\"schema\":{\"$ref\":\"#/components/schemas/Error\"}}},\"description\":\"unexpected error\"}},\"summary\":\"Create a pet\",\"tags\":[\"pets\"]}},\"/pets/{petId}\":{\"get\":{\"operationId\":\"showPetById\",\"parameters\":[{\"description\":\"The id of the pet to retrieve\",\"in\":\"path\",\"name\":\"petId\",\"required\":true,\"schema\":{\"type\":\"string\"}}],\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"$ref\":\"#/components/schemas/Pet\"}}},\"description\":\"Expected response to a valid request\"},\"default\":{\"content\":{\"application/json\":{\"schema\":{\"$ref\":\"#/components/schemas/Error\"}}},\"description\":\"unexpected error\"}},\"summary\":\"Info for a specific pet\",\"tags\":[\"pets\"]}}},\"servers\":[{\"url\":\"http://petstore.swagger.io/v1\"}]}")

func Spec(w http.ResponseWriter, _ *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Write(spec)
}