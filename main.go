package main

import (
	"fmt"
	"os"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"golang.org/x/crypto/ssh"
)

type model struct {
	username   string
	password   string
	host       string
	port       string
	focusIndex int
	choices    []string
	cursor     int
	connected  bool
	outputMsg  string
}

func initialModel() model {
	return model{
		choices:    []string{"Connect", "Quit"},
		focusIndex: 0,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyUp, tea.KeyShiftTab:
			m.focusIndex--
			if m.focusIndex < 0 {
				m.focusIndex = len(m.choices) + 3
			}
		case tea.KeyDown, tea.KeyTab:
			m.focusIndex++
			if m.focusIndex > len(m.choices)+3 {
				m.focusIndex = 0
			}
		case tea.KeyEnter:
			if m.focusIndex == len(m.choices)+3 {
				// Connect button
				m.connected = false
				m.outputMsg = ""
				session, err := connectSSH(m.username, m.password, m.host, m.port)
				if err != nil {
					m.outputMsg = fmt.Sprintf("Failed to connect: %s", err)
					return m, nil
				}
				defer session.Close()
				output, err := session.CombinedOutput("ls -al")
				if err != nil {
					m.outputMsg = fmt.Sprintf("Failed to run command: %s", err)
					return m, nil
				}
				m.connected = true
				m.outputMsg = string(output)
			} else if m.focusIndex == len(m.choices)+4 {
				// Quit button
				return m, tea.Quit
			}
		case tea.KeyBackspace, tea.KeyDelete:
			if m.focusIndex <= 3 {
				switch m.focusIndex {
				case 0:
					if len(m.username) > 0 {
						m.username = m.username[:len(m.username)-1]
					}
				case 1:
					if len(m.password) > 0 {
						m.password = m.password[:len(m.password)-1]
					}
				case 2:
					if len(m.host) > 0 {
						m.host = m.host[:len(m.host)-1]
					}
				case 3:
					if len(m.port) > 0 {
						m.port = m.port[:len(m.port)-1]
					}
				}
			}
		case tea.KeyRunes:
			if m.focusIndex <= 3 {
				switch m.focusIndex {
				case 0:
					m.username += string(msg.Runes)
				case 1:
					m.password += string(msg.Runes)
				case 2:
					m.host += string(msg.Runes)
				case 3:
					m.port += string(msg.Runes)
				}
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	s := "SSH Client\n\n"

	s += fmt.Sprintf("Username: %s%s\n", m.username, getInputCursor(m.focusIndex == 0))
	s += fmt.Sprintf("Password: %s%s\n", m.password, getInputCursor(m.focusIndex == 1))
	s += fmt.Sprintf("Host: %s%s\n", m.host, getInputCursor(m.focusIndex == 2))
	s += fmt.Sprintf("Port: %s%s\n\n", m.port, getInputCursor(m.focusIndex == 3))

	for i, choice := range m.choices {
		cursor := " "
		if m.focusIndex == i+4 {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	s += "\nPress q to quit.\n"

	if m.connected {
		s += "\nConnected! Output:\n" + m.outputMsg
	} else if m.outputMsg != "" {
		s += "\nError: " + m.outputMsg
	}

	return s
}

func getInputCursor(focused bool) string {
	if focused {
		return "█"
	}
	return ""
}

func connectSSH(username, password, host, portStr string) (*ssh.Session, error) {
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, fmt.Errorf("invalid port: %s", err)
	}

	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	address := fmt.Sprintf("%s:%d", host, port)
	client, err := ssh.Dial("tcp", address, config)
	if err != nil {
		return nil, err
	}

	session, err := client.NewSession()
	if err != nil {
		return nil, err
	}

	return session, nil
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
