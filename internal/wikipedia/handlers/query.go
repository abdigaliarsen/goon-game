package handlers

import (
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	desc "goon-game/pkg/proto/wikipedia"
)

func (s *Server) SetLanguage(ctx context.Context, req *desc.SetLanguageRequest) (*desc.EmptyResponse, error) {
	err := s.wikipediaService.SetLanguage(req.GetLanguage())
	if err != nil {
		return nil, err
	}

	return &desc.EmptyResponse{}, nil
}

func (s *Server) GetLanguageUpdates(ctx context.Context, req *desc.EmptyRequest) (*desc.GetLanguageUpdatesResponse, error) {
	languageUpdates, err := s.wikipediaService.GetLanguageUpdates()
	if err != nil {
		return nil, err
	}

	resp := &desc.GetLanguageUpdatesResponse{
		Updates: make([]*desc.GetLanguageUpdatesResponse_Data, 0, len(languageUpdates)),
	}

	for _, update := range languageUpdates {
		resp.Updates = append(resp.Updates, &desc.GetLanguageUpdatesResponse_Data{
			Language:  update.Language,
			UpdatedAt: timestamppb.New(update.UpdatedAt),
		})
	}

	return resp, nil
}
