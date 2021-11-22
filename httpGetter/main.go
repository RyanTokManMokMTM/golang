package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/urfave/cli"
	//"gorm.io/driver/postgres"
	"github.com/teamlint/opencc"
	"gorm.io/gorm"
	"httpGetter/GzFileDownloader"
	"httpGetter/webCrawler"
	"os"

	"io/ioutil"
	"log"
	"strconv"
	"time"
)

const (
	host string = "https://api.themoviedb.org/3"
	apiKey string = "29570e7acc52b3e085ab46f6a60f0a55"

	allMovieURI string ="/discover/movie"
	genreAllURI string = "/genre/movie/list"

	peoplePopular string = "/person/popular"

	//JSON GZ
	fileHost string = "http://files.tmdb.org/p/exports"

	//sqlHOST string = "127.0.0.1"
	//userName string = "postgres"
	//password string = "jackson"
	//port int = 5432
	//db string = "movie"
)

var (
	sqlHOST string = "127.0.0.1"
	userName string = "postgres"
	password string = ""
	port int = 5432
	db string = "TMDB"
	moviePath string = ""
	PersonPath string = ""
	migration bool = false
)


var (
	year ,month,day = time.Now().Date()
	movieGZ string = fmt.Sprintf("/movie_ids_%d_%d_%d.json.gz",int(10),31,year)
	peopleGZ string = fmt.Sprintf("/person_ids_%d_%d_%d.json.gz",int(10),31,year)
)

func dbConfigure() string{
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s",userName,password,sqlHOST,port,db)
	//return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d ",sqlHOST,userName,password,db,port)
}

func main(){

	//readArgc()
	//if PersonPath == "" || moviePath == ""{
	//	log.Fatalln("FilePath can't be empty")
	//}
	//
	//log.Println("Configuring the database...")
	//config := dbConfigure()
	//db, err := gorm.Open(postgres.Open(config),&gorm.Config{
	//})
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//log.Println("DB Configuration Done...")
	//
	//if migration {
	//	log.Println("Creating table...")
	//	db.AutoMigrate(&webCrawler.GenreInfo{})
	//	db.AutoMigrate(&webCrawler.MovieInfo{})
	//	db.AutoMigrate(&webCrawler.GenresMovies{})
	//	db.AutoMigrate(&webCrawler.PersonInfo{})
	//	db.AutoMigrate(&webCrawler.MovieCharacter{})
	//	db.AutoMigrate(&webCrawler.PersonCrew{})
	//
	//	if err := db.Exec("ALTER TABLE genres_movies DROP CONSTRAINT genres_movies_pkey").Error ; err != nil {
	//		log.Println(err)
	//		return
	//	}
	//
	//	if err := db.Exec("ALTER TABLE genres_movies ADD CONSTRAINT  genres_movies_unique UNIQUE(genre_info_id,movie_info_id)").Error; err != nil{
	//		log.Println(err)
	//		return
	//	}
	//
	//	if err := db.Exec("ALTER TABLE genres_movies ADD CONSTRAINT genres_movies_pkey PRIMARY KEY (id)").Error ; err != nil{
	//		log.Println(err)
	//		return
	//	}
	//
	//}
	////TODO - Get Genre And Movie
	//movieCrawlerProcedure(db)
	////
	//////TODO - Get ALL person
	//personCrawlerProcedure(db)

}


func readArgc(){
	app := cli.NewApp()
	app.Name = "TMDB Web Crawler"
	app.Usage = "Fetch Movies and person etc..."
	app.Action = run
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "dbHost",
			Usage: "Postgres DB Host IP(Default:127.0.0.1)",
		},
		cli.StringFlag{
			Name: "dbUser,u",
			Usage: "Postgres DB Username(Default:postgres)",
		},
		cli.StringFlag{
			Name: "dbPw",
			Usage: "Postgres DB password(Default:null)",
		},
		cli.StringFlag{
			Name: "db",
			Usage:"Postgres DB database(Default:null)",
		},
		cli.StringFlag{
			Name: "dbPort,p",
			Usage: "Postgres DB port(Default:5432)",
		},
		cli.StringFlag{
			Name : "moviePath,mf",
			Usage: "Data to store in(Default:null)",
		},
		cli.StringFlag{
			Name : "personPath,pf",
			Usage: "Data to store in(Default:null)",
		},
		cli.StringFlag{
			Name : "createTable,c",
			Usage: "Auto Creating the db Table(0:False,1:True)(Default:false)",
		},
	}
	app.Run(os.Args)
}

func run(c *cli.Context) error{
	if c.String("dbHost") != ""{
		sqlHOST = c.String("dbHost")
	}

	if c.String("dbPort") != ""{
		p,err := strconv.Atoi(c.String("dbPort"))
		if err != nil{
			log.Fatalln(err)
		}
		port = p
	}

	if c.String("dbUser") != ""{
		userName = c.String("dbUser")
	}

	if c.String("dbPw") != ""{
		password = c.String("dbPw")
	}

	if c.String("db") != ""{
		db = c.String("db")
	}

	if c.String("moviePath") != ""{
		moviePath = c.String("moviePath")
	}

	if c.String("personPath") != ""{
		PersonPath = c.String("personPath")
	}

	if c.String("createTable") != ""{
		code,err := strconv.Atoi(c.String("createTable"))
		if err != nil{
			log.Fatalln(err)
		}

		if code == 0{
			migration = false
		} else if code == 1{
			migration = true
		}
	}

	return nil
}

func movieCrawlerProcedure(db *gorm.DB){
	genreAndMoviesAll(db)
	insertJSONsToDB(moviePath,db,"movie")
}

func personCrawlerProcedure(db *gorm.DB){
	err := fetchPersonVisID()
	if err != nil {
		log.Println(err)
		return
	}
	insertJSONsToDB(PersonPath,db,"person")
}

func fetchMovieViaID(moviePath string) error {
	uri := fileHost + movieGZ
	var uris []int
	moviesData, err := GzFileDownloader.DownloadGZFile(uri)
	if err != nil {
		log.Println(err)
		return err
	}

	for _,movie := range *moviesData{
		uris = append(uris,movie.Id)
	}

	webCrawler.FetchMovieInfosViaIDS(uris,moviePath)

	return nil
}

func fetchPersonVisID() error {
	uri := fileHost + peopleGZ
	var uris []int
	personData, err := GzFileDownloader.DownloadGZFile(uri)
	if err != nil {
		log.Println(err)
		return err
	}

	for _,person := range *personData{
		uris = append(uris,person.Id)
	}
	webCrawler.FetchPersonInfosViaIDS(uris,PersonPath)
	return nil
}

func genreAndMoviesAll(db *gorm.DB){
	apiURL := host + genreAllURI +"?api_key=" + apiKey + "&language=zh-TW"

	//TODO - Insert Data to Database
	_ ,err := webCrawler.GenreTableCreate(apiURL,db)
	if err != nil{
		log.Fatalln(err)
		return
	}

	fetchMovieViaID(moviePath)
}

func uriGenerator(uri string,page int) []string{
	var uris []string
	for i := 0;i<page;i++{
		newURI := uri + "&page=" + strconv.Itoa(i+1)
		uris = append(uris, newURI)
	}

	return uris
}

func insertJSONsToDB(dirPath string,db *gorm.DB,jsonType string){
	dir, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return
	}

	if jsonType == "movie"{
		for _,file := range dir{
			err := movieJsonToDB(db, dirPath, file.Name())
			if err != nil {
				log.Fatalln(err)
			}
		}
	}else if jsonType == "person"{
		for _,file := range dir{
			err := personJsonToDB(db,dirPath,file.Name())
			if err != nil {
				log.Fatalln(err)
			}
		}
	}

}

func movieJsonToDB(db *gorm.DB,dirPath string,fileName string)error{
	var movieInfo webCrawler.MovieInfo
	location := fmt.Sprintf("%s/%s",dirPath,fileName)
	jsonsData, err := ioutil.ReadFile(location)
	if err != nil {
		log.Println(err)
		return err
	}

	err = json.Unmarshal(jsonsData, &movieInfo)
	if err != nil {
		log.Println(err)
		return err
	}

	str, err := json.MarshalIndent(&movieInfo,"","\t")
	if err != nil {
		return err
	}

	ioutil.WriteFile(location,str,0666)

	if err := db.Where("id = ?",movieInfo.Id).First(&webCrawler.MovieInfo{});err !=nil{
		if errors.Is(err.Error,gorm.ErrRecordNotFound){
			//not found the record
			//insert to db
			db.Create(&movieInfo)
		}else{
			fmt.Println("???")
		}
	}
	return nil
}

func personJsonToDB(db *gorm.DB,dirPath string,fileName string) error {
	var personInfo webCrawler.PersonInfo
	location := fmt.Sprintf("%s/%s",dirPath,fileName)

	jsonData, err := ioutil.ReadFile(location)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonData, &personInfo)
	if err != nil {
		return err
	}

	if personInfo.ProfilePath == "" || len(personInfo.MovieCredits.Cast) == 0 && len(personInfo.MovieCredits.Crew) == 0{
		fmt.Println(personInfo.Name)
		return nil
	}



	if dbErr := db.Where("id = ?",personInfo.Id).First(&webCrawler.PersonInfo{});dbErr != nil{
		if errors.Is(dbErr.Error,gorm.ErrRecordNotFound){
			//TODO - ForEach cast need to check the movie info is our
			var newMovieCast []webCrawler.MovieCharacter
			var newMovieCrew []webCrawler.PersonCrew

			for _,castData := range personInfo.MovieCredits.Cast{
				//if current cast movie is existed
				if dbInsertErr := db.Where("id = ?",castData.MovieID).First(&webCrawler.MovieInfo{});dbInsertErr!=nil{
					if !errors.Is(dbInsertErr.Error,gorm.ErrRecordNotFound){
						//existed
						newMovieCast = append(newMovieCast,castData)
					}
				}
			}

			for _,crewData := range personInfo.MovieCredits.Crew{
				if dbInsertErr := db.Where("id = ?",crewData.MovieID).First(&webCrawler.MovieInfo{});dbInsertErr!=nil{
					if !errors.Is(dbInsertErr.Error,gorm.ErrRecordNotFound){
						//existed
						newMovieCrew = append(newMovieCrew,crewData)
					}
				}
			}

			personInfo.MovieCharacter = newMovieCast
			personInfo.PersonCrew = newMovieCrew
			db.Create(&personInfo)
		}
	}

	return nil
}
