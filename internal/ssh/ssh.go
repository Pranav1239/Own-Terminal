package ssh

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
)

// SigninToSSH connects to an SSH server using the provided host, user, and password.
func SigninToSSH(host, user, password string) error {
	// Load private key from environment variable or file
	privateKey := []byte(os.Getenv("SSH_PRIVATE_KEY"))

	// Parse the private key
	signer, err := ssh.ParsePrivateKey(privateKey)
	if err != nil {
		return fmt.Errorf("failed to parse private key: %v", err)
	}

	// Load the known_hosts file
	hostKeyCallback, err := knownhosts.New("~/.ssh/known_hosts")
	if err != nil {
		return fmt.Errorf("failed to load known_hosts file: %v", err)
	}

	// SSH client configuration
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: hostKeyCallback,
	}

	conn, err := ssh.Dial("tcp", host, config)
	if err != nil {
		return fmt.Errorf("failed to dial: %v", err)
	}
	defer conn.Close()

	session, err := conn.NewSession()
	if err != nil {
		return fmt.Errorf("failed to create session: %v", err)
	}
	defer session.Close()

	stdin, err := session.StdinPipe()
	if err != nil {
		return fmt.Errorf("failed to create stdin pipe: %v", err)
	}

	stdout, err := session.StdoutPipe()
	if err != nil {
		return fmt.Errorf("failed to create stdout pipe: %v", err)
	}

	stderr, err := session.StderrPipe()
	if err != nil {
		return fmt.Errorf("failed to create stderr pipe: %v", err)
	}

	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Printf("error reading stdout: %v\n", err)
		}
	}()

	go func() {
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Printf("error reading stderr: %v\n", err)
		}
	}()

	if err := session.Shell(); err != nil {
		return fmt.Errorf("failed to start shell: %v", err)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("$ ")
		if !scanner.Scan() {
			break
		}
		cmd := scanner.Text()
		if _, err := io.WriteString(stdin, cmd+"\n"); err != nil {
			return fmt.Errorf("failed to write to stdin: %v", err)
		}
	}

	return session.Wait()
}
