package http

import (
	"net/http"
	"strings"
	"github.com/open-falcon/mail-provider/config"
	"github.com/toolkits/web/param"
	"github.com/go-gomail/gomail"
	"strconv"
	"log"
)

func configProcRoutes(debug bool) {

	http.HandleFunc("/sender/mail", func(w http.ResponseWriter, r *http.Request) {
		cfg := config.Config()
		token := param.String(r, "token", "")
		if cfg.Http.Token != token {
			if debug{
				http.Error(w, "no privilege:cfg.Http.Token != token", http.StatusForbidden)
			}
			log.Println("no privilege:cfg.Http.Token != token")
			return
		}

		tos := strings.Split(param.MustString(r, "tos"), ",")
		subject := param.MustString(r, "subject")
		content := param.MustString(r, "content")

		Addr:=strings.Split(cfg.Smtp.Addr,":")
		port,error:=strconv.Atoi(Addr[1])
		if error != nil{
			if debug {
				http.Error(w, "The address is error to converted Int ", http.StatusForbidden)
			}
			log.Println("The address is error to converted Int ")
			return
		}
		d := gomail.NewDialer(Addr[0], port, cfg.Smtp.Username, cfg.Smtp.Password)
		s, err := d.Dial()
		if err != nil {
			if debug{
				http.Error(w, "Could not send email,smtp configure is Incorrect.", http.StatusForbidden)
			}
			log.Println("Could not send email,smtp configure is Incorrect.")
			return
		}
		m := gomail.NewMessage()
		for _, to := range tos {
			m.SetHeader("From", cfg.Smtp.From)
			m.SetHeader("To", to)
			m.SetHeader("Subject", subject)
			m.SetBody("text/html", content)

			if err := gomail.Send(s, m); err != nil {
				if debug{
					http.Error(w, "Could not send email to "+to, http.StatusForbidden)
				}
				log.Println("Could not send email to "+to)
			}else {
				if debug{
					http.Error(w, "success", http.StatusOK)
					log.Println("success send email to "+to)
				}

			}
			m.Reset()
		}
	})

}
