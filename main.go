package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/tcpassembly"
	"github.com/ttrnecka/bb2_alert/bb2"
	"github.com/ttrnecka/bb2_alert/config"
)

var debugLog = flag.Bool("debug", false, "If true, the app will log verbose debugging information")

var (
	snapshot_len int32 = 1600
	promiscuous  bool  = true
	err          error
	pcapTimeout  time.Duration = pcap.BlockForever
	bb2Timeout   time.Duration = 90 * time.Second
	handle       *pcap.Handle
	cyanideIPs   []string = []string{"46.105.61.38", "46.105.61.201"}
	cfg          config.Config
)

func main() {
	flag.Parse()
	// f, perr := os.Create("cpu.pprof")
	// if perr != nil {
	// 	panic(perr)
	// }
	// pprof.StartCPUProfile(f)
	// defer pprof.StopCPUProfile()

	cfg.InitConf("conf.yaml")

	bb2 := &bb2.StreamFactory{Cfg: cfg}
	pool := tcpassembly.NewStreamPool(bb2)
	asm := tcpassembly.NewAssembler(pool)

	// flushes assembler for packets older than 2 seconds every 2 seconds - this is because BB2 does not have Keep Alive connection and assembler expect SYN to start, which never comes
	go func() {
		for {
			time.Sleep(time.Second * 2)
			asm.FlushWithOptions(tcpassembly.FlushOptions{CloseAll: false, T: time.Now().Add(-time.Second * 2)})
		}
	}()

	devices, err := pcap.FindAllDevs()

	if err != nil {
		log.Fatal(err)
	}

	wg := new(sync.WaitGroup)
	// this will listen on every iface available, however should timeout if packet will not arrive in 90s, that should leave only one open
	for _, device := range devices {
		wg.Add(1)
		handle, err := pcap.OpenLive(device.Name, snapshot_len, promiscuous,
			pcapTimeout)
		if err != nil {
			panic(err)
		}
		defer handle.Close()

		filter := fmt.Sprintf("src host %s or src host %s", cyanideIPs[0], cyanideIPs[1])
		if err := handle.SetBPFFilter(filter); err != nil {
			panic(err)
		}

		packets := gopacket.NewPacketSource(
			handle, handle.LinkType()).Packets()

		go func(packet chan gopacket.Packet, i pcap.Interface) {
			timedout := false
			found := false
			for {
				select {
				case pkt := <-packets:
					// ctrl+c causes this to be nil
					if pkt != nil {
						if !found {
							if *debugLog {
								log.Printf("BB2 found on %s\n", i.Description)
							}
							found = true
						}
						if tcp := pkt.Layer(layers.LayerTypeTCP); tcp != nil {
							asm.AssembleWithTimestamp(
								pkt.TransportLayer().TransportFlow(),
								tcp.(*layers.TCP),
								pkt.Metadata().Timestamp)
						}
					}
				case <-time.After(bb2Timeout):
					if *debugLog {
						log.Printf("BB2 not found on %s or timed out, closing pcap\n", i.Description)
					}
					timedout = true
				}
				if timedout {
					break
				}
			}
			wg.Done()
		}(packets, device)
	}
	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan struct{})
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		<-signalChan
		fmt.Println("Received an interrupt, stopping service...")
		close(cleanupDone)
	}()
	go func() {
		wg.Wait()
		close(cleanupDone)
	}()
	if *debugLog {
		log.Println("Looking for BB2 communication")
	}
	fmt.Println("Press ctrl+c to close")
	<-cleanupDone

}
