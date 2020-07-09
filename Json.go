package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)
var goods []Goods
var file filePath

type Goods struct {
	ID int
	Name string
	Manufacturer string
	Count int
	Price int
}

type filePath struct {
	Path string
}

func changeFile(path string)  {
	file.Path = path // ТУТ!!!!

	b1, _ := json.Marshal(file)
	err := ioutil.WriteFile("fileName.json", b1, 0777)
	if err != nil {
		fmt.Println(err)
	}
}

func saveFile()  {
	b1, _ := json.Marshal(goods)
	err := ioutil.WriteFile(file.Path, b1, 0777)
	if err != nil {
		fmt.Println(err)
	}
}

func readFile()  {
	fileName, err := ioutil.ReadFile("fileName.json")
	_ = json.Unmarshal(fileName, &file)
	fmt.Print("Текущий файл: " + file.Path + "\n")
	data, err := ioutil.ReadFile(file.Path)
	if err != nil {
		fmt.Println(err)
	}
	qwe := json.Unmarshal(data, &goods)
	if qwe != nil {
		fmt.Println(err)
	}
}

func addGood() {
	sc := bufio.NewScanner(os.Stdin)
	var ID int
	var Name string
	var Manufacturer string
	var Count int
	var Price int
	var err error
	ID = setID()

	fmt.Println("Введите наименование: ")
	sc.Scan()
	Name = sc.Text()

	fmt.Println("Введите производителя: ")
	sc.Scan()
	Manufacturer = sc.Text()

	fmt.Println("Введите количество: ")
	for {
		sc.Scan()
		Count, err = strconv.Atoi(sc.Text())
		if err == nil {
			break
		}
		fmt.Print("Некорректный ввод данных. Повторите попытку")
	}

	fmt.Println("Введите цену: ")
	for {
		sc.Scan()
		Price, err = strconv.Atoi(sc.Text())
		if err == nil {
			break
		}
		fmt.Print("Некорректный ввод данных. Повторите попытку")
	}

	g := Goods{
		ID:           ID,
		Name:         Name,
		Manufacturer: Manufacturer,
		Count:        Count,
		Price:        Price,
	}
	goods = append(goods, g)
}

func setID() int {
	ID := len(goods)
	return ID
}

func editGood(GoodID, fieldNum int, newValue string) []Goods  {
	if fieldNum == 1 {
		goods[GoodID].Name = newValue
	} else if fieldNum == 2 {
		goods[GoodID].Manufacturer = newValue
	} else if fieldNum == 3 {
		goods[GoodID].Count, _ = strconv.Atoi(newValue)
	} else if fieldNum == 4 {
		goods[GoodID].Price, _ = strconv.Atoi(newValue)
	}
	return goods
}

func showAll()  {
	fmt.Println("----------------------------------------------------------------------")
	fmt.Printf("|%-4s|%-20s|%-20s|%-8s|%-12s|\n", " ID", "      Название", "   Производитель", "  Цена", " Количество ")
	for i := 0; i < len(goods); i++ {
		fmt.Println("+----+--------------------+--------------------+--------+------------+")
		fmt.Printf("|%-4d|%-20s|%-20s|%8d|%12d|\n", goods[i].ID, goods[i].Name, goods[i].Manufacturer, goods[i].Price, goods[i].Count)
	}
	fmt.Println("+----+--------------------+--------------------+--------+------------+")
}

func deleteGood(ID int)  {
	goods = append(goods[:ID], goods[ID + 1:]...)
	showAll()
}

func main() {
	readFile()
	sc := bufio.NewScanner(os.Stdin)
	var args []string
	sc.Scan()
	line := sc.Text()
	args = strings.Fields(line)
	if args[0] == "list" {
		showAll()
	}
	if args[0] == "add" {
		addGood()
		saveFile()
	}
	if args[0] == "edit" {
		ID, _ := strconv.Atoi(args[1])
		fieldNum, _ := strconv.Atoi(args[2])
		newValue := args[3]
		editGood(ID, fieldNum, newValue)
		saveFile()
	}
	if args[0] == "del" {
		goodID, _ := strconv.Atoi(args[1])
		deleteGood(goodID)
		saveFile()
	}
	if args[0] == "open" {
		saveFile()
		changeFile(args[1])
	}
}

