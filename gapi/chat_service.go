package gapi

import (
	"context"
	"database/sql"

	"github.com/AhmetSBulbul/quarterback-server/pb/chatpb"
	"google.golang.org/grpc/codes"
)

type ChatService struct {
	db *sql.DB
	// Manages connections
	channel map[int32]chan *chatpb.ChatMessage
	chatpb.UnimplementedChatServiceServer
}

func (s *ChatService) GetChatMessages(ctx context.Context, _ *chatpb.Channel) (*chatpb.ChatMessageListResponse, error) {
	// todo extract from token
	sub_id := getUserIdFromCtx(ctx)
	query := "SELECT ID, senderID, receiverID, content, isReceived FROM chat WHERE senderID = ? OR receiverID = ?"

	rows, err := s.db.QueryContext(ctx, query, sub_id, sub_id)

	if err != nil {
		return nil, gerr(codes.Internal, err)
	}

	defer rows.Close()

	var messages []*chatpb.ChatMessage

	for rows.Next() {
		var msg chatpb.ChatMessage
		err := rows.Scan(&msg.Id, &msg.SenderId, &msg.ReceiverId, &msg.Content, &msg.IsReceived)
		if err != nil {
			return nil, gerr(codes.Internal, err)
		}

		messages = append(messages, &msg)
	}

	query = "UPDATE chat SET isReceived=1 WHERE receiverID = ?"
	_, err = s.db.Exec(query, sub_id)
	if err != nil {
		return nil, gerr(codes.Internal, err)
	}

	return &chatpb.ChatMessageListResponse{
		Messages: messages,
	}, nil

}

// connect with id that extracted from jwt token
func (s *ChatService) Connect(ch *chatpb.Channel, msgStream chatpb.ChatService_ConnectServer) error {
	sub_id := ch.Id
	msgChannel := make(chan *chatpb.ChatMessage)
	s.channel[int32(sub_id)] = msgChannel

	for {
		select {
		case <-msgStream.Context().Done():
			delete(s.channel, int32(sub_id))
			return nil
		case msg := <-msgChannel:
			msgStream.Send(msg)
		}
	}
}

func (s *ChatService) SendMessage(ctx context.Context, msg *chatpb.ChatMessageRequest) (*chatpb.ChatMessage, error) {
	// TODO: extract from token
	sub_id := getUserIdFromCtx(ctx)
	receiverChan, ok := s.channel[msg.ReceiverId]

	query := "INSERT INTO chat (senderID, receiverID, content, isReceived) VALUES (?, ?, ?, ?)"
	res, err := s.db.ExecContext(ctx, query, sub_id, msg.ReceiverId, msg.Content, ok)
	if err != nil {
		return nil, gerr(codes.Internal, err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, gerr(codes.Internal, err)
	}

	new_msg := &chatpb.ChatMessage{
		Id:         int32(id),
		SenderId:   int32(sub_id),
		ReceiverId: msg.ReceiverId,
		Content:    msg.Content,
		IsReceived: ok,
	}

	if ok {
		receiverChan <- new_msg
	}

	return new_msg, nil
}
