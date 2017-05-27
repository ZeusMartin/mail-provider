package http

import (
	"github.com/open-falcon/mail-provider/config"
	"github.com/toolkits/web/param"
	"log"
	"net/http"
	"net/smtp"
	"strings"
)

func SendMail(addr string, from string, to []string, msg []byte) error {
	c, err := smtp.Dial(addr)
	if err != nil {
		return err
	}
	defer c.Close()
	if err = c.Mail(from); err != nil {
		return err
	}
	for _, addr := range to {
		if err = c.Rcpt(addr); err != nil {
			return err
		}
	}
	w, err := c.Data()
	if err != nil {
		return err
	}
	_, err = w.Write(msg)
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	return c.Quit()
}

func SendEmail(Addr string, From string, sendTo []string, subject string, content string) error {

	msg := "Subject: " + subject + "\r\n" + "Content-Type: text/html; charset=UTF-8" + "\r\n\r\n" + content

	if config.Config().Debug{
		log.Printf("msg: %s",msg)
	}

	done := make(chan error, 1024)
	go func() {
		defer close(done)
		for _, v := range sendTo {

			err := SendMail(
				Addr,
				From,
				[]string{v},
				[]byte(msg),
			)
			done <- err
		}
	}()

	for i := 0; i < len(sendTo); i++ {
		<-done
	}

	return nil
}

func configProcRoutes() {

	http.HandleFunc("/sender/mail", func(w http.ResponseWriter, r *http.Request) {

		cfg := config.Config()
		debug := cfg.Debug

		token := param.String(r, "token", "")
		if cfg.Http.Token != token {
			if debug {
				http.Error(w, "no privilege:cfg.Http.Token != token", http.StatusForbidden)
			}
			log.Println("no privilege:cfg.Http.Token != token")
			return
		}

		subject := param.MustString(r, "subject")
		content := param.MustString(r, "content")
		sendTo := strings.Split(param.MustString(r, "tos"), ",")

		error := SendEmail(cfg.Smtp.Addr, cfg.Smtp.From, sendTo, subject, content)

		if error != nil {
			if debug{
				http.Error(w, error.Error(), http.StatusInternalServerError)
			}
			log.Println("send email faild: %v",error)
		} else {
			if debug{
				http.Error(w, "success", http.StatusOK)
				log.Printf("success send email to %v",sendTo)
			}

		}

	})

}
