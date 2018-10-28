package faithdroid

import (
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/p2p"
)

var (
	protocol = p2p.Protocol{
		Name:    "xchat",
		Version: 42,
		Length:  1,
	}
	s *p2p.Server
)

func (a *MainActivity) OnCreate() {
	ConstraintLayout(a).DeferShow().Size(-2, -2).Append(
		TextView(a).Size(-2, -1).SetId("t").Text("t").onClick(func() {
			Clipboard(a).ClipData(GetTextViewById("t").GetText())
			ShowToast(a, "copied")
		}),
		Button(a).Size(-2, -1).SetId("b").Text("boot").Top2BottomOf("t").OnClick(func() {
			s = newServer(2345, func(p *p2p.Peer, rw p2p.MsgReadWriter) error {
				b, e := rw.ReadMsg()
				if e != nil {
					fmt.Println("readMsg error:", e)
					RunOnUIThread(a, func() {
						ShowToast(a, "read:"+e.Error())
					})
					return e
				}

				var inmsg string
				e = b.Decode(&inmsg)
				if e != nil {
					fmt.Println("decode error:", e)
					return e
				}
				fmt.Println("#in : ", inmsg)
				RunOnUIThread(a, func() {
					ShowToast(a, inmsg)
				})
				return nil
			})
			e := s.Start()
			if e != nil {
				fmt.Println("start error:", e)
				RunOnUIThread(a, func() {
					ShowToast(a, e.Error())
				})
				return
			}
			GetTextViewById("t").Text(s.Self().String())
		}),
		Button(a).Text("stop").Size(-2, -1).Top2BottomOf("b").OnClick(func() {
			s.Stop()
		}))
}
func newServer(port int, f func(p *p2p.Peer, rw p2p.MsgReadWriter) error) *p2p.Server {
	key, e := crypto.GenerateKey()
	if e != nil {
		fmt.Println("genKey error:", e)
		return nil
	}
	protocol.Run = f
	cfg := p2p.Config{
		PrivateKey:      key,
		Name:            "name",
		MaxPeers:        10,
		Protocols:       []p2p.Protocol{},
		EnableMsgEvents: true,
	}
	if port > 0 {
		cfg.ListenAddr = fmt.Sprintf(":%d", port)
	}
	server := &p2p.Server{
		Config: cfg,
	}
	return server
}
