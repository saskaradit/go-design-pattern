package main

import "strings"

type email struct {
	from, to, subj, body string
}

type EmailBuilder struct {
	email email
}

func (b *EmailBuilder) From(from string) *EmailBuilder {
	if !strings.Contains(from, "@") {
		panic("email should be valid")
	}
	b.email.from = from
	return b
}
func (b *EmailBuilder) To(to string) *EmailBuilder {
	if !strings.Contains(to, "@") {
		panic("email should be valid")
	}
	b.email.to = to
	return b
}
func (b *EmailBuilder) Subject(subject string) *EmailBuilder {
	b.email.subj = subject
	return b
}
func (b *EmailBuilder) Body(body string) *EmailBuilder {
	b.email.body = body
	return b
}

func sendMailImpl(email *email) {

}

type build func(*EmailBuilder)

func SendEmail(action build) {
	builder := EmailBuilder{}
	action(&builder)
	sendMailImpl(&builder.email)
}

func main() {
	SendEmail(func(b *EmailBuilder) {
		b.From("rad@rad.com").
			To("sas@sas.com").
			Subject("Hehe").
			Body("this is the body")
	})
}
