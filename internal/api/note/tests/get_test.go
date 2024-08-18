package tests

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
	"testing"
	"week/internal/api/note"
	"week/internal/model"
	"week/internal/service"
	serviceMoks "week/internal/service/mocks"
	desc "week/pkg/note_v1"
)

func TestGet(t *testing.T) {
	type noteServiceMockFunc func(mc *minimock.Controller) service.NoteService
	type args struct {
		ctx context.Context
		req *desc.GetRequest
	}
	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id        = gofakeit.Int64()
		title     = gofakeit.Animal()
		content   = gofakeit.Animal()
		createdAt = gofakeit.Date()
		updatedAt = gofakeit.Date()

		serviceErr = fmt.Errorf("service error")

		req = &desc.GetRequest{
			Id: id,
		}

		serviceRes = &model.Note{
			ID: id,
			Info: model.NoteInfo{
				Title:   title,
				Content: content,
			},
			CreatedAt: createdAt,
			UpdatedAt: sql.NullTime{
				Time:  updatedAt,
				Valid: true,
			},
		}

		res = &desc.GetResponse{
			Note: &desc.Note{
				Id:        id,
				Info:      &desc.NoteInfo{Title: title, Content: content},
				CreatedAt: timestamppb.New(createdAt),
				UpdatedAt: timestamppb.New(updatedAt),
			},
		}
	)

	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name        string
		args        args
		want        *desc.GetResponse
		err         error
		noteService noteServiceMockFunc
	}{
		{
			name: "success case: get note by id",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			noteService: func(mc *minimock.Controller) service.NoteService {
				mock := serviceMoks.NewNoteServiceMock(mc)
				mock.GetMock.Expect(ctx, id).Return(serviceRes, nil)
				return mock
			},
		},
		{
			name: "error case: get note by id",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  serviceErr,
			noteService: func(mc *minimock.Controller) service.NoteService {
				mock := serviceMoks.NewNoteServiceMock(mc)
				mock.GetMock.Expect(ctx, id).Return(nil, serviceErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			noteServiceMock := tt.noteService(mc)
			api := note.NewImplementation(noteServiceMock)

			newId, err := api.Get(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, newId)
		})
	}

}
