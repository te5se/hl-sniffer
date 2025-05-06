package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
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

type CacheLocal struct {
	Cache map[string]time.Time
}

var localCache = CacheLocal{
	Cache: make(map[string]time.Time),
}

func (c *CacheLocal) ShouldNotify(x int, y int) bool {
	v, ok := c.Cache[string(x)+string(y)]
	if !ok {
		c.Cache[string(x)+string(y)] = time.Now()
		return true
	}

	if time.Since(v) > time.Minute {
		c.Cache[string(x)+string(y)] = time.Now()
		return true
	}

	return false
}

func (s *Sniffer) ProcessRouges(body []byte) {
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

		isOverlord := person.Tid == 1338

		isDragon := person.Tid == 1554

		descriptionArea := ""
		for _, rogueArea := range rogueAreas {
			isRogueAreaLocal := person.X >= rogueArea.Left && person.X <= rogueArea.Right && person.Y >= rogueArea.Top && person.Y <= rogueArea.Bottom

			if isRogueAreaLocal {
				isInRogueArea = true
				descriptionArea = rogueArea.Description
				break
			}
		}

		if isHero && !isRealPerson && !isInBoat && isInRogueArea && localCache.ShouldNotify(person.X, person.Y) {
			notifyMsg := fmt.Sprintf("Rogue hero at %v:%v, %v", person.X, person.Y, descriptionArea)

			/* notification := toast.Notification{
				AppID:   "Microsoft.Windows.Shell.RunDialog",
				Title:   "title",
				Message: notifyMsg,
				Icon:    "C:\\path\\to\\your\\logo.png", // The file must exist
				Actions: []toast.Action{},
			}
			err := notification.Push()
			if err != nil {
				log.Fatalln(err)
			} */

			fmt.Println(notifyMsg)

			err = s.tgService.Notify(notifyMsg)
			if err != nil {
				fmt.Println("error", err.Error())
			}
		}

		if isHero && !isRealPerson && isOverlord && person.X < 100 && person.Y < 100 && localCache.ShouldNotify(person.X, person.Y) {
			/* play() */
			notifyMsg := fmt.Sprint("Mutare is up")

			/* notification := toast.Notification{
				AppID:   "Microsoft.Windows.Shell.RunDialog",
				Title:   "title",
				Message: notifyMsg,
				Icon:    "C:\\path\\to\\your\\logo.png", // The file must exist
				Actions: []toast.Action{},
			}
			err := notification.Push()
			if err != nil {
				log.Fatalln(err)
			} */

			fmt.Println(notifyMsg)

			err = s.tgService.Notify(notifyMsg)
			if err != nil {
				fmt.Println("error", err.Error())
			}
		}

		if isDragon && localCache.ShouldNotify(person.X, person.Y) {
			notifyMsg := fmt.Sprintf("Dragon at at %v:%v", person.X, person.Y)

			/* play() */
			/* notification := toast.Notification{
				AppID:   "Microsoft.Windows.Shell.RunDialog",
				Title:   "title",
				Message: notifyMsg,
				Icon:    "C:\\path\\to\\your\\logo.png", // The file must exist
				Actions: []toast.Action{},
			}
			err := notification.Push()
			if err != nil {
				log.Fatalln(err)
			} */

			fmt.Printf(notifyMsg)

			err = s.tgService.Notify(notifyMsg)
			if err != nil {
				fmt.Println("error", err.Error())
			}
		}
	}
}
