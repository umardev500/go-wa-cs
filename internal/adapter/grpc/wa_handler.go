package grpc

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/umardev500/chat/api/proto"
	"github.com/umardev500/chat/configs"
	"github.com/umardev500/chat/internal/domain"
	grpcManager "github.com/umardev500/chat/internal/grpc"
	"github.com/umardev500/chat/internal/repository"
	"github.com/umardev500/chat/internal/usecase"
	"github.com/umardev500/chat/pkg/utils"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type WaHandler struct {
	repo   repository.WaRepo
	chatUc usecase.ChatUsecase
	proto.UnimplementedWhatsAppServiceServer
}

var (
	mu sync.Mutex
)

func NewWaHandler(repo repository.WaRepo, chatUc usecase.ChatUsecase) *WaHandler {
	return &WaHandler{
		repo:   repo,
		chatUc: chatUc,
	}
}

func (w *WaHandler) getCsid(remoteJid string) (string, error) {
	csid, err := w.repo.FindActiveChat(remoteJid)
	if err != nil && err != mongo.ErrNoDocuments {
		return "", err
	}

	return csid, nil
}

func (w *WaHandler) SubscribeProfilePic(stream proto.WhatsAppService_SubscribeProfilePicServer) error {
	// Create a new client with its own message queue
	client := &grpcManager.PicClient{
		Stream:     stream,
		MsgChan:    make(chan *proto.SubscribeProfilePicResponse, 1), // Buffered channel
		ResultChan: make(chan string, 1),                             // Buffered channel
	}

	// Register the client connection
	grpcManager.AddPicClient(client)

	log.Info().Msgf("📤 Added pic client: %v", client.Stream)

	// Start a dedicated sender goroutine for this client
	go func(c *grpcManager.PicClient) {
		for msg := range c.MsgChan {
			if err := c.Stream.Send(msg); err != nil {
				log.Err(err).Msg("Failed to send message, removing client")
				grpcManager.RemovePicClient(c)
			}
		}
	}(client)

	for {
		msg, err := stream.Recv()
		if err != nil {
			log.Err(err).Msg("failed to receive pic message")
			log.Info().Msg("remove pic client")
			grpcManager.RemovePicClient(client)
			return err
		}

		client.ResultChan <- msg.Url
		log.Info().Msgf("📤 Pic message sent: %v", msg)
	}
}

func (w *WaHandler) SubscribePresense(stream proto.WhatsAppService_SubscribePresenseServer) error {
	// Create a new client with its own message queue
	client := &grpcManager.PresenceClient{
		Stream:  stream,
		MsgChan: make(chan *proto.SubscribePresenseResponse, 1), // Buffered channel
	}

	// Register the client connection
	grpcManager.AddPresenceClient(client)

	log.Info().Msgf("📤 Added presence client: %v", client.Stream)

	// Start a dedicated sender goroutine for this client
	go func(c *grpcManager.PresenceClient) {
		for msg := range c.MsgChan {
			if err := c.Stream.Send(msg); err != nil {
				log.Err(err).Msg("Failed to send message, removing client")
				grpcManager.RemovePresenceClient(c)
				return
			}
		}
	}(client)

	for {
		msg, err := stream.Recv()
		if err != nil {
			log.Err(err).Msg("failed to receive presense message")
			log.Info().Msg("remove presence client")
			grpcManager.RemovePresenceClient(client)
			return err
		}

		log.Info().Msgf("Received message: %s", msg)
	}
}

func (w *WaHandler) SendTextMessage(ctx context.Context, req *proto.TextMessageRequest) (*proto.CommonMessageResponse, error) {
	// Print the full request
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}
	log.Info().Msgf("Received request: %s", string(jsonData))

	// find the csid who active chat with the user
	csid, err := w.getCsid(req.Metadata.RemoteJid)
	if err != nil {
		return nil, err
	}

	if csid == "" {
		// TODO: do logic for selecting csid
		// if not cs is active chating then assign new cs
		csid = "xyz"
	}

	err = w.chatUc.PushChat(ctx, csid, &domain.PushChat{
		Data: domain.PushChatData{
			TextMessage: req,
		},
	})
	if err != nil {
		return nil, err
	}

	return &proto.CommonMessageResponse{
		Status: "success",
	}, nil
}

func (w *WaHandler) SendExtendedTextMessage(ctx context.Context, req *proto.ExtendedTextMessageRequest) (*proto.CommonMessageResponse, error) {
	// Print the full request
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}
	log.Info().Msgf("Received request: %s", string(jsonData))

	w.repo.SaveMessage(req)

	return &proto.CommonMessageResponse{
		Status: "success",
	}, nil
}

func (w *WaHandler) UploadMedia(stream proto.WhatsAppService_UploadMediaServer) error {
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
			response := &proto.MediaUploadResponse{
				FileUrl:  "https://yourserver.com/" + fileName,
				FilePath: filePath,
			}

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

func (w *WaHandler) StoreFileMetadata(ctx context.Context, req *proto.FileMetadataRequest) (*proto.FileMetadataResponse, error) {
	fileName := req.FileName
	ext := filepath.Ext(fileName)

	// Define the new file path with message id
	newFilePath := filepath.Join(filepath.Dir(fileName), fmt.Sprintf("%s%s", req.Metadata.Id, ext))

	// Rename the file
	if err := os.Rename(fileName, newFilePath); err != nil {
		return nil, fmt.Errorf("failed to rename file: %w", err)
	}

	log.Info().Msgf("Renamed file: %s", newFilePath)
	req.FileName = newFilePath

	// Print the full request
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}
	log.Info().Msgf("Received request: %s", string(jsonData))

	w.repo.SaveMessage(req)

	return &proto.FileMetadataResponse{
		Status: "success",
	}, nil
}

func (w *WaHandler) SendOnlineUser(ctx context.Context, req *proto.SendOnlineUserRequest) (*proto.CommonMessageResponse, error) {
	wsClients := utils.WsGetClients()

	var status = configs.PresenseOnlineText
	if req.Presence == string(configs.PresenseOffline) {
		status = configs.PresenseOfflineText
	}

	var data = map[string]interface{}{
		"mt": "status",
		"data": []map[string]interface{}{
			{
				"status":    status,
				"remotejid": req.Jid,
			},
		},
	}

	for _, client := range wsClients {
		mu.Lock()
		err := client.WriteJSON(data)
		mu.Unlock()

		if err != nil {
			log.Err(err).Msg("failed to write json to the websocket client")
			return nil, err
		}
	}

	return &proto.CommonMessageResponse{
		Status: "success",
	}, nil
}

func (w *WaHandler) SendTyping(ctx context.Context, req *proto.SendTypingRequest) (*proto.CommonMessageResponse, error) {
	wsClients := utils.WsGetClients()

	var data = map[string]interface{}{
		"mt": configs.MessageTypeTyping,
		"data": []map[string]interface{}{
			{
				"presence":  req.Presence,
				"remotejid": req.Jid,
			},
		},
	}

	for _, client := range wsClients {
		err := client.WriteJSON(data)
		if err != nil {
			log.Err(err).Msg("failed to write typing json data to the websocket client")
			return nil, err
		}
	}

	return &proto.CommonMessageResponse{
		Status: "success",
	}, nil
}
