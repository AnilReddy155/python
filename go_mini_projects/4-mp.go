package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-message/mail"
	"github.com/jung-kurt/gofpdf"
)

func main() {
	// log.Println("Connecting to server...")
	var body string
	// Connect to server
	c, err := client.DialTLS("imap.gmail.com:993", nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected")

	// Don't forget to logout
	defer c.Logout()

	// Login
	if err := c.Login("reddyanil1445@gmail.com", "October@462"); err != nil {
		log.Fatal(err)
	}
	log.Println("Logged in")

	// Select INBOX
	mbox, err := c.Select("INBOX", false)
	if err != nil {
		log.Fatal(err)
	}

	// Get the last message
	if mbox.Messages == 0 {
		log.Fatal("No message in mailbox")
	}
	seqSet := new(imap.SeqSet)
	seqSet.AddNum(mbox.Messages)

	// Get the whole message body
	var section imap.BodySectionName
	items := []imap.FetchItem{section.FetchItem()}

	messages := make(chan *imap.Message, 1)
	go func() {
		if err := c.Fetch(seqSet, items, messages); err != nil {
			log.Fatal(err)
		}
	}()

	msg := <-messages
	if msg == nil {
		log.Fatal("Server didn't returned message")
	}

	r := msg.GetBody(&section)
	if r == nil {
		log.Fatal("Server didn't returned message body")
	}

	// Create a new mail reader
	mr, err := mail.CreateReader(r)
	if err != nil {
		log.Fatal(err)
	}

	// Print some info about the message
	header := mr.Header
	date := header.Get("Date")

	if err == nil {
		log.Println("Date:", date)
	}
	from := header.Get("From")
	if err == nil {
		log.Println("From:", from)
	}

	to := header.Get("To")

	subject, err := header.Subject()
	if err == nil {
		log.Println("Subject:", subject)
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	// pdf.Cell(100, 10, date)
	// pdf.Cell(100, 10, from)
	// pdf.Cell(100, 10, to)
	// pdf.Cell(100, 10, sub)
	pdf.CellFormat(100, 10, "Date : "+date, "0", 1, "L", false, 0, "")
	pdf.CellFormat(100, 10, "From : "+from, "0", 1, "L", false, 0, "")
	pdf.CellFormat(100, 10, "To : "+to, "0", 1, "L", false, 0, "")
	pdf.CellFormat(100, 10, "Subject : "+subject, "0", 1, "L", false, 0, "")

	// Process each message's part
	for {
		p, err := mr.NextPart()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		switch h := p.Header.(type) {
		case *mail.InlineHeader:
			// This is the message's text (can be plain-text or HTML)
			t, _, _ := h.ContentType()
			if t == "text/plain" {
				b, _ := ioutil.ReadAll(p.Body)
				body = string(b)
				log.Println("Got text: %v", body)
			}
		case *mail.AttachmentHeader:
			// This is an attachment
			filename, err := h.Filename()
			if err != nil {
				log.Printf("ERROR] imap: %v, skip", err)
				continue
			}
			log.Printf("[INFO] imap: found attachment: %v", filename)
			outFile := filepath.Join(".", filename)

			f, err := os.Create(outFile)
			if err != nil {
				log.Printf("[ERROR] imap: %v, skip", err)
				continue
			}
			_, err = io.Copy(f, p.Body)
			if err != nil {
				log.Printf("[ERROR] imap: %v, skip", err)
				continue
			}
			f.Close()
			file_path, err := filepath.Abs(filename)
			if err != nil {
				log.Printf("Not able to find Absolute path")
			}
			fmt.Println(file_path)

			pdf.CellFormat(100, 10, file_path, "0", 1, "L", false, 0, filename)
		}
	}

	// Generating PDF.
	pdf.CellFormat(100, 10, "Attachments: ", "0", 1, "L", false, 0, "")
	pdf.MultiCell(100, 10, body, "", "L", false)
	// pdf.CellFormat(100, 10, string(data), "0", 1, "CM", false, 0, "")
	pdf.OutputFileAndClose("email.pdf")
	log.Println("Done!")
}
