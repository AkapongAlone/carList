package usecases

import (
	"Daveslist/models"
	"Daveslist/responses"
	"reflect"
	"testing"
	"time"
)

func Test_praseModelMessageToRespMessage(t *testing.T) {
	type args struct {
		model []models.Message
	}
	timeNow := time.Now()
	mockModels := []models.Message{}
	mockModel := models.Message{}
	mockModel.ID = 1
	mockModel.FromUserID = 1
	mockModel.ToUserID = 2
	mockModel.Message = "Hello World"
	mockModel.CreatedAt = timeNow
	mockModels = append(mockModels, mockModel)
	tests := []struct {
		name     string
		args     args
		expected []responses.MessageResp
	}{
		{
			name: "Test with one messages",
			args: args{
				model: mockModels,
			},
			expected: []responses.MessageResp{
				{ Message: "Hello World", FromUserID: 1,ToUserID: 2,CreatedAt: timeNow},
			},
		},
		{
			name: "Test with empty messages",
			args: args{
				model: []models.Message{},
			},
			expected: []responses.MessageResp{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := praseModelMessageToRespMessage(tt.args.model)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("praseModelMessageToRespMessage() = %v, expected %v", got, tt.expected)
			}
		})
	}
}
