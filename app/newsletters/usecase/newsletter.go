package usecase

import (
	"context"
	"fmt"
	"net/smtp"
	"os"
	"time"

	"github.com/hammer-code/lms-be/domain"
)

func (us *usecase) Subscribe(ctx context.Context, email string) error {

	// check the email first
	nws, err := us.newsLetterRepo.GetByEmail(ctx, email)
	if err != nil {
		return err
	}

	if nws != nil {
		return nil
	}

	createdNewsletter := domain.Newsletter{
		Email:     email,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	err = us.newsLetterRepo.Subscribe(ctx, createdNewsletter)
	if err != nil {
		return err
	}

	htmlTmpl, err := os.ReadFile("./assets/subscribe_template.html")
	if err != nil {
		fmt.Println(err)
		return err
	}

	subject := "Thank You for Your Interest in PDD24!"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	message := []byte(fmt.Sprintf("To: %s\r\n"+
		"Subject: %s\r\n"+
		"%s\r\n"+
		"%s\r\n", email, subject, mime, string(htmlTmpl)))

	auth := smtp.PlainAuth("", us.cfg.SMTP_EMAIL, us.cfg.SMTP_PASSWORD, us.cfg.SMTP_HOST)

	host := fmt.Sprintf("%s:%s", us.cfg.SMTP_HOST, us.cfg.SMTP_PORT)
	if err := smtp.SendMail(host, auth, us.cfg.SMTP_HOST, []string{email}, message); err != nil {
		return err
	}

	return nil
}
