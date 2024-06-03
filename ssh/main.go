package main

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"os"

	yaml "gopkg.in/yaml.v3"

	"golang.org/x/crypto/ssh"
)


type Config struct{
    Host struct {
        Name string `yaml:"name"`
        IPAddr string `yaml:"ipaddr"`
        Port string `yaml:"port"`
        User string `yaml:"user"`
        Key string `yaml:"key"`
    }
}

//e.g. output, err := remoteRun("root", "MY_IP", "PRIVATE_KEY", "ls")
func remoteRun(user string, addr string, privateKey []byte, cmd string) (string, error) {
    // privateKey could be read from a file, or retrieved from another storage
    // source, such as the Secret Service / GNOME Keyring
    key, err := ssh.ParsePrivateKey(privateKey)
    if err != nil {
        return "", err
    }
    // Authentication
    config := &ssh.ClientConfig{
        User: user,
        // https://github.com/golang/go/issues/19767 
        // as clientConfig is non-permissive by default 
        // you can set ssh.InsercureIgnoreHostKey to allow any host 
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
        Auth: []ssh.AuthMethod{
            ssh.PublicKeys(key),
        },
        //alternatively, you could use a password
        /*
            Auth: []ssh.AuthMethod{
                ssh.Password("PASSWORD"),
            },
        */
    }
    // Connect
    client, err := ssh.Dial("tcp", net.JoinHostPort(addr, "22"), config)
    if err != nil {
        return "", err
    }
    // Create a session. It is one session per command.
    session, err := client.NewSession()
    if err != nil {
        return "", err
    }
    defer session.Close()
    var b bytes.Buffer  // import "bytes"
    session.Stdout = &b // get output
    // you can also pass what gets input to the stdin, allowing you to pipe
    // content from client to server
    //      session.Stdin = bytes.NewBufferString("My input")

    // Finally, run the command
    err = session.Run(cmd)
    return b.String(), err
}


func main() {

//     var cfgFromSample Config
//     sampleData := `
// host:
//   -name: "FQDNorIP"
//   ipaddr: "192.168.XX.XX"
//   port: "22"
//   user: "remoteUserName"
//   key: "/Users/userName/.ssh/id_rsa"
// `
    // sampleDataConfig := yaml.Unmarshal([]byte(sampleData), &cfgFromSample)
    // fmt.Println(cfgFromSample.Host.Name, cfgFromSample.Host.IPAddr, cfgFromSample.Host.Port, cfgFromSample.Host.User, cfgFromSample.Host.Key)

	// Read config file from config.yaml
	config, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatal(err)
	}

    var cfg Config
	err = yaml.Unmarshal(config, &cfg)

	if err != nil {
		log.Fatal(err)
	}
    
    fmt.Println(cfg.Host.Name, cfg.Host.IPAddr, cfg.Host.Port, cfg.Host.User, cfg.Host.Key)

    // Check if a file exists
    if _, err := os.Stat(cfg.Host.Key); os.IsNotExist(err) {
        log.Fatal("[private key file] Key file does not exist @ ", cfg.Host.Key)
    }

    // Read the private key file
    privateKey, err := os.ReadFile(cfg.Host.Key)
    if err != nil {
        log.Fatal(err)
    }

    output, err := remoteRun(cfg.Host.User, cfg.Host.IPAddr, privateKey, "ls")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(output)

}