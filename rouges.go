package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-toast/toast"
	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto/v2"
)

type RogueArea struct {
	Left        int
	Top         int
	Right       int
	Bottom      int
	Description string
}

var rogueAreas = []RogueArea{
	RogueArea{Left: 231, Top: 130, Right: 240, Bottom: 140, Description: "Below puri tulang"},
	RogueArea{Left: 169, Top: 177, Right: 184, Bottom: 272, Description: "Left from cartref"},
	RogueArea{Left: 191, Top: 266, Right: 210, Bottom: 276, Description: "Right from cartref"},
	RogueArea{Left: 198, Top: 95, Right: 209, Bottom: 105, Description: "Top dark"},
	RogueArea{Left: 84, Top: 170, Right: 108, Bottom: 181, Description: "Top neutrals"},
	RogueArea{Left: 97, Top: 187, Right: 111, Bottom: 200, Description: "Bottom neutrals"},
}

func ProcessRouges(body []byte) {
	var hlMap HeroeslandMap

	if !bytes.Contains(body, []byte(`"a":"all"`)) {
		return
	}

	err := json.Unmarshal(body, &hlMap)
	if err != nil {
		//fmt.Println("error", err)
	}

	if len(hlMap) < 1 || hlMap[0].A != "all" {
		return
	}

	for _, person := range hlMap[0].P.D {
		isHero := person.Tid > 1300 && person.Tid < 1400

		isRealPerson := person.T == 3

		isInBoat := person.Tw == 1

		isInRogueArea := false
		descriptionArea := ""
		for _, rogueArea := range rogueAreas {
			isRogueAreaLocal := person.X >= rogueArea.Left && person.X <= rogueArea.Right && person.Y >= rogueArea.Top && person.Y <= rogueArea.Bottom

			if isRogueAreaLocal {
				isInRogueArea = true
				descriptionArea = rogueArea.Description
				break
			}
		}

		if isHero && !isRealPerson && !isInBoat && isInRogueArea {
			/* play() */
			notification := toast.Notification{
				AppID:   "Microsoft.Windows.Shell.RunDialog",
				Title:   "title",
				Message: msg,
				Icon:    "C:\\path\\to\\your\\logo.png", // The file must exist
				Actions: []toast.Action{},
			}
			err := notification.Push()
			if err != nil {
				log.Fatalln(err)
			}
			play()

			fmt.Println(msg)
		}
	}
}

func play() error {
	f, err := os.Open("beep.mp3")
	if err != nil {
		return err
	}
	defer f.Close()

	d, err := mp3.NewDecoder(f)
	if err != nil {
		return err
	}

	c, ready, err := oto.NewContext(d.SampleRate(), 2, 2)
	if err != nil {
		return err
	}
	<-ready

	p := c.NewPlayer(d)
	defer p.Close()
	p.Play()

	fmt.Printf("Length: %d[bytes]\n", d.Length())
	for {
		time.Sleep(time.Second)
		if !p.IsPlaying() {
			break
		}
	}

	return nil
}
