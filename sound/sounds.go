package sound

import (
	"log"
	"os"
	"time"

	"github.com/gopxl/beep/v2"
	"github.com/gopxl/beep/v2/mp3"
	"github.com/gopxl/beep/v2/speaker"
)

func PlaySound(soundName string) {
	f, err := os.Open("sound/mp3/" + soundName + ".mp3")
	if err != nil {
		log.Println(err)
		return
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
}

// sound mapping

// interface for sounds

// Part of the program calls playSound for a different parameter, perhaps 0 if error, 1 if start, 2 if GCP...
// Can introduce a list in a const later.
// Then check which option (switch case)
// Directly play the corresponding sound to the speaker and then end the function as nothing else needs to happen.
