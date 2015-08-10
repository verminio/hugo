package provider

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"

	"github.com/pkg/sftp"
)

const sftpDefaultPort string = "22"

type SftpProvider struct {
	target *url.URL
	host   string
	port   string
	conf   *ssh.ClientConfig
}

func (p *SftpProvider) Publish(target *url.URL) error {
	p.configure(target)
	p.connect()
	return nil
}

func (p *SftpProvider) configure(target *url.URL) {
	host := target.Host
	p.host = host[:strings.Index(host, ":")]
	port := host[strings.Index(host, ":")+1:]

	if port == "" || port == "0" {
		p.port = sftpDefaultPort
	} else {
		p.port = port
	}

	p.conf = &ssh.ClientConfig{
		User: target.User.Username(),
		Auth: []ssh.AuthMethod{
			ssh.PasswordCallback(func() (string, error) {
				if pass, def := target.User.Password(); def {
					return pass, nil
				} else {
					fmt.Print("Enter password: ")
					pass, _ := terminal.ReadPassword(int(os.Stdin.Fd()))
					return string(pass), nil
				}
			}),
		},
	}
}

func (p *SftpProvider) connect() {
	cli, err := ssh.Dial("tcp", p.host+":"+p.port, p.conf)

	if err != nil {
		panic("Error connecting via SFTP: " + err.Error())
	}

	client, err := sftp.NewClient(cli)
	defer client.Close()

	p.upload(client)
}

func (p *SftpProvider) upload(client *sftp.Client) {

}

func init() {
	registerProvider("sftp", &SftpProvider{})
}
