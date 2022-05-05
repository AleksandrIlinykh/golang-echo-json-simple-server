package main

import (
	"api/models"

	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type magazinesSlice []models.Magazine

func readJson() magazinesSlice {
	file, err := ioutil.ReadFile("./document.json")
	if err != nil {
		log.Fatal(err)
	}
	m := magazinesSlice{}
	err = json.Unmarshal(file, &m)
	if err != nil {
		log.Fatal(err)
	}

	return m
}

func writeJson(data magazinesSlice) {
	json_to_file, _ := json.Marshal(data)
	err := ioutil.WriteFile("./document.json", json_to_file, 4)
	if err != nil {
		log.Fatal(err)
	}
}

//-----------------------------------------------------------------

type documentsSlice []models.Document

func readJsonFile() documentsSlice {
	file, err := ioutil.ReadFile("./db.json")
	if err != nil {
		log.Fatal(err)
	}
	d := documentsSlice{}
	err = json.Unmarshal(file, &d)
	if err != nil {
		log.Fatal(err)
	}
	return d
}
func writeJsonFile(data documentsSlice) {
	json_to_file, _ := json.Marshal(data)
	err := ioutil.WriteFile("./db.json", json_to_file, 4)
	if err != nil {
		log.Fatal(err)
	}
}

//-----------------------------------------------------------------
func main() {
	e := echo.New()

	//----------------------------------------------My-server--------------------------
	e.GET("/docs", func(c echo.Context) error {
		documents := readJsonFile()
		return c.JSON(http.StatusOK, documents)
	})

	e.GET("/docs/:id", func(c echo.Context) error {
		documents := readJsonFile()
		for _, document := range documents {
			if c.Param("id") == strconv.Itoa(document.Id) {
				return c.JSON(http.StatusOK, document)
			}
		}
		return c.String(http.StatusNotFound, "Not found.")
	})

	//----------------------------------------------My-server--------------------------

	e.GET("/magazines", func(c echo.Context) error {
		magazines := readJson()
		return c.JSON(http.StatusOK, magazines)
	})

	e.GET("/magazines/:id", func(c echo.Context) error {
		magazines := readJson()

		for _, magazine := range magazines {
			if c.Param("id") == strconv.Itoa(magazine.Id) {
				return c.JSON(http.StatusOK, magazine)
			}
		}
		return c.String(http.StatusNotFound, "Not found.")
	})

	e.POST("/magazines", func(c echo.Context) error {
		magazines := readJson()

		new_magazine := new(models.Magazine)
		err := c.Bind(new_magazine)
		if err != nil {
			return c.String(http.StatusBadRequest, "Bad request.")
		}

		magazines = append(magazines, *new_magazine)
		writeJson(magazines)

		return c.JSON(http.StatusOK, magazines)
	})

	e.PUT("/magazines/:id", func(c echo.Context) error {
		magazines := readJson()

		updated_magazine := new(models.Magazine)
		err := c.Bind(updated_magazine)
		if err != nil {
			return c.String(http.StatusBadRequest, "Bad request.")
		}

		for i, magazine := range magazines {
			if strconv.Itoa(magazine.Id) == c.Param("id") {
				magazines = append(magazines[:i], magazines[i+1:]...)
				magazines = append(magazines, *updated_magazine)

				writeJson(magazines)

				return c.JSON(http.StatusOK, magazines)
			}
		}

		return c.String(http.StatusNotFound, "Not found.")
	})

	e.DELETE("/magazines/:id", func(c echo.Context) error {
		magazines := readJson()

		for i, magazine := range magazines {
			if strconv.Itoa(magazine.Id) == c.Param("id") {
				magazines = append(magazines[:i], magazines[i+1:]...)
				writeJson(magazines)

				return c.JSON(http.StatusOK, magazines)
			}
		}
		return c.String(http.StatusNotFound, "Not found.")
	})

	e.Start(":5000")
}
