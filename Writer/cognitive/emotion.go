package cognitive

import (
	"github.com/ahmdrz/microsoft-cognitive-services/emotion"
)

func RecognizeEmotion(img string) ([]emotion.EmotionDetail, error) {
	emo, err := emotion.New("cf2602463bd1438fb3ce5f799fd4c5da")
	if err != nil {
		panic(err)
	}

	return emo.Recognize(img)
}
