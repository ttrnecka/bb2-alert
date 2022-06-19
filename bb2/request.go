package bb2

import (
	"bufio"
	"bytes"
	"compress/zlib"
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/google/gopacket"
	"github.com/google/gopacket/tcpassembly"
	"github.com/google/gopacket/tcpassembly/tcpreader"
	"github.com/ttrnecka/bb2_alert/config"
	"github.com/ttrnecka/bb2_alert/discord"
)

var xmlLogger *log.Logger

func init() {
	f, err := os.OpenFile("data.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}
	xmlLogger = log.New(f, "", log.LstdFlags)
}

var debugLog = flag.Bool("bb2_debug_log", false, "If true, the bb2 library will log verbose debugging information")

// StreamFactory is a Stream Factory for gopacker assembler that calls request/packet handler
// bb2 := &bb2.StreamFactory{}
// pool := tcpassembly.NewStreamPool(bb2)
// asm := tcpassembly.NewAssembler(pool)
type StreamFactory struct {
	Cfg config.Config
}

// New meets assembler StreamFactory interface
func (f *StreamFactory) New(a, b gopacket.Flow) tcpassembly.Stream {
	r := tcpreader.NewReaderStream()
	go handlePackets(&r, f.Cfg)
	return &r
}

// handlePackets
// r io.Reader will contain payload data from ApplicationLayer
func handlePackets(r io.Reader, cfg config.Config) {
	var readBuffer = make([]byte, 1024)
	var Header XmlHeader
	var messageBuffer bytes.Buffer
	var dataBuffer bytes.Buffer
	var header_start, header_end bool = false, false

	sniff := bufio.NewReader(r)

	wh := discord.WebHook{ServerId: cfg.Discord.ServerID, Token: cfg.Discord.Token, UserId: cfg.Discord.UserID}

	for {
		n, err := sniff.Read(readBuffer)
		if err == io.EOF {
			if *debugLog {
				log.Println("no more data to read")
			}
			break
		}
		if err != nil {
			panic(err)
		}
		messageBuffer.Write(readBuffer[:n])

		if *debugLog {
			log.Printf("buffer:\n%s\n", &messageBuffer)
		}

		if !header_start && bytes.Contains(messageBuffer.Bytes(), []byte(HEADER_START)) {
			if *debugLog {
				log.Println("found header start")
			}
			start := bytes.Index(messageBuffer.Bytes(), []byte(HEADER_START))
			messageBuffer.Next(start)
			header_start = true
		}
		if !header_end && bytes.Contains(messageBuffer.Bytes(), []byte(HEADER_END)) {
			if *debugLog {
				log.Println("found header end")
			}
			end := bytes.Index(messageBuffer.Bytes(), []byte(HEADER_END))
			header := messageBuffer.Next(end + len(HEADER_END))
			data := fmt.Sprintf("header:\n%s\n", string(header))
			if *debugLog {
				log.Printf(data)
			}

			xmlLogger.Printf(data)
			err = xml.Unmarshal(header, &Header)
			if err != nil {
				panic(err)
			}
			if *debugLog {
				log.Printf("%+v\n", Header)
			}
			dataBuffer.Reset()
			header_end = true
		}
		if header_start && header_end {
			if *debugLog {
				log.Printf("buffer size: %d\n", messageBuffer.Len())
			}
			if messageBuffer.Len() >= Header.Data.Size {
				dataBuffer.Write(messageBuffer.Next(Header.Data.Size))
				header_start, header_end = false, false
				var payload string
				if *debugLog {
					fmt.Println("data:")
				}
				if Header.Data.Zipped == "yes" {
					gzr, err := zlib.NewReader(&dataBuffer)
					if err != nil {
						panic(err)
					}
					output, err := ioutil.ReadAll(gzr)
					if err != nil {
						panic(err)
					}
					payload := string(output)
					payload = strings.ReplaceAll(payload, "&#060;", "<")
					payload = strings.ReplaceAll(payload, "&#062;", ">")
					if *debugLog {
						log.Println(payload)
					}
					xmlLogger.Printf(payload)
				} else {
					payload = dataBuffer.String()
					if *debugLog {
						log.Println(dataBuffer.String())
					}
					xmlLogger.Printf(payload)
				}
				// header and full message is read, can parse and process
				if Header.Data.MessageName == MESSAGE_NAME_MATCH_FOUND && cfg.BB2.MatchFoundAlarm {
					var MFM NotificationAutomatchOpponentFound
					err = xml.Unmarshal([]byte(payload), &MFM)
					if err != nil {
						panic(err)
					}
					opponent := fmt.Sprintf("Match found: Coach %s, Team %s, %s, TV %s", MFM.OpponentName,
						Races[MFM.Specific.MessageData.BBGameSession.RowTeamOpponent.IdRaces],
						MFM.Specific.MessageData.BBGameSession.RowTeamOpponent.Name,
						MFM.Specific.MessageData.BBGameSession.RowTeamOpponent.Value)
					if *debugLog {
						log.Printf("Sending discord message: %s", opponent)
					}
					wh.Send(opponent)
				}
				// if Header.Data.MessageName == "KeepAlive" {
				// 	wh.Send("keepalive")
				// }
			}
		}
	}
}
