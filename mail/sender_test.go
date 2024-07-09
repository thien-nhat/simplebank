package mail

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thien-nhat/simplebank/util"
)

func TestSendEmailWithGmail(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	config, err := util.LoadConfig("..")
	require.NoError(t, err)

	sender := NewGmailSender(config.EmailSenderName,
		config.EmailSenderAddress, config.EmailSenderPassword)
	
	subject := "A test Email"
	content := `
	<h1> Hello World</h1>
	<p> This is a message from Simple Bank service</p>
	`
	to := []string{"ngonhatthien02@gmail.com"}
	attachFiles := []string{"../README.md"}

	err = sender.SendEmail(subject, content, to, nil, nil, attachFiles)
	require.NoError(t, err)
}