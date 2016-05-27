// by setting package as main, Go will compile this as an executable file.
// Any other package turns this into a library
package main

// These are your imports / libraries / frameworks
import (
	// this is Go's built-in sql library
	"database/sql"
	"log"
	"net/http"
	"os"
	//"strconv"

	// this allows us to run our web server
	"github.com/gin-gonic/gin"
	// this lets us connect to Postgres DB's
	_ "github.com/lib/pq"
)

var (
	// this is the pointer to the database we will be working with
	// this is a "global" variable (sorta kinda, but you can use it as such)
	db *sql.DB
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	var errd error
	// here we want to open a connection to the database using an environemnt variable.
	// This isn't the best technique, but it is the simplest one for heroku
	db, errd = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if errd != nil {
		log.Fatalf("Error opening database: %q", errd)
	}
	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("html/*")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.GET("/ping", func(c *gin.Context) {
		ping := db.Ping()
		if ping != nil {
			// our site can't handle http status codes, but I'll still put them in cause why not
			c.JSON(http.StatusOK, gin.H{"error": "true", "message": "db was not created. Check your DATABASE_URL"})
		} else {
			c.JSON(http.StatusOK, gin.H{"error": "false", "message": "db created"})
		}
	})

	router.GET("/allTrips", func(c *gin.Context) {
	    rows, err := db.Query("SELECT member.username, member.picture, trip.name, trip.description, trippoint.date FROM trippoint JOIN trip ON trippoint.tripid = trip.id JOIN member ON member.id = trip.memberid WHERE trippoint.id IN (SELECT id FROM (SELECT DISTINCT ON (tripid) tripid, id FROM trippoint) AS tripid) ORDER BY date DESC")
	        if err != nil {
	            c.AbortWithError(http.StatusInternalServerError, err)
	            return
	        }
	        // if you are simply inserting data you can stop here. I'd suggest returning a JSON object saying "insert successful" or something along those lines.
	        // get all the columns. You can do something with them here if you like, such as adding them to a table header, or adding them to the JSON
	        cols, _ := rows.Columns()
	        if len(cols) == 0 {
	            c.AbortWithStatus(http.StatusNoContent)
	            return
	        }
	        // This will hold an array of all values
	        // makes an array of size 1, storing strings (replace with int or whatever data you want to store)
	        output := make([]string, 1)

	    // The variable(s) here should match your returned columns in the EXACT same order as you give them in your query
	        var username string
	        var picture string
	        var name string
	        var description string
	        var date string

	        for rows.Next() {
	            rows.Scan(&username, &picture, &name, &description, &date)
	            // VERY important that you store the result back in output
	            output = append(output, username, picture, name, description, date)
	        }
	        //Finally, return your results to the user:
	    c.JSON(http.StatusOK, gin.H{"result": output})
	})

	router.POST("/trip", func(c *gin.Context) {
		triptitle := c.PostForm("triptitle")
	    rows, err := db.Query("SELECT * FROM trip_essentials WHERE trip = $1 ORDER BY trippointdate", triptitle)
	        if err != nil {
	            c.AbortWithError(http.StatusInternalServerError, err)
	            return
	        }
	        // if you are simply inserting data you can stop here. I'd suggest returning a JSON object saying "insert successful" or something along those lines.
	        // get all the columns. You can do something with them here if you like, such as adding them to a table header, or adding them to the JSON
	        cols, _ := rows.Columns()
	        if len(cols) == 0 {
	            c.AbortWithStatus(http.StatusNoContent)
	            return
	        }
	        // This will hold an array of all values
	        // makes an array of size 1, storing strings (replace with int or whatever data you want to store)
	        output := make([]string, 1)

	    // The variable(s) here should match your returned columns in the EXACT same order as you give them in your query
	        var member string
	        var pic string
	        var trip string
	        var tripinfo string
	        var trippointdate string
	        var trippointinfo string
	        var country string
	        var city string
	        var transportation string
	        for rows.Next() {
	            rows.Scan(&member, &pic, &trip, &tripinfo, &trippointdate, &trippointinfo, &country, &city, &transportation)
	            // VERY important that you store the result back in output
	            output = append(output, member, pic, trip, tripinfo, trippointdate, trippointinfo, country, city, transportation)
	        }
	        //Finally, return your results to the user:
	    c.JSON(http.StatusOK, gin.H{"result": output})
	})

	router.POST("/addtrip", func(c *gin.Context) {
		name := c.PostForm("name")
		description := c.PostForm("description")
		_, err := db.Query("INSERT INTO trip(memberid, name, description) VALUES(1, $1, $2)", name, description)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, gin.H{"result": "success", "trip": name, "user": "gcampbell0"})
	})


	router.POST("/addtrippoint", func(c *gin.Context) {
		date := c.PostForm("date")
		trippointdescription := c.PostForm("trippointdescription")
		address1 := c.PostForm("address1")
		city := c.PostForm("city")
		country := c.PostForm("country")
		transportationtype := c.PostForm("transportationtype")
		transportationcost := c.PostForm("transportationcost")
		transportation := c.PostForm("transportation")
		_, err := db.Query("SELECT * FROM set_trippoint($1, $2, $3, $4, $5, $6, $7, $8)", date, trippointdescription, address1, city, country, transportationtype, transportationcost, transportation)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, gin.H{"result": "success", "trippoint": trippointdescription, "user": "gcampbell0", "location": city + ", " + country})
	})


	// NO code should go after this line. it won't ever reach that point
	router.Run(":" + port)
}

/*
Example of processing a GET request
// this will run whenever someone goes to last-first-lab7.herokuapp.com/EXAMPLE
router.GET("/EXAMPLE", func(c *gin.Context) {
    // process stuff
    // run queries
    // do math
    //decide what to return
    c.JSON(http.StatusOK, gin.H{
        "key": "value"
        }) // this returns a JSON file to the requestor
    // look at https://godoc.org/github.com/gin-gonic/gin to find other return types. JSON will be the most useful for this
})
*/