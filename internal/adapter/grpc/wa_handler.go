package grpc

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/umardev500/chat/api/proto"
	"github.com/umardev500/chat/pkg/utils"
)

type waHandler struct {
	proto.UnimplementedWhatsAppServiceServer
}

func NewWaHandler() *waHandler {
	return &waHandler{}
}

func (w *waHandler) SendTextMessage(ctx context.Context, req *proto.TextMessageRequest) (*proto.CommonMessageResponse, error) {
	return &proto.CommonMessageResponse{
		Status: "success",
	}, nil
}

func (w *waHandler) SendExtendedTextMessage(ctx context.Context, req *proto.ExtendedTextMessageRequest) (*proto.CommonMessageResponse, error) {
	return &proto.CommonMessageResponse{
		Status: "success",
	}, nil
}

func (w *waHandler) UploadMedia(stream proto.WhatsAppService_UploadMediaServer) error {
	// Define media storage directory
	dirPath := "uploads"
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	var mimeType, filePath, fileName string
	var file *os.File

	fmt.Println("📂 Receiving media...")

	// Read the file chunks from the stream
	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("✅ Upload complete:", filePath)
			response := &proto.MediaUploadResponse{FileUrl: "https://yourserver.com/" + fileName}

			return stream.SendAndClose(response)
		}
		if err != nil {
			return fmt.Errorf("failed to receive chunk: %w", err)
		}

		if mimeType == "" {
			mimeType = http.DetectContentType(chunk.Chunk[:512])
			// Map MIME type to file extension
			ext, exists := utils.MimeToExtension[mimeType]
			if !exists {
				ext = ".bin" // Default file extension if MIME type is unknown
			}

			// Create a unique file
			fileName = fmt.Sprintf("%d_uploaded%s", time.Now().UTC().Unix(), ext) // Unique file name
			filePath = filepath.Join(dirPath, fileName)

			// Open file for writing
			file, err = os.Create(filePath)
			if err != nil {
				return fmt.Errorf("failed to create file: %w", err)
			}
			defer file.Close()
		}

		// Save the chunk to the file
		_, err = file.Write(chunk.Chunk)
		if err != nil {
			return fmt.Errorf("failed to write chunk: %w", err)
		}
	}
}

func (w *waHandler) StoreFileMetadata(ctx context.Context, req *proto.FileMetadataRequest) (*proto.FileMetadataResponse, error) {
	return &proto.FileMetadataResponse{
		Status: "success",
	}, nil
}
