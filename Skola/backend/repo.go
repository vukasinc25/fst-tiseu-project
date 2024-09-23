package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
	"strings"
)

type Repo struct {
	db *sql.DB
}

func NewRepository() (*Repo, error) {
	db, err := sql.Open("mysql",
		"root:root@tcp(mysql:3306)/skola")
	if err != nil {
		return nil, err
	}

	return &Repo{
		db: db,
	}, nil
}

func (rp *Repo) Disconnect() error {
	err := rp.db.Close()
	if err != nil {
		return err
	}
	return nil
}

// select d.id from Users u, Students s, Diplomas d where u.id = s.userId and s.id = d.studentId;
func (rp *Repo) GetDiplomaByStudent(userid string) ([]Diploma, error) {
	rows, err := rp.db.Query("select d.id from Users u, Students s, Diplomas d "+
		"where u.id = s.userId and s.id = d.studentId and u.id = ?;", userid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ids []int
	for rows.Next() {
		var diplomaId int
		err := rows.Scan(&diplomaId)
		if err != nil {
			return nil, err
		}
		ids = append(ids, diplomaId)
	}
	var temp []string
	var idString string
	for i := 0; i < len(ids); i++ {
		temp = append(temp, strconv.Itoa(ids[i]))
	}
	idString = strings.Join(temp, ",")
	//log.Println(idString)
	var queryString = fmt.Sprintf("%s%s%s", "select d.averageGrade, s.TotalHighPoints, d.yearFinished from Diplomas d, Students s where d.studentId = s.id and d.id in (", idString, ");")
	rows, err = rp.db.Query(queryString)

	var diplomas []Diploma
	for rows.Next() {
		var diploma Diploma
		err := rows.Scan(&diploma.AverageGrade, &diploma.TotalHighPoints, &diploma.YearFinished)
		if err != nil {
			return nil, err
		}
		diplomas = append(diplomas, diploma)
	}
	for i := 0; i < len(ids); i++ {
		rows, err = rp.db.Query("select grades, finalGrade from HasSubjects where diplomaId = ?", ids[i])

		for rows.Next() {
			var subject Subject
			err := rows.Scan(&subject.Grades, &subject.FinalGrade)
			if err != nil {
				return nil, err
			}
			diplomas[i].Subjects = append(diplomas[i].Subjects, subject)
		}

	}

	//log.Println(diplomas)
	return diplomas, nil
}

func (rp *Repo) CreateData() error {

	_, err := rp.db.Exec("CREATE DATABASE IF NOT EXISTS skola")
	if err != nil {
		return err
	}

	_, err = rp.db.Exec("USE skola")
	if err != nil {
		return err
	}

	_, err = rp.db.Exec("CREATE TABLE IF NOT EXISTS Users(" +
		"id varchar(30) primary key," +
		"firstName varchar(50)," +
		"lastName varchar(50)," +
		"jmbg varchar(50)" +
		")")
	if err != nil {
		return err
	}

	_, err = rp.db.Exec("CREATE TABLE IF NOT EXISTS Students(" +
		"id int AUTO_INCREMENT primary key," +
		"currentSchoolYear int," +
		"TotalHighPoints float," +
		"userId varchar(30)," +
		"FOREIGN KEY (userId) REFERENCES Users(id)" +
		")")
	if err != nil {
		return err
	}

	_, err = rp.db.Exec("CREATE TABLE IF NOT EXISTS Professors(" +
		"id int AUTO_INCREMENT primary key," +
		"userId varchar(30)," +
		"FOREIGN KEY (userId) REFERENCES Users(id)" +
		")")
	if err != nil {
		return err
	}

	_, err = rp.db.Exec("CREATE TABLE IF NOT EXISTS Diplomas(" +
		"id int AUTO_INCREMENT primary key," +
		"averageGrade float," +
		"yearFinished int," +
		"studentId int," +
		"FOREIGN KEY (studentId) REFERENCES Students(id)" +
		")")
	if err != nil {
		return err
	}

	_, err = rp.db.Exec("CREATE TABLE IF NOT EXISTS Subjects(" +
		"id int AUTO_INCREMENT primary key," +
		"name varchar(50)" +
		")")
	if err != nil {
		return err
	}

	_, err = rp.db.Exec("CREATE TABLE IF NOT EXISTS HasSubjects(" +
		"diplomaId int," +
		"subjectId int," +
		"grades varchar(20)," +
		"finalGrade int," +
		"PRIMARY KEY (diplomaId, subjectId)," +
		"FOREIGN KEY (diplomaId) REFERENCES Diplomas(id)," +
		"FOREIGN KEY (subjectId) REFERENCES Subjects(id)" +
		")")
	if err != nil {
		return err
	}

	_, err = rp.db.Exec("CREATE TABLE IF NOT EXISTS TeachesSubjects(" +
		"professorId int," +
		"subjectId int," +
		"PRIMARY KEY (professorId, subjectId)," +
		"FOREIGN KEY (professorId) REFERENCES Professors(id)," +
		"FOREIGN KEY (subjectId) REFERENCES Subjects(id)" +
		")")
	if err != nil {
		return err
	}

	rows, err := rp.db.Query("SELECT * FROM Users")
	if err != nil {
		return err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.id, &user.firstName, &user.lastName, &user.jmbg)
		if err != nil {
			return err
		}
		users = append(users, user)
	}

	log.Println(users)

	if len(users) == 0 {
		_, err := rp.db.Exec("INSERT INTO Users"+
			"(id, firstName, lastName, jmbg) VALUES (?, ?, ?, ?)"+
			"", "113289113681502612644", "Pera", "Peric", " ")
		if err != nil {
			return err
		}

		_, err = rp.db.Exec("INSERT INTO Users"+
			"(id, firstName, lastName, jmbg) VALUES (?, ?, ?, ?)"+
			"", "117316411448317157040", "Zika", "Zikic", " ")
		if err != nil {
			return err
		}

		_, err = rp.db.Exec("INSERT INTO Students"+
			"(currentSchoolYear, TotalHighPoints, userId) VALUES (?, ?, ?)"+
			"", 2, 43.20, "113289113681502612644")
		if err != nil {
			return err
		}

		_, err = rp.db.Exec("INSERT INTO Students"+
			"(currentSchoolYear, TotalHighPoints, userId) VALUES (?, ?, ?)"+
			"", 4, 45.20, "117316411448317157040")
		if err != nil {
			return err
		}

		_, err = rp.db.Exec("INSERT INTO Subjects"+
			"(name) VALUES (?)"+
			"", "Srpski")
		if err != nil {
			return err
		}

		_, err = rp.db.Exec("INSERT INTO Subjects"+
			"(name) VALUES (?)"+
			"", "Matematika")
		if err != nil {
			return err
		}

		_, err = rp.db.Exec("INSERT INTO Subjects"+
			"(name) VALUES (?)"+
			"", "Fizika")
		if err != nil {
			return err
		}

		_, err = rp.db.Exec("INSERT INTO Subjects"+
			"(name) VALUES (?)"+
			"", "Hemija")
		if err != nil {
			return err
		}

		_, err = rp.db.Exec("INSERT INTO Diplomas"+
			"(averageGrade, yearFinished, studentId) VALUES (?, ?, ?)"+
			"", 4.30, 2021, 1)
		if err != nil {
			return err
		}

		_, err = rp.db.Exec("INSERT INTO Diplomas"+
			"(averageGrade, yearFinished, studentId) VALUES (?, ?, ?)"+
			"", 4.55, 2022, 1)
		if err != nil {
			return err
		}

		_, err = rp.db.Exec("INSERT INTO HasSubjects"+
			"(diplomaId, subjectId, grades, finalGrade) VALUES (?, ?, ?, ?)"+
			"", 1, 1, "2,2,3", 0)
		if err != nil {
			return err
		}

		_, err = rp.db.Exec("INSERT INTO HasSubjects"+
			"(diplomaId, subjectId, grades, finalGrade) VALUES (?, ?, ?, ?)"+
			"", "1", "2", "4,5,3", "0")
		if err != nil {
			return err
		}

		_, err = rp.db.Exec("INSERT INTO HasSubjects"+
			"(diplomaId, subjectId, grades, finalGrade) VALUES (?, ?, ?, ?)"+
			"", "1", "3", "3,5,4", "0")
		if err != nil {
			return err
		}

		_, err = rp.db.Exec("INSERT INTO HasSubjects"+
			"(diplomaId, subjectId, grades, finalGrade) VALUES (?, ?, ?, ?)"+
			"", 2, 1, "2,4,4", 0)
		if err != nil {
			return err
		}

		_, err = rp.db.Exec("INSERT INTO HasSubjects"+
			"(diplomaId, subjectId, grades, finalGrade) VALUES (?, ?, ?, ?)"+
			"", 2, 2, "1,5,4", "0")
		if err != nil {
			return err
		}

		_, err = rp.db.Exec("INSERT INTO HasSubjects"+
			"(diplomaId, subjectId, grades, finalGrade) VALUES (?, ?, ?, ?)"+
			"", 2, 3, "3,2,4", "0")
		if err != nil {
			return err
		}

		_, err = rp.db.Exec("INSERT INTO Diplomas"+
			"(averageGrade, yearFinished, studentId) VALUES (?, ?, ?)"+
			"", 4.30, 2021, 2)
		if err != nil {
			return err
		}

		_, err = rp.db.Exec("INSERT INTO Diplomas"+
			"(averageGrade, yearFinished, studentId) VALUES (?, ?, ?)"+
			"", 4.55, 2022, 2)
		if err != nil {
			return err
		}

		_, err = rp.db.Exec("INSERT INTO Diplomas"+
			"(averageGrade, yearFinished, studentId) VALUES (?, ?, ?)"+
			"", 4.30, 2023, 2)
		if err != nil {
			return err
		}

		_, err = rp.db.Exec("INSERT INTO Diplomas"+
			"(averageGrade, yearFinished, studentId) VALUES (?, ?, ?)"+
			"", 4.55, 2024, 2)
		if err != nil {
			return err
		}

	}

	return nil
}
