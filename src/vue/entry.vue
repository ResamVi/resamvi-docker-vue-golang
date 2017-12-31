<!-- 
    An Entry consists of 
    - a date string in format DD.MONTHNAME.YYYY
    - a title string
    - the HTML body string

    "entry" is the json object with Date, Title, Text, Number properties.
    Names are case sensitive
-->
<template>
  <section>
    <span class="entry-date">{{ entry.Date }}</span>
    <h1 class="entry-title" v-html="entry.Title"></h1>
    <div class="entry-content" v-html="entry.Text"></div>
    <div class="footer">
        <a href="gaestebuch.html" id="button">
          <div title="Zum Gästebuch">
            <img src="img/book.jpg" alt="door">
          </div>
        </a>
        <a href="impressum.html" id="button">
          <div title="Zurück zur Homepage">
            <img src="img/door.jpg" alt="door">
          </div>
        </a>
      </div>
  </section>
</template>

<!--
    The data object consists of a single "entry" object
    of which the data is fetched from the database
    and stored in the appropriate properties "Date", "Title", "Text"

    By sending an HTTP POST request to localhost:8080 the associated go server
    responsible for fetching entries off the mongodb database is called and
    returns the date/title/content
-->
<script>
export default
{
    data: function()
    {
        /**
         * Get entry data from database.
         * Access the count variable from the parent (entire entry blog)
         * to get the next entry in line to load
         */
        const DONE = 4;
        const SUCCESS = 200;

        /**
         * Send a http post request to the service responsible for
         * fetching entries out of the database.
         */
        var http = new XMLHttpRequest();
        var url = "http://localhost:8080/";
        var params = this.$parent.count;

        http.open("POST", url);
        http.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
        
        /*
        * The response is a JSON object with 
        * Number, Text, Date, Title properties (see server.go)
        */
        http.onload = function ()
        {
            if (http.readyState == DONE && http.status == SUCCESS)
            {
                var entry = JSON.parse(http.responseText);
                this.entry.Date = entry.Date;
                this.entry.Title = entry.Title;
                this.entry.Text = entry.Text;
            }
        }.bind(this);
        
        http.send(params);
        
        /**
         * While the entry is fetched from database (the code above is async)
         * display nothing.
         */
        return {
            entry : {Date: "Loading", Title:"Loading", Text:"Loading"},
        }
    }
}
</script>

<style>
    .entry-date {
        font-family: 'Cinzel';
        color: #696969;
    }

    .entry-title {
        margin-top:0px;
    }
    
    .entry-content {
        font-family: 'Open Sans';
        margin-bottom: 10%;
        font-size: 1.5em;
    }
</style>