package noteup

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"time"
	"fmt"
	"strconv"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/charmbracelet/lipgloss"


)

type item struct {
	Account string
	Password string
	CreatedAt time.Time
    LastUpdatedAt time.Time
}

type Passd []item
var (
	HeaderStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Bold(true)
	EvenRowStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	OddRowStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("255"))
)
func (p *Passd) Add(account, password string) {
	newItem := item{
		Account: account,
		Password: password,
		CreatedAt: time.Now(),
		LastUpdatedAt: time.Now(),
	}
	*p = append(*p, newItem)
}

func (p *Passd) Update(index int,  password string) error {
	ls:= *p
	if index <= 0 || index > len(*p) {
		return errors.New("index out of range")
	}
	(ls)[index-1].Password = password
	(ls)[index-1].LastUpdatedAt = time.Now()

	return nil
}

func (p *Passd) Delete(index int) error {
	ls := *p
	fmt.Println(ls)

	if index <= 0 || index > len(*p) {
		return errors.New("index out of range")
	}

	// Create a temporary slice without the specified element
	tempSlice := append(ls[:index-1], ls[index:]...)

	// Update the original slice
	*p = tempSlice

	return nil
}


func (p *Passd) Load(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	if len(file) == 0 {
		return err
	}
	err= json.Unmarshal(file, p)
	if err != nil {
		return err
	}
	return nil
}

func (p *Passd) Store(filename string)error{
	data, err := json.Marshal(p)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, data, 0644)
}

func (p *Passd) List() string {
	istLocation, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		panic(err)
	}
	rows := [][]string{}
	for idx, v := range *p {
		createdAtIST := v.CreatedAt.In(istLocation)
		updatedAtIST := v.LastUpdatedAt.In(istLocation)
		rows = append(rows, []string{strconv.Itoa(idx+1),v.Account, v.Password, createdAtIST.Format("2006-01-02 15:04:05"), updatedAtIST.Format("2006-01-02 15:04:05")})
	}

	t := table.New().
	Border(lipgloss.NormalBorder()).
	BorderStyle(lipgloss.NewStyle().
	Foreground(lipgloss.Color("99"))).
	StyleFunc(func(row, col int) lipgloss.Style {
        switch {
        case row == 0:
            return HeaderStyle
        case row%2 == 0:
            return OddRowStyle
        default:
            return OddRowStyle
        }
    }).
    Headers("ID", "Account", "Password", "Created At", "Last Updated").
    Rows(rows...)


return t.Render()
}
