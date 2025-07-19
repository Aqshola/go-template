package controller_main

import model_main "go-template/src/model/main"

func TransformGetMainResponse(modelData []model_main.TestTable) []GetMainResponse {
	var parsedData = []GetMainResponse{}

	for _, v := range modelData {
		parsedData = append(parsedData, GetMainResponse{
			Id:   v.Id,
			Name: v.Name,
		})
	}

	return parsedData
}

func TransformGetDetailMainResponse(modelData model_main.TestTable) GetDetailMainResponse {
	var parsedData = GetDetailMainResponse{
		Id:   modelData.Id,
		Name: modelData.Name,
		Detail: ChildDetailMainResponse{
			Id:       modelData.Detail.Id,
			IsDetail: modelData.Detail.IsDetail,
		},
	}

	return parsedData
}
