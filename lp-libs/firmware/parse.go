package firmware

import (
	"bufio"
	"encoding/binary"
	"encoding/hex"
	"hash/crc32"
	"strconv"
	"strings"

	"github.com/kierdavis/ihex-go"
)

const (
	HEX_ADDR_APPFIRM_START   = 0x08008000
	HEX_ADDR_APPFIRM_VERSION = 0x08010004
	HEX_ADDR_APPFIRM_MARK    = 0x08010000
	HEX_BLOCK_SIZE           = 1024
)

type FirmwareInfo struct {
	Version   uint32
	Size      uint32
	CRC32     uint32
	DataBlock []HexBlock
	Data      []byte
}

type HexBlock struct {
	StartAddress uint32
	Data         [HEX_BLOCK_SIZE]byte
}

func (block *HexBlock) Init() {
	for i, _ := range block.Data {
		block.Data[i] = 0xff
	}
}

func GetFirmwareInfoFromString(hexString string) (*FirmwareInfo, error) {
	info := &FirmwareInfo{}
	scanner := bufio.NewScanner(strings.NewReader(hexString))
	data := []byte{}

	var extAddr, addr uint32
	for scanner.Scan() {
		line := scanner.Text()
		record, err := ihex.DecodeRecordHex(line)
		if err == nil {
			if record.Type == ihex.Data {
				bytesAddr := make([]byte, 2)
				binary.BigEndian.PutUint16(bytesAddr, record.Address)
				i64, _ := strconv.ParseInt(hex.EncodeToString(bytesAddr), 16, 0)
				addr = extAddr<<16 + uint32(i64)
				if addr >= HEX_ADDR_APPFIRM_START {
					data = append(data, record.Data...)
				}
			} else {
				if record.Type == ihex.ExtendedLinearAddress {
					i64, _ := strconv.ParseInt(hex.EncodeToString(record.Data), 16, 0)
					extAddr = uint32(i64)
				}
			}
		} else {
			return info, err
		}
	}
	//version
	indexVerison := HEX_ADDR_APPFIRM_VERSION - HEX_ADDR_APPFIRM_START
	version := binary.LittleEndian.Uint32(data[indexVerison : indexVerison+4])
	info.Version = version

	//block
	blockCnt := (len(data) + HEX_BLOCK_SIZE - 1) / HEX_BLOCK_SIZE
	for i := 0; i < blockCnt; i++ {
		indexBlock := i * HEX_BLOCK_SIZE
		block := HexBlock{}
		block.Init()

		for j := 0; j < HEX_BLOCK_SIZE; j++ {
			if indexBlock+j < len(data) {
				block.Data[j] = data[indexBlock+j]

			}
		}
		info.DataBlock = append(info.DataBlock, block)
	}
	//crc
	crcBuf := []byte{}
	for _, block := range info.DataBlock {
		info.Data = append(info.Data, block.Data[:]...)
		crcBuf = append(crcBuf, block.Data[:]...)
	}
	info.Size = uint32(len(info.Data))
	info.CRC32 = crc32.ChecksumIEEE(crcBuf)
	return info, nil
}
