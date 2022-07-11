package participants

import (
	"encoding/json"
	"net/http"
	"regexp"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/Final-Project-Kelompok-3/participants/internal/factory"
	"github.com/Final-Project-Kelompok-3/participants/internal/model"

	// "github.com/Final-Project-Kelompok-3/participants/internal/factory"
	util_mock "github.com/Final-Project-Kelompok-3/participants/pkg/util/mock"
	"github.com/labstack/echo/v4"
)

func TestFindAll(t *testing.T) {
	// setup database
	db, mock := util_mock.DBConnection()

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

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "participants" WHERE "participants"."deleted_at" IS NULL`)).WillReturnRows(rows)

	// setup context
	e := echo.New()
	f := factory.NewFactory(db)
	h := NewHandler(f)
	h.Route(e.Group("/participants"))

	echoMock := util_mock.HttpMock{E: e}
	c, rec := echoMock.NewRequest(http.MethodGet, "/participants", nil)
	c.SetPath("")

	h.Get(c)

	var data map[string]interface{}

	err := json.Unmarshal(rec.Body.Bytes(), &data)
	if err != nil {
		t.Fatalf("Error unmarshaling %v", err)
	}

	if rec.Code != http.StatusOK {
		t.Fatalf("got status code %d want 200 and error message : %v", rec.Code, rec.Body)
	}

	if len(data) != 2 {
		t.Fatalf("got %d data, want 2 data", len(data))
	}

}
