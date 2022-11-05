package cmd

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"sync"

	"github.com/spf13/cobra"

	"github.com/myhro/feeds/autossegredos"
	"github.com/myhro/feeds/copasa"
	"github.com/myhro/feeds/errormap"
	"github.com/myhro/feeds/liquipedia"
)

type Feed struct {
	Command string
	XML     func() (string, error)
}

func (f *Feed) Generate() {
	log.Print("Generating feed ", f.Command)

	xml, err := f.XML()
	if err != nil {
		log.Printf("%v: %v", f.Command, err)
		return
	}

	list := errormap.List(f.Command)
	if len(list) > 0 {
		for _, err := range list {
			log.Printf("%v: %v", f.Command, err)
		}

		return
	}

	err = f.Save(xml)
	if err != nil {
		log.Printf("%v: %v", f.Command, err)
		return
	}

	log.Print("Finished feed ", f.Command)
}

func (f *Feed) Save(xml string) error {
	folder := "dist/"

	err := os.MkdirAll(folder, 0755)
	if err != nil {
		return fmt.Errorf("os.MkdirAll: %w", err)
	}

	file := path.Join(folder, f.Command+".xml")

	dest, err := os.Create(file)
	if err != nil {
		return fmt.Errorf("os.Create: %w", err)
	}
	defer dest.Close()

	_, err = io.WriteString(dest, xml)
	if err != nil {
		return fmt.Errorf("io.WriteString: %w", err)
	}

	return nil
}

func Generate(cmd *cobra.Command, args []string) {
	var wg sync.WaitGroup

	feeds, err := cmd.Flags().GetStringSlice("feed")
	if err != nil {
		log.Fatal("cmd.Flags().GetStringSlice: ", err)
	}

	wg.Add(len(feeds))

	for _, c := range feeds {
		switch c {
		case autossegredos.Command:
			go func() {
				f := Feed{
					Command: autossegredos.Command,
					XML:     autossegredos.XML,
				}
				f.Generate()
				wg.Done()
			}()
		case copasa.Command:
			go func() {
				f := Feed{
					Command: copasa.Command,
					XML:     copasa.XML,
				}
				f.Generate()
				wg.Done()
			}()
		case liquipedia.Command:
			go func() {
				f := Feed{
					Command: liquipedia.Command,
					XML:     liquipedia.XML,
				}
				f.Generate()
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	log.Print("Waiting for feeds to be generated")
	wg.Wait()
	log.Print("Done")
}
