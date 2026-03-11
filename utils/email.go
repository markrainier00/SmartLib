package utils

import (
	"fmt"
	"os"

	gomail "gopkg.in/gomail.v2"
)

func SendResetEmail(toEmail, token string) error {
	resetLink := fmt.Sprintf("%s/reset-password?token=%s", os.Getenv("APP_URL"), token)

	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("GMAIL_USER"))
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", "Reset SmartLib Password")
	m.SetBody("text/html", fmt.Sprintf(`
		<div style="background:#f5f6fa; padding:40px 20px;">
			<div style="max-width:520px; margin:auto; background:#ffffff; border-radius:16px; box-shadow:0 10px 30px rgba(0,0,0,0.08); padding:32px 28px;">
				<div style="font-family: DM Sans, sans-serif; max-width: 480px; margin: auto; padding: 40px 24px;">
						<h2 style="color: #1a2744;">Reset Your Password</h2>
						<p style="color: #475569;">A password reset was requested for your account.
						<br>If this was you, click the button below to create a new password.<br>
						<br>This expires in <strong>1 hour</strong>.</p>
				<div style="text-align:center; margin-top:16px;">
					<a href="%s" style="display:inline-block; padding:12px 28px; background:#1a2744; color:#fff; border-radius:10px; text-decoration:none; font-weight:600;">
					Reset Password
					</a>
				</div>
						<p style="margin-top: 24px; font-size: 12px; color: #8a8ea8;">If you didn't request this, you can safely ignore this email.</p>
					</div>
					<div style="border-top:1px solid #e5e7eb; margin-top:24px; padding-top:16px; text-align:center;">
						<div style="font-family: 'DM Serif Display', serif; font-weight:bold; font-size:26px; color:#1a2744;">SmartLib</div>
						<div style="font-size:12px; color:#8a8ea8;">School Library Management Portal</div>
					</div>
			</div>
		</div>
	`, resetLink))

	port := 587
	d := gomail.NewDialer("smtp.gmail.com", port, os.Getenv("GMAIL_USER"), os.Getenv("GMAIL_PASS"))

	return d.DialAndSend(m)
}
