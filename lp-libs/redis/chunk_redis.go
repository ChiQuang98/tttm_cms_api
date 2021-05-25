package redis

import (
	"encoding/json"
	"fmt"
	"strconv"

	"tttm_cms_api/lp-libs/models"
)

var chunkInfoFields = []string{
	Record_Chunk_Index,
	Record_Chunk_State,
	Record_Chunk_BufferSize,
	Record_Chunk_TimeLive,
	Record_Chunk_AudioLink,
	Record_Chunk_Priority,
	Record_Chunk_AudioCodec,
	Record_Chunk_ReceiptType,
	Record_Chunk_ReceiptIds,
}

func UpdateLiveStreamChunkInfo(chunk *models.RecLiveStream) error {
	chunkIndex := string(chunk.RecChunkIndex)
	chunkAudioCodec := string(chunk.RecAudioCodec)
	chunkState := string(chunk.RecState)
	chunkBuffer := string(chunk.RecBufferSize)
	chunkTimeLive := string(chunk.RecLiveTime)
	chunkPriority := string(chunk.RecPriority)
	chunkReceiptType := string(chunk.RecReceiptType)
	chunkReceiptIds, err := json.Marshal(chunk.RecReceiptIds)
	if err != nil {
		return err
	}
	chunkId := fmt.Sprintf("%s_%d", strconv.FormatInt(chunk.RecId, 10), chunk.RecChunkIndex)

	return lClient.HMSet(chunkId,
		Record_Chunk_Index, chunkIndex,
		Record_Chunk_State, chunkState,
		Record_Chunk_BufferSize, chunkBuffer,
		Record_Chunk_TimeLive, chunkTimeLive,
		Record_Chunk_AudioLink, chunk.RecAudio,
		Record_Chunk_AudioCodec, chunkAudioCodec,
		Record_Chunk_Priority, chunkPriority,
		Record_Chunk_ReceiptType, chunkReceiptType,
		Record_Chunk_ReceiptIds, string(chunkReceiptIds),
	).Err()
}
