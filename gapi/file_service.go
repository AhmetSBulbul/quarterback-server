package gapi

import (
	"context"
	"database/sql"
	"fmt"
	"mime"
	"os"
	"regexp"

	"github.com/AhmetSBulbul/quarterback-server/pb/filepb"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
)

var (
	publicDir = "public"
)

type FileService struct {
	db *sql.DB
	filepb.UnimplementedFileServiceServer
}

// TODO: Add object storage
func write(data []byte, extension string) (string, error) {
	_guid, _ := uuid.NewUUID()
	guid := _guid.String()

	path := fmt.Sprintf("%s/%s%s", publicDir, guid, extension)
	err := os.WriteFile(path, data, 0644)

	if err != nil {
		return "", err
	}

	return path, nil
}

func (s *FileService) saveFile(ctx context.Context, path string, name string, contentType string) (int, error) {
	query := "INSERT INTO file (path, name, type) VALUES (?, ?, ?);"
	result, err := s.db.Exec(query, path, name, contentType)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (s *FileService) Upload(ctx context.Context, in *filepb.UploadRequest) (*filepb.GetFileResponse, error) {
	var publicHost = os.Getenv("PUBLIC_HOST")
	data := in.GetData()
	name := in.GetName()

	regexp.MustCompile(`\s+`).ReplaceAllString(name, "_")
	extension := regexp.MustCompile(`\.[a-zA-Z0-9]+$`).FindString(name)
	contentType := mime.TypeByExtension(extension)

	path, err := write(data, extension)

	if err != nil {
		return nil, gerr(codes.Internal, err)
	}

	id, err := s.saveFile(ctx, path, name, contentType)

	if err != nil {
		// deleteFile(guid)
		return nil, gerr(codes.Internal, err)
	}

	return &filepb.GetFileResponse{
		Id:          int32(id),
		Path:        publicHost + "/" + path,
		ContentType: contentType,
	}, nil
}

func (s *FileService) GetFile(ctx context.Context, in *filepb.FileId) (*filepb.GetFileResponse, error) {
	row := s.db.QueryRowContext(ctx, "SELECT (id, path, type) FROM file WHERE id = ?", in.Id)

	var file filepb.GetFileResponse

	err := row.Scan(&file.Id, &file.Path, &file.ContentType)

	if err != nil {
		return nil, gerr(codes.Internal, err)
	}

	return &file, nil
}
