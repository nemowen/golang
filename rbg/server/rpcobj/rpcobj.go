package rpcobj

import (
	"log"
	"os"
	"sync"
)

type Obj struct {
	Date                string
	Time                string
	ID                  string
	Type                string
	FaceValue           int
	Version             int
	SerialNumberInTimes int
	Number              string
	Ima                 []byte
}

const (
	BMP_SAVE_PATH string = "D:/"
)

var w sync.WaitGroup

func (o *Obj) SendToServer(obj Obj, replay *string) error {
	saveObj(obj)
	*replay = "ok"
	return nil
}

func saveObj(obj Obj) {
	w.Add(2)
	go func() {
		s := obj.Number + "|" + obj.Date + "|" + string(obj.ID)
		saveTo, e := os.Create(BMP_SAVE_PATH + obj.ID + ".txt")
		if e != nil {
			log.Println(e)
		}
		defer saveTo.Close()
		_, err := saveTo.WriteString(s)
		if err != nil {
			log.Println("保存信息失败", obj.ID, err)
		}
		w.Done()
	}()

	go func() {
		f, err := os.Create(BMP_SAVE_PATH + obj.ID + ".bmp")
		if err != nil {
			log.Println("保存bmp失败：", obj.ID)
		}
		defer f.Close()
		f.Write(obj.Ima)
		w.Done()
	}()
	w.Wait()
}
