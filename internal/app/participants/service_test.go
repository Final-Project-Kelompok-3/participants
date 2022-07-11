package participants

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Final-Project-Kelompok-3/participants/internal/dto"
	"github.com/Final-Project-Kelompok-3/participants/internal/factory"
	"github.com/Final-Project-Kelompok-3/participants/internal/model"
	"github.com/Final-Project-Kelompok-3/participants/pkg/util/mock"
	"github.com/stretchr/testify/assert"
)

// func TestFindAll(t *testing.T) {

// 	// setup database
// 	db, mock := mock.DBConnection()

// 	participants := []model.Participants{
// 		{
// 			// Model: model.Model{
// 			// 	ID:        1,
// 			// 	CreatedAt: time.Now(),
// 			// 	UpdatedAt: time.Now(),
// 			// },
// 			Name:             "test",
// 			Email:            "test@test.com",
// 			Address:          "Jl. test123 semesta 1",
// 			NISN:             "1001000123",
// 			FinalReportScore: 80,
// 			FileRequirement:  "File1.pdf",
// 		},
// 		{
// 			// Model: model.Model{
// 			// 	ID:        2,
// 			// 	CreatedAt: time.Now(),
// 			// 	UpdatedAt: time.Now(),
// 			// },
// 			Name:             "test2",
// 			Email:            "test2@test.com",
// 			Address:          "Jl. test123 semesta 2",
// 			NISN:             "100123232310123",
// 			FinalReportScore: 80,
// 			FileRequirement:  "File2.pdf",
// 		},
// 	}

// 	rows := sqlmock.NewRows([]string{"name", "email", "address", "nisn", "final_report_score", "file_requirement"}).
// 		AddRow(participants[0].Name, participants[0].Email, participants[0].Address, participants[0].NISN, participants[0].FinalReportScore, participants[0].FileRequirement).
// 		AddRow(participants[1].Name, participants[1].Email, participants[1].Address, participants[1].NISN, participants[1].FinalReportScore, participants[1].FileRequirement)
// 	fmt.Println("Error disini")

// 	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "participants" WHERE "participants"."deleted_at" IS NULL`)).WillReturnRows(rows)

// 	f := factory.NewFactory(db)
// 	service := NewService(f)

// 	dto, err := service.FindAll(context.Background(), nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if len(dto.Datas) != 2 {
// 		t.Fatalf("got %d data, want 2 data", len(dto.Datas))
// 	}

// }

func TestFindByID(t *testing.T) {
	db, mock := mock.DBConnection()
	// if err != nil {
	// 	t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	// }

	participants := []model.Participants{
		{
			Model: model.Model{
				ID:        1,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Name:             "test",
			Email:            "test@test.com",
			Address:          "Jl. test123 semesta 1",
			NISN:             "1001000123",
			FinalReportScore: 80,
			FileRequirement:  "File1.pdf",
		},
		{
			Model: model.Model{
				ID:        2,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Name:             "test2",
			Email:            "test2@test.com",
			Address:          "Jl. test123 semesta 2",
			NISN:             "100123232310123",
			FinalReportScore: 80,
			FileRequirement:  "File2.pdf",
		},
	}

	rows := sqlmock.NewRows([]string{"id", "name", "email", "address", "nisn", "final_report_score", "file_requirement", "created_at", "updated_at"}).
		AddRow(participants[0].ID, participants[0].Name, participants[0].Email, participants[0].Address, participants[0].NISN, participants[0].FinalReportScore, participants[0].FileRequirement, participants[0].CreatedAt, participants[0].UpdatedAt).
		AddRow(participants[1].ID, participants[1].Name, participants[1].Email, participants[1].Address, participants[1].NISN, participants[1].FinalReportScore, participants[1].FileRequirement, participants[1].CreatedAt, participants[1].UpdatedAt)

	// mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "participants" WHERE "participants"."deleted_at" IS NULL`)).WillReturnRows(rows)
	fmt.Println("Ini rows", rows)
	num := uint(participants[0].ID)
	println("ini num", participants[0].ID)
	query := "select * from participants where pariticipants.id = ? "

	mock.ExpectQuery((query)).WillReturnRows(rows)
	f := factory.NewFactory(db)
	a := NewService(f)

	anArticle, err := a.FindByID(context.TODO(), &dto.ByIDRequest{ID: num})
	assert.NoError(t, err)
	if err != nil {
		fmt.Println("Error article", err)
	}
	if anArticle != nil {
		fmt.Println(anArticle)
	}
	assert.NotNil(t, anArticle)
}
